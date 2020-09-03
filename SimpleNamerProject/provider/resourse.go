package provider

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type NamesResource struct {
	LastNames  []string `yaml:"lastNames"` //struct Tag本身是一個字符串，但字符串中卻是：以空格分隔的 key:value
	FirstNames []string `yaml:"firstNames"`
}

func GetNames(file string) (ns NamesResource, err error) {
	rc, err := ioutil.ReadFile(file)
	if err == nil {
		err = yaml.Unmarshal(rc, &ns)
	}
	return ns, err

}
