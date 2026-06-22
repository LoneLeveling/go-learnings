package main

func addHundred(x int) int {
        return x + 100
}
func partialSum(x ...int) func() int {
        sum := 0
        for _, value := range x {
                sum += value
        }
        return func() int {
                return addHundred(sum)
        }
}
func main() {
        partial := partialSum(1, 2, 3)
        partial()
}