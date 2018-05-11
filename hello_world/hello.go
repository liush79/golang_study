package main

import (
	"fmt"
)

func main() {
	// // for i := 1; i < 100; i++ {
	// for i := 1; i < 10; i++ {
	// 	switch {
	// 	case i%3 == 0 && i%5 == 0:
	// 		fmt.Println("FizzBuzz")
	// 	case i%3 == 0:
	// 		fmt.Println("Fizz")
	// 	case i%5 == 0:
	// 		fmt.Println("Buzz")
	// 	default:
	// 		fmt.Println(i)
	// 	}
	// }

	// // for i := 9; i > 0; i-- {
	// // 	fmt.Printf("%d bottles of beer on the wall, %d bottles of beer.\n", i, i)
	// // 	fmt.Printf("Take one down, pass it around, ")
	// // 	if i > 1 {
	// // 		fmt.Printf("%d bottles of beer on the wall.\n", i)
	// // 	} else {
	// // 		fmt.Println("No more bottles of beer on the wall.")
	// // 		fmt.Println("No more bottles of beer on the wall, No more bottles of beer.")
	// // 	}
	// // }
	// a := []int{1, 2, 3}
	// a = append(a, 4, 5)
	// fmt.Println(a)

	// b := make([]int, 3)
	// copy(b, a)
	// b[0] = 9
	// fmt.Println(a)
	// fmt.Println(b)

	// c := []int{1, 2, 3, 4}
	// fmt.Println(len(c), cap(c))
	// c = append(c, 5, 6)
	// fmt.Println(len(c), cap(c))

	// var aa map[string]int = make(map[string]int)
	// var bb = make(map[string]int)
	// cc := make(map[string]int)
	// fmt.Println(aa, bb, cc)

	// aaa := map[string]int{"hello": 10, "world": 20}
	// fmt.Println(aaa)

	// solarSystem := make(map[string]float32) // 키는 string, 값은 float32인 맵 생성 및 공간 할당

	// solarSystem["Mercury"] = 87.969 // 맵[키] = 값
	// solarSystem["Venus"] = 224.70069
	// solarSystem["Earth"] = 365.25641
	// solarSystem["Mars"] = 686.9600
	// solarSystem["Jupiter"] = 4333.2867
	// solarSystem["Saturn"] = 10756.1995
	// solarSystem["Uranus"] = 30707.4896
	// solarSystem["Neptune"] = 60223.3528

	// fmt.Println(solarSystem["Earth"]) // 365.25641

	// for k, v := range solarSystem {
	// 	fmt.Println(k, v)
	// }

	// fmt.Println("DDDDDDDDDDDDD")
	// fmt.Println(SumAndDiff(12, 23))

	// var num int = 1
	// var numPtr *int = &num
	// var numPtr2 *int

	// numPtr2 = &num

	// fmt.Println(numPtr)
	// fmt.Println(numPtr2)

	// go gohello()
	// // for i := 0; i < 100; i++ {
	// // 	go goRoutineTest(i)
	// // }

	// var input string
	// fmt.Println("EEEEEEE")
	// fmt.Scanln(&input)
	// fmt.Println("FFFFFF")
	// fmt.Println(input)
	// fmt.Scanln()
	// channel()
	// channelBuffering()
	// time.Sleep(time.Second * 1)
	// fmt.Println("----------------- multi consumer test ---------------")
	// // consumer 1
	// ccc := make(chan int)
	// go createConsumer(ccc, 1)
	// // consumer 2
	// go createConsumer(ccc, 2)
	// // producer
	// for i := 0; i < 10; i++ {
	// 	ccc <- i
	// }
	// fmt.Println("NOT Support multi consumer -------------------------")
	// fmt.Println("----------------select test ------------------")
	// testSelect()
	// fmt.Println("--------------------------------------")
	// time.Sleep(time.Second * 1)
	// testSync()
	// fmt.Println("-------------------- rw Mutex -------")
	// time.Sleep(time.Second * 1)
	fmt.Println("-------------------- file  -------")
	testFiles()
	fmt.Println("Finish !")
}

func SumAndDiff(a int, b int) (c int, d int) {
	c = a - b
	d = a + b
	return
}

func channel() {
	c := make(chan int)
	go goSum(1, 2, c)
	n := <-c
	fmt.Println(n)
	n = <-c
	fmt.Println(n)
}

func channelBuffering() {
	fmt.Println("---------------------------------")
	done := make(chan bool, 2)
	count := 5
	go func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("go routing: ", i)
		}
	}()

	for i := 0; i < count; i++ {
		fmt.Println("Main function: ", i, <-done)
	}
}

func createConsumer(c chan int, num int) {
	for {
		fmt.Println("consumer ", num, ",", <-c)
	}
}
