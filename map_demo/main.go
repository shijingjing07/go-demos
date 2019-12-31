package main


import "fmt"


func main() {
	a := make(map[string]string)
	a["name"] = "shijingjing"
	a["sex"] = "male"
	for key := range a {
		fmt.Println(key, a[key])
	}
	age, ok := a["age"]
	if ok {
		fmt.Println("existed")
		fmt.Println(age)
	} else {
		fmt.Println("not existed")
	}
}