package cdef

import (
	"fmt"
	"os"
	"path"
	"strings"
)




func CompareDirs(dir1, dir2 string) (diffs []Diff, err error){
	var(
		dirs []os.DirEntry
		diff Diff
		
	)
	if dirs, err = os.ReadDir(dir1); err != nil{
		err = fmt.Errorf("cannot read directory d1: %s %w", dir1, err)
		return
	}
	for _, dir := range dirs{
		if dir.IsDir(){
			continue
		}
		name := dir.Name()
		path2 := path.Join(dir2, name)
		if _, fileErr := os.Stat(path2); os.IsNotExist(fileErr){
			fmt.Printf("\nFile name: %s - Does not exists in %s\n", name, dir2)
			fmt.Printf("%s\n",  strings.Repeat("=", 81))
			continue
		}
		if diff, err = Compare(dir1, dir2, name); err != nil{
			fmt.Printf("---- %s -----\n", name)
			fmt.Printf("Error comparing files %s\n", name)
			continue
		}
		if !diff.Empty(){
			diffs = append(diffs, diff)
		}

	}
	return
}

func PrintDiff(printDefs1,printDefs2 bool, diffs... Diff){
	for _, diff := range diffs{
		diff.Print(printDefs1, printDefs2)
	}
}

func Compare(dir1, dir2, fileName string)(diff Diff,err error){
	var (
		feature1 Feature
		feature2 Feature
	)
	
	path1 := path.Join(dir1, fileName)
	path2 := path.Join(dir2, fileName)
	if feature1, err = Load(path1); err != nil{
		err = fmt.Errorf("error loading %s (%w)", path1, err)
		return
	}
	if feature2, err = Load(path2); err != nil{
		err = fmt.Errorf("error loading %s (%w)", path2, err)
		return
	}
	diff = compareFeatures(feature1, feature2, dir1, dir2, fileName)
	
	return
}

func compareFeatures(f1 Feature, f2 Feature, dir1, dir2, fileName string)(diff Diff){
	diff = Diff{
		dir1: dir1,
		dir2: dir2,
		name: fileName,
	}
	diff.TypeDiff = make(map[string][]string)
	for _, field1 := range f1.Fields{
		name1 := field1.Name()
		type1 := field1.Type()
		found := false
		for _, field2 := range f2.Fields{
			name2 := field2.Name()
			type2 := field2.Type()
			if name1 == name2 && type1 == type2{
				found = true
				break
			}
			if name1 == name2 && type1 != type2{
				diff.TypeDiff[name1] = []string{type1,type2}
				found = true
				break
			}
		}
		if !found {
			diff.Fields1 = append(diff.Fields1, field1)
		}
	}
	for _, field2 := range f2.Fields{
		name2 := field2.Name()
		found := false
		for _, field1 := range f1.Fields{
			name1 := field1.Name()

			if name1 == name2{
				found = true
				break
				
			}
		}
		if !found {
			diff.Fields2 = append(diff.Fields2, field2)
		}
	}
	return
}