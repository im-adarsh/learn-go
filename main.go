package main

import (
	"encoding/json"
	"fmt"
	"github.com/adarsh-carousell/learn-go/tempconv"
	"sort"
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

	callSpeak()

	creatingSAHumanType()

	anonymousFunction()

	funcExpression()

	callingReturnFunction()

	pointers()

	jsonMarshal()

	customSort()

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
	Firstname string
	Lastname string
}

type secretAgent struct {
	person
	ltk bool
}

func printStruct() {
	p1 := person{
		Firstname:"Adarsh",
		Lastname:"Kumar",
	}

	p2 := person{
		Firstname: "James",
		Lastname:"Bond",
	}

	secretAgent1 := secretAgent{
		person: person{
			Firstname:"James",
			Lastname:"Bond",

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

/*
* functions
*/

func (sa secretAgent) speak() { // notice this chaining attached to secretAgent
	fmt.Print("I am secret agent. My name is : ", sa.Firstname, sa.Lastname)
}

func callSpeak()  {
	sa := secretAgent{
		person : person{
			Firstname : "James ",
			Lastname : "Bond",
		},
		ltk:true,

	}

	sa.speak(); //
}


/*
* Interfaces and polymorphism
*/

//  keyword identifier type (KIT)
type human interface {
	speak()
}

func creatingSAHumanType() {
	sa1 := secretAgent{
		person{"world", "hello"},
		false,
	}

	sa1.speak()
}

/*
* Anonymus function
*/

func anonymousFunction(){
	foo()
	func() {
		fmt.Println("hello anonymous")
	}()

	func(x int){
		fmt.Println("anonymous function with param ", x)
	}(42)

}

func foo() {
	fmt.Println("foo ran")
}

/*
* Function expression
*/

func funcExpression(){
	fmt.Println( "Hello func expression")

	f :=func(){
		fmt.Println("func expression ")
	}
	f()

	g:= func(x int) {
		fmt.Println("func expression with param ", x)
	}
	g(43)
}

/*
* return function from a function
*/


func callingReturnFunction() {
	f := returnFunction(); // calling a function that returns function
	fmt.Println(f()) // calling the returned function and printing
}

func returnFunction() func() int {
	return func() int {
		return 4
	}
}

/*
* callback [WIP]
*/


func sumTest(x ...int) int {
	sum :=0
	for _,v := range x {
		sum += v
	}
	return sum
}

/*
* Pointers
*/

func pointers(){
	x:= 42
	fmt.Println(x) // 42
	fmt.Println(&x) // 0xc00001a118
	fmt.Println(*(&x)) // 42

	y := &x
	fmt.Println(y) // 0xc00001a118

	var z *int = &x
	fmt.Println(z) // 0xc00001a118

	fmt.Println(*&x) // 42

	changeValueAtAddress(&x)

	fmt.Println(x) // 43
}

func changeValueAtAddress(y *int){
	fmt.Println("Before changing the value ", *y)
	*y = 43
	fmt.Println("after changing the value ",*y)
}

/*
* JSON marshal
*/

func jsonMarshal() {
	p1:= person{
	"adarsh",
	"kumar",
	}

	p2:= person{
		"James",
		"Bond",
	}

	ps :=  []person{p1, p2}
	m, err:=json.Marshal(ps)
	if err != nil {

	}
	fmt.Println(string(m))
	jsonUnmarshal(string(m))

}

/*
* JSON Unmarshal
*/

type person1 struct {
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
}

func jsonUnmarshal(s string) {

	bs := []byte(s)
	var people1 []person1
	err := json.Unmarshal(bs, &people1)
	if err != nil {

	}
	fmt.Println(people1)

}

/*
* custom sort
*/

type Child struct {
	Name string
	Age int
}

// sort by age
type ByAge []Child

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i,j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i,j int) bool  {
	return a[i].Age < a[j].Age
}

func customSort() {
	c1 := Child{"A", 2}
	c2 := Child{"A", 1}
	c3 := Child{"A", 7}
	c4 := Child{"A", 4}

	children := []Child{c1, c2, c3, c4}

	fmt.Println(children)
	sort.Sort(ByAge(children))
	fmt.Println(children)
}




