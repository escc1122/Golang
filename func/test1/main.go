package main

import (
	"fmt"
	"math"
)

/*
Go 没有类。不过你可以为结构体类型定义方法。

方法就是一类带特殊的 接收者 参数的函数。

方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。

在此例中，Abs 方法拥有一个名为 v，类型为 Vertex 的接收者。
*/
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
你也可以为非结构体类型声明方法。

在此例中，我们看到了一个带 Abs 方法的数值类型 MyFloat。

你只能为在同一包内定义的类型的接收者声明方法，而不能为其它包内定义的类型（包括 int 之类的内建类型）的接收者声明方法。

（译注：就是接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法。）
*/
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	abs := v.Abs()
	fmt.Println(abs)

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
