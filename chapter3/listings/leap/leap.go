package main

import "fmt"

func main() {
	var year = 2016
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)
	fmt.Println("Leap year?", leap)
}
