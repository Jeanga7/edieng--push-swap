package main

import (
	"fmt"
	"os"
	"strconv"
)

type Stack []interface {
}

// push the element at the top of the stack
func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

// remove the top element of the stack
func (s *Stack) Pop() (item interface{}) {
	item, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]
	return
}

// verify if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// return the top element of the stack otherwise return nil
func (s *Stack) Peek() (item interface{}) {
	if s.IsEmpty() {
		return nil
	}
	return (*s)[len(*s)-1]
}

// return the length of the stack
func (s *Stack) Size() int {
	return len(*s)
}

//////////////////////////////////////////////////

var S Stack

// var A []int
// var B []int

func Sa2(a []interface{}) []interface{} {
	if len(a) > 1 {
		res := append([]interface{}{a[1]}, a[0])
		for _, v := range a[2:] {
			res = append(res, v)
		}
		return res
	}
	return a
}

func Sb2(a []interface{}) []interface{} {
	if len(a) > 1 {
		res := append([]interface{}{a[1]}, a[0])
		for _, v := range a[2:] {
			res = append(res, v)
		}
		return res
	}
	return a
}

func Pb2(a *Stack, b *Stack) (Stack, Stack) {
	x := a.Pop()
	b.Push(x)
	return *a, *b
}

func Pa2(a *Stack, b *Stack) (Stack, Stack) {
	x := b.Pop()
	a.Push(x)
	return *a, *b
}

// func Ss2(a []int) []int {
// 	res := Sa2(a)
// 	return Sb2(res)
// }

// func Sa(argument *string) string {
// 	if len(*argument) > 1 {
// 		return string((*argument)[1]) + string((*argument)[0]) + (*argument)[2:]
// 	}
// 	return *argument
// }

// func Sb(argument *string) string {
// 	if len(*argument) > 1 {
// 		return string((*argument)[1]) + string((*argument)[0]) + (*argument)[2:]
// 	}
// 	return *argument
// }

// func Ss(argument *string) string {
// 	res := Sa(argument)
// 	return Sb(&res)
// }

// func Pb(argument *string) string {
// 	// B = string((*argument)[:1])
// 	if len(*argument) > 1 {
// 		return string((*argument)[1:])
// 	}
// 	return *argument
// }

// func Pa(argument *string) string {
// 	// A = string((*argument)[:1])
// 	if len(*argument) > 1 {
// 		return string((*argument)[1:])
// 	}
// 	return *argument
// }

func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		var A_number Stack
		var B_number Stack
		for _, v := range arg {
			n, err := strconv.Atoi(string(v))
			if err != nil {
				fmt.Println("Error")
				return
			}
			A_number = append(A_number, n)
		}

		// sort.Ints(A_number)

		for _, v := range A_number {
			S.Push(v)
		}
		for i, n := range A_number {
			fmt.Printf("%d\n", n)
			if i == (A_number.Size())-1 {
				fmt.Println("=\t=")
				fmt.Println("a\tb")
			}
		}
		fmt.Println("//////////////////////////////////////////////////////")
		fmt.Println("After popping the top element ")
		// fmt.Println(S)
		// a := A_number.Pop()
		// fmt.Println("pppppppppppppppppppppp")
		// fmt.Println(a)
		// B_number = append(B_number, a.(int))
		// B_number.Push(a)
		// fmt.Println("This is the stack after removing the top element")
		// fmt.Println(S)
		// fmt.Println("********************")
		// fmt.Println(res_number)
		// A = arg
		// fmt.Println("")
		// for _, n := range arg {
		// 	if (n < '0' || n > '9') && n != 32 {
		// 		fmt.Println("Error")
		// 		return
		// 	}
		// }
		// fmt.Println(A_number)
		// fmt.Println(B_number)

		Pb2(&A_number, &B_number)

		for i, n := range A_number {
			var B_nb string
			if i < len(B_number) {
				B_nb = fmt.Sprintf("%d", B_number[i])
			} else {
				B_nb = ""
			}
			fmt.Printf("%d\t%s\n", n, B_nb)
			if i == (A_number.Size())-1 {
				fmt.Println("=\t=")
				fmt.Println("a\tb")
			}
		}
		fmt.Println(Sa2(A_number))
		// res := Sa(&arg)
		// fmt.Println(res)
	} else {
		return
	}
}
