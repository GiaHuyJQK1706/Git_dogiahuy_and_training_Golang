package main
/*
 * Trong go, 2 file deu bi loi khi co it nhat 1 ham co ten chung trong cung 1 package
 * Package main de chuong trinh Golang chay tu dong
 * De khac phuc tinh trang loi trung ten ham, neu muon chay tu dong chuong trinh a.go, cac flie *.go (khac file a.go) khac cung package thi khac phuc bang // package main
*/

import (
	// "bufio"
	"fmt"
	// "os"
	"strings"
	"unicode"
)

func analyzeText(text string) (int, map[string]int) {
	// Tach tu
	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	// Dem so tu
	wordCount := len(words)

	// Tinh tan suat xuat hien cua tu
	wordFrequency := make(map[string]int)
	for _, word := range words {
		word = strings.ToLower(word) // Khong phan biet hoa/thuong
		wordFrequency[word]++
	}
	return wordCount, wordFrequency
}

/* read from file
func readFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content.WriteString(scanner.Text() + " ")
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content.String(), nil
}
*/

/* read from stdin
func readFromStdin() string {
	fmt.Println("Nhap van ban (Nhan 2 lan Enter de ket thuc): ")
	var content strings.Builder
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { // Dong trong de ket thuc nhap
			break
		}
		content.WriteString(line + " ")
	}
	return content.String()
}
*/

func main() {
	text := `Lorem Ipsum is simply dummy text of the printing and typesetting industry. 
	Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, 
	when an unknown printer took a galley of type and scrambled it to make a type 
	specimen book. It has survived not only five centuries, but also the leap into 
	electronic typesetting, remaining essentially unchanged. It was popularised in the 
	1960s with the release of Letraset sheets containing Lorem Ipsum passages, and 
	more recently with desktop publishing software like Aldus PageMaker including 
	versions of Lorem Ipsum.`

	/*
	text, err = readFromFile("input.txt")
	if err != nil {
		fmt.Printf("Read file error: %v\n", err)
		return
	}
	*/

	/*
	text = readFromStdin()
	*/

	// Phan tich van ban
	totalWords, wordFrequency := analyzeText(text)

	// Xuat ket qua
	fmt.Printf("Tong so tu: %d\n", totalWords)
	fmt.Println("So lan xuat hien cua moi tu:")
	for word, freq := range wordFrequency {
		fmt.Printf("%s: length = %d, freq = %d\n", word, len(word), freq)
	}
}