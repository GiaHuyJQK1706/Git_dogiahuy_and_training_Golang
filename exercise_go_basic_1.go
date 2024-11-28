package main
/*
 * Trong go, 2 file deu bi loi khi co it nhat 1 ham co ten chung trong cung 1 package
 * Package main de chuong trinh Golang chay tu dong
 * De khac phuc tinh trang loi trung ten ham, neu muon chay tu dong chuong trinh a.go, cac flie *.go (khac file a.go) khac cung package thi khac phuc bang // package main
*/

import (
	"fmt"
	"math"
)

/*Not write: 
	import {
		...
	}
*/

func calculateCircumference(r int) float64 {
	return 2 * math.Pi * float64(r)
}

func main() {
	var rad int
	fmt.Print("Hay nhap chu vi: ")
	fmt.Scan(&rad) // Not fmt.Scan(rad)
	fmt.Printf("Chu vi hinh tron ban kinh %d la: %.6f", rad, calculateCircumference(rad)) //Printf nhu trong C
}