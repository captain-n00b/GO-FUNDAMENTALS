package main

import (
	"fmt"
	"strings"
)

func main() {
	// var name string = "Arkabh"
	// var score int = 110
	name, score := "Arkabh", 110

	var details string = `
	Maths: 50
	Physics: 60
	`
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println("Variable and Data types")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println("Student scores")
	fmt.Println(name, details) 
	fmt.Println("Total: ", score)

	fmt.Println(strings.Repeat("-", 14))
	fmt.Println("Arrays and Slices ")
	fmt.Println(strings.Repeat("-", 14))
	students := []string{"Roh",
	 "Pra"}

	scores := []int{50, 60}
	for i := 0; i < len(students); i++ {
		fmt.Println(students[i], " : ", scores[i])
	}

	fmt.Println(strings.Repeat("-", 14))
	fmt.Println("Map ")
	fmt.Println(strings.Repeat("-", 14))
	scoreMap := map[string]int{
		students[0] : scores[0],
		students[1] : scores[1],
	}

	fmt.Println(scoreMap["Roh"])
	fmt.Println(scoreMap["Pra"])
	value, ok := scoreMap["roh"]
	fmt.Println("Does ",value, " exits? ", ok)

	fmt.Println(strings.Repeat("-", 14))
	fmt.Println("Struct")
	fmt.Println(strings.Repeat("-", 14))
	type studentScore struct {
		name string
		score int
	}

	studentDetails := []studentScore{
		{name: "Ork", score: 50},
		{name: "Rkk", score: 60},
	}
	for i := range studentDetails {
		fmt.Println(studentDetails[i])
	}
	for i := range studentDetails {
		fmt.Println(studentDetails[i].name, studentDetails[i].score)
	}

	fmt.Println(strings.Repeat("-", 14))
	fmt.Println("CONTROL FLOWS")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println("IF STATEMENT")
	fmt.Println(strings.Repeat("-", 14))
	var option string
	fmt.Scanln(&option)

	var index int
	if option == "1" {
		index = 0
	} else  if option == "2" {
		index = 1
	} else {
		fmt.Println("unknown option, defaulting to 1")
		index = 1
	}
	fmt.Println(studentDetails[index].name, ":", studentDetails[index].score)


	fmt.Println(strings.Repeat("-", 14))
	fmt.Println("SWITCH STATMENTS")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Scanln(&option)
	switch option {
	case "1":
		index = 0
	case "2":
		index = 1
	default:
		fmt.Println("unknown option, defaulting to 1")
		index = 1
	}
	fmt.Println(studentDetails[index].name, ":", studentDetails[index].score)

	fmt.Println(strings.Repeat("-", 14))
	fmt.Println("FUNCTIONS")
	fmt.Println(strings.Repeat("-", 14))

	fmt.Println("Addition: ", add(10,20))
	val, ok := divide(10, 5)
	fmt.Println("Division: ", val, ok)
	val, ok = remainder(20, 5)
	fmt.Println("Remainder: ", val, ok)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func ratio(num1, num2 int) {
	fmt.Println(num1, ":", num2)
}

func divide(num1, num2 int) (int, bool) {
	if num2 == 0 {
		return 0, false
	}
	return num1 / num2, true
}

func remainder(num1, num2 int) (remainder int, ok bool) {
	if num1 > num2 {
		remainder = num1 % num2
	} else {
		remainder = num2 % num1
	}
	ok = true
	return
}