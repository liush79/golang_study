package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func testSync() {
	fmt.Println("---------------- test sync ---------------")
	fmt.Println("NumCPU:", runtime.NumCPU())
	_noMutex()
	_mutex()
	rwMutex()
	fmt.Println(" condition -------")
	testCond()
	fmt.Println(" test Once -------")
	testOnce()
	testPool()
	fmt.Println("--------------- finish test sync -----------")
}

func _noMutex() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용
	var data = []int{}                   // int형 슬라이스 생성

	go func() { // 고루틴에서
		for i := 0; i < 10000; i++ { // 1000번 반복하면서
			data = append(data, 1) // data 슬라이스에 1을 추가

			runtime.Gosched() // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	go func() { // 고루틴에서
		for i := 0; i < 10000; i++ { // 1000번 반복하면서
			data = append(data, 1) // data 슬라이스에 1을 추가

			runtime.Gosched() // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	time.Sleep(2 * time.Second) // 2초 대기

	fmt.Println(len(data)) // data 슬라이스의 길이 출력
}

func _mutex() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용
	var data = []int{}
	var mutex = new(sync.Mutex)

	go func() { // 고루틴에서
		for i := 0; i < 10000; i++ { // 1000번 반복하면서
			mutex.Lock()           // 뮤텍스 잠금, data 슬라이스 보호 시작
			data = append(data, 1) // data 슬라이스에 1을 추가
			mutex.Unlock()         // 뮤텍스 잠금 해제, data 슬라이스 보호 종료

			runtime.Gosched() // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	go func() { // 고루틴에서
		for i := 0; i < 10000; i++ { // 1000번 반복하면서
			mutex.Lock()           // 뮤텍스 잠금, data 슬라이스 보호 시작
			data = append(data, 1) // data 슬라이스에 1을 추가
			mutex.Unlock()         // 뮤텍스 잠금 해제, data 슬라이스 보호 종료

			runtime.Gosched() // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	time.Sleep(2 * time.Second) // 2초 대기

	fmt.Println(len(data)) // data 슬라이스의 길이 출력
}

func rwMutex() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data int = 0
	var rwMutex = new(sync.RWMutex) // 읽기, 쓰기 뮤텍스 생성

	go func() { // 값을 쓰는 고루틴
		for i := 0; i < 3; i++ {
			rwMutex.Lock()                    // 쓰기 뮤텍스 잠금, 쓰기 보호 시작
			data += 1                         // data에 값 쓰기
			fmt.Println("write  : ", data)    // data 값을 출력
			time.Sleep(10 * time.Millisecond) // 10 밀리초 대기
			rwMutex.Unlock()                  // 쓰기 뮤텍스 잠금 해제, 쓰기 보호 종료
		}
	}()

	go func() { // 값을 읽는 고루틴
		for i := 0; i < 3; i++ {
			rwMutex.RLock()                // 읽기 뮤텍스 잠금, 읽기 보호 시작
			fmt.Println("read 1 : ", data) // data 값을 출력(읽기)
			time.Sleep(1 * time.Second)    // 1초 대기
			rwMutex.RUnlock()              // 읽기 뮤텍스 잠금 해제, 읽기 보호 종료
		}
	}()

	go func() { // 값을 읽는 고루틴
		for i := 0; i < 3; i++ {
			rwMutex.RLock()                // 읽기 뮤텍스 잠금, 읽기 보호 시작
			fmt.Println("read 2 : ", data) // data 값을 출력(읽기)
			time.Sleep(2 * time.Second)    // 2초 대기
			rwMutex.RUnlock()              // 읽기 뮤텍스 잠금 해제, 읽기 보호 종료
		}
	}()

	time.Sleep(10 * time.Second) // 10초 동안 프로그램 실행
}

func testCond() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var mutex = new(sync.Mutex)    // 뮤텍스 생성
	var cond = sync.NewCond(mutex) // 뮤텍스를 이용하여 조건 변수 생성

	c := make(chan bool, 3) // 비동기 채널 생성

	for i := 0; i < 3; i++ {
		go func(n int) { // 고루틴 3개 생성
			mutex.Lock() // 뮤텍스 잠금, cond.Wait() 보호 시작
			c <- true    // 채널 c에 true를 보냄
			fmt.Println("wait begin : ", n)
			cond.Wait() // 조건 변수 대기
			fmt.Println("wait end : ", n)
			mutex.Unlock() // 뮤텍스 잠금 해제, cond.Wait() 보호 종료

		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c // 채널에서 값을 꺼냄, 고루틴 3개가 모두 실행될 때까지 기다림
	}

	for i := 0; i < 3; i++ {
		mutex.Lock() // 뮤텍스 잠금, cond.Signal() 보호 시작
		fmt.Println("signal : ", i)
		cond.Signal()  // 대기하고 있는 고루틴을 하나씩 깨움
		mutex.Unlock() // 뮤텍스 잠금 해제, cond.Signal() 보고 종료
	}

	fmt.Scanln()
}

func helloOnce() {
	fmt.Println("helloOnce world")
}

func testOnce() {
	once := new(sync.Once)
	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Println("goroutine: ", n)
			once.Do(helloOnce)
		}(i)
	}
}

type Data struct { // Data 구조체 정의
	tag    string // 풀 태그
	buffer []int  // 데이터 저장용 슬라이스
}

func testPool() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	pool := sync.Pool{ // 풀 할당
		New: func() interface{} { // Get 함수를 사용했을 때 호출될 함수 정의
			data := new(Data)             // 새 메모리 할당
			data.tag = "new"              // 태그 설정
			data.buffer = make([]int, 10) // 슬라이스 공간 할당
			return data                   // 할당한 메모리(객체) 리턴
		},
	}

	for i := 0; i < 10; i++ {
		go func() { // 고루틴 10개 생성
			data := pool.Get().(*Data) // 풀에서 *Data 타입으로 데이터를 가져옴
			for index := range data.buffer {
				data.buffer[index] = rand.Intn(100) // 슬라이스에 랜덤 값 저장
			}
			fmt.Println(data) // data 내용 출력
			data.tag = "used" // 객체가 사용되었다는 태그 설정
			pool.Put(data)    // 풀에 객체를 보관
		}()
	}

	for i := 0; i < 10; i++ {
		go func() { // 고루틴 10개 생성
			data := pool.Get().(*Data) // 풀에서 *Data 타입으로 데이터를 가져옴
			n := 0
			for index := range data.buffer {
				data.buffer[index] = n // 슬라이스에 짝수 저장
				n += 2
			}
			fmt.Println(data) // data 내용 출력
			data.tag = "used" // 객체가 사용되었다는 태그 설정
			pool.Put(data)    // 풀에 객체 보관
		}()
	}

	fmt.Scanln()
}
