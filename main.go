package main

import "github.com/tkanos/gonfig"

type Configuration struct {
	Port         string
	Pass         string
	DatabaseName string
	UserName     string
}

func main() {
	configuration := Configuration{}
	err := gonfig.GetConf("./Configuration.json", &configuration)

	if err != nil {

	} else {
		a := App{}
		// You need to set your Username and Password here
		a.Initialize(configuration.UserName, configuration.Pass, configuration.DatabaseName)
		a.Run(configuration.Port)
	}
}
