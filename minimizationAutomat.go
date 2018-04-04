package main

import (
	"fmt"
)

func makeSet (v int, parent []int, depth []int) {
	parent[v] = v
	depth[v] = 0
}

func findSet (v int, parent_ptr *[]int) int {
	parent := *(parent_ptr)
	/*

		fmt.Print("FIND Parents :: \n")
		fmt.Print(*(parent_ptr), "\n")
		fmt.Print(parent, "\n")
	*/
	if v == parent[v] {
		//		fmt.Print("return v = ", v, "\n-----------------------\n")
		return v
	}
	/*
	fmt.Print("FIND parent : ", parent)
	fmt.Println()
	parent[v] = findSet(parent[v], &parent)
	fmt.Print("FIND new parent : ", parent)
	fmt.Println()
	fmt.Println()
	*/
	//	fmt.Print("return parent[v] = ", v ,"\n------------------------\n")
	return parent[v]
}

func UnionSet (a, b int, parent_ptr *[]int, depth_ptr *[]int) {
	parent := *(parent_ptr)
	depth := *(depth_ptr)

	/*
	fmt.Print("UNION parent : ", parent)
	fmt.Println()
	fmt.Print("UNION depth : ", depth)
	fmt.Println()
	*/
	a = findSet(a, &parent)
	b = findSet(b, &parent)
	if a != b {
		if depth[a] < depth[b] {
			a, b = b, a
		}
		parent[b] = a
		if depth[a] == depth[b] {
			depth[a] += 1
		}
	}
	/*
	fmt.Print("UNION new parent : ", parent)
	fmt.Println()
	fmt.Print("UNION new depth : ", depth)
	fmt.Println()
	fmt.Println()
	*/
}

func Split1 (q [][] int, out [][] string, nAlf int) (m int, Pi []int) {
	m = len(q)
	Pi = make ([] int, m)

	var parent = make ([] int, m)
	var depth = make ([] int, m)

	for i := range q {
		parent[i] = i
		depth[i] = 0
	}

	for i := range q {
		for j := range q {
			//fmt.Print("i , j = ", i, j, "\n")
			//fmt.Print("REAL parent\n", parent, "\n\n")
			if findSet(i, &parent) != findSet(j, &parent) {
				eq := true
				for k := 0; k < nAlf; k++ {
					if out[i][k] != out[j][k] {
						eq = false
						break
					}
				}
				if eq {
					UnionSet(i, j, &parent, &depth)
					m -= 1
				}
			}
		}
	}

	for i := range q {
		Pi[i] = findSet(i, &parent)
	}
	return
}


func Split (q [][] int, out [][] string, nAlf int, Pi_ptr *[]int) (int, []int) {
	m := len(q)

	Pi := *(Pi_ptr)
	var parent = make ([] int, m)
	var depth = make ([] int, m)

	for i := range q {
		parent[i] = i
		depth[i] = 0
	}

	for i := range q {
		for j := range q {
			//fmt.Print("i , j = ", i, j, "\n")
			//fmt.Print("REAL parent\n", parent, "\n\n")
			if Pi[i] == Pi[j] && findSet(i, &parent) != findSet(j, &parent) {
				eq := true
				for k := 0; k < nAlf; k++ {
					w1 := q[i][k]
					w2 := q[j][k]
					if Pi[w1] != Pi[w2] {
						eq = false
						break
					}
				}
				if eq {
					UnionSet(i, j, &parent, &depth)
					m -= 1
				}
			}
		}
	}
	for i := range q {
		Pi[i] = findSet(i, &parent)
	}
	return m, Pi
}


func containsSym(symbol string, Alf []string) bool {
	for i := range Alf {
		if symbol == Alf[i] {
			return true
		}
	}
	return false
}

func containsQ(q int, arr []int) bool {
	for i := range arr {
		if q == arr[i] {
			return true
		}
	}
	return false
}

func indexOf(q int, mapp map[int] []int) int {
	for i := range mapp {
		if containsQ(q, mapp[i]) {
			return i
		}
	}
	return -1
}

func print2DSlice (sl[][] int) {
	for i := range sl {
		for j := range sl[i] {
			fmt.Print(sl[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func print2DString (sl[][] string) {
	for i := range sl {
		for j := range sl[i] {
			fmt.Print(sl[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}


func canonize (q [][]int,  out [][]string, nCond int, nAlf int, start int) {
	preorder := make([]int, 0, nCond)
	mark := make([]bool, nCond)
	Stack := make([]int, 0, 50)

	ptr := start

	mapp := make(map[int]int, nCond)
	k := 0
	flag := false
	for len(preorder) != nCond {
		if !mark[ptr] {
			mark[ptr] = true
			preorder = append(preorder, ptr)
			mapp[ptr] = k
			k++
			//Stack = append(Stack, ptr)
			for i := len(q[ptr]) - 1; i >=0; i-- {
				if !mark[q[ptr][i]] {
					Stack = append(Stack, q[ptr][i])
				}
			}
			ptr = q[ptr][0]
		} else {
			flag = false
			for i := len(Stack) - 1; i >=0; i-- {
				if !mark[Stack[i]] {
					ptr = Stack[i]
					flag = true
					break
				}
			}
			if !flag {
				break
			}
		}
	}

	size := len(preorder)
	canon := make([][]int, size)
	for i := range canon {
		canon[i] = make([]int, nAlf)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < nAlf; j++ {
			canon[i][j] = mapp[q[preorder[i]][j]]
		}
	}
	/*
	start = 0
	fmt.Println(size)
	fmt.Println(nAlf)
	fmt.Println(start)

	for i := 0; i < size; i++ {
		for k := 0; k < nAlf; k++ {
			fmt.Print(canon[i][k])
			fmt.Print(" ")
		}
		fmt.Println()
	}

	for i := 0; i < size; i++ {
		for k := 0; k < nAlf; k++ {
			fmt.Print(out[preorder[i]][k])
			fmt.Print(" ")
		}
		fmt.Println()
	}
	*/
	alf := make([]string, 26)
	for i, j := 0, 97; i < 26; i++ {
		alf[i] = string(j)
		j++
	}
	fmt.Print("digraph {\n")
	fmt.Print("rankdir = LR\n")
	fmt.Print("dummy [label = \"\", shape = none]\n")
	for i := range canon {
		fmt.Print(i, " [shape = circle]\n")
	}
	fmt.Print("dummy -> 0\n")
	for i := 0; i < len(canon); i++ {
		for j := 0; j < nAlf; j++ {
			fmt.Print(i, " -> ", canon[i][j])
			fmt.Print(" [label = \"", alf[j], "(",out[preorder[i]][j],")","\"]")
			fmt.Println()
		}
	}
	fmt.Print("}")

}


func main() {
	var nCond int
	fmt.Scan(&nCond)

	var nAlf int
	fmt.Scan(&nAlf)

	var base int
	fmt.Scan(&base)

	q := make ([][] int, nCond)
	out := make ([][] string, nCond)
	for i := range q {
		q[i] = make ([] int, nAlf)
		out[i] = make ([] string, nAlf)
	}

	for i := 0; i < nCond; i++ {
		for j := range q[i] {
			fmt.Scan(&q[i][j])
		}
	}

	for i := 0; i < nCond; i++ {
		for j := range out[i] {
			fmt.Scan(&out[i][j])
		}
	}

	m, Pi := Split1(q, out, nAlf)
	mk := -1
	for ; ; {
		mk, Pi = Split(q, out, nAlf, &Pi)
		if m == mk {
			break
		}
		m = mk
	}

	/*
	fmt.Print("\n", mk, "\n")
	fmt.Print(Pi)
	*/

	base = Pi[base]

	Q := make ([]int, 0)

	qNew := make ([][] int, nCond)
	outNew := make ([][] string, nCond)
	for i := range qNew {
		qNew[i] = make ([] int, nAlf)
		outNew[i] = make ([] string, nAlf)
	}

	for i := range q {
		qTemp := Pi[i]
		//fmt.Print(Q, "\n")
		if !containsQ(qTemp, Q) {
			Q = append(Q, qTemp)
			for k := 0; k < nAlf; k++ {
				//fmt.Print("\nqTemp, x: ", qTemp, ",", out[i][k], " -> ", "Pi[q.i, x] ", Pi[q[i][k]], ",", out[i][k], "\n")
				qNew[qTemp][k] = Pi[q[i][k]]

				//fmt.Print("qTemp, x: ", qTemp, ",", out[i][k], " -> ", "out[q.i, x] ", out[i][k], ", ", out[i][k], "\n")

				outNew[qTemp][k] = out[q[i][k]][k]
			}
			//	fmt.Print("------------------------")
		}
	}


	/*
	fmt.Println()
	print2DSlice(qNew)
	fmt.Println()
	*/
	canonize(qNew, out, nCond, nAlf, base)


}