package main

import "miniomp"
import "fmt"

func calc(args []interface{}) interface{} {
		var a int = args[0].(int)
		var b int = args[1].(int)
		fmt.Print("Value of task ", a+b, "\n")
		return a+b
}

func master_func(){
	a := []int{ 1,2,3,4,5,6,78,9,9,9,99,9,9,9,9,9,9,9}
	b := []int{ 1,2,3,4,5,6,78,9,9,9,99,9,9,9,9,9,9,9}

	for i,_ := range a {
		miniomp.NewTask(calc, []interface{}{ a[i], b[i] })
	}
	fmt.Println("Number of tasks ", len(a))
}

func main() {
	miniomp.Init(master_func)
}
