package main

import "testing"

func TestProblem001(t *testing.T) {
	output := make(chan intresult)
	go problem001(1000, output)
	result := <-output
	if result.result != 233168 {
		t.Error("problem #1: Wrong Answer", result.result)
	}
}

func TestProblem001TestCase(t *testing.T) {
	output := make(chan intresult)
	go problem001(10, output)
	result := <-output
	if result.result != 23 {
		t.Error("Problem #1 Test Case: Wrong Answer", result.result)
	}
}

func BenchmarkProblem001(b *testing.B) {
	output := make(chan intresult)
	go problem001(1000000, output)
	_ = <-output
}

func TestProblem002(t *testing.T) {
	output := make(chan intresult)
	go problem002(4000000, output)
	result := <-output
	if result.result != 4613732 {
		t.Error("Problem #2: Wrong Answer", result.result)
	}
}

func BenchmarkProblem002(b *testing.B) {
	output := make(chan intresult)
	go problem002(400000000, output)
	_ = <-output
}

func TestProblem002FirstTenFib(t *testing.T) {
	thelist := []int{}
	output := make(chan int)
	go fib(1, 2, output)
	for i := 0; i < 10; i++ {
		thelist = append(thelist, <-output)
	}
	good := true
	for i := 0; i < 10; i++ {
		if thelist[i] != [10]int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}[i] {
			good = false
		}
	}
	if !good {
		t.Error("Problem #2: Fibonacci Generator Broken", thelist)
	}
}

func TestProblem003(t *testing.T) {
	output := make(chan intresult)
	go problem003(600851475143, output)
	result := <-output
	if result.result != 6857 {
		t.Error("Problem #3: Wrong Answer", result.result)
	}
}

func TestProblem003TestCase(t *testing.T) {
	output := make(chan intresult)
	go problem003(13195, output)
	result := <-output
	if result.result != 29 {
		t.Error("Problem #3 Test Case: Wrong Answer", result.result)
	}
}

func TestProblem003TestFactors(t *testing.T) {
	thelist := factors(13195)
	if len(thelist) != 4 {
		t.Error("Problem #3: Wrong Number of Factors", thelist)
	}
	good := true
	for i := 0; i < 4; i++ {
		if thelist[i] != [4]int{5, 7, 13, 29}[i] {
			good = false
		}
	}
	if !good {
		t.Error("Problem #3: Factor Function Broken")
	}
}

func TestProblem004(t *testing.T) {
	output := make(chan intresult)
	go problem004(3, output)
	result := <-output
	if result.result != 906609 {
		t.Error("Problem #4: Wrong Answer", result.result)
	}
}

func TestProblem004TestCase(t *testing.T) {
	output := make(chan intresult)
	go problem004(2, output)
	result := <-output
	if result.result != 9009 {
		t.Error("Problem #4 Test Case: Wrong Answer", result.result)
	}
}

func TestProblem005(t *testing.T) {
	output := make(chan intresult)
	go problem005(20, output)
	result := <-output
	if result.result != 232792560 {
		t.Error("Problem #5: Wrong Answer", result.result)
	}
}

func TestProblem005TestCase(t *testing.T) {
	output := make(chan intresult)
	go problem005(10, output)
	result := <-output
	if result.result != 2520 {
		t.Error("Problem #5 Test Case: Wrong Answer", result.result)
	}
}
