package main

import (
	"fmt"
	"sync"
)

type letterCount struct {
	counts map[string]int
	mu sync.Mutex
}

func (lc *letterCount) countLetters(word string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, letter := range word {
		lc.mu.Lock()
		lc.counts[string(letter)]++
		lc.mu.Unlock()
	}
}

func main() {
	//var words = []string{"quick", "brown", "foxy", "lazy", "dog"}
	var wordList []string
	var numberOfWords int
	fmt.Print("Enter number of words: ")
	fmt.Scan(&numberOfWords)
	fmt.Println("Enter list of words:")
	for i:=0; i<numberOfWords; i++ {
		var word string
		fmt.Scan(&word)
		wordList = append(wordList, word)
	}
	//fmt.Println(words)
	var wg sync.WaitGroup
	var letterCounter = letterCount{ counts: make(map[string]int)}
	for _, word := range wordList {
		wg.Add(1)
		go letterCounter.countLetters(word, &wg)
	}
	wg.Wait()
	fmt.Println(letterCounter.counts)
}
