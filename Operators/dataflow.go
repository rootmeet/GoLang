package main
import "fmt"

func main(){
	num1 := 6
	num2 := 2

	sum := num1 + num2
	fmt.Printf("%d + %d = %d/n", num1, num2, sum)

	diff := num1 - num2
	fmt.Printf("%d - %d = %d/n", num1, num2, diff)

	multi := num1 * num2
	fmt.Printf("%d * %d = %d/n", num1, num2, multi)
}