package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeComments([]string{"a/*/b//*c", "blank", "d/*/e*//f"}))
	fmt.Println(removeComments([]string{"a/*/b//*c", "blank", "d//*e/*/f"}))
	fmt.Println(removeComments([]string{"a//*b//*c", "blank", "d/*/e*//f"}))
	fmt.Println(removeComments([]string{"main() {", "   func(1);", "   /** / more comments here", "   float f = 2.0", "   f += f;", "   cout << f; */", "}"}))
	fmt.Println(removeComments([]string{"void func(int k) {", "// this function does nothing /*", "   k = k*2/4;", "   k = k/2;*/", "}"}))
	fmt.Println(removeComments([]string{"struct Node{", "    /*/ declare members;/**/", "    int size;", "    /**/int val;", "};"}))
	fmt.Println(removeComments([]string{"class test{", "public: ", "   int x = 1;", "   /*double y = 1;*/", "   char c;", "};"}))
	fmt.Println(removeComments([]string{"a/*comment", "line", "more_comment*/b"}))
	fmt.Println(removeComments([]string{"/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test", "   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"}))
}

//todo has bug
func removeComments(sl []string) []string {
	var r []string

	for i := 0; i < len(sl); {
		if index := strings.Index(sl[i], "//"); index != -1 {
			if gIndex := strings.Index(sl[i], "/*"); gIndex == -1 || gIndex > index {
				if sl[i][:index] != "" {
					r = append(r, sl[i][:index])
				}
				i++
				continue
			}
		}

		leftIndex := strings.Index(sl[i], "/*")
		if leftIndex == -1 {
			r = append(r, sl[i])
			i++
			continue
		}

		rightIndex := strings.LastIndex(sl[i], "*/")
		if rightIndex == leftIndex+1 {
			rightIndex = -1
		}
		if rightIndex != -1 {
			ts := sl[i][:leftIndex]
			if len(sl[i])-rightIndex > 2 {
				ts += sl[i][rightIndex+2:]
			}
			if ts != "" {
				r = append(r, ts)
			}
			i++
			continue
		}

		j := i + 1
		for j < len(sl) && strings.Index(sl[j], "*/") == -1 {
			j++
		}

		if j >= len(sl) {
			r = append(r, "[token error]")
			break
		}

		rightIndex = strings.Index(sl[j], "*/")
		ts := sl[i][:leftIndex]
		if len(sl[j])-rightIndex > 2 {
			ts += sl[j][rightIndex+2:]
			if index := strings.Index(ts, "//"); index != -1 {
				ts = ts[:index]
			}
		}

		if ts != "" {
			r = append(r, ts)
		}

		i = j + 1
	}

	return r
}
