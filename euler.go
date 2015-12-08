package main

import "fmt"
import "flag"

type intresult struct {
	problemnumber int
	result        int
}

func problem001(upperlimit int, output chan intresult) {
	total := 0
	for i := 0; i < upperlimit; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	output <- intresult{1, total}
}

func fib(a int, b int, output chan int) {
	output <- a
	output <- b
	for {
		c := a + b
		output <- c
		a = b
		b = c
	}
}

func problem002(upperlimit int, output chan intresult) {
	total := 0
	results := make(chan int)
	go fib(1, 2, results)
	temp := <-results
	for temp < upperlimit {
		if temp%2 == 0 {
			total += temp
		}
		temp = <-results
	}
	output <- intresult{2, total}
}

func factors(i int) []int {
	result := []int{}
	for j := 2; i > 1; j++ {
		for i%j == 0 {
			result = append(result, j)
			i /= j
		}
	}
	return result
}

func problem003(i int, output chan intresult) {
	output <- intresult{3, factors(i)[len(factors(i))-1]}
}

func main() {
	var problemnumber int
	flag.IntVar(&problemnumber, "problem", 0, "problem number to solve")
	flag.Parse()
	count := 0
	intanswers := make(chan intresult)
	if problemnumber == 0 || problemnumber == 1 {
		count += 1
		go problem001(1000, intanswers)
	}
	if problemnumber == 0 || problemnumber == 2 {
		count += 1
		go problem002(4000000, intanswers)
	}
	if problemnumber == 0 || problemnumber == 3 {
		count += 1
		go problem003(600851475143, intanswers)
	}
	for count > 0 {
		temp := <-intanswers
		fmt.Println(temp.problemnumber, temp.result)
		count -= 1
	}
}
