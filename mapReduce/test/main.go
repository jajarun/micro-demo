package main

import (
	"fmt"
	"microDemo/mapReduce"
	"time"
)

func main() {
	mp := new(mapReduce.MapReduce)
	mp.AddMap(func() any {
		fmt.Println("part 1 start")
		time.Sleep(time.Second * 2)
		fmt.Println("part 1 end")
		return (1 + 3 + 5)
	})

	mp.AddMap(func() any {
		fmt.Println("part 2 start")
		time.Sleep(time.Second * 2)
		fmt.Println("part 2 end")
		return (2 + 7 + 8)
	})

	mp.Reduce(func(mapResults []any) {
		fmt.Println("reduce  start")
		sum := 0
		for _, mapResult := range mapResults {
			sum += mapResult.(int)
		}
		fmt.Println(sum)
	})
}
