package config

import (
  "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "strings"
  "path/filepath"
  "os"
)

func LoadConfigFromFile(file string) *Config {
     var fcfg = &Config{}

     yamlFile, err := ioutil.ReadFile(file)
     if err != nil {
         fmt.Println("yamlFile.Get err: ", err)
         panic(err)
     }

     err = yaml.Unmarshal(yamlFile, fcfg)
     if err != nil {
        fmt.Println("Unmarshal error: ", err)
        panic(err)
     }
     fmt.Println("Load config: ",fcfg.Version, ", Released: ", fcfg.Released)
     return fcfg
}

func LoadKey(root string,addr string) []byte {
    var files []string
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
               files = append(files, path)
               return nil
           })
    if err != nil {
         panic(err)
    }
    for _, file := range files {
         fmt.Println("File:", file)
         list := strings.Split(file,"--")
         if len(list) == 3 {
					   account := list[2]
						 if account == strings.TrimPrefix(addr,"0x") {
							 keyjson, err := ioutil.ReadFile(file)
							 if err != nil {
										fmt.Println("Error in read file: ", file )
										return nil
							 }
							 return keyjson
						 }
         }
    }
		return nil
}
