//1’den 1000’e kadar olan çift ve tek sayılar toplamlarını ve çift ve tek sayıların
//adetlerini bulan ve bunları yazıcıya yazan algoritmayı yapınız.

package main

import "fmt"

func main()  {
		toplamCift:=0
		toplamTek:=0
		i:=0
		adetTek:=0
		adetCift:=0
		for ;i<100;i++{
			if i%2==0{
				toplamCift=toplamCift+i
				adetCift=adetCift+1
			}
			if i%2!=0{
				toplamTek=toplamTek+i
				adetTek=adetTek+1
			}
		}
		fmt.Println(adetTek ,"Adet Tek Sayı Vardır Toplamları da:", toplamTek,"dır")
		fmt.Println(adetCift, "Adet Çift Sayı Vardır Toplamları da:",toplamCift,"dir")

}
