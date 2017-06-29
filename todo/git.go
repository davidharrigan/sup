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
	"gopkg.in/src-d/go-git.v4"
	"strings"
)

func SearchCurrentCommit(repo *git.Repository, author string, result SearchResults) (SearchResults, error) {
	searchResults := make(map[string][]SearchResult)

	ref, err := repo.Head()
	if err != nil {
		// log
		return nil, err
	}

	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		// log
		return nil, err
	}

	for f := range result {
		gitResult := []SearchResult{}
		blame, err := git.Blame(commit, f)

		if err == nil {

			for i, line := range blame.Lines {
				if strings.Contains(line.Text, "TODO") {
					match := SearchResult{line: i, file: f, content: strings.TrimSpace(line.Text), author: line.Author}
					gitResult = append(gitResult, match)
				}
			}
		}

		searchResults[f] = gitResult
	}

	return searchResults, nil
}
