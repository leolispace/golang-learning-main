package main

import (
	"fmt"
	"time"
)

// Q1: 指针 - 增加值
// 编写一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10
func test_AddTen() {
	fmt.Println("======================= Q1: 指针 - 增加值 ============================")
	num := 5
	fmt.Println("加之前:", num) // 5

	// 传入变量的地址
	AddTen(&num)

	fmt.Println("加之后:", num) // 15
}

// Q2: 指针 - 切片元素乘2
// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
func test_DoubleSlice() {
	fmt.Println("=======================  Q2: 指针 - 切片元素乘2 ============================")
	s := []int{1, 2, 3, 4}
	fmt.Println("修改前:", s)

	// 传入切片地址
	DoubleSlice(&s)

	fmt.Println("修改后:", s) // [2 4 6 8]
}

// Q3: Goroutine - 奇偶数打印
// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 提示：为了测试能看到输出，可以使用 sync.WaitGroup 确保主程序等待协程结束
func test_PrintOddEven() {
	fmt.Println("=======================  Q3: Goroutine - 奇偶数打印 ============================")
	PrintOddEven()
}

// Q4: Goroutine - 任务调度
// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
func test_TaskScheduler() {
	fmt.Println("=======================  Q4: Goroutine - 任务调度 ============================")
	// 定义3个测试任务
	tasks := []func(){
		func() { time.Sleep(200 * time.Millisecond) }, // 任务1
		func() { time.Sleep(300 * time.Millisecond) }, // 任务2
		func() { time.Sleep(150 * time.Millisecond) }, // 任务3
	}

	// 调用调度器
	TaskScheduler(tasks)
}

// Q5: 面向对象 - 接口
func test_PrintShapeInfo() {
	fmt.Println("=======================  Q5: 面向对象 - 接口 ============================")
	// 不同形状
	shapes := []Shape{
		&Rectangle{Width: 10, Height: 5},
		&Circle{Radius: 5},
	}

	names := []string{"矩形", "圆形", "三角形"}

	for i, shape := range shapes {
		PrintShapeInfo(shape, names[i])
	}
}

// Q6: 面向对象 - 组合
func test_EmployeePrintInfo() {
	fmt.Println("=======================  Q6: 面向对象 - 组合 ============================")
	emp := Employee{
		Person: Person{
			Name: "张三",
			Age:  25,
		},
		EmployeeID: "1001",
	}

	fmt.Println(emp.PrintInfo())
}

// Q7: Channel - 生产者消费者
// 编写一个程序，使用通道实现两个协程之间的通信。
// 一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 为了方便测试，请将接收到的数字返回
func test_ProducerConsumer() {
	fmt.Println("=======================  Q7: Channel - 生产者消费者 ============================")
	res := ProducerConsumer()
	fmt.Println("最终返回的数字：", res)
}

// Q8: Channel - 缓冲通道
// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
func test_BufferedChannel() {
	fmt.Println("=======================  Q8: Channel - 缓冲通道 ============================")
	BufferedChannel()
}

// Q9: 锁机制 - Mutex
// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func test_MutexCounter() {
	fmt.Println("======================= Q9: 锁机制 - Mutex ============================")
	println(MutexCounter()) // 输出：10000
}

// Q10: 锁机制 - Atomic
// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func test_AtomicCounter() {
	fmt.Println("======================= Q10: 锁机制 - Atomic ============================")
	println(AtomicCounter()) // 输出：10000
}

func main() {
	test_AddTen()
	test_DoubleSlice()
	test_PrintOddEven()
	test_TaskScheduler()
	test_PrintShapeInfo()
	test_EmployeePrintInfo()
	test_ProducerConsumer()
	test_BufferedChannel()
	test_MutexCounter()
	test_AtomicCounter()
}
