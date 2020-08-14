package finder

import (
	"io/ioutil"
    "path/filepath"
	"gopkg.in/yaml.v2"
	"github.com/blackarrowsec/fozar/internal/log"
)

type config struct {
    Rules                 []string `yaml:"rules"`
}

type fileExceptions struct {
    Fexceptions           []string `yaml:"avoid"`
}


func ParseRules(rules_file string) []string{
	filename, err := filepath.Abs(rules_file)
    handleError(err)

	log.DebugPrint("Opening "+filename)

    yamlFile, err := ioutil.ReadFile(filename)
    handleError(err)

    var rules config
    err = yaml.Unmarshal(yamlFile, &rules)
    handleError(err)

    for _, rule := range rules.Rules {
    	 log.DebugPrint(" +found rule: "+rule)
    }

    return rules.Rules
}


func ParseFileExceptions(rules_file string) []string{
    filename, err := filepath.Abs(rules_file)
    handleError(err)
    log.DebugPrint("Opening "+filename)

    yamlFile, err := ioutil.ReadFile(filename)
    handleError(err)


    var rules fileExceptions
    err = yaml.Unmarshal(yamlFile, &rules)
    handleError(err)

    for _, rule := range rules.Fexceptions {
         log.DebugPrint(" +found rule: "+rule)
    }
    
    return rules.Fexceptions
}


func handleError(err error){
    if err != nil {
        log.Fatal(err)
        return
    }
}
