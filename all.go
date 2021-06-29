package main

import "fmt"

func main() {

	//hello world
	fmt.Println("hello world!")

	// var dan const
	var (
		firstName = "dimas"
		lastName = "aditya"
	)
	fmt.Println(firstName, lastName)

	const (
		id = 12029492894
		nama = "aditya"
	)
	fmt.Println(id)
	fmt.Println(nama)

	// string boolean
	var world string = "this is string" //string menggunakan dua tanda kutip
	var pelajar = true //output boolean adalah true dan false
	fmt.Println(world)
	fmt.Println(pelajar)

	//type dec dan operasi matematika

	type str string
	type angka int

	var name str = "wardi"
	var ktp angka = 1213131
	fmt.Println(name)
	fmt.Println(ktp)

	var a = 2
	var b = 7
	var c = a * b
	fmt.Println(c)

	var i = 10
	i += 1
	fmt.Println(i)

	var j = 10
	j++
	fmt.Println(j)

	var result = 100 * 12
	fmt.Println(result)

	//perbandingan dan conversion

	var x = 10
	var y = 10
	var z = x == y
	fmt.Println(z)

	var nilai32 int32 = 100000
	var nilai8 int8 = int8(nilai32)
	fmt.Println(nilai8)

	var android string = "os11"
	var ios string = string(android)
	fmt.Println(ios)

	//type data number

	var nilai1 int8 = 126
	var nilai2 int32 = 2147483647
	var nilai3 int64 = 9223372036854775807
	var nilai4 int = 121212 //sama dengan nilai int32 dan int64

	fmt.Println(nilai1)
	fmt.Println(nilai2)
	fmt.Println(nilai3)
	fmt.Println(nilai4)
}
