// Klavyeden girilen üç sayıdan en büyüğünü bulan ve ekranda görüntüleyen algoritmayı yapınız.

package main

import "fmt"

func main() {
	var x [3]int
	for i := 0; i <= 2; i++ {
		fmt.Print(i+1, ". sayıyı giriniz:")
		fmt.Scanln(&x[i])
		}
		enBuyuk:=x[0]
	if enBuyuk<x[1] {
		enBuyuk=x[1]

	}
	if enBuyuk<x[2]{
		enBuyuk=x[2]
	}
	fmt.Println("EN BÜYÜK SAYI:",enBuyuk)
}