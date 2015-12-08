package main

import "fmt"
import "flag"
import "strconv"

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

func problem006(upperlimit int, output chan intresult) {
	sumofsquares := make(chan int)
	squareofsums := make(chan int)
	go func(upperlimit int, output chan int) {
		result := 0
		for i := 1; i <= upperlimit; i++ {
			result += i * i
		}
		output <- result
	}(upperlimit, sumofsquares)
	go func(upperlimit int, output chan int) {
		result := 0
		for i := 1; i <= upperlimit; i++ {
			result += i
		}
		output <- result * result
	}(upperlimit, squareofsums)
	output <- intresult{6, (<-squareofsums) - (<-sumofsquares)}
}

func is_prime(i int) bool {
	if i < 2 {
		return false
	}
	for j := 2; j*j <= i; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}

func problem007(upperlimit int, output chan intresult) {
	primes := []int{}
	for i := 2; len(primes) < upperlimit; i++ {
		if is_prime(i) {
			primes = append(primes, i)
		}
	}
	output <- intresult{7, primes[len(primes)-1]}
}

func problem008(length int, output chan intresult) {
	fullnumber := "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"
	answer := 0
	for i := 0; i+length < len(fullnumber); i++ {
		temp := 1
		for j := 0; j < length; j++ {
			temp2, _ := strconv.Atoi(fullnumber[i+j : i+j+1])
			temp *= temp2
		}
		if temp > answer {
			answer = temp
		}
	}
	output <- intresult{8, answer}
}

func problem009(target int, output chan intresult) {
	for a := 1; a <= target; a++ {
		for b := a + 1; a+b <= target; b++ {
			for c := b + 1; a+b+c <= target; c++ {
				if (a+b+c == target) && (a*a+b*b == c*c) {
					output <- intresult{9, a * b * c}
				}
			}
		}
	}
}

func primesunder(upperlimit int) []int {
	array := []int{}
	for i := 0; i < upperlimit; i++ {
		array = append(array, 0)
	}
	array[0] = 1
	array[1] = 1
	result := []int{}
	for i := 2; i < upperlimit; i++ {
		if array[i] == 0 {
			result = append(result, i)
			for j := i * i; j < upperlimit; j += i {
				array[j] = 1
			}
		}
	}
	return result
}

func sum(thelist []int) int {
	total := 0
	for i := 0; i < len(thelist); i++ {
		total += thelist[i]
	}
	return total
}

func problem010(upperlimit int, output chan intresult) {
	output <- intresult{10, sum(primesunder(upperlimit))}
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
	if problemnumber == 0 || problemnumber == 6 {
		count += 1
		go problem006(100, intanswers)
	}
	if problemnumber == 0 || problemnumber == 7 {
		count += 1
		go problem007(10001, intanswers)
	}
	if problemnumber == 0 || problemnumber == 8 {
		count += 1
		go problem008(13, intanswers)
	}
	if problemnumber == 0 || problemnumber == 9 {
		count += 1
		go problem009(1000, intanswers)
	}
	if problemnumber == 0 || problemnumber == 10 {
		count += 1
		go problem010(2000000, intanswers)
	}
	for count > 0 {
		temp := <-intanswers
		fmt.Println(temp.problemnumber, temp.result)
		count -= 1
	}
}
