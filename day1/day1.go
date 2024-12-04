package main 

import "os"
import "log"
import "bufio"
import "fmt"
import "strconv"
import "bytes" 
import "sort"

func abs(x int) int {
    if x > 0{
        return x
    }
    return -x
}

func calcSimilar(x, y []int) int {
    m := make(map[int]int)
    for i := 0; i < len(y); i++ {
        m[y[i]] = m[y[i]] + 1
    }

    sum := 0
    for i := 0; i < len(x); i++ {
        sum += x[i] * m[x[i]]
    }
    return sum
}

func main() {
    input, err := os.ReadFile("input")
    if err != nil {
        log.Fatal(err)
    }
    var leftV []int
    var rightV []int

    scanner := bufio.NewScanner(bytes.NewReader(input))
    for scanner.Scan() {

        line := scanner.Text()
        numbers := bytes.Fields([]byte(line))

        if len(numbers) == 2 {
            leftN, _ := strconv.Atoi(string(numbers[0]))
            rightN, _ := strconv.Atoi(string(numbers[1]))
            leftV = append(leftV, leftN)
            rightV = append(rightV, rightN)
        }
    }
    fmt.Println("The similarity is ", calcSimilar(leftV, rightV))// part two of day 1

    sort.Ints(leftV)
    sort.Ints(rightV)
    sum := 0

    for i := 0; i < len(leftV); i++ {
        sum += abs(leftV[i]-rightV[i])
    }
    fmt.Println(sum)
}
