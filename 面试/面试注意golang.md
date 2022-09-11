
# 一些笔试题链接

https://jishuin.proginn.com/p/763bfbd4cfd5


# 需要仔细看看的问题

- select

- channel



# 注意点

- 如果一个结构里边有map,优先考虑,map没有初始化就使用

- map 的使用必须初始化, 大括号初始化,或者make

- 函数定义和传参数必须要一致,接收指针,或者接收变量要看清


# 特殊代码

```go
func main() {
	// slice ,map 如果没有用两个变量接收,都取的是index,
	x := []string{"a", "b", "c"}
	for _, v := range x {
		fmt.Print(v)
	}
}

/*---------------------*/
func main() {
	strs := []string{"one", "two", "three"}

	for _, s := range strs {
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Printf("%s ", s)
		}()
	}
	time.Sleep(3 * time.Second)
}
// result:
// three three three

/*---------------------*/
type Slice []int
func NewSlice() Slice {
         return make(Slice, 0)
}
func (s* Slice) Add(elem int) *Slice {
         *s = append(*s, elem)
         fmt.Print(elem)
         return s
}
func main() {  
         s := NewSlice()
         defer s.Add(1).Add(2).Add(3)
         s.Add(4)
}
// result:1243 
// 最后一条语句总是在倒数第二个位置出现,我也不知道为什么

```