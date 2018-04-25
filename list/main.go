package main

import "fmt"

const MAXSIZE = 20

type List struct {
	Element [MAXSIZE]int
	length  int
}

//初始化线性表
func (l *List) InitList(d, p int) {
	l.Element[p] = d
	l.length++
}

//插入元素
//d:插入的数据
//p:插入位置
func (l *List) Insert(d int, p int) bool {
	if p < 0 || p >= MAXSIZE || l.length >= MAXSIZE {
		return false
	}
	if p < l.length {
		for k := l.length - 1; k >= p; k-- {
			l.Element[k+1] = l.Element[k]
		}
		l.Element[p] = d
		l.length++
		return true
	} else {
		l.Element[l.length] = d
		l.length++
		return true
	}
}

//删除元素
//p:删除元素的位置
func (l *List) Delete(p int) bool {
	if p < 0 || p > l.length || p >= MAXSIZE {
		return false
	}
	for ; p < l.length-1; p++ {
		l.Element[p] = l.Element[p+1]
	}
	l.Element[l.length-1] = 0
	l.length--
	return true
}
func main() {
	var l List
	i := 0
	b := 1
	//初始化一个线性表
	for i < 15 {
		l.InitList(b, i)
		i++
		b++
	}
	//插入一个元素
	l.Insert(110, 3)
	//删除一个元素
	l.Delete(5)
	fmt.Println(l)
}
