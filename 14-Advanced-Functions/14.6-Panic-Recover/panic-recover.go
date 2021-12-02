package main

import "fmt"

func recoverRun() {
	fmt.Println("Attempt to recover execution")
	if r := recover(); r != nil {
		fmt.Println("Run retrieved successfully")
	}
}

func studentIsApproved(n1, n2 float64) bool {
	defer recoverRun()
	average := (n1 + n2) / 2

	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}

	panic("The average is exactly 6")
}

func main() {
	fmt.Println("----------- Panic Recover -----------")
	fmt.Println(studentIsApproved(6, 6))
	fmt.Println("Pós execution")
}
