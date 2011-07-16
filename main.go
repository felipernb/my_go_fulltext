package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"container/list"
)

type Indexer struct {
	dictionary map[string]*list.List
}

func NewIndexer() {
	indexer := new(Indexer)
	indexer.dictionary = make(map[string]*list.List)
}

func (indexer *Indexer) add(ref int, content string) {		
}

func (indexer *Indexer) search(query string) *list.List {
	return indexer.dictionary[query]
}

func initialize() {
	flag.Parse();

	files := make([]string,flag.NArg())
	contents := make([]string, flag.NArg())
	index := NewIndexer()

	fmt.Printf("Indexing %d files...\n", flag.NArg())
	for i := 0; i < flag.NArg(); i++ {
		fmt.Printf("Indexing %s: ",flag.Arg(i))
		content, e := ioutil.ReadFile(flag.Arg(i))
		
		if e != nil {
			fmt.Printf("FAIL! %s couldn't be open.\n", flag.Arg(i))
			os.Exit(1)
		}
		
		fmt.Printf("OK! %d/%d\n", i+1, flag.NArg())

		files[i] = flag.Arg(i)
		contents[i] = string(content)

	}
	index.index(contents)

}

func main() {
	initialize();
}
