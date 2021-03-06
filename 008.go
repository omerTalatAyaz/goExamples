// Girilen bir sayının palidrom sayı olup olmadığını ekrana yazan algoritmayı yapınız.

package main

import (
	"fmt"
)

func main() {
	var num, remainder, temp int
	var reverse int = 0

	fmt.Println("Pozitif Sayı Gir : ")
	fmt.Scan(&num)

	temp = num

	for {
		remainder = num % 10
		reverse = reverse*10 + remainder
		num /= 10

		if num == 0 {
			// Break the loop
			break
		}
	}

	if temp == reverse {
		fmt.Printf("%d is a Palindrome\n", temp)
	} else {
		fmt.Printf("%d is not a Palindrome\n", temp)
	}

}
