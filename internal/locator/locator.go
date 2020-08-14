package locator

/*
Identifies result under a given directory
*/

import (
	"github.com/blackarrowsec/fozar/internal/log"
	"github.com/blackarrowsec/fozar/internal/custom_datatypes"
    "github.com/blackarrowsec/fozar/internal/separator"
    "os"
    "path/filepath"
    "strings"
)


func LocateRepositories(path string) custom_datatypes.Result {
	var result custom_datatypes.Result
    var found_repos map[string]bool
    found_repos = make(map[string]bool)

    err := filepath.Walk(path,
        func(path string, info os.FileInfo, err error) error {
        handleError(err)        
        if info.Name() == ".git"{
             log.DebugPrint("Path: "+path)
            splitted_path := strings.Split(path,separator.PATH_SEPARATOR)
            repo_name := splitted_path[len(splitted_path)-2]
            if found_repos[repo_name] == false { 
                log.DebugPrint("Found "+repo_name)
                var current_repo custom_datatypes.Repo
                current_repo.RepoName = repo_name
                current_repo.RepoPath = filepath.Dir(path)
                current_repo.AlreadyFound = make(map[string]bool)
                result.Repo = append(result.Repo,current_repo)
                found_repos[repo_name] = true
            }
        }
        return nil
    })
    handleError(err)
	return result
}

func handleError(err error){
    if err != nil {
        log.PrintError(err)
    }
}