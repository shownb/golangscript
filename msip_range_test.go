package main

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
)

//ip来自https://www.microsoft.com/en-us/download/confirmation.aspx?id=41653
//转换http://dfsq.github.io/xml2json
type msjson struct {
	AzurePublicIpAddresses struct{
		Region []struct{
			Name string `json:"@Name"`
			IpRange []struct{
				Subnet string `json:"@Subnet"`
			}
		}
	}
}

func main() {
	s := new(msjson)
	file, err := os.Open("ms.json") // For read access.		
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	json.Unmarshal(data, s)
	//fmt.Printf("%+v",*s)
	for _,v:= range s.AzurePublicIpAddresses.Region{
		//fmt.Printf("%v\n",v.Name)
		for _,i:= range v.IpRange{
			fmt.Printf("%v\n",i.Subnet)
		}
	}
}
