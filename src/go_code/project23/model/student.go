package model

// 定义一个结构体


type Student struct{
	Name string
	Score int
}

/**定义为私有的，需要通过公有的getset方法获取*/
type persion struct{
	name string
	age int
}


/**这里使用这个函数，返回我们的对象
**/
func NewPersion(n string,a int) *persion{
	return &persion{
		name : n,
		age : a,
	}
}

/*提供get，set方法！！！！1*/
func (p *persion)SetName(n string) {
    (*p).name = n
}


func (p *persion)SetAge(a int){
	(*p).age = a
}
func (p *persion)GetName() string{
    return (*p).name
}


func (p *persion)GetAge()int {
	return (*p).age
}