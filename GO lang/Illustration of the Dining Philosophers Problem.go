package main
import (
"fmt"
"sync"
"time"
)
type Philosopher struct {
id int
leftFork, rightFork *sync.Mutex
diningCycles int
}
type DiningTable struct {
philosophers []*Philosopher
waiter *sync.Mutex
}
func (p *Philosopher) think() {
fmt.Printf("Philosopher %d is thinking\n", p.id)
time.Sleep(time.Second)
}
func (p *Philosopher) eat() {
p.leftFork.Lock()
p.rightFork.Lock()
fmt.Printf("Philosopher %d is eating\n", p.id)
time.Sleep(time.Second)
p.rightFork.Unlock()
p.leftFork.Unlock()
p.diningCycles++
}
func (p *Philosopher) dine(table *DiningTable, maxDiningCycles int) {
for p.diningCycles < maxDiningCycles {
p.think()
table.waiter.Lock()
p.eat()
table.waiter.Unlock()
}
}
func main() {
var numPhilosophers, maxDiningCycles int
fmt.Print("Enter the number of philosophers: ")
fmt.Scan(&numPhilosophers)
fmt.Print("Enter the maximum dining cycles: ")
fmt.Scan(&maxDiningCycles)
table := &DiningTable{
philosophers: make([]*Philosopher, numPhilosophers),
waiter: &sync.Mutex{},
}
forks := make([]*sync.Mutex, numPhilosophers)
for i := 0; i < numPhilosophers; i++ {
forks[i] = &sync.Mutex{}
}
for i := 0; i < numPhilosophers; i++ {
table.philosophers[i] = &Philosopher{
id: i,
leftFork: forks[i],
rightFork: forks[(i+1)%numPhilosophers],
}
}
// Start dining
var wg sync.WaitGroup
wg.Add(numPhilosophers)
for i := 0; i < numPhilosophers; i++ {
go func(p *Philosopher) {
defer wg.Done()
p.dine(table, maxDiningCycles)
}(table.philosophers[i])
}
wg.Wait()
}