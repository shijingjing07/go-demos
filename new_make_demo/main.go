package main


import "fmt"


func main() {
	a := make(map[string]string)
	a["name"] = "shijingjing"
	a["sex"] = "male"
	for key := range a {
		fmt.Println(key, a[key])
	}

	b := make([]string, 5)
	b[0] = "item1"
	b[1] = "item2"
	fmt.Println(b)

	p := new(int)
	*p = 1
	fmt.Println(p)
	fmt.Println(*p)
}