package main

import (
	"fmt"
	"sort"
)

func main() {
	c := Constructor()

	//c.PostTweet(1, 5)
	//fmt.Println(c.GetNewsFeed(1))
	//
	//c.Follow(1, 2)
	//c.PostTweet(2, 6)
	//fmt.Println(c.GetNewsFeed(1))
	//
	//c.Unfollow(1, 2)
	//fmt.Println(c.GetNewsFeed(1))

	c.PostTweet(1, 4)
	c.PostTweet(2, 5)
	c.Unfollow(1, 2)
	fmt.Println(c.GetNewsFeed(1))
}

type Twitter struct {
	time int
	mp   map[int][]Node
	ma   map[int][]int
}

type Node struct {
	time  int
	value int
}

func Constructor() Twitter {
	r := Twitter{
		time: 0,
		mp:   make(map[int][]Node),
		ma:   make(map[int][]int),
	}
	return r
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	t.time++

	t.mp[userId] = append(t.mp[userId], Node{
		time:  t.time,
		value: tweetId,
	})
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	set := make(map[Node]bool)
	for _, v := range t.mp[userId] {
		set[v] = true
	}

	for _, u := range t.ma[userId] {
		for _, v := range t.mp[u] {
			set[v] = true
		}
	}

	l := make([]Node, 0)
	for k := range set {
		l = append(l, k)
	}

	sort.Slice(l, func(a, b int) bool {
		return l[a].time > l[b].time
	})

	r := make([]int, 0)
	for _, v := range l {
		r = append(r, v.value)
	}

	if len(l) < 10 {
		return r
	}
	return r[:10]
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	tma := append(t.ma[followerId], followeeId)
	sort.Ints(tma)
	t.ma[followerId] = tma
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	tma := t.ma[followerId]
	if len(tma) == 0 {
		return
	}

	i := sort.SearchInts(tma, followeeId)
	if tma[i] != followeeId {
		return
	}

	if i == len(tma)-1 {
		t.ma[followerId] = tma[:i]
		return
	}

	tma = append(tma[:i], tma[i+1:]...)
	t.ma[followerId] = tma
}
