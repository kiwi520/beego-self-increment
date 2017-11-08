package tool
/**
 * Created by GoLand.
 * User: hushunhui
 * Date: 17-11-8
 * 自增ID方法
 * Time: 上午10:11
 */

 //定义自增结构体
type AutoInc struct {
	start, step int
	queue chan int
	running bool
}
//生成自增方法
func New(start, step int) (ai *AutoInc) {
	ai = &AutoInc{
		start: start,
		step: step,
		running: true,
		queue: make(chan int, 4),
	}
	go ai.process()
	return
}
//把数据写入channel
func (ai *AutoInc) process() {
	defer func() {recover()}()
	for i := ai.start; ai.running ; i=i+ai.step {
		ai.queue <- i
	}
}
//返回channel中的数据
func (ai *AutoInc) Id() int {
	return <-ai.queue
}
//关闭chan
func (ai *AutoInc) Close() {
	ai.running = false
	close(ai.queue)
}