package common

import (
	"encoding/json"
	"io/ioutil"
)

func SaveJson(filePath string) {
	ResultsMap.Lock()
	file, _ := json.MarshalIndent(ResultsMap, "", " ")
	_ = ioutil.WriteFile(filePath, file, 0644)
	ResultsMap.Unlock()
}
