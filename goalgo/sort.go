package goalgo

import (
	"fmt"
)

func SortTest() {
	var test1 = []int{4, 1, 31, 65, 34, 90, 15, 4, 34, 56, 90, 101, 35, 54, 83, 3, 40, 30}
	result1 := quicksort(test1)
	result2 := mergesort(test1)
	fmt.Println(test1)
	fmt.Println(result1)
	fmt.Println(result2)
}

func mergesort(arr []int) []int {
	var length = len(arr)
	var sorted = make([]int, length)
	copy(sorted, arr)
	var temp = make([]int, length)
	var merge = func(left, middle, right int) {
		i, j, k := left, middle+1, 0
		for i <= middle && j <= right {
			if sorted[i] < sorted[j] {
				temp[k] = sorted[i]
				i++
			} else {
				temp[k] = sorted[j]
				j++
			}
			k++
		}
		for i <= middle {
			temp[k] = sorted[i]
			i++
			k++
		}
		for j <= right {
			temp[k] = sorted[j]
			j++
			k++
		}
		k = 0
		for left <= right {
			sorted[left] = temp[k]
			left++
			k++
		}

	}
	var sort func(int, int)
	sort = func(left, right int) {
		if left < right {
			var middle int = (left + right) / 2
			sort(left, middle)
			sort(middle+1, right)
			merge(left, middle, right)
		}
	}
	sort(0, length-1)
	return sorted
}

//快速排序
func quicksort(arr []int) []int {
	/*
	 *[]int 是指针
	 *所以我这里拷贝复制一下入参的数组参数
	 */
	var hold = make([]int, len(arr))
	copy(hold, arr)

	//定义函数
	swap := func(a, b *int) {
		//指针，和C类似，语法不一样
		//*pointer 取指针指向内存的值 ,&a,获取保存变量a的地址的指针
		var tem int = *a
		*a = *b
		*b = tem

	}

	//也可以这样定义函数
	var procedure func(int, int)
	procedure = func(left, right int) {
		var i, j, pilot int
		if left > right {
			return
		}
		pilot = hold[left]
		i, j = left, right
		for i != j {
			for hold[j] >= pilot && i < j {
				j--
			}
			for hold[i] <= pilot && i < j {
				i++
			}
			if i < j {
				swap(&hold[i], &hold[j]) //这里用到了 &
			}
		}
		hold[left] = hold[i]
		hold[i] = pilot
		//递归调用
		procedure(left, i-1)
		procedure(i+1, right)
	}

	procedure(0, len(hold)-1)
	return hold
}
