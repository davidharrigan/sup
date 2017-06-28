// Copyright © 2017 David Harrigan <dave.t.harrigan@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package todo

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// SearchResult is a struct to store search result
type SearchResult struct {
	file    string
	line    int64
	content string
	author  string
}

// Search will walk through all directories starting at the given path and will find
// all files that match the search criteria
func Search(path string) []SearchResult {
	searchResults := []SearchResult{}
	filepath.Walk(path, visit(&searchResults))
	return searchResults
}

func visit(searchResults *[]SearchResult) filepath.WalkFunc {
	searchTerm := "TODO"
	return func(path string, f os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if f.IsDir() {
			return nil
		}
		*searchResults = append(*searchResults, SearchFile(path, []byte(searchTerm))...)
		return nil
	}
}

// Search for given pattern on the file
func SearchFile(file string, pat []byte) []SearchResult {
	line := int64(1)
	searchResults := []SearchResult{}

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if bytes.Contains(scanner.Bytes(), pat) {
			match := SearchResult{line: line, file: file, content: strings.TrimSpace(scanner.Text())}
			searchResults = append(searchResults, match)
		}
		line++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return searchResults
}

// PrintSearchResults will print the given searchResults to stdout
func PrintSearchResults(searchResults []SearchResult) {
	fmt.Printf("Found %d outstand TODOs!\n\n", len(searchResults))
	for _, result := range searchResults {
		fmt.Printf("%s [%d] %s\n", result.file, result.line, result.content)
	}
}
