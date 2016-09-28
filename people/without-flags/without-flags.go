package main

import (
	"log"
	"time"
)

// START OMIT
func main() {
	s := 2

	t := time.Now() // OMIT
	for i := 0; i < 10; i++ {
		doWork(s) // HL
	}

	d := time.Now().Sub(t)         // OMIT
	log.Println("Total time: ", d) // OMIT
}

func doWork(seconds int) { // HL
	time.Sleep(time.Duration(seconds) * time.Second)
}

// END OMIT
