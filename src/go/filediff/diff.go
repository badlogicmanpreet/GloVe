package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type compare [][]byte

func (c compare) Len() int {
	return (len(c))
}

func (c compare) Less(i, j int) bool {
	return bytes.Compare(c[i], c[j]) < 0
}

func (c compare) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// read, sort, compare, diffs operations for files
func main() {
	/**
	//=============================== read & sort =================================
	customVocabFile, err1 := ioutil.ReadFile("/Users/manpreet.singh/git/deeplearning/dl/src/github.com/GloVe/src/go/mapreduce/result.txt")
	if err1 != nil {
		fmt.Println("Problem reading file", err1)
	}
	fmt.Println(len(customVocabFile))
	lines1 := bytes.Split(customVocabFile, []byte{'\n'})
	sort.Sort(compare(lines1))
	text1 := bytes.Join(lines1, []byte{'\n'})
	ioutil.WriteFile("/Users/manpreet.singh/git/deeplearning/dl/src/github.com/GloVe/src/sorted_ms.txt", text1, 0644)

	chrisVocabFile, err2 := ioutil.ReadFile("/Users/manpreet.singh/git/deeplearning/dl/src/github.com/GloVe/src/vocab.txt")
	if err2 != nil {
		fmt.Println("Problem reading file", err2)
	}
	fmt.Println(len(chrisVocabFile))
	lines := bytes.Split(chrisVocabFile, []byte{'\n'})
	sort.Sort(compare(lines))
	text := bytes.Join(lines, []byte{'\n'})
	ioutil.WriteFile("/Users/manpreet.singh/git/deeplearning/dl/src/github.com/GloVe/src/sorted_chris.txt", text, 0644)

	//=============================== compare =================================
	sortedMS, err3 := ioutil.ReadFile("/Users/manpreet.singh/git/deeplearning/dl/src/github.com/GloVe/src/sorted_ms.txt")
	if err3 != nil {
		fmt.Println("Problem reading file", err3)
	}

	fmt.Println(len(sortedMS))

	sortedChris, err4 := ioutil.ReadFile("/Users/manpreet.singh/git/deeplearning/dl/src/github.com/GloVe/src/sorted_chris.txt")
	if err4 != nil {
		fmt.Println("Problem reading file", err4)
	}

	fmt.Println(len(sortedChris))

	fmt.Println(bytes.Compare(sortedMS, sortedChris))

	*/
	//================================ diffs ================================ 253854, 253854
	storage := make(map[string]string)

	f, _ := os.Open("/Users/manpreet.singh/git/deeplearning/dl/src/github.com/GloVe/src/sorted_chris.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		var keyValue = strings.Split(line, " ")
		storage[keyValue[0]] = keyValue[1]
	}
	fmt.Println(len(storage))

	f1, _ := os.Open("/Users/manpreet.singh/git/deeplearning/dl/src/github.com/GloVe/src/sorted_ms.txt")
	scanner1 := bufio.NewScanner(f1)

	for scanner1.Scan() {
		line1 := scanner1.Text()
		var keyValue1 = strings.Split(line1, " ")
		_, ok := storage[keyValue1[0]]
		if ok {
			delete(storage, keyValue1[0])
		}
	}
	fmt.Println(len(storage))
	for key, value := range storage {
		fmt.Println("Key:", key, "Value:", value)
	}
}

func sorter() {

}
