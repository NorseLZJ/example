package main

import "fmt"

func main() {
	cur := make([]int, 0, 5)
	cur = append(cur, 2, 5, 6, 4, 7)
	val := uint64(0)
	for _, v := range cur {
		val |= 1 << uint64(v)
	}
	fmt.Printf("val bit %b\n", val)
	for _, v := range cur {
		fmt.Printf("idx : %d ret : %d  \n", v, val&(1<<uint64(v)))
	}
	// unDefine
	fmt.Printf("idx : %d ret : %d  \n", 10, val&(1<<uint64(10)))
	fmt.Printf("idx : %d ret : %d  \n", 1, val&(1<<uint64(1)))
	fmt.Printf("idx : %d ret : %d  \n", 16, val&(1<<uint64(16)))

	// max
	val |= 1 << 63
	fmt.Printf("bit %b\n", val)
	fmt.Printf("idx : %d ret : %d  \n", 63, val&(1<<uint64(63)))

	// err overflows uint64
	//fmt.Printf("idx : %d ret : %d  \n", 64, val&(1<<uint64(64)))
}
