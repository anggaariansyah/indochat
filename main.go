package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Print(i, "*", j, "=", i*j)
			fmt.Print("\t")
		}
		fmt.Println()
	}
}