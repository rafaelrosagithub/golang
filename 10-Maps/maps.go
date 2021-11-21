package main

import "fmt"

func main() {
	fmt.Println("-------------- Maps --------------")

	user1 := map[string]string{
		"name":    "Peter",
		"surname": "Keppler",
	}

	fmt.Println(user1)
	fmt.Println(user1["name"])

	user2 := map[string]map[string]string{
		"name": {
			"first": "Mike",
			"last":  "Tod",
		},
		"course": {
			"name":   "Engineering",
			"campus": "Campus A",
		},
	}

	fmt.Println(user2)

	delete(user2, "name")
	fmt.Println(user2)

	user2["signo"] = map[string]string{
		"name": "Twins",
	}

	fmt.Println(user2)
}
