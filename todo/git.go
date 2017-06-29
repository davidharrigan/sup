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
	"os/exec"
	"strings"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func LookupGitUser() string {
	cmdName := "git"
	cmdArgs := []string{"config", "--global", "user.email"}
	user := ""

	if out, err := exec.Command(cmdName, cmdArgs...).Output(); err == nil {
		user = strings.TrimSpace(string(out))
	}

	return user
}

func GetCommitObject(path string) (*object.Commit, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		// log
		return nil, err
	}

	ref, err := repo.Head()
	if err != nil {
		// log
		return nil, err
	}

	commit, err := repo.CommitObject(ref.Hash())
	return commit, err
}

func SearchCurrentCommit(path string, commit *object.Commit, author string, searchTerm string) ([]SearchResult, error) {
	blame, err := git.Blame(commit, path)
	if err != nil {
		// log
		return nil, err
	}

	gitResult := []SearchResult{}

	for i, line := range blame.Lines {
		if strings.Contains(line.Text, searchTerm) {

			if author == "" || author == line.Author {
				match := SearchResult{line: i, file: path, content: strings.TrimSpace(line.Text), author: line.Author}
				gitResult = append(gitResult, match)
			}

		}
	}

	return gitResult, nil
}
