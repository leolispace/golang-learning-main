package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

// Q1: 指针 - 增加值
// 编写一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10
func AddTen(val *int) {
	// 解引用 *val 拿到原变量，直接 +10
	*val = *val + 10
}

// Q2: 指针 - 切片元素乘2
// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
func DoubleSlice(slice *[]int) {
	// 解引用得到切片，然后遍历修改
	for i := range *slice {
		(*slice)[i] *= 2
	}
}

// Q3: Goroutine - 奇偶数打印
// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 提示：为了测试能看到输出，可以使用 sync.WaitGroup 确保主程序等待协程结束
func PrintOddEven() {
	// 等待组：等待两个协程完成
	var wg sync.WaitGroup
	wg.Add(2) // 登记 2 个协程

	// 协程1：打印 1,3,5,7,9
	go func() {
		defer wg.Done() // 执行完自动减1
		for i := 1; i <= 10; i += 2 {
			fmt.Println("奇数:", i)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// 协程2：打印 2,4,6,8,10
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Println("偶数:", i)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	wg.Wait() // 主协程等待
	fmt.Println("打印完成！")
}

// Q4: Goroutine - 任务调度
// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
func TaskScheduler(tasks []func()) {
	var wg sync.WaitGroup

	// 遍历所有任务
	for i, task := range tasks {
		wg.Add(1)

		// 启动协程执行任务
		go func(index int, t func()) {
			defer wg.Done()

			// 记录开始时间
			start := time.Now()

			// 执行任务
			t()

			// 计算耗时
			cost := time.Since(start)
			fmt.Printf("任务 %d 执行完成，耗时：%v\n", index, cost)
		}(i, task)
	}

	// 等待所有任务执行完毕
	wg.Wait()
	fmt.Println("所有任务已执行完成！")
}

// Q5: 面向对象 - 接口
// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

// TODO: 为 Rectangle 实现 Shape 接口
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 矩形实现 Perimeter
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

// TODO: 为 Circle 实现 Shape 接口
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// PrintShapeInfo 多态函数
func PrintShapeInfo(s Shape, name string) {
	fmt.Printf("%s: Area=%.2f, Perimeter=%.2f\n",
		name, s.Area(), s.Perimeter())
}

// Q6: 面向对象 - 组合
// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段
type Person struct {
	Name string
	Age  int
}

// 再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段
type Employee struct {
	Person
	EmployeeID string
}

// TODO: 为 Employee 结构体实现一个 PrintInfo() string 方法，返回员工的信息 (格式自定，包含Name, Age, EmployeeID)
func (e Employee) PrintInfo() string {
	return fmt.Sprintf("姓名：%s，年龄：%d，员工ID：%s", e.Name, e.Age, e.EmployeeID)
}

// Q7: Channel - 生产者消费者
// 编写一个程序，使用通道实现两个协程之间的通信。
// 一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 为了方便测试，请将接收到的数字返回
func ProducerConsumer() []int {
	var result []int
	ch := make(chan int) // 通道
	var wg sync.WaitGroup

	// 消费者：接收数字，存入 result
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch {
			result = append(result, num)
			fmt.Println("消费者打印：", num)
		}
	}()

	// 生产者：发送 1~10
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch) // 发完关闭通道
	}()

	wg.Wait()     // 等待消费者结束
	return result // 返回收到的数字
}

// Q8: Channel - 缓冲通道
// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
func BufferedChannel() {
	// 1. 创建一个带缓冲的通道，缓冲大小可以设为 10（随便设，只要>0就是缓冲通道）
	ch := make(chan int, 10)

	var wg sync.WaitGroup

	// 生产者：发送 100 个整数
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		// 发送完关闭通道，消费者才能正常退出循环
		close(ch)
		fmt.Println("✅ 生产者发送完成")
	}()

	// 消费者：接收并打印
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 遍历通道接收数据
		for num := range ch {
			fmt.Printf("消费者接收：%d\n", num)
		}
		fmt.Println("✅ 消费者接收完成")
	}()

	// 等待两个协程结束
	wg.Wait()
	fmt.Println("🎉 缓冲通道演示全部完成！")
}

// Q9: 锁机制 - Mutex
// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func MutexCounter() int {
	var count int
	var mu sync.Mutex     // 互斥锁
	var wg sync.WaitGroup // 等待协程完成

	// 启动 10 个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// 每个协程累加 1000 次
			for j := 0; j < 1000; j++ {
				mu.Lock()   // 加锁
				count++     // 修改共享变量
				mu.Unlock() // 解锁
			}
		}()
	}

	wg.Wait() // 等待所有协程结束
	return count
}

// Q10: 锁机制 - Atomic
// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func AtomicCounter() int32 {
	var count int32
	var wg sync.WaitGroup

	// 启动 10 个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 每个协程 1000 次原子递增
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&count, 1) // 原子操作 +1
			}
		}()
	}

	wg.Wait()
	return count // 一定返回 10000
}
