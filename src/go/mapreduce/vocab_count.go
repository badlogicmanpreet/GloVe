package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func readInputs() []string {
	var names []string
	names = append(names, "/Users/manpreet.singh/git/deeplearning/dl/src/github.com/GloVe/src/corpus.txt")
	return names
}

func line(filename string) chan string {
	output := make(chan string)

	go func() {
		file, err := os.Open(filename)

		if err != nil {
			return
		}

		defer file.Close()
		reader := bufio.NewReader(file)

		for {
			line, err := reader.ReadString('\n')
			output <- line

			if err == io.EOF {
				break
			}
		}

		close(output)
	}()

	return output
}

func find_files(dirname string) chan interface{} {
	output := make(chan interface{})

	go func() {
		_find_files(dirname, output)
		close(output)
	}()

	return output
}

func _find_files(dirname string, output chan interface{}) {
	dir, _ := os.Open(dirname)
	dirnames, _ := dir.Readdirnames(-1)

	for i := 0; i < len(dirnames); i++ {
		fullpath := dirname + "/" + dirnames[i]
		file, _ := os.Stat(fullpath)

		if file.IsDir() {
			_find_files(fullpath, output)
		} else {
			output <- fullpath
		}
	}
}

// Split in words
func mapper(file string, value string) (res []KeyValue) {
	words := regexp.MustCompile(`[A-Za-z0-9_]*`)

	for line := range line(file) {
		fmt.Fprintln(fileptr, "No. of lines")
		for _, match := range words.FindAllString(line, -1) {
			//fmt.Println("match is : " + match)
			kv := KeyValue{match, "1"}
			res = append(res, kv)
		}
	}
	return res
}

func reducer(key string, values []string) string {
	var sum int
	sum = 0
	for _, i := range values {
		val, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println("got an error")
		}
		sum = sum + val
	}
	fmt.Fprintln(fileptr, key+" "+strconv.Itoa(sum))
	return ""
}

var fileptr, err = os.Create("/Users/manpreet.singh/git/deeplearning/dl/src/github.com/GloVe/src/go/mapreduce/result.txt")

func main() {
	nReduce := 50
	mr := Sequential("test", readInputs(), nReduce, mapper, reducer)
	mr.Wait()
}
