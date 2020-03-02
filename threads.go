package miniomp

type Thread struct {
	tid int
	tasks chan Task
}

type ThreadPool struct {
	numThreads int
	pool []Thread
}

func (t *Thread) GetTid() int {
	return t.tid
}

func (t *Thread) SetTid(tid int) {
	t.tid = tid
}

func (t *Thread) SetTasks(tasks chan Task){
	t.tasks = tasks
}

func CreatePool(nt int) *ThreadPool {
	ret := &ThreadPool{
		numThreads: nt,
		pool: make([]Thread, nt),
	}

	for i, t := range ret.pool {
		(&t).SetTid(i)
	}

	return ret
}

func (t *Thread) Work() {
	for true {
		task := <-t.tasks
		task.f(task.params)
		miniomp_wg.Add(-1)
	}
}
