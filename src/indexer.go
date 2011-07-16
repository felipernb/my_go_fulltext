package indexer

import (
	"container/list"
	"strings"
	"regexp"
)

/*
  The indexer contains a hashmap where the words are the keys, 
  and the values are other hashmaps where the keys are the file names
  and the values are linked-list containing the position the word appears
  in that file

  @author Felipe Ribeiro <felipernb@gmail.com>
*/
type Indexer struct {
	dictionary map[string]map[int] *list.List
}

func New() *Indexer {
	indexer := new(Indexer)
	indexer.dictionary = make(map[string]map[int] *list.List)
	return indexer
}

/*
  Adds a file content to the index.
  The text is tokenized and stored in the Indexer hashmap
*/
func (indexer *Indexer) Add(ref int, content string) {
	re, err := regexp.Compile("([a-z]+)")
	if err != nil {
		return
	}
	content = strings.ToLower(content)
	
	words := re.FindAllString(content, -1)
	indexes := re.FindAllStringIndex(content, -1)

	for i, w := range words {
		if indexer.dictionary[w] == nil {
			indexer.dictionary[w] = make(map[int] *list.List)
		}
		if indexer.dictionary[w][ref] == nil {
			indexer.dictionary[w][ref] = list.New()
		}
		indexer.dictionary[w][ref].PushBack(indexes[i][0])
	}
}

func (indexer *Indexer) Search(query string) map[int] *list.List {
	return indexer.dictionary[query]
}

