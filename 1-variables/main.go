package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
)

func printTypes(vars ...any) {
	for _, v := range vars {
		if reflect.TypeOf(v).Kind() == reflect.Int {
			fmt.Printf("значение = %d, тип = %T\n", v, v)
		} else if reflect.TypeOf(v).Kind() == reflect.Float64 {
			fmt.Printf("значение = %.2f, тип = %T\n", v, v)
		} else if reflect.TypeOf(v).Kind() == reflect.String {
			fmt.Printf("значение = %s, тип = %T\n", v, v)
		} else if reflect.TypeOf(v).Kind() == reflect.Bool {
			fmt.Printf("значение = %t, тип = %T\n", v, v)
		} else if reflect.TypeOf(v).Kind() == reflect.Complex64 {
			fmt.Printf("значение = %v, тип = %T\n", v, v)
		}
	}
}

func convertToString(vars ...any) string {
	var result string
	for _, v := range vars {
		result += fmt.Sprintf("%v", v)
	}

	return result
}

func convertToRuneSlice(str string) []rune {
	return []rune(str)
}

func hashRuneSliceWithSalt(runeSlice []rune, salt string) string {
	const defaultSalt = "go-2024"
	if salt == "" {
		salt = defaultSalt
	}

	middleSlice := len(runeSlice) / 2

	result := make([]rune, 0, len(runeSlice)+len(salt))
	result = append(result, runeSlice[:middleSlice]...)
	result = append(result, []rune(salt)...)
	result = append(result, runeSlice[middleSlice:]...)

	hash := sha256.New()
	hash.Write([]byte(string(result)))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func main() {
	var numDecimal int = 11
	var numOctal int = 013
	var numHexadecimal int = 0x11
	var pi float64 = 0.11
	var name string = "Python"
	var isActive bool = false
	var complexNum complex64 = 11 + 5i

	printTypes(
		numDecimal,
		numOctal,
		numHexadecimal,
		pi,
		name,
		isActive,
		complexNum)
	fmt.Println("--------------------------------")

	str := convertToString(
		numDecimal,
		numOctal,
		numHexadecimal,
		pi,
		name,
		isActive,
		complexNum)
	fmt.Println(str)
	fmt.Println("--------------------------------")

	runeSlice := convertToRuneSlice(str)
	fmt.Println(runeSlice)
	fmt.Println("--------------------------------")

	hash := hashRuneSliceWithSalt(runeSlice, "")
	fmt.Println(hash)
}
