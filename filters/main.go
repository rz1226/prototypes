package main

import (
	"fmt"
	"github.com/rz1226/prototypes/filters/filter"
)

func main() {
	testbloom()
}

func testbloom() {
	n := 301

	for i := 0; i < n; i++ {
		filter.SetBloom("a_" + fmt.Sprint(i))

	}

	for i := 0; i < n; i++ {
		res := filter.GetBloom("a_" + fmt.Sprint(i))
		if !res {
			fmt.Println("找不到", i)
		} else {
			fmt.Println("找得到", res)
		}

	}
}

func testcache() {
	n := 301

	for i := 0; i < n; i++ {
		filter.Set(fmt.Sprint(i), "a_"+fmt.Sprint(i))

	}

	for i := 0; i < n; i++ {
		res, err := filter.Get(fmt.Sprint(i))
		if err != nil {
			fmt.Println("找不到", i)
		} else {
			fmt.Println("找得到", res)
		}

	}

}
