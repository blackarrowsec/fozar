package fozar

import (
	"flag"
	"sync"
	"time"
	"github.com/blackarrowsec/fozar/internal/log"
	"github.com/blackarrowsec/fozar/internal/locator"
	"github.com/blackarrowsec/fozar/internal/finder"
	"github.com/blackarrowsec/fozar/internal/reader"
	"github.com/blackarrowsec/fozar/internal/custom_datatypes"
	"github.com/blackarrowsec/fozar/internal/regex_searcher"
	"github.com/blackarrowsec/fozar/internal/report"

)

type repository custom_datatypes.Repo

func Run(){
	log.PrintBanner()

	// Parse flags
	var work_path string
	var rules_file string
	flag.StringVar(&work_path, "path","", "Folder from which to start searching")
	flag.BoolVar(&(log.Debug),"debug",false, "Print debug information (slower!)")
	flag.StringVar(&rules_file,"config","config/config.yml","Yaml config location (default \"config/config.yml\")")
	flag.StringVar(&(report.TextFile),"ot","","Filename for the Markdown file output")
	flag.IntVar(&(regex_searcher.Previous_lines),"B",1,"Previous lines to show on match")
	flag.IntVar(&(regex_searcher.Following_lines),"A",1,"Following lines to show on match")
	flag.StringVar(&(report.FancyFile),"of","","Filename for the html file output")
	flag.Parse()

	// Check for path
	if work_path == "" {
    	flag.PrintDefaults()
    	return
	}

	// Locate repos
	result := locator.LocateRepositories(work_path)
	result.PageTitle = "Fozar Result"
	
	// Parse Regex
	config := finder.ParseRules(rules_file)
	avoid := finder.ParseFileExceptions(rules_file)

	// Time measure
	start := time.Now()

	// Multithreading
	repo_number := len(result.Repo)
	var wg sync.WaitGroup
	wg.Add(repo_number)

	// Search in repos
	for index,repo := range result.Repo {
		do_work(&wg,&repo,config,avoid)
		result.Repo[index] = repo

	}

	// Multithreading
	wg.Wait()
	
	elapsed := time.Since(start)
	log.RawPrint("Execution time: "+elapsed.String())

	if report.TextFile != "" {
		report.WriteTextReport(result)
	}
	if report.FancyFile != ""{
		report.WriteFancyReport(result)
	}
}

func do_work (wg *sync.WaitGroup,repo * custom_datatypes.Repo, rules []string,avoid []string){
	defer wg.Done()
	log.DebugPrint("# Searching "+repo.RepoName)
	reader.SearchInRepo(rules,avoid,repo)
}