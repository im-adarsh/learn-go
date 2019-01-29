package array

import(
	"fmt"
	"sort"
)


func Print(text string) {
	for _, val := range text {
		fmt.Println(val)
	}
}

func Sort(arr []string){
	sort.Strings(arr)
}