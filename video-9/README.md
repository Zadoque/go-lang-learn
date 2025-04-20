# Video 9: Function Basics
## Video Information
**URL**: https://www.youtube.com/playlist?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=9  
**Focus**: Function declaration and usage

## Key Concepts
- Function parameters
- Return types
- Named returns
- Basic error handling

```go
func add(a int, b int) int {
    return a + b
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```