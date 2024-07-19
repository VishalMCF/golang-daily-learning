package algo

import (
	"fmt"
)

func merge(data []int, start, end, mid int) {
	p1 := start
	p2 := mid + 1

	newData := make([]int, end-start+1)

	cnt := 0

	for p1 <= mid && p2 <= end {
		if data[p1] <= data[p2] {
			newData[cnt] = data[p1]
			p1++
		} else {
			newData[cnt] = data[p2]
			p2++
		}
		cnt++
	}

	for p1 <= mid {
		newData[cnt] = data[p1]
		p1++
		cnt++
	}

	for p2 <= end {
		newData[cnt] = data[p2]
		p2++
		cnt++
	}

	data = newData
}

type sort interface {
	sortSlice(data []int, start, end int)
}

type MergeSort struct {
	data []int
}

func (m MergeSort) sortSlice(start, end int) {
	data := m.data
	if start >= end {
		return
	}
	mid := start + (end-start)/2
	m.sortSlice(start, mid)
	m.sortSlice(mid+1, end)
	merge(data, start, end, mid)
}

func main() {
	fmt.Println("Hello World")
	myData := []int{18, 17, 15, 11, 7, 6, 12, 19, 3, 10}

	for d := range myData {
		fmt.Println("element :=> ", d)
	}
}
