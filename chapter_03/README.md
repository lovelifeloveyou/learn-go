##  内建容器

###  一、数组

- 数组定义

  ```go
  var arr1 [5]int // 声明数组
  arr2 := [3]int{1, 3, 5}  // 声明数组并赋值
  arr3 := [...]int{2, 4, 6, 8, 10} // 不输入数组长度，让编译器来计算长度
  var grid [4][5]int // 二维数组
  ```

- 数量写在类型前

- 可通过 _ 来省略变量，不仅仅是 range，任何地方都可通过 _ 来省略变量

  ```go
  sum := 0
  for _, v := range numbers {
      sum += v
  }
  ```

- 如果只要下标 i，可写成 for i := range numbers
- 数组是值类型
  - [10]int 和 [20]int 是不同类型
  - 调用 func f(arr [10]int) 会 拷贝数组
  - 在 go 语言中一般不直接使用数组（指针），使用切片

```go
package main

import "fmt"

func printArray(arr [5]int) {
	fmt.Println("Traversal i")
	for i := range arr {
		fmt.Println(arr[i])
	}
	fmt.Println("Traversal (i, val)")
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Println("Traversal (_, val)")
	for _, v := range arr {
		fmt.Println(v)
	}
}

// 值类型传递 函数中修改入参数组，函数外不会被修改
func modifyArray1(arr [5]int) {
	arr[0] = 100
	fmt.Println(arr)
}

// 值类型传递 函数中修改入参数组，函数外不会被修改
func modifyArray2(arr *[5]int) {
	(*arr)[0] = 100
	arr[1] = 101
	fmt.Println(*arr)
}

func main() {
	fmt.Printf("-------------Define arr test-----------\n")
	var arr1 [5]int
	var arr2 = [5]int{2, 4, 6, 7, 8}
	arr3 := [3]int{1, 3, 5}
	arr4 := [...]int{2, 4, 6, 7, 8}
	var arr5 [3][5]int

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)

	fmt.Printf("-------------Print arr2 test: %d-------------\n", arr4)
	printArray(arr2)

	fmt.Printf("-------------Modify arr4 test: %d-------------\n", arr4)
	modifyArray1(arr4)
	fmt.Println(arr4)
	modifyArray2(&arr4)
	fmt.Println(arr4)
}
```

输出结果：

```go
-------------Define arr test-----------
[0 0 0 0 0]
[2 4 6 7 8]
[1 3 5]
[2 4 6 7 8]
[[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]
-------------Print arr2 test: [2 4 6 7 8]-------------
Traversal i
2
4
6
7
8
Traversal (i, val)
0 2
1 4
2 6
3 7
4 8
Traversal (_, val)
2
4
6
7
8
-------------Modify arr4 test: [2 4 6 7 8]-------------
[100 4 6 7 8]
[2 4 6 7 8]
[100 101 6 7 8]
[100 101 6 7 8]
```

###  二、切片 Slice

- Slice 本身没有数据，是对底层 array 的一个 view

  ![img](https://raw.githubusercontent.com/lovelifeloveyou/somePic/master/learn-go/202308232344901.png)

- 切片的定义方法

  ```go
  // 方法一： 通过对数组的切片获取
  arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
  s := arr[2:6]
  // 方法二： 直接定义切片，底层会自动为其分配数组
  var s1 []int
  var s2 = []int{1, 2, 3}
  s3 := []int{1, 2, 3}
  // 方法三：通过make 定义 Slice
  s4 := make([]int, 8)   // len=8, cap=8
  s5 := make([]int, 10, 16) // len=10, cap=16
  ```

- Slice 的扩展

  ```go
  arr := [...]int{0,1,2,3,4,5,6,7}
  s1 := arr[2:6]
  s2 := s1[3:5]
  
  // s1的值为？
  // s2的值为？
  ```

  ![img](https://raw.githubusercontent.com/lovelifeloveyou/somePic/master/learn-go/202308232345843.png)

  - s1 的值为[2,3,4,5]，s2 的值为[5,6]
  - slice 可以向后扩展，不可以向前扩展
  - s[i] 不可以超越 len(s)，向后扩展不可以超越底层数组 cap(s)

- 向 Slice 添加元素

  ```go
  arr := [...]int{0,1,2,3,4,5,6,7}
  s1 := arr[2:6]
  s2 := s1[3:5]
  s3 := append(s2, 10)
  s4 := append(s3, 11)
  s5 := append(s4, 12)
  // s3, s4, s5的值为？arr的值为？
  fmt.Println("s3, s4, s5 =", s3, s4, s5)  // s3, s4, s5 = [5 6 10] [5 6 10 11] [5 6 10 11 12]
  fmt.Println("arr =", arr)  // arr = [0 1 2 3 4 5 6 10] 
  // s4 and s5 不再 view arr，而是新的array
  ```

  - 添加元素时如果超越 cap，系统会重新分配更大的底层数组
  - 由于值传递的关系，必须接收 append 的返回值 `s = append(s, val)`

- Slice 的完整测试代码

  ```go
  package main
  
  import "fmt"
  
  func printSlice(sliceName string, s []int) {
  	fmt.Printf("Slice %s=%v, len=%d, cap=%d\n", sliceName, s, len(s), cap(s))
  }
  
  func appendSlice() {
  	var s []int
  	for i := 0; i < 10; i++ {
  		printSlice("s", s)
  		s = append(s, i*2)
  		s = append(s, i*2+1)
  	}
  }
  
  func main() {
  	arr := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
  	// Slice本身没有数据，是对底层 array 的一个 view
  	fmt.Println("------------定义切片----------")
  	fmt.Println("arr[2:6] = ", arr[2:6])
  	fmt.Println("arr[:6] = ", arr[:6])
  	fmt.Println("arr[2:] = ", arr[2:])
  	fmt.Println("arr[:] = ", arr[:])
  
  	// 切片的切片 Re_slice(Slice the Slice)
  	fmt.Println("------------切片的切片----------")
  	printSlice("arr", arr[:])
  	s1 := arr[2:6]
  	printSlice("s1", s1)
  	s2 := s1[3:5]
  	printSlice("s2", s2)
  
  	// 切片超出 cap, 将会报错：panic: runtime error: slice bounds out of range [:7] with capacity 6
  	fmt.Println("------------切片超出 cap----------")
  	printSlice("arr", arr[:])
  	s1 = arr[2:6]
  	printSlice("s1", s1)
  	//s2 = s1[3:7]
  	//printSlice("s2", s2)
  
  	// 切片添加元素 append
  	fmt.Println("------------切片添加元素 append----------")
  	printSlice("arr", arr[:])
  	s1 = arr[2:7]
  	printSlice("s1", s1)
  	s1 = append(s1, 10)
  	printSlice("s1", s1)
  	printSlice("arr", arr[:])
  	// 超出 arr 的 cap, 系统将会定义一个新的arr来对应 s1, 旧的arr在系统没有使用的情况下会被垃圾回收
  	s1 = append(s1, 11)
  	printSlice("s1", s1)
  	printSlice("arr", arr[:])
  
  	// 使用 make 定义切片 （可以指定切片的 len 和 cap）
  	fmt.Println("------------使用 make 定义切片----------")
  	makeS1 := make([]int, 10, 32) // Slice make_s1=[0 0 0 0 0 0 0 0 0 0], len=10, cap=32
  	printSlice("make_s1", makeS1)
  
  	// 拷贝slice到另一个slice
  	fmt.Println("------------拷贝slice到另一个slice----------")
  	arr = [8]int{0, 1, 2, 3, 4, 5, 6, 7}
  	s1 = arr[2:7]
  	s2 = make([]int, 10, 16)
  	printSlice("s1", s1)
  	printSlice("s2", s2)
  	num := copy(s2, s1) // func copy(dst, src []Type) int
  	fmt.Println("拷贝 s1 到另一个 s2, 总计 copy 个数: ", num)
  	printSlice("s1", s1)
  	printSlice("s2", s2)
  
  	// 删除 slice 中的元素, 没有内建方法，需要通过 copy 覆盖掉要删除的元素
  	fmt.Println("------------ 删除 slice 中 第3个元素 ----------")
  	srcS1 := []int{0, 1, 2, 3, 4, 5, 6, 7}
  	printSlice("srcS1", srcS1)
  	srcS1 = append(srcS1[:3], srcS1[4:]...)
  	printSlice("srcS1", srcS1)
  
  	// Pop slice 中的第一元素, Pop slice 中的最后一元素
  	fmt.Println("------------ Pop slice 中的第一元素, Pop slice 中的最后一元素 ----------")
  	fmt.Println("Pop slice 中的第一元素")
  	srcS1 = []int{0, 1, 2, 3, 4, 5, 6, 7}
  	printSlice("srcS1", srcS1)
  	fmt.Println("Pop slice first element: ", srcS1[0])
  	srcS1 = srcS1[1:]
  	printSlice("srcS1", srcS1)
  
  	fmt.Println("Pop slice last element: ", srcS1[len(srcS1)-1])
  	srcS1 = srcS1[:len(srcS1)-1]
  	printSlice("srcS1", srcS1)
  
  	// 测试 slice 的自动扩容
  	fmt.Println("------------ 测试 slice 的自动扩容 ----------")
  	appendSlice()
  }
  ```

  输出结果：

  ```go
  ------------定义切片----------
  arr[2:6] =  [2 3 4 5]
  arr[:6] =  [0 1 2 3 4 5]
  arr[2:] =  [2 3 4 5 6 7]
  arr[:] =  [0 1 2 3 4 5 6 7]
  ------------切片的切片----------
  Slice arr=[0 1 2 3 4 5 6 7], len=8, cap=8
  Slice s1=[2 3 4 5], len=4, cap=6
  Slice s2=[5 6], len=2, cap=3
  ------------切片超出 cap----------
  Slice arr=[0 1 2 3 4 5 6 7], len=8, cap=8
  Slice s1=[2 3 4 5], len=4, cap=6
  ------------切片添加元素 append----------
  Slice arr=[0 1 2 3 4 5 6 7], len=8, cap=8
  Slice s1=[2 3 4 5 6], len=5, cap=6
  Slice s1=[2 3 4 5 6 10], len=6, cap=6
  Slice arr=[0 1 2 3 4 5 6 10], len=8, cap=8
  Slice s1=[2 3 4 5 6 10 11], len=7, cap=12
  Slice arr=[0 1 2 3 4 5 6 10], len=8, cap=8
  ------------使用 make 定义切片----------
  Slice make_s1=[0 0 0 0 0 0 0 0 0 0], len=10, cap=32
  ------------拷贝slice到另一个slice----------
  Slice s1=[2 3 4 5 6], len=5, cap=6
  Slice s2=[0 0 0 0 0 0 0 0 0 0], len=10, cap=16
  拷贝 s1 到另一个 s2, 总计 copy 个数:  5
  Slice s1=[2 3 4 5 6], len=5, cap=6
  Slice s2=[2 3 4 5 6 0 0 0 0 0], len=10, cap=16
  ------------ 删除 slice 中 第3个元素 ----------
  Slice srcS1=[0 1 2 3 4 5 6 7], len=8, cap=8
  Slice srcS1=[0 1 2 4 5 6 7], len=7, cap=8
  ------------ Pop slice 中的第一元素, Pop slice 中的最后一元素 ----------
  Pop slice 中的第一元素
  Slice srcS1=[0 1 2 3 4 5 6 7], len=8, cap=8
  Pop slice first element:  0
  Slice srcS1=[1 2 3 4 5 6 7], len=7, cap=7
  Pop slice last element:  7
  Slice srcS1=[1 2 3 4 5 6], len=6, cap=7
  ------------ 测试 slice 的自动扩容 ----------
  Slice s=[], len=0, cap=0
  Slice s=[0 1], len=2, cap=2
  Slice s=[0 1 2 3], len=4, cap=4
  Slice s=[0 1 2 3 4 5], len=6, cap=8
  Slice s=[0 1 2 3 4 5 6 7], len=8, cap=8
  Slice s=[0 1 2 3 4 5 6 7 8 9], len=10, cap=16
  Slice s=[0 1 2 3 4 5 6 7 8 9 10 11], len=12, cap=16
  Slice s=[0 1 2 3 4 5 6 7 8 9 10 11 12 13], len=14, cap=16
  Slice s=[0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15], len=16, cap=16
  Slice s=[0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17], len=18, cap=32
  ```

###  三、Map

1. Map 的操作

   - Map：`map[k]v, map[k1]map[k2]V（复合 map）`
   - 创建：`make(map[string]int)`
   - 获取元素：`m[key]`
   - key 不存在时，获得 value 类型的初始值（Zero value）
   - 用 `value, ok := m[key]` 来判断是否存在 key
   - 用 delete 删除一个 key

2. map 的遍历

   - 使用 range 遍历 key，或者遍历 key, value 对
   - 不保证遍历顺序，如需顺序，需手动对 key 排序
   - 使用 len 获取元素个数

3. map 中的 key

   - map 使用哈希表，必须可以比较相等
   - 除 slice，map，function 外的内建类型都可以作为 key
   - Struct 类型不包含上述字段，也可作为 key

4. Map 的基本操作示例

   ```go
   package main
   
   import "fmt"
   
   func main() {
   	// map 的定义
   	fmt.Println("------------ map 的定义 ----------")
   	m := map[string]string{
   		"course": "golang",
   		"site":   "imooc",
   	}
   	m1 := make(map[string]int)
   	var m2 map[string]int
   	fmt.Println("m=", m)
   	fmt.Println("m1=", m1)
   	fmt.Println("m2=", m2)
   
   	// 定义 map 中嵌套 map
   	mm := map[string]map[string]string{
   		"name": {
   			"first":  "Wang",
   			"second": "er"},
   		"age": {"last_year": "18"}}
   	for k, v := range mm {
   		fmt.Println(k, v)
   	}
   
   	// map 的遍历
   	fmt.Println("------------ map 的遍历 ----------")
   	for k, v := range m {
   		fmt.Println(k, v)
   	}
   	fmt.Println("只遍历 map 中的 key")
   	for k := range m { // 只遍历 map 中的 key
   		fmt.Println(k)
   	}
   	fmt.Println("只遍历 map 中的 val")
   	for _, v := range m { // 只遍历 map 中的 val
   		fmt.Println(v)
   	}
   
   	// 获取 map 中的值
   	fmt.Println("------------ 修改和删除map中的值 ----------")
   	fmt.Println("len(m) = ", len(m))
   	fmt.Println("m['course'] = ", m["course"])
   	// 如果获取的key不存在，就会返回默认值
   	fmt.Println("m['name'] = ", m["name"]) // "golang true"
   	val, ok := m["course"]                 // " false"
   	fmt.Println(val, ok)
   	val, ok = m["name"]
   	fmt.Println(val, ok)
   	// 判断 m 中是否有 key = "name" 的元素
   	if v, ok := m["name"]; ok {
   		fmt.Println("m['name'] = ", v)
   	} else {
   		fmt.Println("The element[name] dose not exist.")
   	}
   
   	// 修改和删除 map 中的值
   	fmt.Println("------------ 修改和删除map中的值 ----------")
   	delete(m, "course")
   	fmt.Println("m=", m)
   
   }
   ```

   输出结果：

   ```go
   ------------ map 的定义 ----------
   m= map[course:golang site:imooc]
   m1= map[]
   m2= map[]
   age map[last_year:18]
   name map[first:Wang second:er]
   ------------ map 的遍历 ----------
   course golang
   site imooc
   只遍历 map 中的 key
   course
   site
   只遍历 map 中的 val
   golang
   imooc
   ------------ 修改和删除map中的值 ----------
   len(m) =  2
   m['course'] =  golang
   m['name'] =  
   golang true
    false
   The element[name] dose not exist.
   ------------ 修改和删除map中的值 ----------
   m= map[site:imooc]
   ```

5. map 例题

###  四、字符和字符串处理

1. rune

   - rune 相当于 go 的 char

   - 使用 range 遍历字符串 pos 时，rune 是按照真实字符的 pos

     ```go
     s := "You弄啥了!"
     for i, ch := range s {      // ch is a rune
     	//fmt.Printf("(%d, %d, %c) ", i, ch, ch)
     	fmt.Printf("(%d, %x) ", i, ch)          // ((0, 59) (1, 6f) (2, 75) (3, 5f04) (6, 5565) (9, 4e86) (12, 21)
     }
     fmt.Println()
     for i, ch := range []rune(s) {
     	//fmt.Printf("(%d, %x, %c) ", i, ch, ch)
     	fmt.Printf("(%d, %x) ", i, ch)          // (0, 59) (1, 6f) (2, 75) (3, 5f04) (4, 5565) (5, 4e86) (6, 21)
     }
     ```

   - 使用 `utf8.RuneCountInString` 获得字符数量是正确的， len 获取字符串包含中文的字符数是不准确的

     ```go
     s := "You弄啥了!"
     fmt.Println(s, "len=", len(s))   // 输出: You弄啥了! len= 13
     fmt.Println(s, "Rune count=", utf8.RuneCountInString(s))  // 输出: You弄啥了! Rune count= 7
     ```

   - 测试代码

     ```go
     package main
     
     import (
         "fmt"
         "unicode/utf8"
     )
     
     func main() {
         s := "You弄啥了!"
         fmt.Println(s, "len=", len(s))   // 输出: You弄啥了! len= 13
         fmt.Println(s, "Rune count=", utf8.RuneCountInString(s))  // 输出: You弄啥了! Rune count= 7
         fmt.Println("--------------分割线----------------")
         fmt.Println("[]byte(s) = ", []byte(s))
         fmt.Println("[]rune(s) = ", []rune(s))
         fmt.Println()
         fmt.Println("--------------分割线----------------")
         for i, ch := range s {      // ch is a rune
             //fmt.Printf("(%d, %d, %c) ", i, ch, ch)
             fmt.Printf("(%d, %x) ", i, ch)          // ((0, 59) (1, 6f) (2, 75) (3, 5f04) (6, 5565) (9, 4e86) (12, 21)
         }
         fmt.Println()
         for i, ch := range []rune(s) {
             //fmt.Printf("(%d, %x, %c) ", i, ch, ch)
             fmt.Printf("(%d, %x) ", i, ch)          // (0, 59) (1, 6f) (2, 75) (3, 5f04) (4, 5565) (5, 4e86) (6, 21)
         }
         fmt.Println()
         fmt.Println("--------------分割线----------------")
         s = "You弄啥了!"
         myBytes := []byte(s)
         for len(myBytes) > 0 {
             ch, size := utf8.DecodeRune(myBytes)
             myBytes = myBytes[size:]
             fmt.Printf("%c ", ch)   // Y o u 弄 啥 了 !
         }
     }
     ```

     输出结果：

     ```go
     You弄啥了! len= 13
     You弄啥了! Rune count= 7
     --------------分割线----------------
     []byte(s) =  [89 111 117 229 188 132 229 149 165 228 186 134 33]
     []rune(s) =  [89 111 117 24324 21861 20102 33]
     
     --------------分割线----------------
     (0, 59) (1, 6f) (2, 75) (3, 5f04) (6, 5565) (9, 4e86) (12, 21) 
     (0, 59) (1, 6f) (2, 75) (3, 5f04) (4, 5565) (5, 4e86) (6, 21) 
     --------------分割线----------------
     Y o u 弄 啥 了 ! 
     ```

2. 字符串的内建操作

   - Fields, Split, Join
   - Contains, Index
   - ToLower, ToUpper
   - Trim, TrimRight, TrimLeft