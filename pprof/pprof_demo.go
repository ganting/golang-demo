package main

import (
	"github.com/pkg/profile"
	"fmt"
	"time"
)

func main()  {
	c := profile.Start(profile.CPUProfile, profile.MemProfile, profile.ProfilePath("D:\\git_repo\\github\\ganting\\bin"))
	defer c.Stop()
	start := time.Now()
	defer func() {
		fmt.Printf("elapsed %d ns\n", time.Since(start).Nanoseconds())
	}()


	primes := sieve()
	for {
		p := <- primes
		if p > 50000 {
			break
		}
	}
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generator()
		for {
			prime := <-ch
			out<- prime
			ch = filter(ch, prime)
		}
	}()
	return out
}

func generator() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i % prime != 0 {
				out <- i
			}
		}
	}()
	return out
}