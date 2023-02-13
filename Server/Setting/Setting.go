package Config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Conf struct {
	DBUser     string `yaml:"DBUser"`
	DBPassword string `yaml:"DBPassword"`
	DBHost     string `yaml:"DBHost"`
	DBPort     string `yaml:"DBPort"`
	DBName     string `yaml:"DBName"`
}

func (c *Conf) getConf() *Conf {
	yamlFile, err := ioutil.ReadFile("./Config/conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
