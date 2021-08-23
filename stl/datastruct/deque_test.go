package datastruct

import (
	"container/list"
	"fmt"
)

func ExampleDeque() {
	// list.New()は双方向の連結リストを生成する→dequeとして利用できる
	deque := list.New()

	// 先頭に要素を追加する
	deque.PushFront(-5) // [-5]
	deque.PushFront(10) // [10, -5]

	// 先頭の要素を取得する
	// Valueはinterface型なので，型を意識した操作 (構造体の属性へのアクセス等) を
	// したい場合にはキャストする必要がある
	fmt.Println(deque.Front().Value) // 10

	// 先頭の要素を削除する
	// Goの実装ではPopFrontは存在しないため，
	// 先頭要素へのポインタをRemoveの引数として与える必要がある
	deque.Remove(deque.Front()) // [-5]

	// 末尾に要素を追加する
	deque.PushBack(3)  // [-5, 3]
	deque.PushBack(-1) // [-5, 3, -1]

	// 末尾の要素を取得する
	fmt.Println(deque.Back().Value) // -1

	// 末尾の要素を削除する
	// Goの実装ではPopBackは存在しないため，
	// 末尾要素へのポインタをRemoveの引数として与える必要がある
	deque.Remove(deque.Back()) // [-5, -3]

	// Removeは削除する要素の値を返却するので，要素の取得と削除を一括で行うこともできる
	fmt.Println(deque.Remove(deque.Back())) // -3 ; [-5]

	// 特定要素の前後に要素を追加することもできる
	// 末尾の要素の1つ手前に新しい要素2を追加する
	e1 := deque.InsertBefore(2, deque.Back()) // 2 ; [2, -5]
	// 新しく追加した要素の後ろに-2を追加する
	deque.InsertAfter(-2, e1) // [2, -2, -5]

	// 先頭から順に走査したい場合にはNextを利用する
	for e := deque.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // [2, -2, -5]
	}

	// 末尾から順に走査したい場合にはPrevを利用する
	for e := deque.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value) // [-5, -2, 2]
	}

	// dequeの要素が無くなるまで処理をしたいときはLenで判定できる
	for deque.Len() > 0 {
		fmt.Println(deque.Remove(deque.Front())) // [2, -2, -5]
	}

	// Output:
	// 10
	// -1
	// 3
	// 2
	// -2
	// -5
	// -5
	// -2
	// 2
	// 2
	// -2
	// -5
}
