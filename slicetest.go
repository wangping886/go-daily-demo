package main

import "fmt"

type Slice []int

func (A2 Slice) Append(value int) {
	A1 := append(A2, value)
	fmt.Println(A1)
	fmt.Printf("slice %p\n%p\n", A2, A1)
	fmt.Println("a a1 ", A2, A1)
	/*运行代码我们会发现两个Slice的Data不再一样了。

	A  Data:824633835680,Len:10,Cap:10
	A1 Data:824634204160,Len:11,Cap:20
	这是因为在append的时候，发现Cap不够，生成了一个新的Data数组，用于存储新的数据，并且同时扩充了Cap容量。*/
}

type Student struct {
	Age  int
	Name string
}
type aaa2 struct{}

func ParseStruct() map[string]*Student {
	data := make(map[string]*Student, 0)
	stus := []Student{
		{Age: 1, Name: "111"},
		{Age: 2, Name: "222"},
	}
	for k, v := range stus {
		fmt.Println(&k, &v)
		fmt.Println(k, v)
		data[v.Name] = &v
	}
	fmt.Println(data)
	return data
}
func main() {
	b := &aaa2{}
	if b == nil {
		fmt.Println("nil")
	}
	stu := ParseStruct()
	for k, v := range stu {
		fmt.Sprintf("%s%V", k, v)
		fmt.Println(&k, &v)
	}
}
