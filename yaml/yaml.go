package yaml

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

//解析文件，取出所有参数
func GetYamlConfig(yamlFile string) map[interface{}]interface{} {

	data, err := ioutil.ReadFile(yamlFile)
	//将解析出的参数转为map的形式
	m := make(map[interface{}]interface{})
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal([]byte(data), &m)

	return m
}
