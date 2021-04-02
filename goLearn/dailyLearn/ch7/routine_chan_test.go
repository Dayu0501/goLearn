package ch7

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	pipline := make(chan int, 10)
	fmt.Printf("信道可缓冲 %d 个数据\n", cap(pipline))
	pipline <- 1
	fmt.Printf("信道中当前有 %d 个数据", len(pipline))
}

func TestChan1(t *testing.T) {
	pipline := make(chan int)

	go func() {
		fmt.Println("准备发送数据: 100")
		pipline <- 100
		fmt.Println("end !!")
	}()

	go func() {
		num := <-pipline
		fmt.Printf("接收到的数据是: %d", num)
	}()
	// 主函数sleep，使得上面两个goroutine有机会执行
	time.Sleep(time.Second)
}

func TestChan2(t *testing.T) {
	var pipline = make(chan int)
	/* 定义只读信道 */
	type Receiver = <-chan int // 关键代码：定义别名类型
	var receiver Receiver = pipline
	fmt.Println(len(receiver))
}

func TestChan3(t *testing.T) {
	var pipline = make(chan int)
	type Sender = chan<- int // 关键代码：定义别名类型
	var sender Sender = pipline
	fmt.Println(len(sender))
}

//定义只写信道类型
type Sender = chan<- int

//定义只读信道类型
type Receiver = <-chan int

func TestChan4(t *testing.T) {
	var pipline = make(chan int)

	go func() {
		var sender Sender = pipline
		fmt.Println("准备发送数据: 100")
		sender <- 100
	}()

	go func() {
		var receiver Receiver = pipline
		num := <-receiver
		fmt.Printf("接收到的数据是: %d", num)
	}()
	// 主函数sleep，使得上面两个goroutine有机会执行
	time.Sleep(time.Second)
}

func fibonacci(mychan chan int) {
	n := cap(mychan)
	x, y := 1, 1
	for i := 0; i < n; i++ {
		mychan <- x
		x, y = y, x+y
	}
	// 记得 close 信道
	// 不然主函数中遍历完并不会结束，而是会阻塞。
	close(mychan)
}

func TestChan5(t *testing.T) {
	pipline := make(chan int, 10)

	go fibonacci(pipline)

	for k := range pipline {
		fmt.Println(k)
	}
}

// 由于 x=x+1 不是原子操作
// 所以应避免多个协程对x进行操作
// 使用容量为1的信道可以达到锁的效果
func increment(ch chan bool, x *int) {
	ch <- true
	*x = *x + 1
	<-ch
}

func TestChan6(t *testing.T) {
	// 注意要设置容量为 1 的缓冲信道
	pipline := make(chan bool, 1)

	var x int
	for i := 0; i < 1000; i++ {
		go increment(pipline, &x)
	}

	// 确保所有的协程都已完成
	// 以后会介绍一种更合适的方法（Mutex），这里暂时使用sleep
	time.Sleep(time.Second)
	fmt.Println("x 的值：", x)
}

func TestChan7(t *testing.T) {
	done := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		done <- true
	}()
	<-done
}

func add(count *int, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		*count = *count + 1
	}
	wg.Done()
}

func TestChan8(t *testing.T) {
	var wg sync.WaitGroup
	count := 0
	wg.Add(3)
	go add(&count, &wg)
	go add(&count, &wg)
	go add(&count, &wg)

	wg.Wait()
	fmt.Println("count 的值为：", count)
}

func add1(count *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		*count = *count + 1
		lock.Unlock()
	}
	wg.Done()
}

func TestChan9(t *testing.T) {
	var wg sync.WaitGroup
	lock := &sync.Mutex{}
	count := 0
	wg.Add(3)
	go add1(&count, &wg, lock)
	go add1(&count, &wg, lock)
	go add1(&count, &wg, lock)

	wg.Wait()
	fmt.Println("count 的值为：", count)
}

func TestChan10(t *testing.T) {
	lock := &sync.RWMutex{}
	lock.Lock()

	for i := 0; i < 4; i++ {
		go func(i int) {
			fmt.Printf("第 %d 个协程准备开始... \n", i)
			lock.RLock()
			fmt.Printf("第 %d 个协程获得读锁, sleep 1s 后，释放锁\n", i)
			time.Sleep(time.Second)
			lock.RUnlock()
		}(i)
	}

	time.Sleep(time.Second * 2)

	fmt.Println("准备释放写锁，读锁不再阻塞")
	// 写锁一释放，读锁就自由了
	lock.Unlock()

	// 由于会等到读锁全部释放，才能获得写锁
	// 因为这里一定会在上面 4 个协程全部完成才能往下走
	lock.Lock()
	fmt.Println("程序退出...")
	lock.Unlock()
}

func TestChan11(t *testing.T) {
	pipeLine := make(chan string)
	go func() {
		time.Sleep(2)
		pipeLine <- "hello world"
		pipeLine <- "hello China"
		close(pipeLine)
	}()
	for data := range pipeLine {
		fmt.Println(data)
	}
}

type Pool struct {
	work chan func()   // 任务
	sem  chan struct{} // 数量
}

func New(size int) *Pool {
	return &Pool{
		work: make(chan func()),
		sem:  make(chan struct{}, size),
	}
}

func (p *Pool) NewTask(task func()) {
	select {
	case p.work <- task: 				/* 在case中，无缓冲channel当没有接受者的时候，是不会阻塞的（也不会写进去，因为需要接收端），如果不是在case中，会阻塞的 */
	case p.sem <- struct{}{}: 			//struct{}是类型
		fmt.Println("hahaha")
		go p.worker(task)
	}
}

func (p *Pool) worker(task func()) {
	defer func() { <-p.sem }()
	for {
		fmt.Println("wwwww")
		task()	//执行任务
		task = <-p.work //已启动的协程，接受新的任务（有了接受端了0）
	}
}

func TestChan13(t *testing.T) {
	pool := New(2)

	/*for i := 1; i < 5; i++ {
		pool.NewTask(func() {
			fmt.Println("******")
			time.Sleep(10 * time.Second)
			fmt.Println(time.Now())
		})
	}*/

	pool.NewTask(func() {
		fmt.Println("11111")
		//time.Sleep(5 * time.Second)
		fmt.Println(time.Now())
	})

	pool.NewTask(func() {
		fmt.Println("22222")
		//time.Sleep(5 * time.Second)
		fmt.Println(time.Now())
	})

	pool.NewTask(func() {
		fmt.Println("33333")
		//time.Sleep(5 * time.Second)
		fmt.Println(time.Now())
	})

	pool.NewTask(func() {
		fmt.Println("44444")
		//time.Sleep(5 * time.Second)
		fmt.Println(time.Now())
	})

	// 保证所有的协程都执行完毕
	time.Sleep(5 * time.Second)
}
