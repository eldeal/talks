package main

import (
	"flag"
	"log"
	"time"
)

// START OMIT
func main() {
	s := flag.Int("seconds", 2, "an int") // HL
	flag.Parse()                          // HL

	t := time.Now() // OMIT
	for i := 0; i < 10; i++ {
		doWork(*s) // HL
	}

	d := time.Now().Sub(t)         // OMIT
	log.Println("Total time: ", d) // OMIT
}

func doWork(seconds int) { // HL
	time.Sleep(time.Duration(seconds) * time.Second)
}

// END OMIT
