package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5)

	//get the length of the ring
	n:=r.Len()

	//inisialisasi the ring with some integer value
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next() 
	}

	//iterasi through the ring and print its content
	r.Do(func (p any)  {
		fmt.Println(p.(int))
	})
}
