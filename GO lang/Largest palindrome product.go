package main
import (
"fmt"
"strconv"
)
func isPalindrome(n int) bool {
str := strconv.Itoa(n)
length := len(str)
for i := 0; i < length/2; i++ {
if str[i] != str[length-i-1] {
return false
}
}
return true
}
func largestPalindromeProduct(min, max int) (int, int, int) {
largestPalindrome := 0
var multiplicand1, multiplicand2 int
for i := max; i >= min; i-- {
for j := i; j >= min; j-- {
product := i * j
if product < largestPalindrome {
break
}
if isPalindrome(product) && product > largestPalindrome {
largestPalindrome = product
multiplicand1 = i
multiplicand2 = j
}
}
}
return largestPalindrome, multiplicand1, multiplicand2
}
func main() {
var min, max int
fmt.Print("Enter the minimum value: ")
fmt.Scan(&min)
fmt.Print("Enter the maximum value: ")
fmt.Scan(&max)
result, multiplicand1, multiplicand2 := largestPalindromeProduct(min, max)
fmt.Printf("The largest palindrome product between %d and %d is: %d\n", 
min, max, result)
fmt.Printf("The multiplicands are: %d and %d\n", multiplicand1, 
multiplicand2)
}
