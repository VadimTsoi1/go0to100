package main

import "fmt"

//func main() {
//	num := 12
//	num = num + 8 // 8 - неименованная константа
//	fmt.Println(num)
//}

// строковая константа
// func main() {
// 	hello := "Привет,"
// 	greeting := hello + "программист!"

// 	fmt.Println(greeting)
// }

// Именованные константы
// Чтобы определить имя для константы, используют ключевое слово const.
// Вот как можно переписать предыдущие примеры с использованием именованных констант:

// func main() {
// 	const eight = 8 // constant

// 	num := 12 // variable

// 	num = num + eight

// 	fmt.Println(num)
// }

// В примере была определена константа eight, которая впоследствии была использована в операции сложения.
// То же самое можно сделать со строковой константой:
// func main() {
// 	const developer = "Developer!"

// 	hello := "Privet, "

// 	greeting := hello + developer

// 	fmt.Println(greeting)

// 	const num = 8
// 	num = 10         // нельзя менять const - будет ошибка
// 	fmt.Println(num) // cannot assign to num (neither addressable nor a map index expression)
// }
// func main() {
// 	const ocean = 12

// 	intNumber := 8 //type int

// 	floatNumber := 11.91 //type float64

// 	result1 := intNumber + ocean //const ocean now int

// 	result2 := floatNumber + ocean // const ocean now float64

// 	fmt.Println(result1, result2)
// }

//ЗАДАЕМ ТИП КОНСТАНТЕ
// func main() {
// 	const twelve int = 12 //assigned type int when declaring const

// 	intNumber := 8 //type int

// 	floatNumber := 11.91 //type float64

// 	result1 := intNumber + twelve // will work

// 	result2 := floatNumber + twelve // will show mistake cause two different types

// 	fmt.Println(result1)

// 	fmt.Println(result2)
// }

const oceansTwelve = 12 // declared outside of the function = can be seen globally

func main() {
	intNumber := 9

	floatNumber := 11.91

	result1 := intNumber + oceansTwelve

	result2 := floatNumber + oceansTwelve

	fmt.Println(result1)
	fmt.Println(result2)
} //still works even though oceansTwelve is outside of the function
