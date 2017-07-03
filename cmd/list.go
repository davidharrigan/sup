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

package cmd

import (
	"fmt"

	"github.com/davidharrigan/sup/todo"
	"github.com/spf13/cobra"

	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list <path>",
	Short: "List outstanding TODO items in current or specified directory",
	Long:  "List outstanding TODO items in current or specified directory",
	Run:   listRun,
}

var skipAuthor bool
var skipGit bool
var email string

const LookupTemplate = "Looking up commits by %s...\n\n"

func listRun(cmd *cobra.Command, args []string) {
	root := "./"
	if len(args) > 0 {
		root = args[0]
	}

	var commit *object.Commit
	var err error

	author := email

	// determine if this is a git directory
	if !skipGit && !skipAuthor {
		commit, err = todo.GetCommitObject(root)
		if err == nil {
			// log error
			if author == "" {
				author = todo.LookupGitUser()
			}
			if author != "" {
				fmt.Printf(LookupTemplate, author)
			}
		}
	} else {
		author = ""
	}

	result := todo.Search(root, commit, author)
	todo.PrintSearchResults(result)
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&skipGit, "skip-git", "g", false, "Skip current git commit lookup")
	listCmd.Flags().BoolVarP(&skipAuthor, "skip-author", "a", false, "Skip author lookup")
	listCmd.Flags().StringVarP(&email, "email", "e", "", "Override author look up value (default git config --global user.email")
}
