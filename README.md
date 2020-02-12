<!--
 * @Description: 文档
 * @Author: leo
 * @Date: 2020-02-12 14:11:43
 * @LastEditors  : leo
 * @LastEditTime : 2020-02-12 19:23:06
 -->
#### go 的基本数据类型

```go
  bool
  string 字符型
  int int8 int16 int32 int64 整型
  uint uint8 uint16 uint32 uint64 uintptr 无符号整型
  byte // alias for uint8 无符号8位整型
  rune // alias for int32,represents a unicode code point 字符编码
  float32 float64 浮点类型
  complex64 complex128 复数类型
```

#### go 的类型转化

```go
  GO 语言的类型转换非常严格：

  1、Go 语言不允许隐式类型转换
  2、别名和原有类型也不能进行隐式类型转换
```

#### go 指针类型

```go

  1、支持指针类型，但不支持指针运算
  2、string 是值类型， 其默认的初始化值为空字符串，而不是 nil
```

#### go 算术运算符

```go
  GO 语言没有前置的 ++, -- (++a) // 不支持
  GO 语言支持后置的 ++, -- (a++) // 支持

  数组之间如果长度不一样是不能进行比较的 ，会直接报错
```

#### go 循环

```go

  GO 语言仅支持循环关键字 for
  for i := 0; i < count; i++ {
    // 循环体
	}
  
```

#### if 条件

```go
  1、condition 表达式结果必须为 布尔值 eg // if xxx 必须为bool

  2、支持变量赋值：
    if var declaration; condition {
      // ...
    }
```

#### switch 条件

```go
  1、条件表达式 不限制 为常量或者整数

  2、单个 case 中，可以出现多个结果选项，使用逗号分隔

  3、与 C 语言等规则相反，GO 语言不需要用break来明确退出一个 case
  
  4、可以不设定 switch 之后的条件表达式，在此种情况下，整个 switch 结构与 多个 if...else... 的逻辑作用等同

```

#### 数组的声明

```go

var a [3] int // 声明并初始化为默认值
a[0] = 1

b := [3]int{1, 2, 3} // 声明同时初始化
c := [2][2]int{{1, 2}, {3, 4}} // 多维数组初始化
```

#### 数组的截取

```go
a[开始索引(包含), 结束索引(不包含)]

a := [...]int{1, 2, 3, 4, 5}
a[1:2] // 2
a[1:3] // 2, 3
a [1:len(a)] // 2, 3, 4, 5
a [1:] // 2, 3, 4, 5
a [:3] // 1, 2, 3
```

#### 切片的声明

```go
  var s0 []int

  s0 = append(s0, 1) // 添加

  s := []int{}
   
  s1 := []int{1, 2, 3}

  s2 := make([]int, 2, 4)
  /*
    []type, len, cap
    其中 len 个元素会被初始化为默认零值，未初始化元素不可以访问
  */
  切片就像数组的引用， 它不存储任何数据， 
  更改切片的元素 会修改 其底层数组中对应的元素, 而且与之共享底层数组的切片都会被修改

  切片文法类似于没有长度的数组文法：
    
    [1]bool{true}
  
  切片文法：
    []bool{false}
  
  切片的默认行为
    切片下界的默认值为 0，上界则是改切片的长度。
    
    a[0:10] == a[:10] == a[0] == a[:]
  
  用make 内置函数 创建 切片
    a := make([]int, 0, 5) // 第一个参数是类型 第二个参数是长度，第三个参数是容量
```

#### 切片 vs 数组

```go
  1、切片容量可以伸缩，数组不行
  
  2、数组可以进行比较, 切片之间不能比较 ，只能 跟 nil 比较
```


#### Map声明

```go

  m := map[string]int{"one": 1, "two": 2, "three": 3}

  m1 := map[string]int{}

  m2["one"] = 1

  m3 := make(map[string]int, 10) // 第一个参数 map 类型 第二个参数是 容量
  // 为什么不初始化len  切片len初始化的时候会默认给为 0 值，而map无法


  注意：
  map 在访问 key 不存在时，仍会返回零值，不能通过返回 nil 来判断元素是否存在!
```

#### 实现 Set

```go
Go 的内置集合中没有 Set 实现，可以 map[type]bool

1、元素的唯一性

2、基本操作
  1) 添加元素
  2) 判断元素是否存在
  3) 删除元素
  4) 元素个数
```

#### 字符串 string

```go
1、string 是数据类型，不是引用或指针类型

2、string 是只读的 byte slice, len 函数可以它所包含的 byte 数

3、string 的 byte 是数组可以 存放 任何数据
```

#### Unicode UTF8

```go
1、Unicode 是一种字符集 (code point)

2、UTF8 是 unicode 的存储实现 (转换为字节序列的规则)
```

#### 函数 一等公民

```
1、可以有多个返回值

2、所有参数都是值传递：slice，map，channel 会有传引用的错觉

3、函数可以作为变量的值

4、函数可以作为参数和返回值
```

#### 函数的可变参数

```go
  func sum (ops ...int) int {
    s := 0
    for _, op := range ops {
      s += op
    }
    return s
  }
```


#### defer 函数

```go
func TestDefer(t *testing.T) {
  defer func() {
    t.Log("Clear resources")
  }()
  t.Log("Started")
  panic("Fatal error") // defer 仍会执行
}
```

#### 结构体定义

```go
  type Employee struct {
    Id string
    Name string
    Age int
  }
```

#### 行为方法 定义

```go
与其他主要编程语言的差异

type Employee struct {
	Id string
	Name string
	Age      int
}
// 第一种定义方式 在实例对应方法被调用时，实例的成员会进行值赋值
func (e Employee) String() string {
  return fmt.Sprintf("Id:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

// 通常情况为了避免内存拷贝我们使用第二种定义方式
func (e *Employee) String() string {
  return fmt.Sprintf("Id:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}
```

#### 接口实现

```go
接口定义
type Programmer interface {
  WriteHelloWorld() Code
}

接口实现
type GoProgrammer struct {

}

func (p *GoProgrammer) WriteHelloWorld() Code {
  return "fmt.Println(\"Hello World!\")"
}
特点：
1、接口为非入侵性的，实现不依赖于接口定义

2、所以接口的定义可以包含在接口使用者包内
```

#### 接口变量

```go
  var prog Coder = &GoProgrammer{} // prog 是变量 &GoProgrammer{} 是内容
  type GoProgrammer struct {

  } // 这个是类型
```