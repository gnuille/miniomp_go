
package main

import "miniomp"
import "fmt"

func calcIt( i int, data []interface{} ){
		a  := data[0].([]interface{})
		b  := data[1].([]interface{})
		fmt.Print("Value is ", a[i].(int)+b[i].(int), "\n")
}

func master_func(){

	A := []int{ 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18 }
	B := []int{ 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18 }

	miniomp.TaskLoop(calcIt, len(A), []interface{} { miniomp.IntSliceToInterface(A), miniomp.IntSliceToInterface(B) } )

}

func main() {
	miniomp.Init(master_func)
}
