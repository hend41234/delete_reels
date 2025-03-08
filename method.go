package main

import "fmt"

func MethodeDelete() string {
	methode := map[string]string{
		"1": "default 25",
		"2": "quantity",
		"3": "since",
		"4": "all",
	}
	for key, value := range methode {
		fmt.Println(key, " : ", value)
	}
	var method string
	fmt.Scanln(&method)
	fmt.Println("You choose : ", methode[method])
	return methode[method]
}

func IsNotDefault(method string) string {
	if method == "quantity" {
		var quantity string
		fmt.Scanln(&quantity)
		return quantity
	} else if method == "since" {
		var since string
		fmt.Println("since [ 2024-01-01 ]-[format: YYYY-MM-DD ] : ")
		fmt.Scanln(&since)
		return since
	} else {
		return "all"
	}
}
