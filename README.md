# Functools

[![Go Reference](https://pkg.go.dev/badge/github.com/cuishu/functools.svg)](https://pkg.go.dev/github.com/cuishu/functools)
[![Go Report Card](https://goreportcard.com/badge/github.com/cuishu/functools)](https://goreportcard.com/report/github.com/cuishu/functools)

这是一个简单的 Go 语言辅助函数库，提供了一些常用的切片操作函数。

## Map

对切片中的每个元素应用给定的函数，返回一个新的切片。

示例：

```go
a = Map(func(i int) int {
    return i + 1
}, a)
```

## Filter

根据给定的条件函数过滤切片，返回满足条件的元素组成的新切片。

示例：

```go
a = Filter(func(i int) bool {
    return i < 50
}, a)
```

## Reduce

使用给定的归并函数将切片归约为单个值。

示例：

```go
n := Reduce(func(x, y int) int {
    return x + y
}, a, 0)
```

## Reverse

反转切片中的元素顺序。

示例：

```go
a = Reverse(a)
```

## Unique

去除切片中的重复元素，返回一个新切片（保持原顺序）。

示例：

```go
a = Unique(a)
```