package libStruct

import "fmt"

type person struct {
	firstname string
	lastname string
}

func main() {
	p1 := person{
		firstname:"Adarsh",
		lastname:"Kumar",
	}

	p2 := person{
		firstname: "James",
		lastname:"Bond",
	}

	fmt.Println(p1,p2)
}