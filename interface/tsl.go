package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var gs [5]struct { // 模拟实现tsl
		id     int // 编号
		result int // 返回值
	}

	for i := 0; i < len(gs); i++ {
		wg.Add(1)

		go func(id int) { // 使用参数传递，避免闭包传递延迟
			gs[id].id = id
			gs[id].result = (id + 1) * 100
			wg.Done() // 避免延迟调用的性能损耗
		}(i)
	}

	wg.Wait()
	fmt.Printf("%+v\n", gs)

}
