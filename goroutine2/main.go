package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// waitGroup

func f() {

	rand.Seed(time.Now().Unix())
	for i := 0; i < 5; i++ {
		r2 := rand.Intn(10)
		fmt.Println(r2)
	}
}

func f1(i int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	fmt.Println(i)
	defer wg.Done()
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go f1(i)
	}
	wg.Wait()
}
