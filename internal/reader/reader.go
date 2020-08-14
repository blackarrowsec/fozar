package reader


/*
Given a repo location traverses all files and commits
and applies an operation to each of them.
*/

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"github.com/blackarrowsec/fozar/internal/log"
	"github.com/blackarrowsec/fozar/internal/regex_searcher"
	"github.com/blackarrowsec/fozar/internal/custom_datatypes"
)


func SearchInRepo(regex []string,avoid []string,repo * custom_datatypes.Repo){
	r, err := git.PlainOpen(repo.RepoPath)
	handleError(err)

	ref, err := r.Head()
	handleError(err)

	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	handleError(err)

	log.PrintRepoName(repo.RepoName)
	err = cIter.ForEach(func(c *object.Commit) error {
		commit, err := r.CommitObject((*c).Hash)
		handleError(err)
		tree, err := commit.Tree()
		handleError(err)
		tree.Files().ForEach(func(f *object.File) error {
			if (regex_searcher.IsFileToBeParsed(avoid,f.Name)){
				log.DebugPrint("Checking: "+f.Name)
				regex_searcher.FindMatchOnFile(f,regex,repo)
			}
			return nil
		})
		regex_searcher.FindMatchOnMessage(commit.Message,commit.Hash,regex,repo)
		return nil
	})
	handleError(err)
}


func handleError(err error){
	if err != nil {
		log.Fatal(err)
	}
}