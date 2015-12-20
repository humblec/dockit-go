//
// Copyright (c) 2015
//
// Author "Humble Chirammal" < hchiramm@redhat.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"flag"
	"fmt"
	aploprov "github.com/humblec/aplo-prov/aploprov"
	heketicli "github.com/humblec/aplo-prov/heketicli"
	"os"
	"time"

	// Below repo have an error
	// client "github.com/mdevilliers/kubernetes/pkg/client"
	// k8s.io/kubernetes/pkg/util/parsers
	//../src/k8s.io/kubernetes/pkg/util/parsers/parsers.go:30: undefined: parsers.ParseRepositoryTag
)

var (
	APLO_VERSION = "1.0"
	configfile   string
	image        string
	mode         string
	showVersion  bool
)

func init() {
	flag.StringVar(&image, "image", "", "Docker Image Name")
	flag.StringVar(&mode, "mode", "", "Mode of Operation")
	flag.StringVar(&configfile, "config", "", "Configuration file")
	flag.BoolVar(&showVersion, "version", false, "Show version")
}

func main() {

	fmt.Println("Info: Aplo Provisioner ")

	flag.Parse()

	/*
	   if image == " " {
	       fmt.Println("No image provided, working on gluster/gluster-centos")
	       image := "gluster-centos"
	   }
	*/

	if mode == "docker" {
		fmt.Println("Selected Mode: Docker ")
		go aploprov.Dockermode()
		time.Sleep(10000 * time.Millisecond)
		os.Exit(1)

	}

	if mode == "kube" {
		fmt.Println("Selected Mode: Kubernetes ")
		go aploprov.Kubemode()
		time.Sleep(10000 * time.Millisecond)
		os.Exit(1)

	}

	heketicli.Connect()
}
