// Most of the examples in this lesson are from the following blog:
// https://go.dev/blog/slices-intro

package main

import "fmt"

func main() {
	// ------------------------------------------------------------
	// SLICES
	// ------------------------------------------------------------

	// Slices in go are data structues built on top of array and similar to array
	// but offers more flexibility as size of the slice is not fixed on initzalization.

	// Lets create a slice with a slice literal.
	// As you will notice, the syntax for creation is similar to array
	// but you don't have to specify the size

	num := []int{1, 2, 3, 4, 5}
	fmt.Println(num)

	// you can get length of a slice with the inbuilt `len` function
	fmt.Printf("Size of slice %v of type %T is %v\n", num, num, len(num))
	// Slice is backed by an array and is just a representation of the `slice` of the array.
	// You can get the capacity of the underlying array by using the inbuilt funciton `cap`

	// The initial capcity will be similar to the size of he slice
	fmt.Printf("Capacity of array is %v\n", cap(num))

	// You can create new slices from existing arrays or slices by "slicing" them.
	// The "half open" [:] operator slices the array/slice.
	// The half open operator means that new slice created will be inclusive of the first specified index
	// and exlcusive of the last specified index.
	num1 := num[1:3] // [2 3]
	fmt.Println(num1)

	// both the first and the last index are optional, defaulting to 0 and slice length.
	num2 := num[:3] // [1 2 3]
	fmt.Println(num2)

	num3 := num[3:] // [4 5]
	fmt.Println(num3)

	// ------------------------------------------------------------
	// IMPORTANT
	// ------------------------------------------------------------
	// num1, num2, num3 all point to the same underlying array.
	// From the above code num1[0] == num2[1]
	// so if we change any of those value, all the over slices will be affected

	num1[0] = 10
	fmt.Printf("num\t %v\t capacity = %v\n", num, cap(num))      // [1 10 3 4 5]
	fmt.Printf("num1\t %v\t\t capacity = %v\n", num1, cap(num1)) // [10 3]
	fmt.Printf("num2\t %v\t capacity = %v\n", num2, cap(num2))   // [1 10 3]

	// Even though you have slice that has an array elements more than its index,
	// You cannot access the index that is out of bound of the slice.
	// fmt.Println("num1[3]", num1[3])

	// Both num and newNum point to the same array
	// newNum := num[:]
	// You can use the `make` function to create a new slice.

	ranks := make([]int, 10) // all the elements will be initalized to zero-value of int that is 0
	fmt.Println("ranks", ranks)

	// ------------------------------------------------------------
	// GROWING A SLICE
	// ------------------------------------------------------------

	// You can grow a slice by creating a new slice with a bigger capacity,
	// Copying the data to the new slice and assiging the slice to the
	// old slice

	a1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice 'a' = %v with capcaity %v\n", a1, cap(a1))

	b1 := make([]int, 10)

	//nolint
	for i := range a1 { //nolint:staticcheck
		b1[i] = a1[i]
	}
	a1 = b1
	fmt.Printf("slice 'a' = %v with capcaity %v\n", a1, cap(a1))

	// There is a `copy` method that copies data from one slice to another
	a2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice 'a' = %v with capcaity %v\n", a2, cap(a2))

	b2 := make([]int, 10)
	copy(b2, a2)
	a2 = b2

	// ------------------------------------------------------------
	// APPEND DATA TO A SLICE
	// ------------------------------------------------------------
	// Look at the AppendData function below.
	// We can write custom append fuction to append data
	// to an existing slice
	s1 := []int{1, 2, 3, 4}
	s1 = AppendData(s1, 5, 6, 7)
	fmt.Println(s1)

	s2 := []int{1, 2, 3, 4}
	s2 = append(s2, 4, 4, 5, 6)
	fmt.Println(s2)

	// You can also use append to combine two slices together
	// by using the '...' operator on the 2nd slice
	s3 := []string{"a", "b", "c", "d"}
	s4 := []string{"e", "f", "g", "h"}
	s3 = append(s3, s4...)
	fmt.Println(s3)

	// ------------------------------------------------------------
	// DELETING FROM A SLICE
	// ------------------------------------------------------------

	// you delete elements from slice using 'slicing' and append
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	limit := len(numbers)
	index := 0

	for index < limit {
		v := numbers[index]
		if v%2 == 0 {
			numbers = DeleteAtIndex(numbers, index)
			limit--
			continue
		}
		index++
	}

	fmt.Println("Odd numbers =", numbers)
}

// Function to delete an element from the slice
func DeleteAtIndex(slice []int, index int) []int {
	//if we are deleting the last element just slice it of
	if index == len(slice)-1 {
		return slice[:index]
	}

	// other wise, we append two slices together
	// slice 1: all elments upto that index(exlcusive)
	// slice 2: all elemnt from the next index (inclusive) till the end of slice
	return append(slice[:index], slice[index+1:]...)
}

// AppendData
func AppendData(slice []int, data ...int) []int {
	m := len(slice)
	n := len(data)

	// If required, increase the capacity of the slice
	if m+n > cap(slice) {
		newSlice := make([]int, (m+n+1)*2) // +1 because m + n can be 0
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[:m+n]   // adjust the lenght of the slice to match the total elements
	copy(slice[m:], data) // Add data
	return slice
}
