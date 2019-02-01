package _map

import (
	"fmt"
)

func MakeMap () map[string]int {
	age := make(map[string]int)
	age["alice"] =  1
	age["bob"] = 2
	return age
}

func DeleteKey (age map[string] int, key string) {
	delete(age, key)
}

func IncrementValue (age map[string] int, key string) {
	age[key] +=1 // age[key]++ will work too
}

func Print (age map[string]int) {
	for _, value := range age {
		fmt.Println(value)
	}
}

func Sort(age map[string]int)  {
	//sort.Strings(age)
}

func Contains(age map[string]int, key string) bool {
	if _, ok := age[key]; !ok {
		return false;
	}
	return true
}