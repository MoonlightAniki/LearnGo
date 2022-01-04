package main

import "fmt"

func main() {
	fmt.Println("Creating map")
	map1 := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	fmt.Println(map1)

	map2 := make(map[string]int) // Zero value is empty map
	fmt.Println(map2, map2 == nil)

	var map3 = make(map[string]int) // Zero value is empty map
	fmt.Println(map3, map3 == nil)

	var map4 map[string]int // Zero value is nil
	fmt.Println(map4, map4 == nil)

	fmt.Println("======================")
	fmt.Println("Traversing map")
	// 遍历的顺序是随机的，说明go语言中的map是HashMap
	for k, v := range map1 {
		fmt.Printf("k=%s, v=%s\n", k, v)
	}

	for k, _ := range map1 {
		fmt.Printf("k=%s, v=%s\n", k, map1[k])
	}

	for _, v := range map1 {
		fmt.Printf("v=%s\n", v)
	}

	fmt.Println("======================")
	fmt.Println("Getting values")
	name := map1["name"]
	course := map1["course"]
	site := map1["site"]
	quality := map1["quality"]
	fmt.Printf("name=%s, course=%s, site=%s, quality=%s\n", name, course, site, quality)

	cause := map1["cause"]
	fmt.Println(cause, cause == "")

	cause, ok := map1["cause"]
	fmt.Println(cause, ok)

	if name, ok := map1["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("Key does not exist")
	}

	fmt.Println("======================")
	fmt.Println("Deleting value")

	name, ok = map1["name"]
	fmt.Println(name, ok)

	delete(map1, "name")

	name, ok = map1["name"]
	fmt.Println(name, ok)

	fmt.Println("======================")
	fmt.Println("Setting values")
	map1["name"] = "guoliang03"
	name, ok = map1["name"]
	fmt.Println(name, ok)
}
