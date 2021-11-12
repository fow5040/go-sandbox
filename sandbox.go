package main

import (
	"fmt"
	"math/rand"
	"time"
)

type FuncIntInt func(int) int

func main() {
	fmt.Println("go-sandbox at ", time.Now(), " o-clock!!")

	randMin := 0
	randMax := 34
	fmt.Println(
		"I'm thinking of a number between ",
		randMin,
		" and ",
		randMax,
		"... what about: ",
		rand.Intn(randMax),
	)

	fib := Fib
	mFib := MemoizedFib()

	eval := 3
	fmt.Println("Calculating the ", eval, "nth fib number")

	s1 := time.Now()
	a1 := fib(eval)
	d1 := time.Since(s1)

	s2 := time.Now()
	a2 := mFib(eval)
	d2 := time.Since(s2)

	fmt.Println("Basic Fib: ", a1, " and took ", d1.Nanoseconds(), " nanoseconds.")
	fmt.Println("Memo Fib: ", a2, " and took ", d2.Nanoseconds(), " nanoseconds.")

	fmt.Println("\nExecuting udemyReviewMain!...")
	udemyReviewMain()

	fmt.Println("\nExecuting udemyTemplatesMain!...")
	udemyTemplatesMain()
}

func MemoizedFib() FuncIntInt {

	memo := make(map[int]int)

	// Need to declare it forward to allow recursion
	var fn func(int) int

	fn = func(num int) int {
		if val, found := memo[num]; found {
			return val
		}

		ans := 1
		if num > 2 {
			ans = fn(num-1) + fn(num-2)
		}

		memo[num] = ans
		return ans
	}

	return fn
}

func Fib(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return (Fib(n-1) + Fib(n-2))
}
