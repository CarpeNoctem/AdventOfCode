package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    var inputs = get_inputs("input.txt")

    offsets := []int{0}

    input_index := 0
    for {
        new_offset := offsets[len(offsets) -1] + inputs[input_index]
        //fmt.Printf("searching %v\n", new_offset)
        for _, prev_offset := range offsets {
            if prev_offset == new_offset {
                //fmt.Printf("Found repeat after %v iterations: %v\n", len(offsets), new_offset)
                fmt.Printf("Found repeat: %v\n", new_offset)
                os.Exit(0)
            }
        }
        offsets = append(offsets, new_offset)
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
