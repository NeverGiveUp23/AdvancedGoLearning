package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	firstName   string
	lastName    string
	age         int
	phoneNumber PhoneHomeCell
}

type PhoneHomeCell struct {
	home string
	cell string
}

// composite type
type example struct {
	flag bool
	// 7 bytes of padding here
	counter int64
	pi      float32
}

// when defining a struct, the compiler automatically adds padding bytes to ensure proper memory alignment.
// to minimize padding and optimize memory usage, arrange struct fields in decreasing order of size. -> really useful only if you need it

type arrangeExample struct {
	B int64 // 8bytes
	D int64 // 8bytes
	A bool  // 1 byte
	C bool  // 1 byte
	// only 6 bytes of padding here
}

func main() {
	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
	}

	// anonymous structs
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "sudo@mail.com",
	}

	fmt.Println(user.email)

	//methods to structs
	fmt.Println(p.fullname())
	fmt.Println(p.getAge())

	fmt.Println(unsafe.Sizeof(arrangeExample{}))
	fmt.Println(unsafe.Sizeof(example{}))

	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	fmt.Println("Flag: ", e2.flag)

	// anonymous structs -> literal struct type
	e3 := struct {
		flag    bool
		counter int64
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	fmt.Println(e3 == e2) // true

	p2 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       49,
		phoneNumber: PhoneHomeCell{
			home: "123456789",
			cell: "123456789",
		},
	}

	fmt.Println(p2.phoneNumber.cell)

	p2.changeCellNumber("535452523")
	fmt.Println(p2.phoneNumber.cell)

}

func (p Person) fullname() string {
	return p.firstName + p.lastName
}

func (p Person) getAge() int {
	return p.age
}

// methods
func (p *Person) changeCellNumber(number string) {
	if p.phoneNumber.cell == "" {
		fmt.Println("Phone number is empty")
		return
	}
	p.phoneNumber.cell = number
}
