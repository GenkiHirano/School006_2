package main

import "fmt"

type Calculation interface {
	total() int
	average() int
}

type Subject struct {
	Math, English int
}

func (r Subject) total() int {
	return r.Math + r.English
}

func (r Subject) average() int {
	return (r.Math + r.English) / 2
}

func Calculation2(g Calculation) {
	fmt.Println("2教科の合計：", g.total())
	fmt.Println("2教科の平均：", g.average())
}

func main() {
	var a, b int
	fmt.Printf("数学の点数：")
	fmt.Scan(&a)
	fmt.Printf("英語の点数：")
	fmt.Scan(&b)
	r := Subject{Math: a, English: b}

	Calculation2(r)
}
