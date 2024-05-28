package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Println("РЕШИТЕ КВАДРАТНОЕ УРАВНЕНИЕ")
	fmt.Println("Введите значение а:")
	fmt.Scanln(&a)
	fmt.Println("Введите значение b:")
	fmt.Scanln(&b)
	fmt.Println("Введите значение c:")
	fmt.Scanln(&c)

	D := (b * b) - 4*(a*c)

	if D > 0 {
		var x1 float64
		var x2 float64
		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)

		fmt.Println("Ваше уравнение имеет 2 корня\nD = " + fmt.Sprint(D))
		fmt.Println("X1: " + fmt.Sprint(x1) + "\nX2: " + fmt.Sprint(x2))
	} else if D == 0 {
		var x float64
		x = (-b) / (2 * a)
		fmt.Println("Ваше уравнение имеет 1 корень\nD = ")
		fmt.Println("X: " + fmt.Sprint(x))
	} else if D < 0 {
		fmt.Println("Ваше уравнение не имеет корней\nD < 0\nD = " + fmt.Sprint(D))
	}

}
