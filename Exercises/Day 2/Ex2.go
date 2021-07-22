package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func addStudentRating(totalRating *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(100))*time.Millisecond)
	atomic.AddInt64(totalRating, int64(rand.Intn(6)))
}

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	var totalRating int64
	for i:=0; i<200; i++ {
		wg.Add(1)
		go addStudentRating(&totalRating, &wg)
	}
	wg.Wait()
	fmt.Printf("Final rating = %f out of 5\n", float64(totalRating)/200.0)
}
