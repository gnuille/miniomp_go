package main

import "miniomp"
import "fmt"
import "math/rand"


var count int64

func calcIt( i int, data []interface{} ){

	a  := data[0].([]interface{})[i].([]interface{})
	b  := data[1].([]interface{})[i].([]interface{})

	var res int = 0
	for it,_ := range a {
		res += a[it].(int) * b[it].(int)
	}

	fmt.Println(res)
}

func randomArray(len int) []int {
    a := make([]int, len)
    for i := 0; i <= len-1; i++ {
        a[i] = rand.Intn(len)
    }
    return a
}

func randomArrayOfArray(len int) [][]int {
	a := make([][]int, len)
	for i := 0; i < len; i++ {
		a[i] = randomArray(len)
	}
	return a
}

func master_func(){


	A := randomArrayOfArray(10000)
	B := randomArrayOfArray(10000)

	miniomp.TaskLoop(calcIt, len(A), []interface{} { miniomp.IntSliceIntSliceToInterface(A), miniomp.IntSliceIntSliceToInterface(B) } )

}

func main() {
	miniomp.Init(master_func)
}
