package array

import(
	"fmt"
	"sort"
)

func getIntArray() []int {
	x := []int{1,2,3,4,}
	return x
}

func getStringArray() []string {
	return []string{"a", "b", "c"}
}

func Print(text string) {
	for _, val := range text {
		fmt.Println(val)
	}

	for k, v :=range getStringArray() {
		fmt.Println(k,v)
	}
}

func StartSlice(arr []int, startIndex int) []int {
	return arr[startIndex:]
}

func EndSlice(arr []int, endIndex int) []int{
	return arr[:endIndex]
}

func RangeSlice(arr []int, startIndex int, endIndex int) []int {
	return arr[startIndex:endIndex]
}

func AppendElements(arr []int, elems ...int) []int {
	return append(arr, elems ...)
}

func AppendArray(arr []int, a []int) []int {
	return append(arr, a...)
}

func Sort(arr []string) {
	sort.Strings(arr)
}

func Length(arr []string) int{
	return len(arr)
}


