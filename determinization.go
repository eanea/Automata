package main

import (
	"fmt"
	"sort"
)

const lambda string = "lambda"

/**
5
8

0 1 a
0 2 a
1 3 b
3 1 lambda
3 3 a
2 4 c
4 4 a
4 4 b

0 0 0 1 1

0
 */

func contains (str string, v []string) bool {
	for i := range v {
		if v[i] == str {
			return true
		}
	}
	return false
}

func isEqual (zr[] int, t[]int) bool {
	sort.Ints(zr)
	sort.Ints(t)
	if len(zr) == len(t) {
		for i := range zr {
			if zr[i] != t[i] {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func Dfs(q [][] []string, vert int, out_ptr *[]int, mark_ptr *[]bool) {
	mark := *(mark_ptr)
	if !mark[vert] {
		*(out_ptr) = append(*(out_ptr), vert)
		mark[vert] = true
		for i := range q[vert] {
			if contains(lambda, q[vert][i]) {
				Dfs(q, i, out_ptr, &mark)
			}
		}
	}
}

func closure (q[][] []string, arr []int) []int  {

	out := make ([] int, 0)	// C
	mark := make ([] bool, len(q))
	for i := range arr {
		Dfs(q, arr[i], &out, &mark)
	}
	return out
}


func edgeGoSymbol (q[][] []string, z[]int, symbol string) []int {
	zr := make ([]int, 0)
	for i := range z {
		for j := range q[z[i]] {
			if contains(symbol, q[z[i]][j]) {
				zr = append(zr, j)
			}
		}
	}
	return zr
}

func index (zr[] int, Q[] []int) int {
	sort.Ints(zr)
	for i := range Q {
		sort.Ints(Q[i])
		if isEqual(Q[i], zr) {
			return i
			break
		}
	}
	return 0
}
/*
func print2DSlice (sl[][] int) {
	for i := range sl {
		for j := range sl[i] {
			fmt.Print(sl[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func printSlice (sl[] int) {
	for i := range sl {
		fmt.Print(sl[i], " ")
	}
	fmt.Println()
}
*/
func printElemOfMap (s[] string) {
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i])
		if i + 1 < len(s) {
			fmt.Print(", ")
		}
	}
}

func main()  {

	s1 := make ([]int, 0)
	s2 := make ([]int, 0)

	s1 = append(s1, 2)
	s1 = append(s1, 5)
	s1 = append(s1, 7)

	s2 = append(s2, 7)
	s2 = append(s2, 2)
	s2 = append(s2, 5)

	var nCond int
	fmt.Scan(&nCond)

	var m int
	fmt.Scan(&m)
	fmt.Scan("\n")

	q := make ([][] []string, nCond)

	for i := 0; i < nCond; i++ {
		q[i] = make ([] []string, nCond)
		//	mapp[i] = make (map[int] []string)
	}

	Alf := make([]string, 0)
	var v1, v2 int
	var label string
	for i := 0; i < m; i++ {
		fmt.Scan(&v1)
		fmt.Scan(&v2)
		fmt.Scan(&label)
		if label != lambda && !contains(label, Alf) {
			Alf = append(Alf, label)
		}
		//	mapp[v1][v2] = append(mapp[v1][v2], label)
		q[v1][v2] = append(q[v1][v2], label)
	}

	fmt.Scan("\n")
	final := make ([]int, nCond)
	for i := 0; i < nCond; i++ {
		fmt.Scan(&final[i])
	}

	fmt.Scan("\n")
	var base int
	fmt.Scan(&base)

	/*
	fmt.Print(nCond, "\n")
	fmt.Print(m, "\n")

	for i := 0; i < nCond; i++ {
		for j := 0; j < nCond; j++ {
			fmt.Print(q[i][j], " ")
		}
		fmt.Println()
	}


	fmt.Println()
	for i := 0; i < m; i++ {
		fmt.Print(i, " -> ", mapp[i])
		fmt.Println()
	}



	for i := 0; i < nCond; i++ {
		fmt.Print(final[i], " ")
	}
	fmt.Println()
	fmt.Print(base, "\n")

	fmt.Print(q[nCond - 1][nCond - 1][0])
	*/


	b := make ([] int, 0)
	b = append(b, base)
	q0 := closure(q, b)	//	q0
	//fmt.Print(q0, "\n")

	Q := make ([] []int, 0)
	Q = append(Q, q0)	//	Q

	F := make ([] []int, 0)

	delta := make ([][] int, nCond)
	for i := range delta {
		delta[i] = make([] int, len(Alf))
	}

	Stack := make([] []int, 0)
	Stack = append(Stack, q0)
	for len(Stack) != 0 {
		var z []int
		z, Stack = Stack[len(Stack)-1], Stack[:len(Stack)-1]

		for i := range z {
			if final[z[i]] == 1 {
				F = append(F, z)
				break
			}
		}

		for i := range Alf {
			zr := edgeGoSymbol(q, z, Alf[i])
			zr = append(closure(q, zr))

			flag := true
			for i:= range Q {
				if isEqual(Q[i], zr) {
					flag = false
					break
				}
			}
			if flag {
				Q = append(Q, zr)
				Stack = append(Stack, zr)
			}
			/*
			fmt.Print("index(z, Q) ", index(z, Q), "\n")
			fmt.Print("index(zr, Q) ", index(zr, Q), "\n")
			fmt.Print("i ", i, "\n")
			*/
			if (index(z, Q) >= nCond) {
				delta = append(delta, make([] int, len(Alf)))
			}
			delta[index(z, Q)][i] = index(zr, Q)
		}
	}

	mapp := make (map[int]map[int] []string, len(delta))
	for i := 0; i < len(delta); i++ {
		mapp[i] = make (map[int] []string)
	}

	for i := range delta {
		for j := range delta[i] {
			mapp[i][delta[i][j]] = append(mapp[i][delta[i][j]], Alf[j])
		}
	}
	/*
		for i := 0; i < m; i++ {
			fmt.Print(i, " -> ", mapp[i])
			fmt.Println()
		}
		*/
	//print2DSlice(delta)
	//fmt.Print("---------------\n")
	//print2DSlice(Q)
	//fmt.Print("---------------\n")
	//print2DSlice(F)

	fmt.Print("digraph {\n")
	fmt.Print("rankdir = LR\n")
	fmt.Print("dummy [label = \"\", shape = none]\n")

	for i := 0; i < len(Q); i++ {
		fmt.Print(i, " [label = \"", Q[i], "\", shape = ")

		flag := false
		for j := range F {
			if isEqual(Q[i], F[j]) {
				flag = true
				break
			}
		}
		if flag {
			fmt.Print("doublecircle]")
		} else  {
			fmt.Print("circle]")
		}

		fmt.Println()
	}
	fmt.Print("dummy -> ", base, "\n")

	for i := 0; i < len(Q); i++ {
		for j:= range mapp[i] {
			fmt.Print(i, " -> ", j, " [label = \"")
			printElemOfMap(mapp[i][j])
			fmt.Print("\"]")
			fmt.Println()
		}
	}
	fmt.Print("}")
}