package custom_datatypes


type Match struct {
    CommitId    string
    FileName    string
    MatchedRule string
    MatchedData []string
}

type Repo struct {
    RepoName    string
    RepoPath    string
    Matches []Match
    AlreadyFound map[string]bool

}

type Result struct {
    PageTitle string
    Repo   []Repo
}
