//Girilen herhangi bir aralıktaki sayma sayılarının çift ve tek sayılar toplamlarını ve
//çift ve tek sayıların adetlerin bulan ve bunları yazıcıya yazan algoritmayı yapınız.

package main

import "fmt"

func main() {
	var bas, son, i int
	tCift := 0;
	tTek := 0
	totalc := 0;
	totalt := 0

	fmt.Scan(&bas, &son)
	i = bas

	for ; i <= son; i++ {
		if i%2 == 0 {
			tCift++
			totalc = totalc + i
		}
		if i%2 != 0 {
			tTek++
			totalt = totalt + i
		}
	}
	fmt.Printf("%v adet tek sayının toplamı:%v sayısıdır,%v adet çift sayının toplamı:%v sayısıdır", tTek, totalt, tCift, totalc)
}