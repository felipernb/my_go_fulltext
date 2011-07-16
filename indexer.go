package indexer

import (
	"container/list"
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
}

func (indexer *Indexer) Search(query string) *list.List {
	return indexer.dictionary[query]
}


