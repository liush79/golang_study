package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Author struct {
	Name  string `json:"name"` // 구조체 필드에 태그 지정
	Email string `json:"email"`
}

type Comment struct {
	Id      uint64 `json:"id"`
	Author  Author `json:"author"`
	Content string `json:"content"`
}

type Article struct {
	Id         uint64    `json:"id"`
	Title      string    `json:"title"`
	Author     Author    `json:"author"`
	Content    string    `json:"content"`
	Recommends []string  `json:"recommends"`
	Comments   []Comment `json:"comments"`
}

func testJson() {
	writeJson()
	readJson()
}

func writeJson() {
	data := make([]Article, 1) // 값을 저장할 구조체 슬라이스 생성

	data[0].Id = 1
	data[0].Title = "Hello, world!"
	data[0].Author.Name = "Maria"
	data[0].Author.Email = "maria@example.com"
	data[0].Content = "Hello~"
	data[0].Recommends = []string{"John", "Andrew"}
	data[0].Comments = make([]Comment, 1)
	data[0].Comments[0].Id = 1
	data[0].Comments[0].Author.Name = "Andrew"
	data[0].Comments[0].Author.Email = "andrew@hello.com"
	data[0].Comments[0].Content = "Hello Maria"

	doc, _ := json.Marshal(data) // data를 JSON 문서로 변환

	err := ioutil.WriteFile("./articles.json", doc, os.FileMode(0644)) // articles.json 파일에 JSON 문서 저장
	if err != nil {
		fmt.Println(err)
		return
	}
}

func readJson() {
	b, err := ioutil.ReadFile("./articles.json") // articles.json 파일의 내용을 읽어서 바이트 슬라이스에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	var data []Article // JSON 문서의 데이터를 저장할 구조체 슬라이스 선언

	json.Unmarshal(b, &data) // JSON 문서의 내용을 변환하여 data에 저장

	fmt.Println(data)        // [{1 Hello, world! {Maria maria@exa... (생략)
	json.Unmarshal(b, &data) // JSON 문서의 내용을 변환하여 data에 저장
}
