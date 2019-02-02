package main

import (
	"fmt"
	"github.com/adarsh-carousell/learn-go/tempconv"
)



func main() {
	fmt.Println("Hello, playground")

	const pi float64 = 3.14
	var name string = "Adarsh"
	var win bool = true

	fmt.Println(len(name))
	fmt.Println(name+" kumar")
	fmt.Printf("%f \n",  pi)

	fmt.Printf("%T \n", name )
	fmt.Println(win)
	fmt.Printf("%b \n", 2)
	fmt.Printf("%x \n", 15)
	fmt.Printf("%e \n", pi)

	forLoop()
	isWin(true)
	isWin(false)
	switchCase(1)
	switchCase(2)
	arrayManipulation()

	fibonacci(7)

	printArray(getStringArray())

	println(tempconv.FahrenhietToCelsius(tempconv.Fahrenhiet(4)))

	printStruct()

	variatic()

}


func forLoop(){
	for i:=1; i<5 ; i++ {

		fmt.Println("Hello")
	}
}

func isWin(is_win bool) {
	if is_win {
		fmt.Println("Win")
	} else {
		fmt.Println("Lose")
	}
}

func switchCase(cases int) {

	switch cases {
	case 1: fmt.Println("I am 1")
	case 2: fmt.Println("I am 2")
	case 3: fmt.Println("I am 3")
	}

}

func arrayManipulation(){

	// assign and print
	var EvenNum[5] int
	EvenNum[0] = 0
	EvenNum[2] = 2
	fmt.Println(EvenNum)

	// print with for loop
	for _, value := range EvenNum {
		fmt.Println(value)
	}
}

func fibonacci(n int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		fmt.Print(x, y);
	}
}


func getStringArray() []string {
	a := []string{"hello","world", "i", "am", "learning", "go"}
	return a
}

func printArray(arr []string) {
	for _, val := range arr {
		fmt.Println(val)
	}
}

/*
*  STRUCT
*
*/

type person struct { // similar to class
	firstname string
	lastname string
}

type secretAgent struct {
	person
	ltk bool
}

func printStruct() {
	p1 := person{
		firstname:"Adarsh",
		lastname:"Kumar",
	}

	p2 := person{
		firstname: "James",
		lastname:"Bond",
	}

	secretAgent1 := secretAgent{
		person: person{
			firstname:"James",
			lastname:"Bond",

		},
		ltk : false,

	}
	fmt.Println(p1,p2)
	fmt.Println(secretAgent1)
}

/*
* Variatic arguments
*/

func variatic()  {
	x := []int{1,2,3}
	sum(x...)
}

func sum(x ...int) int {

	sum :=0;
	for _,v := range x {
		sum +=v
	}
	return sum;
}




