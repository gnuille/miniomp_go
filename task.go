package miniomp

type Fn func([]interface{}) interface{}
type LoopFn func(int, []interface{})

type Task struct {
	f Fn
	params []interface{}
}

func (t *Task) Execute() interface{} {
	return t.f(t.params)
}

func CreateTaskqueue() chan Task {
	return make(chan Task)
}

func NewTask(f Fn, params []interface{}){
	miniomp_wg.Add(1)
	miniomp_taskqueue  <-  Task {
		f: f,
		params: params,
	}
}

func task4TaskLoop(args []interface{}) interface{}{
	fun := args[0].(LoopFn)
	init := args[1].(int)
	fi := args[2].(int)
	loopFnArgs := args[3].([]interface{})

	for i:=init; i < fi; i++ {
		fun(i, loopFnArgs)
	}

	return nil
}

func TaskLoop(f LoopFn, iterations int, loopFnArgs []interface{}){

	var BS int = (iterations/miniomp_nt)/2
	for i:=0; i < iterations; i += BS {
		fin :=  i+BS
		if  fin >= iterations {
			fin = iterations
		}
		NewTask(task4TaskLoop, []interface{}{ f, i, fin, loopFnArgs})
	}

	TaskWait()
}
