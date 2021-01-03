package main

import "fmt"

func main() {
	nbs := []int{0, 1, 2 ,3 ,4 ,5, 6, 7, 8, 9, 10}
	for _, nb := range nbs {
		if nb % 2 == 0 {
			fmt.Println(nb, "is even")
		} else {
			fmt.Println(nb, "is odd")
		}
	}
}