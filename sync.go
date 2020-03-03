package miniomp

func TaskWait(){
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
