package ipvsGo

import (
	"encoding/json"
	"io/ioutil"
)

type IpRange struct {
	Begin string `json: "begin"`
	End	  string `json: "end"`
	Name  string `json: "source"`
}

func ReadDataFromJson(filepath string) []*IpRange {
	var r []*IpRange
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(data, r)
	if err != nil {
		return nil
	}
	return r
}

