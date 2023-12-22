package cdef

import (
	"fmt"
	"strings"
)

type Diff struct {
	dir1, dir2, name string
	Fields1          []Field
	Fields2          []Field
	TypeDiff         map[string][]string
	Error            error
}

func (d Diff) Empty() bool {
	return len(d.Fields1) == 0 && len(d.Fields2) == 0 && len(d.TypeDiff) == 0
}

func (d Diff) Print(printDefs1 bool, printDefs2 bool) {
	fmt.Printf("File name: %s\n", d.name)
	fmt.Printf("%s\n", strings.Repeat("=", 81))
	if err := d.Error; err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("%40s|%40s\n", d.dir1, d.dir2)
	fmt.Printf("%s+%s\n", strings.Repeat("-", 40), strings.Repeat("-", 40))
	for _, field := range d.Fields1 {
		fmt.Printf("%40s|%40s\n", field.Name(), "-")
	}
	if printDefs1 {
		for _, field := range d.Fields1 {
			printDef(field)
			// fmt.Printf("{\n")
			// for k,v := range field{
			// 	fmt.Printf("   \"%s\": \"%v\",\n", k, v)
			// }
			// fmt.Printf("}\n")
		}
	}
	for _, field := range d.Fields2 {
		fmt.Printf("%40s|%40s\n", "-", field.Name())
	}
	if printDefs2 {
		for _, field := range d.Fields2 {
			// fmt.Printf("{\n")
			// for k,v := range field{
			// 	fmt.Printf("   \"%s\": \"%v\",\n", k, v)
			// }
			// fmt.Printf("}\n")
			printDef(field)
		}
	}
	if len(d.TypeDiff) > 0 {
		fmt.Printf("Type changes:\n")
		fmt.Printf("%s+%s\n", strings.Repeat("-", 40), strings.Repeat("-", 40))
		for name, types := range d.TypeDiff {
			fmt.Printf("%40s|%40s\n", fmt.Sprintf("%s(%s)", name, types[0]), fmt.Sprintf("%s(%s)", name, types[1]))
		}
	}
}

func printDef(field Field) {
	intend := "   "
	strs := make([]string, 0, len(field))
	for k, v := range field {
		strs = append(strs, fmt.Sprintf("%s\"%s\": \"%v\"", intend, k, v))
		
	}
	def := fmt.Sprintf("{\n%s\n},", strings.Join(strs, ",\n"))
	fmt.Println(def)
}
