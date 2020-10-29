# 7天掌握go语言助你拿高薪

- [语言介绍](#语言介绍)
  * [语言的起源](#语言的起源)
  * [语言的特性](#语言的特性)
  * [Hello World](#hello-world)
- [语言基础](#语言基础)
  * [常量与变量](#常量与变量)
  * [整型和浮点型](#整型与浮点型)
  * [字符串](#字符串)
  * [字符型](#字符型)
  * [数组和切片](#数组和切片)
  * [包](#包)
- [语言顺序编程](#语言顺序编程)
  * [流程控制](#控制流程)
  * [函数](#函数)
  * [类型转换](#类型转换)
  * [类型断言](#类型断言)
  * [error](#error)
  * [defer和panic](#defer和panic)
- [面向对象编程](#面向对象编程)
  * [自定义类型和结构体](#自定义类型和结构体)
  * [方法](#方法)
  * [组合](#组合)
  * [接口](#接口)
- [并发编程](#并发编程)
  * [并行与并发](#并行与并发)
  * [协程](#协程)
  * [groutine](#groutine)
  * [channel](#channel)
  * [select](#select)

## 语言介绍
### 语言的起源

Go（又称Golang）是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。

罗伯特·格瑞史莫，罗勃·派克（Rob Pike）及肯·汤普逊于2007年9月开始设计Go，稍后Ian Lance Taylor、Russ Cox加入项目。Go是基于Inferno操作系统所开发的。Go于2009年11月正式宣布推出，成为开放源代码项目，支持Linux、macOS、Windows等操作系统。在2016年，Go被软件评价公司TIOBE 选为“TIOBE 2016 年最佳语言”。

Go的语法接近C语言，但对于变量的声明有所不同。Go支持垃圾回收功能。Go的并行计算模型是以东尼·霍尔的通信顺序进程（CSP）为基础，采取类似模型的其他语言包括Occam和Limbo，Go也具有这个模型的特征，比如通道传输。通过goroutine和通道等并行构造可以建造线程池和管道等[8]。在1.8版本中开放插件（Plugin）的支持，这意味着现在能从Go中动态加载部分函数。

与C++相比，Go并不包括如枚举、异常处理、继承、泛型、断言、虚函数等功能，但增加了 切片(Slice) 型、并发、管道、垃圾回收功能、接口等特性的语言级支持。Go 2.0版本将支持泛型，对于断言的存在，则持负面态度，同时也为自己不提供类型继承来辩护。

不同于Java，Go原生提供了关联数组（也称为哈希表（Hashes）或字典（Dictionaries）），就像字符串类型一样。

### 语言的特性

Go有定义如下的撰写风格

- 每行程序结束后不需要撰写分号（;）。
- 大括号（{）不能够换行放置。
- if判断式和for循环不需要以小括号包覆起来。
- 使用 tab 做排版
- 除了第二点外（换行会产生编译错误），在不符合上述规定时，仍旧可以编译，但使用了内置gofmt工具后，会自动整理代码，使之符合规定的撰写风格。

### Hello World

```golang
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```
## 语言基础

在这里我想先和大家普及一个观念就是，编程的本质可以理解为就是在计算，否则为什么我们的电脑又叫做计算机的原因。那它到底又是在计算什么呢？这也就是接下来要和大家一起探讨的问题。

### 常量与变量

在声明完之后可以被改变的量，我们称之为常量；不能被改变的量，我们称之为变量。

给大家举两个例子：

常量：

```golang
const str1 string = "hello world"
```

先解释一下上面那句话是什么意思吧，我们生命了一个名为str1的string类型的常量，并非它赋值为"hello world"，这里需要注意的是常量声明完之后就必须要附上初始值，否则在编译的时候程序就会出错。

常量的关键字是**const**，只要加上这个关键字，str1就为常量了。

变量

```golang

var str2 string = "hello world"

str2 := "hello world"

var str2 string
str2 = "hello world"

```
在上面的代码中，我用三种形式声明了变量，并赋上了值。

变量有三种声明方式：
1. 第一种与声明常量时一样，显示声明了变量的数据类型为**string**，然后赋值上"hello world"；
2. 第二种使用**:=**的方式声明变量，通过这种方法golang的编译器可以自动推断出变量的数据类型;
3. 第三种其实是第一种方式的分步法，写上这种方式是为了突出变量和常量的区别，**常量声明完之后就必须要附上初始值**。

在例子中由于篇幅有限，我只使用了string类型声明常量，读者可以自己做实验，来试验一下变得数据类型。

### 整型与浮点型

Go 语言中，整型可以分为int和uint两种类型。

int 和 uint 的区别就在于一个 u，有 u 说明是无符号，没有 u 代表有符号。

解释这个符号的区别
以 int8 和 uint8 举例，8 代表 8个bit，能表示的数值个数有 2^8 = 256。
uint8 是无符号，能表示的都是正数，0-255，刚好256个数。
int8 是有符号，既可以正数，也可以负数，那怎么办？对半分呗，-128-127，也刚好 256个数。
int8 int16 int32 int64 这几个类型的最后都有一个数值，这表明了它们能表示的数值个数是固定的。
而 int 没有并没有指定它的位数，说明它的大小，是可以变化的，那根据什么变化呢？

当你在32位的系统下，int 和 uint 都占用 4个字节，也就是32位。
若你在64位的系统下，int 和 uint 都占用 8个字节，也就是64位。

出于这个原因，在某些场景下，你应当避免使用 int 和 uint ，而使用更加精确的 int32 和 int64，比如在二进制传输、读写文件的结构描述（为了保持文件的结构不会受到不同编译目标平台字节长度的影响）
不同进制的表示方法
出于习惯，在初始化数据类型为整型的变量时，我们会使用10进制的表示法，因为它最直观，比如这样，表示整数10.
```go
var num int = 10
```
不过，你要清楚，你一样可以使用其他进制来表示一个整数，这里以比较常用的2进制、8进制和16进制举例。
2进制：以0b或0B为前缀
```
var num01 int = 0b1100
```
8进制：以0o或者 0O为前缀
```
var num02 int = 0o14
```
16进制：以0x 为前缀
```
var num03 int = 0xC
```
下面用一段代码分别使用二进制、8进制、16进制来表示 10 进制的数值：12
```go
package main

import (
	"fmt"
)

func main() {
	var num01 int = 0b1100
	var num02 int = 0o14
	var num03 int = 0xC
    
	fmt.Printf("2进制数 %b 表示的是: %d \n", num01, num01)
	fmt.Printf("8进制数 %o 表示的是: %d \n", num02, num02)
	fmt.Printf("16进制数 %X 表示的是: %d \n", num03, num03)
}
```
输出如下
```
2进制数 1100 表示的是: 12 
8进制数 14 表示的是: 12 
16进制数 C 表示的是: 12 
```


以上代码用过了 fmt 包的格式化功能，你可以参考这里去看上面的代码
```
%b    表示为二进制
%c    该值对应的unicode码值
%d    表示为十进制
%o    表示为八进制
%q    该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
%x    表示为十六进制，使用a-f
%X    表示为十六进制，使用A-F
%U    表示为Unicode格式：U+1234，等价于"U+%04X"
%E    用科学计数法表示
%f    用浮点数表示
```

浮点数类型的值一般由整数部分、小数点“.”和小数部分组成。

其中，整数部分和小数部分均由10进制表示法表示。不过还有另一种表示方法。那就是在其中加入指数部分。指数部分由“E”或“e”以及一个带正负号的10进制数组成。比如，3.7E-2表示浮点数0.037。又比如，3.7E+1表示浮点数37。

有时候，浮点数类型值的表示也可以被简化。比如，37.0可以被简化为37。又比如，0.037可以被简化为.037。

有一点需要注意，在Go语言里，浮点数的相关部分只能由10进制表示法表示，而不能由8进制表示法或16进制表示法表示。比如，03.7表示的一定是浮点数3.7。

#### float32 和 float64

Go语言中提供了两种精度的浮点数 float32 和 float64。

**float32**，也即我们常说的单精度，存储占用4个字节，也即4*8=32位，其中1位用来符号，8位用来指数，剩下的23位表示尾数

**float64**，也即我们熟悉的双精度，存储占用8个字节，也即8*8=64位，其中1位用来符号，11位用来指数，剩下的52位表示尾数

**那么精度是什么意思？有效位有多少位？**

精度主要取决于尾数部分的位数。

对于 float32（单精度）来说，表示尾数的为23位，除去全部为0的情况以外，最小为2^-23，约等于1.19*10^-7，所以float小数部分只能精确到后面6位，加上小数点前的一位，即有效数字为7位。

同理 float64（单精度）的尾数部分为 52位，最小为2^-52，约为2.22*10^-16，所以精确到小数点后15位，加上小数点前的一位，有效位数为16位。

### 布尔类型
布尔型的值只可以是常量 true 或者 false。
一个简单的例子：
```go
var b bool = true
```


### 字符串

一个字符串是一个不可改变的字节序列，字符串可以包含任意的数据，但是通常是用来包含可读的文本，字符串是 UTF-8 字符的一个序列（当字符为 ASCII 码表上的字符时则占用 1 个字节，其它字符根据需要占用 2-4 个字节）。

UTF-8 是一种被广泛使用的编码格式，是文本文件的标准编码，其中包括 XML 和 JSON 在内也都使用该编码。由于该编码对占用字节长度的不定性，在Go语言中字符串也可能根据需要占用 1 至 4 个字节，这与其它编程语言如 C++、Java 或者 Python 不同（Java 始终使用 2 个字节）。Go语言这样做不仅减少了内存和硬盘空间占用，同时也不用像其它语言那样需要对使用 UTF-8 字符集的文本进行编码和解码。

字符串是一种值类型，且值不可变，即创建某个文本后将无法再次修改这个文本的内容，更深入地讲，字符串是字节的定长数组。

以下就是定义字符串的方法

```golang
package main

import (
    "fmt"
)

func main() {
    var str = "C语言中文网\nGo语言教程"
    fmt.Println(str)
}
```

### 字符型

字符串中的每一个元素叫做“字符”，在遍历或者单个获取字符串元素时可以获得字符。

Go语言的字符有以下两种：
* 一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
* 另一种是 rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型等价于 int32 类型。

byte 类型是 uint8 的别名，对于只占用 1 个字节的传统 ASCII 编码的字符来说，完全没有问题，例如 var ch byte = 'A'，字符使用单引号括起来。

在 ASCII 码表中，A 的值是 65，使用 16 进制表示则为 41，所以下面的写法是等效的：
```
var ch byte = 65 或 var ch byte = '\x41'      //（\x 总是紧跟着长度为 2 的 16 进制数）
```

另外一种可能的写法是\后面紧跟着长度为 3 的八进制数，例如 \377。

Go语言同样支持 Unicode（UTF-8），因此字符同样称为 Unicode 代码点或者 runes，并在内存中使用 int 来表示。在文档中，一般使用格式 U+hhhh 来表示，其中 h 表示一个 16 进制数。

在书写 Unicode 字符时，需要在 16 进制数之前加上前缀\u或者\U。因为 Unicode 至少占用 2 个字节，所以我们使用 int16 或者 int 类型来表示。如果需要使用到 4 字节，则使用\u前缀，如果需要使用到 8 个字节，则使用\U前缀。

```golang
var ch int = '\u0041'
var ch2 int = '\u03B2'
var ch3 int = '\U00101234'
fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point
```

格式化说明符**%c**用于表示字符，当和字符配合使用时，**%v**或**%d**会输出用于表示该字符的整数，**%U**输出格式为 U+hhhh 的字符串。

Unicode 包中内置了一些用于测试字符的函数，这些函数的返回值都是一个布尔值，如下所示（其中 ch 代表字符）：
* 判断是否为字母：unicode.IsLetter(ch)
* 判断是否为数字：unicode.IsDigit(ch)
* 判断是否为空白符号：unicode.IsSpace(ch)

### 数组和切片

golang中的**数组**是一种由固定长度和固定对象类型所组成的数据类型。例如下面：
```golang
var a [4]int
```
a是一个拥有4个int类型元素的数组。当a一旦被声明之后，元素个数就被固定了下来，在a这个变量的生命周期之内，元素个数不会发生变化。而此时a的类型就是[4]int，如果同时存在一个b变量，为[5]int。即便两个变量仅仅相差一个元素，那么在内存中也占据着完全不同的地址分配单元，a和b就是两个完全不同的数据类型。在golang中，数组一旦被定义了，那么其内部的元素就完成了初始化。也就是时候a[0]等于0。
在golang当中，一个数组就是一个数据实体对象。在golang当使用a时，就代表再使用a这个数组。而在C中，当使用a时，则代表着这个数组第一个元素的指针。

数组声明时候方括号内需要写明数组的长度或者使用(...)符号自动计算长度，切片不需要指定数组的长度。比较规范的声明方式是使用make，大致有两种方式
* 只指定长度，这个时候切片长度和容量相同；
* 同时指定切片的长度和容量。

```golang
var s1 = make([]int, 5)
var s2 = make([]int, 5, 10)
```

由于切片是 **引用类型** ，因此当引用改变其中元素的值时候，其他的所有引用都会改变该值。例如
```
var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
s1 = a[:4]
s2 = a[3:7]
fmt.Println(s1)
fmt.Println(s2)
s1[3] = 100
fmt.Println(s1)
fmt.Println(s2)
```
结果是：

[1 2 3 4]

[4 5 6 7]

[1 2 3 100]

[100 5 6 7]

### 包
Go语言是使用包来组织源代码的，包（package）是多个 Go 源码的集合，是一种高级的代码复用方案。Go语言中为我们提供了很多内置包，如 fmt、os、io 等。

任何源代码文件必须属于某个包，同时源码文件的第一行有效代码必须是package pacakgeName 语句，通过该语句声明自己所在的包。

#### 包的基本概念
Go语言的包借助了目录树的组织形式，一般包的名称就是其源文件所在目录的名称，虽然Go语言没有强制要求包名必须和其所在的目录名同名，但还是建议包名和所在目录同名，这样结构更清晰。

包可以定义在很深的目录中，包名的定义是不包括目录路径的，但是包在引用时一般使用全路径引用。比如在GOPATH/src/a/b/ 下定义一个包 c。在包 c 的源码中只需声明为package c，而不是声明为package a/b/c，但是在导入 c 包时，需要带上路径，例如import "a/b/c"。

包的习惯用法：
* 包名一般是小写的，使用一个简短且有意义的名称。
* 包名一般要和所在的目录同名，也可以不同，包名中不能包含- 等特殊符号。
* 包一般使用域名作为目录名称，这样能保证包名的唯一性，比如 GitHub 项目的包一般会放到GOPATH/src/github.com/userName/projectName 目录下。
* 包名为 main 的包为应用程序的入口包，编译不包含 main 包的源码文件时不会得到可执行文件。
* 一个文件夹下的所有源码文件只能属于同一个包，同样属于同一个包的源码文件不能放在多个文件夹下。

#### 包的导入
要在代码中引用其他包的内容，需要使用 import 关键字导入使用的包。具体语法如下：
```golang
import "包的路径"
```

注意事项：
import 导入语句通常放在源码文件开头包声明语句的下面；
导入的包名需要使用双引号包裹起来；
包名是从GOPATH/src/ 后开始计算的，使用/ 进行路径分隔。

## 语言顺序编程
### 流程控制

流程控制在编程语言中是最伟大的发明了，因为有了它，你可以通过很简单的流程描述来表达很复杂的逻辑。Go中流程控制分三大类：条件判断，循环控制和无条件跳转。

### if

`if`也许是各种编程语言中最常见的了，它的语法概括起来就是：如果满足条件就做某事，否则做另一件事。

Go里面`if`条件判断语句中不需要括号，如下代码所示

```Go
if x > 10 {
	fmt.Println("x is greater than 10")
} else {
	fmt.Println("x is less than 10")
}
```

Go的`if`还有一个强大的地方就是条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了，如下所示

```Go
// 计算获取值x,然后根据x返回的大小，判断是否大于10。
if x := computedValue(); x > 10 {
	fmt.Println("x is greater than 10")
} else {
	fmt.Println("x is less than 10")
}

//这个地方如果这样调用就编译出错了，因为x是条件里面的变量
fmt.Println(x)
```

多个条件的时候如下所示：

```Go
if integer == 3 {
	fmt.Println("The integer is equal to 3")
} else if integer < 3 {
	fmt.Println("The integer is less than 3")
} else {
	fmt.Println("The integer is greater than 3")
}
```

### goto

Go有`goto`语句——请明智地使用它。用`goto`跳转到必须在当前函数内定义的标签。例如假设这样一个循环：

```Go
func myFunc() {
	i := 0
Here:   //这行的第一个词，以冒号结束作为标签
	println(i)
	i++
	goto Here   //跳转到Here去
}
```

>标签名是大小写敏感的。

### for

Go里面最强大的一个控制逻辑就是`for`，它既可以用来循环读取数据，又可以当作`while`来控制逻辑，还能迭代操作。它的语法如下：

```Go
for expression1; expression2; expression3 {
	//...
}
```

`expression1`、`expression2`和`expression3`都是表达式，其中`expression1`和`expression3`是变量声明或者函数调用返回值之类的，`expression2`是用来条件判断，`expression1`在循环开始之前调用，`expression3`在每轮循环结束之时调用。

一个例子比上面讲那么多更有用，那么我们看看下面的例子吧：

```Go
package main

import "fmt"

func main(){
	sum := 0;
	for index:=0; index < 10 ; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)
}
// 输出：sum is equal to 45
```

有些时候需要进行多个赋值操作，由于Go里面没有`,`操作符，那么可以使用平行赋值`i, j = i+1, j-1`


有些时候如果我们忽略`expression1`和`expression3`：

```Go
sum := 1
for ; sum < 1000;  {
	sum += sum
}
```

其中`;`也可以省略，那么就变成如下的代码了，是不是似曾相识？对，这就是`while`的功能。

```Go
sum := 1
for sum < 1000 {
	sum += sum
}
```

在循环里面有两个关键操作`break`和`continue`	,`break`操作是跳出当前循环，`continue`是跳过本次循环。当嵌套过深的时候，`break`可以配合标签使用，即跳转至标签所指定的位置，详细参考如下例子：

```Go
for index := 10; index>0; index-- {
	if index == 5{
		break // 或者continue
	}
	fmt.Println(index)
}
// break打印出来10、9、8、7、6
// continue打印出来10、9、8、7、6、4、3、2、1
```

`break`和`continue`还可以跟着标号，用来跳到多重循环中的外层循环

`for`配合`range`可以用于读取`slice`和`map`的数据：

```Go
for k,v:=range map {
	fmt.Println("map's key:",k)
	fmt.Println("map's val:",v)
}
```

由于 Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错, 在这种情况下, 可以使用`_`来丢弃不需要的返回值
例如

```Go
for _, v := range map{
	fmt.Println("map's val:", v)
}

```

### switch

有些时候你需要写很多的`if-else`来实现一些逻辑处理，这个时候代码看上去就很丑很冗长，而且也不易于以后的维护，这个时候`switch`就能很好的解决这个问题。它的语法如下

```Go
switch sExpr {
case expr1:
	some instructions
case expr2:
	some other instructions
case expr3:
	some other instructions
default:
	other code
}
```

`sExpr`和`expr1`、`expr2`、`expr3`的类型必须一致。Go的`switch`非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；而如果`switch`没有表达式，它会匹配`true`。

```Go
i := 10
switch i {
case 1:
	fmt.Println("i is equal to 1")
case 2, 3, 4:
	fmt.Println("i is equal to 2, 3 or 4")
case 10:
	fmt.Println("i is equal to 10")
default:
	fmt.Println("All I know is that i is an integer")
}
```

在第5行中，我们把很多值聚合在了一个`case`里面，同时，Go里面`switch`默认相当于每个`case`最后带有`break`，匹配成功后不会自动向下执行其他case，而是跳出整个`switch`, 但是可以使用`fallthrough`强制执行后面的case代码。

```Go
integer := 6
switch integer {
case 4:
	fmt.Println("The integer was <= 4")
	fallthrough
case 5:
	fmt.Println("The integer was <= 5")
	fallthrough
case 6:
	fmt.Println("The integer was <= 6")
	fallthrough
case 7:
	fmt.Println("The integer was <= 7")
	fallthrough
case 8:
	fmt.Println("The integer was <= 8")
	fallthrough
default:
	fmt.Println("default case")
}
```

上面的程序将输出

```Go
The integer was <= 6
The integer was <= 7
The integer was <= 8
default case

```

### 函数

函数是Go里面的核心设计，它通过关键字`func`来声明，它的格式如下：

```Go
func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
	//这里是处理逻辑代码
	//返回多个值
	return value1, value2
}
```

上面的代码我们看出

- 关键字`func`用来声明一个函数`funcName`
- 函数可以有一个或者多个参数，每个参数后面带有类型，通过`,`分隔
- 函数可以返回多个值
- 上面返回值声明了两个变量`output1`和`output2`，如果你不想声明也可以，直接就两个类型
- 如果只有一个返回值且不声明返回值变量，那么你可以省略 包括返回值 的括号
- 如果没有返回值，那么就直接省略最后的返回信息
- 如果有返回值， 那么必须在函数的外层添加return语句

下面我们来看一个实际应用函数的例子（用来计算Max值）

```Go
package main

import "fmt"

// 返回a、b中最大值.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	x := 3
	y := 4
	z := 5

	max_xy := max(x, y) //调用函数max(x, y)
	max_xz := max(x, z) //调用函数max(x, z)

	fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
	fmt.Printf("max(%d, %d) = %d\n", x, z, max_xz)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y,z)) // 也可在这直接调用它
}
```

上面这个里面我们可以看到`max`函数有两个参数，它们的类型都是`int`，那么第一个变量的类型可以省略（即 a,b int,而非 a int, b int)，默认为离它最近的类型，同理多于2个同类型的变量或者返回值。同时我们注意到它的返回值就是一个类型，这个就是省略写法。

### 多个返回值

Go语言比C更先进的特性，其中一点就是函数能够返回多个值。

我们直接上代码看例子

```Go
package main

import "fmt"

//返回 A+B 和 A*B
func SumAndProduct(A, B int) (int, int) {
	return A+B, A*B
}

func main() {
	x := 3
	y := 4

	xPLUSy, xTIMESy := SumAndProduct(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
}
```

上面的例子我们可以看到直接返回了两个参数，当然我们也可以命名返回参数的变量，这个例子里面只是用了两个类型，我们也可以改成如下这样的定义，然后返回的时候不用带上变量名，因为直接在函数里面初始化了。但如果你的函数是导出的(首字母大写)，官方建议：最好命名返回值，因为不命名返回值，虽然使得代码更加简洁了，但是会造成生成的文档可读性差。

```Go
func SumAndProduct(A, B int) (add int, Multiplied int) {
	add = A+B
	Multiplied = A*B
	return
}
```

### 变参

Go函数支持变参。接受变参的函数是有着不定数量的参数的。为了做到这点，首先需要定义函数使其接受变参：

```Go
func myfunc(arg ...int) {}
```

`arg ...int`告诉Go这个函数接受不定数量的参数。注意，这些参数的类型全部是`int`。在函数体中，变量`arg`是一个`int`的`slice`：

```Go
for _, n := range arg {
	fmt.Printf("And the number is: %d\n", n)
}
```

### 传值与传指针

当我们传一个参数值到被调用函数里面时，实际上是传了这个值的一份copy，当在被调用函数中修改参数值的时候，调用函数中相应实参不会发生任何变化，因为数值变化只作用在copy上。

为了验证我们上面的说法，我们来看一个例子

```Go
package main

import "fmt"

//简单的一个函数，实现了参数+1的操作
func add1(a int) int {
	a = a+1 // 我们改变了a的值
	return a //返回一个新值
}

func main() {
	x := 3

	fmt.Println("x = ", x)  // 应该输出 "x = 3"

	x1 := add1(x)  //调用add1(x)

	fmt.Println("x+1 = ", x1) // 应该输出"x+1 = 4"
	fmt.Println("x = ", x)    // 应该输出"x = 3"
}
```

看到了吗？虽然我们调用了`add1`函数，并且在`add1`中执行`a = a+1`操作，但是上面例子中`x`变量的值没有发生变化

理由很简单：因为当我们调用`add1`的时候，`add1`接收的参数其实是`x`的copy，而不是`x`本身。

那你也许会问了，如果真的需要传这个`x`本身,该怎么办呢？

这就牵扯到了所谓的指针。我们知道，变量在内存中是存放于一定地址上的，修改变量实际是修改变量地址处的内存。只有`add1`函数知道`x`变量所在的地址，才能修改`x`变量的值。所以我们需要将`x`所在地址`&x`传入函数，并将函数的参数的类型由`int`改为`*int`，即改为指针类型，才能在函数中修改`x`变量的值。此时参数仍然是按copy传递的，只是copy的是一个指针。请看下面的例子

```Go
package main

import "fmt"

//简单的一个函数，实现了参数+1的操作
func add1(a *int) int { // 请注意，
	*a = *a+1 // 修改了a的值
	return *a // 返回新值
}

func main() {
	x := 3

	fmt.Println("x = ", x)  // 应该输出 "x = 3"

	x1 := add1(&x)  // 调用 add1(&x) 传x的地址

	fmt.Println("x+1 = ", x1) // 应该输出 "x+1 = 4"
	fmt.Println("x = ", x)    // 应该输出 "x = 4"
}
```

这样，我们就达到了修改`x`的目的。那么到底传指针有什么好处呢？

- 传指针使得多个函数能操作同一个对象。
- 传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。如果用参数值传递的话, 在每次copy上面就会花费相对较多的系统开销（内存和时间）。所以当你要传递大的结构体的时候，用指针是一个明智的选择。
- Go语言中`channel`，`slice`，`map`这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。（注：若函数需改变`slice`的长度，则仍需要取地址传递指针）

### 类型转换

由于Go语言不允许隐式类型转换。而类型转换和类型断言的本质，就是把一个类型转换到另一个类型。

语法：<结果类型> := <目标类型> ( <表达式> )

类型转换是用来在不同但相互兼容的类型之间的相互转换的方式，所以，当类型不兼容的时候，是无法转换的。如下：

```go
package main

import "fmt"

func main() {
    var var1 int = 7
    fmt.Printf("%T->%v\n", var1, var1)
    var2 := float32(var1)
    var3 := int64(var1)
    //var4 := []int8(var1)
    //var5 := []string(var1)
    fmt.Printf("%T->%v\n", var2, var2)
    fmt.Printf("%T->%v\n", var3, var3)
    //fmt.Printf("%T->%d", var4, var4)
    //fmt.Printf("%T->%d", var5, var5)
}
```

其中，var4和var5处运行会报错。因为类型不兼容。注释后，输出如下：

```
int->7
float32->7
int64->7
```

值得注意的是，如果某些类型可能引起误会，应该用括号括起来转换，如下：

```go
package main

import "fmt"

func main() {
    //创建一个int变量，并获得它的指针
    var1 := new(int32)
    fmt.Printf("%T->%v\n", var1, var1)
    var2 := *int32(var1)
    fmt.Printf("%T->%v\n", var2, var2)
}
```

*int32(var1)相当于*(int32(var1))，一个指针，当然不能直接转换成一个int32类型，所以该表达式直接编译错误。将该表达式改为 (*int32)(var1)就可以正常输出了。

### 类型断言

语法：

　　<目标类型的值>，<布尔参数> := <表达式>.( 目标类型 ) // 安全类型断言

　　<目标类型的值> := <表达式>.( 目标类型 )　　//非安全类型断言

类型断言的本质，跟类型转换类似，都是类型之间进行转换，不同之处在于，类型断言实在接口之间进行，相当于Java中，对于一个对象，把一种接口的引用转换成另一种。

我们先来看一个最简单的错误的类型断言：

```go
func test6() {
    var i interface{} = "kk"
    j := i.(int)
    fmt.Printf("%T->%d\n", j, j)
}
```

var i interface{} = "KK" 某种程度上相当于java中的，Object i = "KK";

现在把这个 i 转换成 int 类型，系统内部检测到这种不匹配，就会调用内置的panic()函数，抛出一个异常。

改一下，把 i 的定义改为：var i interface{} = 99，就没问题了。输出为：

```
int->99
```

以上是不安全的类型断言。我们来看一下安全的类型断言：

```go
package main

import "fmt"

func main() {
    var i interface{} = "TT"
    j, b := i.(int)
    if b {
        fmt.Printf("%T->%d\n", j, j)
    } else {
        fmt.Println("类型不匹配")
    }
}
```

输出“类型不匹配”。

### error

Go 语言通过内置的错误接口提供了非常简单的错误处理机制。

error类型是一个接口类型，这是它的定义：

```go
type error interface {
    Error() string
}
```

我们可以在编码中通过实现 error 接口类型来生成错误信息。

函数通常在最后的返回值中返回错误信息。使用errors.New 可返回一个错误信息：

```go
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // 实现
}
```

在下面的例子中，我们在调用Sqrt的时候传递的一个负数，然后就得到了non-nil的error对象，将此对象与nil比较，结果为true，所以fmt.Println(fmt包在处理error时会调用Error方法)被调用，以输出错误，请看下面调用的示例代码：

```go
result, err:= Sqrt(-1)

if err != nil {
   fmt.Println(err)
}
```

实例代码

```go
package main

import (
    "fmt"
)

// 定义一个 DivideError 结构
type DivideError struct {
    dividee int
    divider int
}

// 实现 `error` 接口
func (de *DivideError) Error() string {
    strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
    return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
    if varDivider == 0 {
            dData := DivideError{
                    dividee: varDividee,
                    divider: varDivider,
            }
            errorMsg = dData.Error()
            return
    } else {
            return varDividee / varDivider, ""
    }

}

func main() {

    // 正常情况
    if result, errorMsg := Divide(100, 10); errorMsg == "" {
            fmt.Println("100/10 = ", result)
    }
    // 当除数为零的时候会返回错误信息
    if _, errorMsg := Divide(100, 0); errorMsg != "" {
            fmt.Println("errorMsg is: ", errorMsg)
    }

}
```





### defer

Go语言中有种不错的设计，即延迟（defer）语句，你可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题。如下代码所示，我们一般写打开一个资源是这样操作的：

```go
func ReadWrite() bool {
	file.Open("file")
// 做一些工作
	if failureX {
		file.Close()
		return false
	}

	if failureY {
		file.Close()
		return false
	}

	file.Close()
	return true
}
```

我们看到上面有很多重复的代码，Go的`defer`有效解决了这个问题。使用它后，不但代码量减少了很多，而且程序变得更优雅。在`defer`后指定的函数会在函数退出前调用。

```go
func ReadWrite() bool {
	file.Open("file")
	defer file.Close()
	if failureX {
		return false
	}
	if failureY {
		return false
	}
	return true
}
```

如果有很多调用`defer`，那么`defer`是采用后进先出模式，所以如下代码会输出`4 3 2 1 0`

```Go
for i := 0; i < 5; i++ {
	defer fmt.Printf("%d ", i)
}
```

## 面向对象编程
### 自定义类型和结构体
### 方法
### 组合
### 接口
## 并发编程
### 并行与并发
### 协程
### groutine
### channel
### select
