package main

import (
	"LearnGo/basics/retriever/mock"
	"LearnGo/basics/retriever/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) string {
	return poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

// 接口的组合
type PosterRetriever interface {
	Retriever
	Poster
	// 还可以增加方法
	//Connect()
}

const url = "https://www.imooc.com"

func session(s PosterRetriever) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
	//s.Connect()
}

func main() {
	var r Retriever
	//r = mock.Retriever{"this is a fake imooc.com"}
	retriever := &mock.Retriever{"this is a fake imooc.com"}
	inspect(retriever)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	// Type Assertion
	//realRetriever := r.(*real.Retriever)
	//fmt.Println(realRetriever.TimeOut)
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println("Contents:", mockRetriever.Contents)
	} else {
		fmt.Println("Not a mock retriever.")
	}
	//fmt.Println(download(r))

	fmt.Println("Try a session")
	fmt.Println(session(retriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Printf(" > Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Conetnts:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}

// 接口变量自带指针

// 接口变量同样采用值传递，几乎不需要使用接口的指针

// 指针接受者实现只能以指针方式使用；值接受者都可。

// 查看接口变量：Type Assertion 、 Type Swtich
// 表示任何类型 interface{}
