package main

import (
	"log"
	"time"
)

// START OMIT
func main() {
	t := time.Now() // OMIT
	for i := 0; i < 10; i++ {
		doWork() // HL
	}
	d := time.Now().Sub(t)         // OMIT
	log.Println("Total time: ", d) // OMIT
}

func doWork() {
	time.Sleep(2 * time.Second)
}

// END OMIT
