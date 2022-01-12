package main

import (
	"LearnGo/data_structures/07-Set-and-Map/bstmap"
	"LearnGo/data_structures/07-Set-and-Map/bstset"
	"LearnGo/data_structures/07-Set-and-Map/linkedlistmap"
	"LearnGo/data_structures/07-Set-and-Map/linkedlistset"
	mymap "LearnGo/data_structures/07-Set-and-Map/map"
	"LearnGo/data_structures/07-Set-and-Map/set"
	"LearnGo/data_structures/utils"
	"fmt"
	"path/filepath"
	"time"
)

func main() {
	//compareSets()
	compareMaps()
}

func compareMaps() {
	filename, _ := filepath.Abs("pride-and-prejudice.txt")
	bstMap := bstmap.NewBSTMap()
	time1 := testMap(bstMap, filename)
	fmt.Printf("BSTMap cost time: %d ms\n", time1.Milliseconds())

	linkedListMap := linkedlistmap.NewLinkedListMap()
	time2 := testMap(linkedListMap, filename)
	fmt.Printf("LinkedListMap cost time: %d ms\n", time2.Milliseconds())
}

func testMap(m mymap.Map, filename string) time.Duration {
	words := utils.ReadFile(filename)
	fmt.Println("Total words:", len(words))
	startTime := time.Now()
	for _, word := range words {
		if m.Contains(word) {
			m.Set(word, m.Get(word).(int)+1)
		} else {
			m.Add(word, 1)
		}
	}
	fmt.Println("Total different words:", m.GetSize())
	fmt.Println("PRIDE frequency:", m.Get("pride"))
	fmt.Println("PREJUDICE frequency:", m.Get("prejudice"))
	return time.Now().Sub(startTime)
}

func compareSets() {
	filename, _ := filepath.Abs("a-tale-of-two-cities.txt")
	bstSet := bstset.NewBSTSet()
	time1 := testSet(bstSet, filename)
	fmt.Printf("BSTSet cost time: %d ms\n", time1.Milliseconds())

	linkedListSet := linkedlistset.NewLinkedListSet()
	time2 := testSet(linkedListSet, filename)
	fmt.Printf("LinkedListSet cost time: %d ms\n", time2.Milliseconds())
}

func testSet(set set.Set, filename string) time.Duration {
	words := utils.ReadFile(filename)
	fmt.Println("Total words:", len(words))
	startTime := time.Now()
	for _, word := range words {
		set.Add(word)
	}
	fmt.Println("Total different words:", set.GetSize())
	return time.Now().Sub(startTime)
}
