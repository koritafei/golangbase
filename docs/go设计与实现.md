# `GO`设计与实现
## 编译原理
抽象语法树(`AST`): 源代码语法结构的一种抽象表示，用树状的方式表示编程语言的语法结构。
抽象语法树中的每个节点表示源代码中的一个元素，每颗子树表示一个语法结构。
`e.g. 2 * 3 + 7`的抽象语法树：
![简单抽象语法树](images/简单抽象语法树.png)  
抽象语法树去除了源代码中的不重要字符，如空格，括号和分号等。
### 静态单赋值
`Static Single Assigment(SSA)`: 中间代码的一个特性。
如果一个代码具有静态单赋值特性，那么每个变量只会被赋值一次。
```go
x := 1
x := 2
y := x
```
如上述代码中，`x`被赋值多次，但`x := 1`在代码中是无用的。
在中间代码中使用`SSA`特性能够使程序实现以下优化：
1. 常数传播(`constant propagation`);
2. 值域传播(`value range propagation`);
3. 稀疏有条件的常数传播(`spare conditional constant propagation`);
4. 消除无用的程式码(`dead code elimination`);
5. 全域数值编号(`global value numbering`);
6. 消除部分冗余(`partial redundancy elimination`);
7. 强度折减(`strength reduction`);
8. 寄存器分配(`register allocation`)。

### 编译器
`go`的编译器源代码存储在`src/cmd/compile`。
![编译器组成](images/编译器组成.png)  
`go`编译器由如下几个部分组成：词法与语法分析器，类型检查和`AST`转换，通用`SSA`生成和机器码生成。
#### 词法与语法分析器
所有的编译过程都是从解析源代码文件开始的。
词法分析的作用：解析源代码文件，将源代码文件中的字符串系列转换为`Token`序列。一般将词法分析器称之为`lexer`。
语法分析的作用: 将词法分析生成的`Token`按照语言定义好的文法`Grammar`自上而下或者自下而上的进行规约，每一个`GO`源代码文件都会被规约成一个`SourceFile`结构。
```go
SourceFile := PackageClause ";" {ImportDecl ";"} {TopLevelDecl ";"}
```
词法分析会返回一个不包含空格，换行等的`Token`序列。例如：`package`, `json`, `import`, `(`... 语法分析将`Token`序列装成有意义的结构体，也就是语法树：
```go
"json.go": SourceFile{
  PackageName: "json",
  ImportDecl: []Import{
    "io",
  },
  TopLevelDecl: ...
}
```
每个抽象语法树(`AST`),包含当前文件属于的包名，定义的常量，结构体和函数等。
**每一个的`AST`都对应一个单独的`GO`文件.**
### 类型检查
检查类型定义，同时将对应的函数展开。
在这个阶段，`make`关键字会被替换成`makeslice`,`makechan`等函数。
![gomake](images/gomake.png)
类型检查之后，通过`complileFunctions`函数，开始编译全部函数。
这些函数会在一个工作队列里被多个协程消费。
![并发编译](images/并发编译过程.png)
### 数组
#### 初始化
```go
arr1 := [3]int{1,2,3}
arr2 := [...]int{1,2,3} // 产生数组大小推导
```
字面量数组处理：
> 1. 长度小与或等于`4`，直接将数组元素放置到栈上；
> 2. 长度大于`4`，将数组元素放置到静态区，并在运行时取出。

数组是一块连续的内存区，所以数组的算法需要以下参数：
> 1. 数组的首地址；
> 2. 数组元素个数；
> 3. 数组元素类型。

### 切片
切片结构：
```go
type SliceHeader struct {
  Data uintptrt // 指向数组的指针
  Len int // 当前切片长度
  Cap int // 当前切片的容量
}
```
![golangslice 结构](./images/slice.png)
`slice`追加操作：
> 1. 容量满足时，直接在data指针地址后追加元素；
> 2. 容量不满足时，进入`growslice`的流程，对原`slice`进行扩容，扩容完成之后，追加元素。

![sliceappend](./images/sliceappend.png)
`slice`扩容策略:
1. 期望容量大于当前容量的`2`倍时，使用期望容量作为新的容量;否则：
2. 如果当前容量小与`1024`时,就会将容量翻倍；
3. 如果当前容量大于`1024`时，每次扩容`25%`,直到新的容量大于期望容量。

#### 切片拷贝
内存拷贝：
![slicecopy](images/slicecopy.png)

内存拷贝是通过拷贝整块内存实现。

### 哈希表
![golanghash](images/hash.png)
冲突解决：
1. 开放寻址法
   `key`发生冲突时，则将依次探索下一个空白地址处。
2. 拉链法
   将冲突的`key`存入到一个链表中。

#### 哈希表实现
```go
type hmap struct {
  count int // 哈希表数量
  flags uint8 
  B uint8 // 哈希表持有的buckets个数
  noverflow uint8
  hash0 uint32 // 哈希种子
  buckets unsafe.Pointer // 哈希表扩容时保存的之前的buckets字段，当前buckets大小的一半
  oldbuckets unsafe.Pointer
  nevacuate uintptr
  extra *mapextra
}
```
![hashtable](images/hashtable.png)
`bmap`对应一个`hash buckets`,每个`bmap`可以存储`8`个元素，超出部分放置到`extra`中。
```go
type bmap struct {
  topbits [8]uint8
  keys [8]keytype
  values [8]valuetype
  pad uintptr
  overflow uintptr
}
```









