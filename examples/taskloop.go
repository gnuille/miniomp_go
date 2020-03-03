package main

import "miniomp"
import "fmt"
import "math/rand"


var count int64

func calcIt( i int, data []interface{} ){
	a  := data[0].([]interface{})
	b  := data[1].([]interface{})
	fmt.Println(a[i].(int)+b[i].(int))
}

func randomArray(len int) []int {
    a := make([]int, len)
    for i := 0; i <= len-1; i++ {
        a[i] = rand.Intn(len)
    }
    return a
}

func master_func(){


	A := randomArray(100000)
	B := randomArray(100000)

	miniomp.TaskLoop(calcIt, len(A), []interface{} { miniomp.IntSliceToInterface(A), miniomp.IntSliceToInterface(B) } )

}

func main() {
	miniomp.Init(master_func)
}
