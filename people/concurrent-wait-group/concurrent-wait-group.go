package main

import (
	"log"
	"sync"
	"time"
)

// START OMIT
var wg sync.WaitGroup // HL

func main() {
	t := time.Now() // OMIT
	for i := 0; i < 10; i++ {
		wg.Add(1)   // HL
		go doWork() // HL
	}
	wg.Wait()                      // HL
	d := time.Now().Sub(t)         // OMIT
	log.Println("Total time: ", d) // OMIT
}

func doWork() {
	time.Sleep(2 * time.Second)
	wg.Done() // HL
}

// END OMIT
