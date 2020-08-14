package report

import (
	"os"
	"github.com/blackarrowsec/fozar/internal/log"
	"github.com/blackarrowsec/fozar/internal/custom_datatypes"
	"html/template"
    "github.com/blackarrowsec/fozar/internal/separator"
    "io"
    "bufio"
)


var (
	FancyFile string = "output"+separator.PATH_SEPARATOR+"fozar.html"
    TextFile string = "output"+separator.PATH_SEPARATOR+"fozar.md"
    OutputFolder string = "output"+separator.PATH_SEPARATOR+"Fozar Report"+separator.PATH_SEPARATOR
    TemplatesLocation string = "templates"+separator.PATH_SEPARATOR
    cssFiles []string = []string{"base_report.css","print_report.css"}
    jsFiles []string = []string{"base_report.js"}
)


/*
    Fancy Output
*/

func WriteFancyReport(result custom_datatypes.Result) {
    createOutputFolderAndCopyTemplates()
    templates,_ := template.ParseFiles(
            TemplatesLocation + "html" + separator.PATH_SEPARATOR + "base_report.html",
            TemplatesLocation + "html" + separator.PATH_SEPARATOR + "base_repo.html",
            TemplatesLocation + "html" + separator.PATH_SEPARATOR + "base_commit.html",
        )
    var filename = OutputFolder + FancyFile
    file,err := os.Create(filename)
    handleError(err)
    defer file.Close()
    templates.Execute(file,result)
    file.Sync()
}


func createOutputFolderAndCopyTemplates(){
    os.MkdirAll(OutputFolder, os.ModePerm)
    os.MkdirAll(OutputFolder+"js", os.ModePerm)
    os.MkdirAll(OutputFolder+"css", os.ModePerm)
    for _,file := range cssFiles {
       copy(TemplatesLocation+"css"+separator.PATH_SEPARATOR+file,OutputFolder+"css"+separator.PATH_SEPARATOR+file)
    }
    for _,file := range jsFiles {
       copy(TemplatesLocation+"js"+separator.PATH_SEPARATOR+file,OutputFolder+"js"+separator.PATH_SEPARATOR+file)
    }
}


// https://opensource.com/article/18/6/copying-files-go
func copy(src, dst string) (int64, error) {

        source, err := os.Open(src)
        handleError(err)
        defer source.Close()

        destination, err := os.Create(dst)
        handleError(err)

        defer destination.Close()
        nBytes, err := io.Copy(destination, source)
        return nBytes, err
}
/*
        Text Output
*/

func WriteTextReport(result custom_datatypes.Result) {
    f, err := os.Create(TextFile)
    handleError(err)
    defer f.Close()
    w := bufio.NewWriter(f)
    w.WriteString("\n# "+result.PageTitle+"\n")
    w.WriteString("\n[TOC]\n")
    for _,repo := range result.Repo {
        w.WriteString("\n## "+repo.RepoName+"\n\t"+repo.RepoPath+"\n")
        for _,match := range repo.Matches {
            w.WriteString("\n### "+match.FileName+"\n\t["+match.CommitId+"]\n\t["+match.MatchedRule+"]"+"\n")
            w.WriteString("```python\n")
            for _,line := range match.MatchedData {
                w.WriteString(line+"\n")
            }
            w.WriteString("```\n")
        }
        w.WriteString("---"+"\n")
    }
    w.Flush()
}



func handleError(err error){
	 if err != nil {
        log.PrintError(err)
        return
	}
}