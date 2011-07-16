package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"./indexer"
)

var index = indexer.New()
var files []string

func initialize() {
	flag.Parse();

	files = make([]string,flag.NArg())

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
		index.Add(i, string(content))
	}

}

func main() {
	initialize();
	
	results := index.Search("package")
	for e := results.Front(); e != nil; e = e.Next() {
		fmt.Printf("%s\n",files[e.Value])
	}
}
