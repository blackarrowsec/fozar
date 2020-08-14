package log

import (
	"fmt"
	"github.com/fatih/color"
)

var (
	Debug bool = false
	Out_File string = ""

	red = color.New(color.FgRed).PrintfFunc()
	green = color.New(color.FgGreen).PrintfFunc()
	blue = color.New(color.FgBlue).PrintfFunc()
	success = color.New(color.Bold, color.FgGreen).PrintfFunc()
)


func Print(data string){
	if Debug {
		DebugPrint(data)
	}
}

func PrintBanner(){
	fmt.Printf(
		" _____\n"+                      
		"|  ___|___  ____ __ _  _ __\n"+
		"| |_  / _ \\|_  // _` || '__|\n"+
		"|  _|| (_) |/ /| (_| || |\n"+  
			"|_|   \\___//___|\\__,_||_|\n")
			fmt.Printf("\t\t\tBy @30vh1 [https://blackarrow.net] [https://tarlogic.com]\n")
		}

func PrintRepoName(name string){
	red("## %s\n",name)
}

func DebugPrint(data string){
	if Debug {
		green("[DEBUG] ")
		fmt.Println(data)
	}
}


func Fatal(err error){
	if Debug {
		red("[ERROR] ")
		fmt.Println(err)
	}
	panic(err)
}

func PrintError(err error){
	if Debug {
		red("[ERROR] ")
		fmt.Println(err)
	}
}

func PrintCommitInfo(hash string,file string){
	green("\n# %s\t",hash)
	blue("%s\n",file)
}

func RawPrint(data string){
	fmt.Println(data)
}

func FoundPrint(data string){
	success("%s\n",data)
}

func handleError(err error){
	if err != nil {
		PrintError(err)
		return
	}
}
