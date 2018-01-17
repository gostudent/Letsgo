---
published: true
author: surasnayak
layout: post
title: Implementation of Sieve of Eratosthenes in GO
---

## Sieve of Eratosthenes

- Sieve of Eratosthenes is a simple and ancient algorithm used to find the prime numbers up to any given limit. 
- It is one of the most efficient ways to find small prime numbers.

- For a given upper limit  the algorithm works by iteratively marking the multiples of primes as composite, starting from 2. 
- Once all multiples of 2 have been marked composite, the muliples of next prime, ie 3 are marked composite. 
- This process continues until  where  is a prime number.

## Implementation :

- In the following algorithm, the number 0 represents a composite number.

- To find out all primes under , generate a list of all integers from 2 to n. (Note: 1 is not prime)
- Start with a smallest prime number, ie :
- Mark all the multiples of  which are less than  as composite. 
- To do this, mark the value of the numbers (multiples of ) in the generated list as 0. 
- Do not mark  itself as composite.
Assign the value of  to the next prime. 
- The next prime is the next non-zero number in the list which is greater than p.
- Repeat the process until .

We are done!


### Proof: To see why Sieve of Eratosthenes generates all the primes

- To first understand why the sieve works, let's look into the definitions of prime decomposition and composite numbers.

- Theorem 1

- Every integer which is greater than 1 can be expressed as product of primes.

- Theorem 2

- If any integer  is composite, then  has a prime divisor less than or equal to .

- Since for any number  in the list, we are looking all prime numbers up to , we are indeed separating all composite numbers.
 Hence Sieve of Eratosthenes generates all primes numbers less than the upper limit. 
 
 ## Code for implementaion of queue data structure in GO :
 
 ```go
 package main

import (
	"fmt"
)

//This program only works for primes smaller or equal to 10e7
func Sieve(upperBound int64) []int64 {
	_SieveSize := upperBound + 10
	//Creates set to mark wich numbers are primes and wich are not
	//true: not primes, false: primes
	//this to favor default initialization of arrays in go
	var bs [10000010]bool
	//creates a slice to save the primes it finds
	primes := make([]int64, 0, 1000)
	
	bs[0] = true
	bs[1] = true
	//looping over the numbers set
	for i := int64(0); i <= _SieveSize; i++ {
		//if one number is found that is unmarked as a compund number, mark all its multiples
		if !bs[i] {
			for j := i * i; j <= _SieveSize; j += i {
				bs[j] = true
			}
			//Add the prime you just find to the slice of primes
			primes = append(primes, i)
		}
	}
	return primes
}

/*func main() {
	//prints first N primes into console
	N := 100
	primes := Sieve(N)
	fmt.Println(primes)
}*/
 
 ```
