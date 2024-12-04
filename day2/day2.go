package main

import "fmt"
import "os"
import "log"
import "strings"
import "strconv"

//7 6 4 2 1
//1 2 7 8 9
//9 7 6 2 1
//1 3 2 4 5
//8 6 4 4 1
//1 3 6 7 9

func isAsc(numbers []int) bool {
    if numbers[0] > numbers[len(numbers)-1] {
        return false
    }
    return true
}
func calculate(numbers []int) int {
    //lastMiss := numbers[0];
    if isAsc(numbers) {
        for i := 0; i < len(numbers) - 1; i++ {
            if (numbers[i+1] - numbers[i] > 3 || numbers[i] - numbers[i+1] == 0) {
                fmt.Println("Broke", numbers[i+1], numbers[i])
                return 0
            }
        }
    } else {
        for i := 0; i < len(numbers) - 1; i++ {
            if (numbers[i] - numbers[i+1] > 3 || numbers[i] - numbers[i+1] == 0) {
                fmt.Println("Broke", numbers[i+1], numbers[i])
                return 0
            }
        }

    }
    return 1
}

//only parses to int array
func calculateSafe(line string) int {
    parsed := strings.Fields(line)

    numbers := []int{}

    for _, element := range parsed {
        num, _:= strconv.Atoi(element)

        numbers = append(numbers, num)
    }
    return calculate(numbers)
}

func main() {

    input, err := os.ReadFile("input")
    if err != nil {
        log.Fatal(err)
    } 

    lines := strings.Split(strings.TrimSpace(string(input)), "\n")

    sum := 0
    for _, element := range lines {
        if calculateSafe(element) == 1 {
            sum += calculateSafe(element)
            continue
        }         
        fmt.Println("BAD", element)
    }
    fmt.Println("The safe reports are", sum)
}
