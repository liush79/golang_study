package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func testFiles() {
	writeFile()
	readFile()
	fmt.Println("------ read line test --------")
	readLine()
}

func writeFile() {
	fp, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	s := "hello world ~~~ \nhello world2\nhello world3\nmulti line test\nhaha"
	n, err := fp.Write([]byte(s))
	fmt.Println(n, "bytes 저장완료")
}

func readFile() {
	file, err := os.Open("hello.txt") // hello.txt 파일을 열기
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	fi, err := file.Stat() // 파일 정보를 가져오기
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fi.Size()) // 파일 크기만큼 바이트 슬라이스 생성

	n, err := file.Read(data) // 파일의 내용을 읽어서 바이트 슬라이스에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 읽기 완료")
	fmt.Println(string(data)) // 문자열로 변환하여 data의 내용 출력
}

func readLine() {
	file, err := os.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
