package main
import (
"fmt"
)
func main() {
myMap := make(map[string]int)
var key string
var value int
for {
fmt.Print("Enter a key (or 'quit' to exit): ")
fmt.Scan(&key)
if key == "quit" {
break
}
fmt.Print("Enter a value: ")
fmt.Scan(&value)
myMap[key] = value
}
fmt.Println("Map after user input:", myMap)
fmt.Print("Enter a key to get its value: ")
fmt.Scan(&key)
if val, ok := myMap[key]; ok {
fmt.Printf("Value of %s: %d\n", key, val)
} else {
fmt.Printf("Key %s not found in the map\n", key)
}
fmt.Print("Enter a key to delete: ")
fmt.Scan(&key)
delete(myMap, key)
fmt.Println("Map after deleting:", myMap)
fmt.Println("Iterating over the map:")
for key, value := range myMap {
fmt.Printf("Key: %s, Value: %d\n", key, value)
}
fmt.Println("Length of the map:", len(myMap))
}
