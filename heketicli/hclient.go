//package main

package heketicli

import (
	"fmt"
	hclient "github.com/heketi/heketi/client/api/go-client"
	//"github.com/heketi/heketi/apps/glusterfs"
	//"crypto/sha256"
	//"encoding/hex"
	//"github.com/heketi/heketi"

	//jwt "github.com/dgrijalva/jwt-go"
	//"github.com/heketi/heketi/apps/glusterfs"
	//"github.com/heketi/utils"
	//"net/http"
	//"io/ioutil"

	//"time"
)

var (
	HEKETI_HOST = "http://10.70.42.29:8080"
	HEKETI_USER = ""
	HEKETI_KEY  = ""
)

func Connect() {

	//func main() {

	fmt.Println("Heketi Client")

	heketi_handler := hclient.NewClient(HEKETI_HOST, HEKETI_USER, HEKETI_KEY)

	//heketi_handler := hclient.NewClient("10.70.1.40", "", "")
	if heketi_handler == nil {
		fmt.Println("Failed to initiate the heketi client")
	}

	fmt.Println(heketi_handler)
	hello := heketi_handler.Hello()
	if hello != nil {
		fmt.Println("Looks like heketi is not running")
	}
	fmt.Println("Heketi Server is listening")
	fmt.Println("Proceeding")
	clusters, _ := heketi_handler.ClusterList()
	fmt.Println(clusters)
	fmt.Println("End of heketi handler")

	/* Output

	[root@humbles-lap heketicli]# go run hclient.go
	Heketi Client
	&{http://10.70.42.29:8080   0xc820072180}
	Heketi Server is listening <nil>

	*/

}
