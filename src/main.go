package main

import (
	"../src/conf"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	"fmt"
)

func main() {

	conf := conf.GetConf("conf/trystack.json")

	fmt.Println(conf)

	opts:= gophercloud.AuthOptions{
		IdentityEndpoint : conf.IdentityEndpoint,
		Username : conf.Username,
		Password: conf.Password,
		TenantID: conf.TenantID,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		panic(err)
	}
	cli, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: "nova",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(cli.TokenID)
}
