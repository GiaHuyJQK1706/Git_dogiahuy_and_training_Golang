package main 
/*
 * Trong go, 2 file deu bi loi khi co it nhat 1 ham co ten chung trong cung 1 package
 * Package main de chuong trinh Golang chay tu dong
 * De khac phuc tinh trang loi trung ten ham, neu muon chay tu dong chuong trinh a.go, cac flie *.go (khac file a.go) khac cung package thi khac phuc bang // package main
*/

import (
	"fmt"
)

func printFibonacci(n int) { //O(n)
	f0 := 0
	f1 := 1
	if n == 1 {
		fmt.Printf("%d,", f0)
	} else if n == 2 {
		fmt.Printf("%d, %d", f0, f1)
	} else {
		fmt.Printf("%d, %d, ", f0, f1)
		for i := 3; i <= n; i++ {
			tmp := f0 + f1
			if i != n {
				fmt.Printf("%d, ",tmp)
			} else {
				fmt.Printf("%d",tmp)
			}
			f0 = f1
			f1 = tmp
		} 
	}
}
func main() {
	var input int
	fmt.Print("Hay nhap input: ")
	fmt.Scan(&input)
	printFibonacci(input)
}