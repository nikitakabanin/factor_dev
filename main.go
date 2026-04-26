package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func factorial(n int) int {
    if n < 0 {
        return -1
    }
    
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    
    fmt.Print("Введите число для вычисления факториала: ")
    scanner.Scan()
    
    input := strings.TrimSpace(scanner.Text())
    n, err := strconv.Atoi(input)
    
    if err != nil {
        fmt.Println("Ошибка: введите корректное целое число")
        return
    }
    
    if n < 0 {
        fmt.Println("Ошибка: факториал отрицательного числа не определен")
    } else {
        fmt.Printf("%d! = %d\n", n, factorial(n))
    }
}