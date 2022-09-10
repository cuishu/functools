# Function Tools

# Map

example:

```go
	a = Map(func(i int) int {
		return i + 1
	}, a)
```

# Filter

example:

```go
	a = Filter(func(i int) bool {
		return i < 50
	}, a)
```

# Reduce

example:

```go
	n := Reduce(func(x, y int) int {
		return x + y
	}, a, 0)
```
