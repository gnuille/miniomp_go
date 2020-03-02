package miniomp

type Fn func([]interface{}) interface{}

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
