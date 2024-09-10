package utils

import "fmt"

func SliceFromArrayDemo() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("---------------------")
	fmt.Println("array:", array)
	slice := array[1:3]
	fmt.Println("slice:", slice)
	fmt.Println("slice length:", len(slice))
	fmt.Println("slice capacity:", cap(slice))
	fmt.Println("array:", array)
	fmt.Println("---------------------")
	slice[0] = 55
	fmt.Println("** modify slice [0] to 55")
	fmt.Println("slice:", slice)
	fmt.Println("slice length:", len(slice))
	fmt.Println("slice capacity:", cap(slice))
	fmt.Println("array:", array)
	fmt.Println("---------------------")

	slice = append(slice, 6, 7, 8)
	fmt.Println("** append 6, 7, 8 to slice")
	fmt.Println("slice:", slice)
	fmt.Println("slice length:", len(slice))
	fmt.Println("slice capacity:", cap(slice))
	fmt.Println("array:", array)
	fmt.Println("---------------------")

}

func SliceMakeDemo() {
	slice := make([]int, 5)
	fmt.Println("slice:", slice)
	fmt.Println("slice length:", len(slice))
	fmt.Println("slice capacity:", cap(slice))
	fmt.Println("---------------------")
	slice = append(slice, 6)
	fmt.Println("** append 6 to slice")
	fmt.Println("slice:", slice)
	fmt.Println("slice length:", len(slice))
	fmt.Println("slice capacity:", cap(slice))
	fmt.Println("---------------------")
	slice2 := append(slice, 7, 8, 9, 10, 11, 12)
	fmt.Println("** append 7, 8, 9, 10, 11, 12 to slice2 using slice as source")
	fmt.Println("slice:\t", slice)
	fmt.Println("slice2:\t", slice2)
	fmt.Println("slice length:", len(slice))
	fmt.Println("slice2 length:", len(slice2))
	fmt.Println("slice capacity:", cap(slice))
	fmt.Println("slice2 capacity:", cap(slice2))
	fmt.Println("---------------------")
	slice[0] = 25
	slice2[0] = 35
	fmt.Println("** modify slice[0] to 25 and slice2[0] to 35")
	fmt.Println("slice:\t", slice)
	fmt.Println("slice2:\t", slice2)
	fmt.Println("slice length:", len(slice))
	fmt.Println("slice2 length:", len(slice2))
	fmt.Println("slice capacity:", cap(slice))
	fmt.Println("slice2 capacity:", cap(slice2))

}
