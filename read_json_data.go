package ipvsGo

import (
	"encoding/json"
	"io/ioutil"
)

type IpRange struct {
	Start string `json: "start"`
	End	  string `json: "end"`
	Source  string `json: "source"`
}

func ReadDataFromJson(filepath string) []*IpRange {
	var r []*IpRange
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		println("文件读取错误", err)
		return nil
	}
	err = json.Unmarshal(data, &r)
	if err != nil {
		println("数据解析错误", err)
		return nil
	}
	return r
}

