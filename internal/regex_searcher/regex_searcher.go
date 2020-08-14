package regex_searcher

import (
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"github.com/blackarrowsec/fozar/internal/custom_datatypes"
	"regexp"
	"strings"
	"crypto/md5"
	"encoding/hex"
)

var (
	Previous_lines int = 1
	Following_lines int = 1
)


func IsFileToBeParsed(avoid []string,filename string) bool{
	for _, av := range avoid {
		rav, _ := regexp.Compile(av)
		if  rav.MatchString(filename){
			return false
		}
	}
	return true
}

func FindMatchOnMessage(message string,hash plumbing.Hash,regex []string,repo * custom_datatypes.Repo){
	for _, rule := range regex {
		r, _ := regexp.Compile(rule)
		content := message
		if r.MatchString(content){
			lines := strings.Split(content,"\n")
			for line_counter,line := range lines {
				if r.MatchString(line) {
					var sum string
					if line_counter > 0 && line_counter < (len(lines)-2){
						md5HashInBytes := md5.Sum([]byte(lines[line_counter-1]+line+lines[line_counter+1]))
						sum = hex.EncodeToString(md5HashInBytes[:])
					}else if line_counter <= 0 && line_counter < (len(lines)-2){
						md5HashInBytes := md5.Sum([]byte(line+lines[line_counter+1]))
						sum = hex.EncodeToString(md5HashInBytes[:])
					}else if line_counter > 0 && line_counter == (len(lines)-1){
						md5HashInBytes := md5.Sum([]byte(lines[line_counter-1]+line))
						sum = hex.EncodeToString(md5HashInBytes[:])
					}else {
						md5HashInBytes := md5.Sum([]byte(line))
						sum = hex.EncodeToString(md5HashInBytes[:])
					}
					if repo.AlreadyFound[sum] == false {
						repo.AlreadyFound[sum] = true
						matched_strings := []string{}

						for _,element := range getPrevious(line_counter,Previous_lines,lines) {
							matched_strings = append(matched_strings,element)
						}
						matched_strings = append(matched_strings,line)
						for _,element := range getFollowing(line_counter,Previous_lines,lines) {
							matched_strings = append(matched_strings,element)
						}

						repo.Matches = append(repo.Matches,custom_datatypes.Match {
							hash.String(),
							"Commit Message",
							rule,
							matched_strings,
						})
					}
				}
			}
		}
	}
}

func FindMatchOnFile(file *object.File,regex []string,repo * custom_datatypes.Repo){
	for _, rule := range regex {
		r, _ := regexp.Compile(rule)
		content,_ := file.Contents()
		if r.MatchString(content){
			lines := strings.Split(content,"\n")
			for line_counter,line := range lines {
				if r.MatchString(line) {
					var sum string
					if line_counter > 0 && line_counter < (len(lines)-2){
						md5HashInBytes := md5.Sum([]byte(lines[line_counter-1]+line+lines[line_counter+1]))
						sum = hex.EncodeToString(md5HashInBytes[:])
					}else if line_counter <= 0 && line_counter < (len(lines)-2){
						md5HashInBytes := md5.Sum([]byte(line+lines[line_counter+1]))
						sum = hex.EncodeToString(md5HashInBytes[:])
					}else if line_counter > 0 && line_counter == (len(lines)-1){
						md5HashInBytes := md5.Sum([]byte(lines[line_counter-1]+line))
						sum = hex.EncodeToString(md5HashInBytes[:])
					}else {
						md5HashInBytes := md5.Sum([]byte(line))
						sum = hex.EncodeToString(md5HashInBytes[:])
					}
					if repo.AlreadyFound[sum] == false {
						repo.AlreadyFound[sum] = true
						
						matched_strings := []string{}

						for _,element := range getPrevious(line_counter,Previous_lines,lines) {
							matched_strings = append(matched_strings,element)
						}
						matched_strings = append(matched_strings,line)
						for _,element := range getFollowing(line_counter,Previous_lines,lines) {
							matched_strings = append(matched_strings,element)
						}

						repo.Matches = append(repo.Matches,custom_datatypes.Match {
							file.Hash.String(),
							file.Name,
							rule,
							matched_strings,
						})
					}
				}
			}
		}
	}
}

func getPrevious(line_number int,count int,lines []string) []string{
	var data []string
	for i := (line_number-count); i < line_number && i < len(lines); i++ {
		if i >= 0 {
			if len(lines[i]) > 0 {
				data = append(data, lines[i])
			}
		}
	}
	return data
}

func getFollowing(line_number int,count int,lines []string) []string{
	var data []string
	for i := (line_number+1); i <= (line_number+count) && i < len(lines); i++ {
		if i < len(lines) && i>=0 {
			if len(lines[i]) > 0 {
				data = append(data, lines[i])
			}
		}
	}
	return data
}