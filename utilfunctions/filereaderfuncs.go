package utilfunctions

import (
	"bytes"
	"os"
	"io/ioutil"
)

func MakeJsonReader(path string) (reader *bytes.Reader, err error) {
	//jsonFile, err := os.Open(path)
	//if err != nil {
	//	return
	//}
	//defer jsonFile.Close()
	//jsonData, err := ioutil.ReadAll(jsonFile)
	jsonData, err := MakeJsonData(path)
	if err != nil {
		return
	}
	reader = bytes.NewReader(jsonData)
	return
}

func MakeJsonData(path string) (jsonData []byte, err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return
	}
	defer jsonFile.Close()
	jsonData, err = ioutil.ReadAll(jsonFile)
	return
}
