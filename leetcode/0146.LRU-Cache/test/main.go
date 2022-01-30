package main

import solution0146 "LearnGo/leetcode/0146.LRU-Cache"

func main() {
	cache := solution0146.Constructor(2)
	cache.Put(1, 1)       // {1=1}
	cache.Put(2, 2)       // {2=2, 1=1}
	println(cache.Get(1)) // 1, {1=1, 2=2}
	cache.Put(3, 3)       // {3=3, 1=1}
	println(cache.Get(2)) // -1, {3=3, 1=1}
	cache.Put(4, 4)       // {4=4, 3=3}
	println(cache.Get(1)) // -1
	println(cache.Get(3)) // 3
	println(cache.Get(4)) // 4
}
