package main
import (
"fmt"
"math"
)
func main() {
var num int
fmt.Print("Enter a number: ")
fmt.Scan(&num)
largestFactor := -1
for num%2 == 0 {
largestFactor = 2
num /= 2
}
for i := 3; i <= int(math.Sqrt(float64(num))); i += 2 {
for num%i == 0 {
largestFactor = i
num /= i
}
}
if num > 2 {
largestFactor = num
}
fmt.Printf("The largest prime factor is: %d\n", largestFactor)
}