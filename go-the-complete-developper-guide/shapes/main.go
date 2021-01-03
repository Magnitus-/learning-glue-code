package main

func main() {
	s := square{sideLength: 10.0}
	printArea(s)

	t := triangle{height: 10.0, base: 5.0}
	printArea(t)
}
