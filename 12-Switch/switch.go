package main

import (
	"fmt"
)

func dayOfTheWeek1(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid number"
	}

}

func dayOfTheWeek2(number int) string {
	var dayWeek string

	switch {
	case number == 1:
		dayWeek = "Sunday"
		fallthrough // Go to the next condition
	case number == 2:
		dayWeek = "Monday"
	case number == 3:
		dayWeek = "Tuesday"
	case number == 4:
		dayWeek = "Wednesday"
	case number == 5:
		dayWeek = "Thursday"
	case number == 6:
		dayWeek = "Friday"
	case number == 7:
		dayWeek = "Saturday"
	default:
		dayWeek = "Invalid number"
	}
	return dayWeek
}

func main() {
	fmt.Println("---------------  Switch ---------------")
	day := dayOfTheWeek1(10)
	fmt.Println(day)
	day = dayOfTheWeek2(7)
	fmt.Println(day)
}
