package actor

// 调度器函数
func worker() {
	ch := make(chan *Actor)
	select {
	case actor <- ch:
		actor.Dispatch()
	default
	}
}
