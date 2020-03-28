package main

import "miniomp"
//import "fmt"
import "math/rand"

func calcIt( i int, data []interface{} ){

	a  := data[0].([]interface{})[i].(int)
	b  := data[1].([]interface{})[i].(int)
	res := 0
	for it := 0; it<200; it++ {
		for it2 := 0; it2 < it*2*2; it2++ {
			res += it*it2
			res -= it/(it2+1)
			res = res*3
			res -= it2/(it+1)
			res = res/a
			res = res*b
			for it3 := 0; it3 <  it2+it; it3++ {
				res--
			}
		}
	}
}

func master_func(){
	const SIZE int = 1000

	a := []int{}
	b := []int{}

	for i := 0; i <  SIZE; i++ {
		a = append(a, rand.Int()+1)
		b = append(b, rand.Int()+1)
	}

	miniomp.TaskLoop(calcIt, len(a), []interface{} { miniomp.IntSliceToInterface(a), miniomp.IntSliceToInterface(b)} )

}

func main() {
	miniomp.Init(master_func)
}
