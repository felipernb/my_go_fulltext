package indexer

import (
	"container/list"
	"strings"
	"regexp"
)

type Indexer struct {
	dictionary map[string]*list.List
}

func New() *Indexer {
	indexer := new(Indexer)
	indexer.dictionary = make(map[string]*list.List)
	return indexer
}

func (indexer *Indexer) Add(ref int, content string) {
	re, err := regexp.Compile("([a-z]+)")
	if err != nil {
		return
	}
	content = strings.ToLower(content)
	
	words := re.FindAllString(content, -1)
	//indexes := re.FindAllStringIndex(content, -1)

	for _,w := range words {
		if indexer.dictionary[w] == nil {
			indexer.dictionary[w] = list.New()
		}
		indexer.dictionary[w].PushBack(ref)
	}
}

func (indexer *Indexer) Search(query string) *list.List {
	return indexer.dictionary[query]
}

