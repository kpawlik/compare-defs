package cdef

import (
	"encoding/json"
	"os"
)


type Feature struct {
	Name string `json:"name"`
	ExternalName string  `json:"external_name"`
	Fields []Field `json:"fields"`
}

type Field map[string]interface{}


func (f Field) Name() string{
	return f["name"].(string)
}

func (f Field) Type() string{
	return f["type"].(string)
}


func Load(filePath string) (f Feature, err error){
	var (
		cont []byte
	)
	if cont, err = os.ReadFile(filePath); err != nil{
		return;
	}
	json.Unmarshal(cont, &f)
	return
}