
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

	a := []int{ 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18 }
	b := []int{ 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18 }

	for i,_ := range a {
		miniomp.NewTask(calc, []interface{}{ a[i], b[i] })
		miniomp.TaskWait()
	}
	//the output should be serialized vs task.go example
}

func main() {
	miniomp.Init(master_func)
}
