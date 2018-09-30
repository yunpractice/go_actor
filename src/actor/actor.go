package actor

type Actor struct {
	id        int
	name      string
	ref_count int
	queue     *Queue
}

func NewActor(id int)*Actor {
	actor := &Actor{
		id: id,
		name: "",
		ref_count: 1,
		queue: new(Queue)
	}
}

func (actor *Actor) Init() {

}

// 注册回调函数
func (actor *Actor)Register(t string,func()interface{}){
	
}

// 使用回调函数处理消息
func (actor *Actor) Dispatch(){
	
}
