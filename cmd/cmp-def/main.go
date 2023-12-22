package main

import (
	"flag"
	"fmt"
	"os"

	"kpawlik.pl/cdef"
)

var (
	featureName string
	dir1 string
	dir2 string
	printDefs1 bool
	printDefs2 bool
)

func init(){
	var help bool
	flag.BoolVar(&help, "h", false, "Print help")
	flag.StringVar(&featureName, "file", "", "name of file to compare")
	flag.StringVar(&dir1, "dir1", "", "Path to dir1")
	flag.StringVar(&dir2, "dir2", "", "Path to dir2")
	flag.BoolVar(&printDefs1, "print1", false, "Print field")
	flag.BoolVar(&printDefs2, "print2", false, "Print field")
	flag.Parse()
	if help{
		fmt.Printf("Version %s\n", cdef.VERSION)
		fmt.Println(`Params -dir1 and -dir2 are mandatory. Will compare all files from dir1 against dir2
Param -file is optional. If provided, then will try to compare only this file.
Parameters:`)
		flag.PrintDefaults()
		os.Exit(0)
	}
	if len(dir1)==0 || len(dir2)==0{
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main(){
	var(
		err error
		diffs []cdef.Diff
		
	)
	fmt.Println(cdef.VERSION)
	if len(featureName) > 0{
		var diff cdef.Diff
		if diff, err = cdef.Compare(dir1, dir2, featureName); err != nil{
			fmt.Println(err)
			return	
		}
		diff.Print(printDefs1, printDefs2)
		return;
		
	}
	if diffs, err = cdef.CompareDirs(dir1, dir2); err != nil{
		fmt.Println(err)
		return
	}
	cdef.PrintDiff(printDefs1, printDefs2, diffs...)
	
}