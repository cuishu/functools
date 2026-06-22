# Functools

[![Go Reference](https://pkg.go.dev/badge/github.com/cuishu/functools.svg)](https://pkg.go.dev/github.com/cuishu/functools)
[![Go Report Card](https://goreportcard.com/badge/github.com/cuishu/functools)](https://goreportcard.com/report/github.com/cuishu/functools)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Functools** 是一个轻量级、类型安全的 Go 语言辅助函数库，提供丰富的切片（slice）操作工具，涵盖函数式编程范式（`Map`/`Filter`/`Reduce`）、集合运算（去重、并集、差集）、数学统计、分组、反转及组合等常见需求。所有函数均基于 Go 泛型实现，确保高性能与编译期类型检查。

---

## 安装

```bash
go get github.com/cuishu/functools
```

---

## 函数索引

### 基础函数式操作
- [`Map`](#map) – 将切片元素逐一转换为新类型
- [`ForEach`](#foreach) – 遍历切片执行副作用
- [`Filter`](#filter) – 按条件筛选元素
- [`Reduce`](#reduce) – 归约切片为单个值

### 集合运算（去重、并集、差集）
- [`Unique`](#unique) / [`UniqueBy`](#uniqueby) – 去重
- [`Union`](#union) / [`UnionBy`](#unionby) – 合并多个切片并去重
- [`Diff`](#diff) / [`DiffBy`](#diffby) – 对称差集
- [`DiffSeparate`](#diffseparate) / [`DiffSeparateBy`](#diffseparateby) – 分别获取仅存在于 A 或 B 中的元素

### 数学与数值工具
- [`Sum`](#sum) – 求和
- [`Min`](#min) / [`Max`](#max) – 比较两个数值
- [`Range`](#range) – 生成整数序列

### 顺序操作
- [`Reverse`](#reverse) / [`ReverseInPlace`](#reverseinplace) – 反转切片

### 分组与组合
- [`GroupBy`](#groupby) – 按键分组
- [`Zip`](#zip) – 将两个切片按位置配对为 `Pair` 结构
- [`Pair`](#pair) – 二元组辅助类型

---

## 详细说明与示例

### `Map`

对切片中的每个元素应用转换函数，返回新切片。

```go
func Map[A, B any](list []A, f func(A) B) []B
```

**示例**：将整数切片转为字符串

```go
nums := []int{1, 2, 3}
strs := functools.Map(nums, func(i int) string { return strconv.Itoa(i) })
// strs: ["1", "2", "3"]
```

---

### `ForEach`

遍历切片并对每个元素执行传入的函数（常用于副作用操作）。

```go
func ForEach[A any](list []A, f func(A))
```

**示例**：打印每个元素

```go
functools.ForEach([]string{"a", "b"}, func(s string) { fmt.Println(s) })
```

---

### `Filter`

根据谓词函数过滤元素，返回满足条件的新切片。

```go
func Filter[A any](list []A, f func(A) bool) []A
```

**示例**：筛选偶数

```go
evens := functools.Filter([]int{1,2,3,4}, func(x int) bool { return x%2==0 })
// evens: [2,4]
```

---

### `Reduce`

使用归约函数将切片累计为单个值（需提供初始值）。

```go
func Reduce[A any](list []A, f func(A, A) A, init A) A
```

**示例**：求和

```go
sum := functools.Reduce([]int{1,2,3,4}, func(x, y int) int { return x+y }, 0)
// sum: 10
```

---

### `Unique` / `UniqueBy`

去除切片中的重复元素，保持首次出现顺序。

```go
func Unique[T comparable](list []T) []T
func UniqueBy[A any, B comparable](list []A, selector func(A) B) []A
```

- `Unique` 直接比较元素自身（需 `comparable`）。
- `UniqueBy` 通过 `selector` 提取比较键。

**示例**：

```go
ints := []int{1,2,2,3}
uniq := functools.Unique(ints) // [1,2,3]

type User struct{ ID int; Name string }
users := []User{{1,"A"}, {2,"B"}, {1,"C"}}
uniqByID := functools.UniqueBy(users, func(u User) int { return u.ID })
// 保留首次出现的 {1,"A"} 和 {2,"B"}
```

---

### `Union` / `UnionBy`

合并多个切片，去除重复元素（结果保持第一个切片的元素顺序）。

```go
func Union[T comparable](collections ...[]T) []T
func UnionBy[A any, B comparable](selector func(A) B, collections ...[]A) []A
```

**示例**：

```go
a := []int{1,2}
b := []int{2,3}
c := []int{3,4}
union := functools.Union(a, b, c) // [1,2,3,4]

// 按名称去重
type Item struct{ Name string }
items1 := []Item{{"x"},{"y"}}
items2 := []Item{{"y"},{"z"}}
unionByName := functools.UnionBy(func(it Item) string { return it.Name }, items1, items2)
// 保留第一个出现的同名元素
```

---

### `Diff` / `DiffBy`

返回两个切片的**对称差集**（即只在其中一个切片中出现的元素），结果去重。

```go
func Diff[T comparable](a, b []T) []T
func DiffBy[T any, K comparable](a, b []T, selector func(T) K) []T
```

**示例**：

```go
a := []int{1,2,3}
b := []int{3,4,5}
diff := functools.Diff(a, b) // [1,2,4,5]（顺序不定）
```

---

### `DiffSeparate` / `DiffSeparateBy`

分别返回 **仅在 A 中** 和 **仅在 B 中** 的元素（去重）。

```go
func DiffSeparate[T comparable](a, b []T) (onlyA, onlyB []T)
func DiffSeparateBy[T any, K comparable](a, b []T, selector func(T) K) (onlyA, onlyB []T)
```

**示例**：

```go
a := []int{1,2,3}
b := []int{3,4,5}
onlyA, onlyB := functools.DiffSeparate(a, b)
// onlyA: [1,2], onlyB: [4,5]
```

---

### `Sum`

计算数值切片的和（支持所有整数、无符号整数及浮点数类型）。

```go
func Sum[T Number](slice []T) T
```

**示例**：

```go
total := functools.Sum([]int{1,2,3}) // 6
```

---

### `Min` / `Max`

返回两个数值中的较小/较大值。

```go
func Min[T cmp.Ordered](a T, args ...T) T
func Max[T cmp.Ordered](a T, args ...T) T
```

**示例**：

```go
m := functools.Min(10, 20, 30) // 10
M := functools.Max(10, 20, 30) // 30
```

---

### `Range`

生成从 `start` 到 `stop`（不包含）的整数序列，步长为 `step`。

```go
func Range[T Int | Uint](start, stop, step T) []T
```

**示例**：

```go
seq := functools.Range(0, 10, 2) // [0,2,4,6,8]
```

> 注意：`stop` 必须大于 `start`，且步长必须为正。

---

### `Reverse` / `ReverseInPlace`

`Reverse` 返回反转后的新切片；`ReverseInPlace` 原地反转（不分配新内存）。

```go
func Reverse[T any](list []T) []T
func ReverseInPlace[T any](list []T)
```

**示例**：

```go
orig := []int{1,2,3}
reversed := functools.Reverse(orig)      // [3,2,1]
functools.ReverseInPlace(orig)           // orig 变为 [3,2,1]
```

---

### `GroupBy`

根据选择器函数返回的键对元素进行分组，结果为一个映射（map）。

```go
func GroupBy[A any, B comparable](list []A, selector func(A) B) map[B][]A
```

**示例**：按字符串长度分组

```go
words := []string{"go", "rust", "java", "zig"}
groups := functools.GroupBy(words, func(s string) int { return len(s) })
// groups[2] = ["go","zig"], groups[4] = ["rust","java"]
```

---

### `Zip`

将两个切片按位置配对，返回 `Pair` 切片。结果长度取较短者。

```go
func Zip[A, B any](a []A, b []B) []Pair[A, B]
```

**示例**：

```go
names := []string{"Alice", "Bob"}
ages := []int{30, 25}
pairs := functools.Zip(names, ages)
// pairs[0] = Pair{First:"Alice", Second:30}
```

---

### `Pair` 辅助类型

`Pair` 是一个简单的二元组结构，提供构造方法和取值方法。

```go
type Pair[A, B any] struct { First A; Second B }
func NewPair[A, B any](a A, b B) Pair[A, B]
func (p Pair[A, B]) Values() (A, B)
```

---

## 性能与兼容性

- **Go 版本**：需 Go 1.18 及以上（泛型支持）。
- **内存效率**：多数函数返回新切片（`ReverseInPlace` 除外），避免修改输入（除特殊说明）。
- **并发安全**：所有函数均为纯函数，不涉及并发，但输入切片在函数外部被修改不影响函数结果（除非传递引用类型元素）。

---

## 贡献

欢迎提交 Issue 或 Pull Request。请确保代码通过 `go test` 并保持测试覆盖率。

---

## 许可证

[MIT](https://opensource.org/licenses/MIT) © [cuishu](https://github.com/cuishu)