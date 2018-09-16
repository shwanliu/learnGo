package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"time"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

var c, python, java bool = true, true, true //这边是一个包的变量的定义 进行赋值
//函数外的语句都必须关键字开始(var func)) :=使用在func里面

var (
	ToBe   bool       = true
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// part1：
//go 在不同类型的项之间的赋值时需要显示转换！！！
// 常量不能使用 := 语法声明 ，前面加上const
//int 类型最大可以存储一个64位的整数，有时会更小

//part2：
//go 语言只有一种循环就是for循环
//for的初始化语句和后置语句可以省去，其分号可以去掉成为while一样
//if 语句 与for循环类似 无需（），{}必须
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

//练习：循环与函数
func Sqrt(x float64) float64 {
	var z float64 = 1.0
	for i := 0; i < 10; i++ {
		z -= (math.Pow(z, 2) - x) / 2.0 * z
		// fmt.Println(z)
	}
	return z
}

func switch_test() {

	//switch语句是，自动提供了break，case的取值无需为常量，且取值不必为整数，没有条件就和switch true一样
	// fmt.Print("Go runs on")
	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("OS X.")
	// case "linux":
	// 	fmt.Println("Linux.")
	// default:
	// 	//freebsd,openbsd
	// 	//plan9,windows...
	// 	fmt.Printf("%s.", os)
	// }
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")

	}

	// defer语句会将函数推迟到外层函数返回之后执行

}

//Go 拥有指针，指针保存了值的内存地址,没有指针运算
//  var p *int
//  i := 42
//  p = &i

// 一个结构体(struct)就是一个字段的集合
//结构体指针，一个指向结构体指针，可以直接使用p.x的形式访问
type Vertex struct {
	X int
	Y int
}

func main() {
	java = true
	// var j int = 1
	defer fmt.Println("1111111111111111111111111111") //defer语句
	fmt.Println(add(11, 55))
	// fmt.Println(swap("前面", "后面"))
	// fmt.Println(i, c, python, java)
	// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// fmt.Printf("Type: %T Value: %v\n", z, z)
	sum := 1
	// for i := 0; i < 10; i++ {
	// sum += i
	// }
	for sum < 50 {
		sum += sum
	}
	fmt.Println(sum)
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(Sqrt(2))
	// switch_test()
	// fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	fmt.Println(v.X, v.Y)

}
