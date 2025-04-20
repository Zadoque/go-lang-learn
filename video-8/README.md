# Video 8: Looping with For
## Video Information
**URL**: https://www.youtube.com/playlist?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=8  
**Focus**: Iteration constructs

## Key Concepts
- Three-component for loops
- While-like loops
- Infinite loops
- Range loops

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

numbers := []int{1, 2, 3}
for index, value := range numbers {
    fmt.Println(index, value)
}
```