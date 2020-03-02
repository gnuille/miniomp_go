package miniomp

import (
	"os"
	"log"
	"runtime"
	"strconv"
	"sync"
)

var miniomp_wg sync.WaitGroup
var miniomp_nt int
var miniomp_pool *ThreadPool
var miniomp_taskqueue chan Task
var err error

func Init(main func() ) {
	var nt_st string
	if nt_st = os.Getenv("OMP_NUM_THREADS"); nt_st == ""{
		miniomp_nt = runtime.NumCPU()
	}else {
		if miniomp_nt, err = strconv.Atoi(nt_st); err != nil {
			log.Fatal("Unable to parse OMP_NUM_THREADS variable. Value is: ", nt_st)
		}
	}

	log.Print("Number of threads is ", miniomp_nt)

	miniomp_pool = CreatePool(miniomp_nt)
	miniomp_taskqueue = CreateTaskqueue()

	for _, t := range miniomp_pool.pool {
		t.SetTasks(miniomp_taskqueue)
		go t.Work()
	}

	main()
	for true {
		select {
		case task := <-miniomp_taskqueue:
			task.f(task.params)
			miniomp_wg.Add(-1)
		default:
			miniomp_wg.Wait()
			return
		}
	}
}


