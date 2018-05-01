package main

import "fmt"
import "math/rand"
import "time"

func gohello() {
	fmt.Println("HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHhhello world!!")
}

func goRoutineTest(n int) {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r))
	fmt.Println("Select", n, ",", r)
}

func goSum(a int, b int, c chan int) {
	c <- a + b
	c <- a - b
}
