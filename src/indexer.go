package indexer

import (
	"container/list"
	"strings"
	"regexp"
)

/*
 * The indexer contains a hashmap where the words are the keys, 
 * and the values are other hashmaps where the keys are the file names
 * and the values are linked-list containing the position the word appears
 * in that file
 *
 * @author Felipe Ribeiro <felipernb@gmail.com>
 */
type Indexer struct {
	dictionary map[string]map[int] *list.List
	numberOfFiles int
}

func New() *Indexer {
	indexer := new(Indexer)
	indexer.dictionary = make(map[string]map[int] *list.List)
	indexer.numberOfFiles = 0
	return indexer
}

/*
 * Adds a file content to the index.
 * The text is tokenized and stored in the Indexer hashmap
 */
func (indexer *Indexer) Add(ref int, content string) {
	re, err := regexp.Compile("([a-z]+)")
	if err != nil {
		return
	}
	content = strings.ToLower(content)
	
	words := re.FindAllString(content, -1)
	indexes := re.FindAllStringIndex(content, -1)
	fileIsAlreadyIncluded := true

	for i, w := range words {
		if indexer.dictionary[w] == nil {
			indexer.dictionary[w] = make(map[int] *list.List)
		}
		if indexer.dictionary[w][ref] == nil {
			fileIsAlreadyIncluded = false
			indexer.dictionary[w][ref] = list.New()
		}
		indexer.dictionary[w][ref].PushBack(indexes[i][0])
	}
	
	if !fileIsAlreadyIncluded {
		indexer.numberOfFiles++
	}
}

func (indexer *Indexer) Search(query string) map[int] *list.List {
	query = strings.TrimSpace(strings.ToLower(query))

	var result map[int] *list.List
	if strings.Contains(query, " ") {
		result = indexer.multiSearch(strings.Split(query, " ", -1))				
	} else {
		result = indexer.dictionary[query]
	}
	return result
}

/*
 *   Makes separate queries for each token and than calculates the intersection
 */
func (indexer *Indexer) multiSearch(subqueries []string) map[int] *list.List {
	tempResults := make(map[string] map[int] *list.List)

	for _,q := range subqueries {
		tempResults[q] = indexer.dictionary[q]
	}
	numberOfUniqueTokens := len(tempResults)
	numberOfTokensFoundInFile := make([]int, indexer.numberOfFiles)
	
	//Count the ocurrences of the files on the search results
	// as we use maps for each token, there's no way for a file occur more than 
	// once for each token, so if # of ocurrences == # of tokens, the file contains
	// all tokens

	for _, res := range tempResults {
		for file, _ := range res {
			numberOfTokensFoundInFile[file]++
		}
	}
	
	results := make(map[int] *list.List)
	
	for _, res := range tempResults {
		for file, occurrences := range res {
			if numberOfTokensFoundInFile[file] == numberOfUniqueTokens {
				if results[file] == nil {
					results[file] = list.New()
				}
				results[file].PushBackList(occurrences)
			}
		}
	}

	return results	
}	

