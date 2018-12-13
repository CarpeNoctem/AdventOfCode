package main

import (
	"bufio"
	"fmt"
	"os"
    //"runtime"
    "strconv"
)

//var num_threads = runtime.NumCPU()

func main() {
    inputs := get_inputs("input.txt")
    num_ids := len(inputs)

    // Did some playing around with different worker thread pool sizes, and
    // to my surprise, the solution was found faster, with sizes above the number
    // of available CPU cores. I suppose because we're doing data partitioning moreso
    // than actual computational work distribution.
    num_threads, _ := strconv.Atoi(os.Getenv("THREADS"));
    if num_threads < 1 { num_threads = 1 }
    fmt.Printf("Comparing %v ids across %v worker threads...\n", num_ids, num_threads)

    found_chan := make(chan string, 1)
    work_chan := make(chan int, num_threads)

    for i := 0; i < num_threads; i++ {
        go compare_ids_worker(inputs, num_ids, work_chan, found_chan)
    }

    for index, _ := range inputs {
        work_chan <- index
    }

    fmt.Println(<- found_chan)
}


func compare_ids_worker(id_list []string, length int, work_chan <-chan int, found_chan chan<- string) {
    for current_index := range work_chan {
        for other_index := current_index + 1; other_index < length; other_index++ {
            found, common_chars := is_closest_pair(id_list[current_index], id_list[other_index])
            if found { found_chan <- common_chars }
        }
    }
}

func is_closest_pair(id1 string, id2 string)(bool, string) {
    common_chars := ""
    for i, char := range id1 {
        if char == rune(id2[i]) { common_chars += string(char) }
    }
    return len(common_chars) == len(id1) -1, common_chars
}


func get_inputs(filename string) []string {
    file, _ := os.Open(filename)
    defer file.Close()

    inputs := []string{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        inputs = append(inputs, scanner.Text())
    }

    return inputs
}
