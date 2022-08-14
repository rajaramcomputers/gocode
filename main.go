package main

import "fmt"

// Constant
const (
	first  = iota
	second = iota
)

// go mod init github.com/pluralsight/webservice
func main() {
	fmt.Println("I am learning Go")
	var i int
	i = 43
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstname := "Krishnakumar"
	fmt.Println(firstname)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	var firstname1 *string = new(string)
	*firstname1 = "Neela"
	fmt.Println(*firstname1)

	ptr := &firstname
	fmt.Println(ptr, *ptr)

	const pi = 3.1415
	fmt.Println(pi)

	fmt.Println(first, second)
	//Array
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr1 := [3]int{1, 2, 5}
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr[1])

	// Slices

	slice := arr[:]
	arr[1] = 15
	slice[2] = 22
	fmt.Println("Array Value", arr, "Slice Value", slice)

	slice1 := []int{11, 12, 13}

	slice1 = append(slice1, 42, 43, 44)

	s1 := slice1[0:]
	s2 := slice1[1:]
	s3 := slice1[:2]
	s4 := slice1[1:2]

	fmt.Println("Direct Assignment", slice1)

	fmt.Println("Sliced Values:", s1, s2, s3, s4)
	//maps

	stu1 := map[string]int{"ram": 10, "krishna": 20, "abi": 30}
	fmt.Println(stu1)
	fmt.Println(stu1["ram"])
	fmt.Println(stu1["krishna"])
	fmt.Println(stu1["abi"])
	stu1["abi"] = 34
	fmt.Println(stu1["abi"])

	delete(stu1, "ram")
	fmt.Println(stu1)

	//struct

	type user struct {
		ID        int
		firstname string
		lastname  string
	}
	var u, u1 user
	u.ID = 1
	u.firstname = "Meera"
	u.lastname = "Ramkumar"
	u1.ID = 2
	u1.firstname = "Neela"
	u1.lastname = "Ramkumar"

	u3 := user{ID: 3, firstname: "Parvatha", lastname: "Vardhini"}
	fmt.Println(u, u1, u3)

}
