//n tane girilen sayı içerisinden pozitif, negatif sayıların ve 0 sayısının adedini ekrana yazan algoritmayı yapınız.

package main

import "fmt"

func main(){
	var n int
	fmt.Print("Kaç Adet Sayı Gireceksiniz:")
	fmt.Scanln(&n)
	pozitif:=0
	negatif:=0
	sıfır:=false
	for i:=1 ;i<=n; i++{
		var x int
		fmt.Print("Bir Sayı Girinz:")
		fmt.Scanln(&x)

		if x<0 {
			negatif++

			fmt.Println(x ,"negatiftir")
		}
		if x>0{
			pozitif++

			fmt.Println(x, "pozitiftir")
		}
		if x==0{
			sıfır=true

		}

		}
		if sıfır ==true {
			fmt.Print("Sıfır değeri mevcuttur ,")
		}
		fmt.Println(negatif ,"negatif sayı ve" , pozitif ,"pozitif sayı vardır")
	}



