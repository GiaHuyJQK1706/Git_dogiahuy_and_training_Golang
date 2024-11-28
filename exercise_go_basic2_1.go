package main
/*
 * Trong go, 2 file deu bi loi khi co it nhat 1 ham co ten chung trong cung 1 package
 * Package main de chuong trinh Golang chay tu dong
 * De khac phuc tinh trang loi trung ten ham, neu muon chay tu dong chuong trinh a.go, cac flie *.go (khac file a.go) khac cung package thi khac phuc bang // package main
*/

import (
	"fmt"
)

// Tinh dinh thuc cua ma tran 3x3
func det(matrix [3][3]float64) float64 {
	return matrix[0][0]*(matrix[1][1]*matrix[2][2]-matrix[1][2]*matrix[2][1]) -
		matrix[0][1]*(matrix[1][0]*matrix[2][2]-matrix[1][2]*matrix[2][0]) +
		matrix[0][2]*(matrix[1][0]*matrix[2][1]-matrix[1][1]*matrix[2][0])
}

func main() {
	var A [3][3]float64

	fmt.Println("Nhap cac phan tu cua ma tran 3x3 (A[dong][cot]): ")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("A[%d][%d]: ", i+1, j+1)
			fmt.Scan(&A[i][j])
		}
	}

	fmt.Println("\nMatrix A:")
	for _, row := range A {
		fmt.Println(row)
	}

	fmt.Printf("\nĐịnh thức của ma trận: %.2f\n", det(A))
}
