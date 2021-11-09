package homework006

import "fmt"

// Triangle 構造体を定義してください。
type Triangle struct {
	Bottom, Height int
}

func (t Triangle) Test_3() int {
	fmt.Println("homework006_3")
	// 第6回スクール 宿題3 メソッドを実装してください。
	// Test_3関数をメソッドにしてください。
	//t := Triangle{Bottom: 25, Height: 20}
	//fmt.Println("Triangle: ", t.Triangle())
	return t.Bottom * t.Height / 2
}
