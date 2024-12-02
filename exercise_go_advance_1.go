package main
/*
 * Trong go, 2 file deu bi loi khi co it nhat 1 ham co ten chung trong cung 1 package
 * Package main de chuong trinh Golang chay tu dong
 * De khac phuc tinh trang loi trung ten ham, neu muon chay tu dong chuong trinh a.go, cac flie *.go (khac file a.go) khac cung package thi khac phuc bang // package main
*/

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

// Struct Person
type Person struct {
	name         string
	birthdayYear int
	age          int
	email        string
	phone        string
}

// Interface PersonInterface
type PersonInterface interface {
	SetName(name string) error
	SetBirthdayYear(year int) error
	SetEmail(email string) error
	SetPhone(phone string) error
}

func (p *Person) SetName(name string) error {
	matched, err := regexp.MatchString(`^[A-Z][a-zA-Z]*$`, name)
	if err != nil {
    		return fmt.Errorf("Error while validating name: %v", err)
	}
	if !matched {
		return errors.New("Ten khong hop le: Phai bat dau bang chu cai in hoa (A-Z)")
	}
	
	p.name = name
	
	return nil
}

func (p *Person) SetBirthdayYear(year int) error {
	currentYear := time.Now().Year()
	if year < 1900 {
		return errors.New("Nam sinh khong hop le: Phai tu 1900 tro lai")
	} else if year > currentYear {
		return errors.New("Nam sinh khong hop le: Khong duoc vuot qua nam nay")
	}
	
	p.birthdayYear = year
	p.age = currentYear - year
	
	return nil
}

func (p *Person) SetEmail(email string) error {
	matched, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
	if err != nil {
    		return fmt.Errorf("Error while validating email: %v", err)
	}
	if !matched {
    		return errors.New("Email khong hop le")
	}

	p.email = email
	
	return nil
}

func (p *Person) SetPhone(phone string) error {
	if len(phone) < 10 {
		return errors.New("SDT khong hop le: Qua ngan")
	}
	if phone[0] == '0' {
		if len(phone) != 10 && len(phone) != 11 {
			return errors.New("SDT khong hop le: Phai co 10-11 ky tu neu bat dau tu 0")
		}
	} else if phone[0] == '+' {
		if len(phone[1:]) < 10 || len(phone[1:]) > 14 {
			return errors.New("SDT khong hop le: Phai co 10-14 ky tu neu bat dau tu + va khong tinh ky tu +")
		}
	} else {
		return errors.New("SDT khong hop le: Phai bat dau bang 0 hoac +")
	}
	
	p.phone = phone
	
	return nil
}

func main() {
	var person Person

	// Test SetName
	err := person.SetName("Alice")
	if err != nil {
		fmt.Println("Loi SetName:", err)
	} else {
		fmt.Println("Ten:", person.name)
	}

	// Test SetBirthdayYear
	err = person.SetBirthdayYear(1990)
	if err != nil {
		fmt.Println("Loi SetBirthdayYear:", err)
	} else {
		fmt.Printf("Nam sinh: %d, tuoi: %d\n", person.birthdayYear, person.age)
	}

	// Test SetEmail
	err = person.SetEmail("alice@example.com")
	if err != nil {
		fmt.Println("Loi SetEmail:", err)
	} else {
		fmt.Println("Email:", person.email)
	}

	// Test SetPhone
	err = person.SetPhone("+84123456789")
	if err != nil {
		fmt.Println("Loi SetPhone:", err)
	} else {
		fmt.Println("SDT:", person.phone)
	}
}

