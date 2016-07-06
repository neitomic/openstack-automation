package conf

import (
	"io/ioutil"
	"encoding/json"
)

type Conf struct {
	IdentityEndpoint string
	Username string
	Password string
	TenantID string
}

func GetConf(jsonFile string) *Conf {
	conf := &Conf{}
	confJson, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic("Read json configuration failed")
	}
	if err = json.Unmarshal(confJson, &conf); err != nil {
		panic("Unmarshal json failed")
	}

	return conf
}
