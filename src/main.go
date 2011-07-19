/*
User I/O
@author Felipe Ribeiro <felipernb@gmail.com>
*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"bufio"
	"./indexer"
)

var index = indexer.New()
var files []string

func initialize() {
	flag.Parse();

	if flag.NArg() == 0 {
		fmt.Printf("You need to specify the files to be indexed as arguments.\ne.g.: $ indexer fileA.txt fileB.txt\n")
		os.Exit(1)
	}

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

func query(q string) {
	fmt.Printf("Searching for \"%s\":\n",q)

	startTimeSec, startTimeNSec,_  := os.Time()
	results := index.Search(q)
	
	occurrences := 0

	for file, positions := range results {

		if positions.Len() == 1 {
			fmt.Printf("%s - 1 occurrence at position: ",files[file])
		} else {
			fmt.Printf("%s - %d occurrences at positions: ",files[file], positions.Len())
		}

		for p := positions.Front(); p != nil; p = p.Next() {
			fmt.Printf("%d ",p.Value.(int))
			occurrences++
		}
		fmt.Printf("\n")
	}
	endTimeSec, endTimeNSec,_ := os.Time()
	
	startTime := float64(startTimeSec) + float64(startTimeNSec)/1e9
	endTime := float64(endTimeSec) + float64(endTimeNSec)/1e9
	fmt.Printf("Total: Found %d occurrences of \"%s\" in %f secs\n\n", occurrences, q, endTime - startTime)
}

func prompt(in *bufio.Reader) string {
	fmt.Printf("query> ")
	 u_q,_,_ := in.ReadLine()
	 return string(u_q)
}

func main() {
	initialize();
	in := bufio.NewReader(os.Stdin)

	for q := prompt(in); q != ""; q = prompt(in) {
		query(q)
	}

}
