package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    var inputs = get_inputs("input.txt")
    new_offset := 0
    offsets := map[int]bool{0: true}

    input_index := 0
    for {
        new_offset = new_offset + inputs[input_index]
        if _, exists := offsets[new_offset]; exists {
            fmt.Printf("Found repeat: %v\n", new_offset)
            os.Exit(0)
        }
        offsets[new_offset] = true
        input_index += 1
        if input_index >= len(inputs) {
            input_index = 0
        }
    }
}

func get_inputs(filename string) []int {
    file, _ := os.Open(filename)
    defer file.Close()

    inputs := []int{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        input, _ := strconv.Atoi(scanner.Text())
        inputs = append(inputs, input)
    }

    return inputs
}
