package main

import "fmt"
import "flag"

type intresult struct {
	problemnumber int
	result        int
}

func problem001a(upperlimit int, output chan intresult) {
	total := 0
	for i := 0; i < upperlimit; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	output <- intresult{1, total}
}

func problem001b(upperlimit int, output chan intresult) {
	upperlimit -= 1
	output <- intresult{1, 3*(upperlimit/3)*(upperlimit/3+1)/2 + 5*(upperlimit/5)*(upperlimit/5+1)/2 - 15*(upperlimit/15)*(upperlimit/15+1)/2}
}

func problem001(upperlimit int, output chan intresult) {
	fastest := make(chan intresult)
	go problem001a(upperlimit, fastest)
	go problem001b(upperlimit, fastest)
	output <- <-fastest
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

func reverse_int(i int) int {
	answer := 0
	for i > 0 {
		answer *= 10
		answer += i % 10
		i /= 10
	}
	return answer
}

func is_palindrome(i int) bool {
	if i == reverse_int(i) {
		return true
	}
	return false
}

// https://groups.google.com/forum/#!msg/golang-nuts/PnLnr4bc9Wo/fvp154Hms2QJ
func Pow(a, b int) int {
	var result int = 1

	for 0 != b {
		if 0 != (b & 1) {
			result *= a
		}
		b >>= 1
		a *= a
	}

	return result
}

func problem004(numberlength int, output chan intresult) {
	answer := 0
	for i := Pow(10, numberlength-1); i < Pow(10, numberlength); i++ {
		for j := i; j < Pow(10, numberlength); j++ {
			if i*j > answer && is_palindrome(i*j) {
				answer = i * j
			}
		}
	}
	output <- intresult{4, answer}
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a int, b int) int {
	return a * b / gcd(a, b)
}

func problem005(upperlimit int, output chan intresult) {
	answer := 1
	for i := 1; i <= upperlimit; i++ {
		answer = lcm(answer, i)
	}
	output <- intresult{5, answer}
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
	if problemnumber == 0 || problemnumber == 4 {
		count += 1
		go problem004(3, intanswers)
	}
	if problemnumber == 0 || problemnumber == 5 {
		count += 1
		go problem005(20, intanswers)
	}
	for count > 0 {
		temp := <-intanswers
		fmt.Println(temp.problemnumber, temp.result)
		count -= 1
	}
}
