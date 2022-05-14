package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	var sum [32]byte
	var mass string 
	fmt.Scan(&mass)
	sum = sha256.Sum256([]byte(mass))

	var b bool 
	b = false
	i := 0
	for b==false {
		if (sum[0]==0)&&(sum[1]==0) {
	//		fmt.Printf("done\n")
			b=true
		}else{
	//		fmt.Printf("no\n")
			sum = sha256.Sum256([]byte(mass + string(i)))
			i++
		}
	//	fmt.Printf("%x  %x\n\n",sum[0],sum[1])	
	}
	fmt.Printf("%x\n", sum)
}