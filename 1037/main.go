package main

func main() {

}

func isBoomerang(points [][]int) bool {
	ax, ay := points[0][0], points[0][1]
	bx, by := points[1][0], points[1][1]
	cx, cy := points[2][0], points[2][1]

	return (ay-by)*(ax-cx) == (ay-cy)*(ax-bx)
}
