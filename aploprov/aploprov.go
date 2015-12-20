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

//package main

package aploprov

import (
	//"flag"
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"github.com/kubernetes/kubernetes/pkg/api"
	client "github.com/kubernetes/kubernetes/pkg/client/unversioned"
	k8api "k8s.io/kubernetes/pkg/api"
	"log"
	"os"
	"time"
	// Below repo have an error
	// client "github.com/mdevilliers/kubernetes/pkg/client"
	// k8s.io/kubernetes/pkg/util/parsers
	//../src/k8s.io/kubernetes/pkg/util/parsers/parsers.go:30: undefined: parsers.ParseRepositoryTag
)

var (
	dockerImage  = "gluster/gluster-centos"
	dockerSocket = "unix:///var/run/docker.sock"
	kubeHost     = "http://10.70.42.184:8080"
	kubeService  = "gluster"
)

func Dockermode() {

	fmt.Printf("Docker Image :%s", dockerImage)

	endpoint := dockerSocket

	fmt.Println("\n Aplo Provisioner Connected to the Docker Deamon")

	client, _ := docker.NewClient(endpoint)

	if client == nil {
		fmt.Println("Failed to connect to the Docker Deamon.. exiting")
		os.Exit(1)
	}

	fmt.Println("Connected to Docker Deamon")

	imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})

	// If we want to list all the images
	//imgs, _ := client.ListImages(docker.ListImagesOptions{All: true})

	if imgs == nil {
		fmt.Println("Listimages: exiting")
		os.Exit(1)
	}

	for _, img := range imgs {
		fmt.Println("ID: ", img.ID)
		fmt.Println("\t RepoTags: ", img.RepoTags)
		// Other Attributes of the images if needed
		//fmt.Println("\t \t Created: ", img.Created)
		//fmt.Println("\t \t Size: ", img.Size)
		//fmt.Println("\t \t \tVirtualSize: ", img.VirtualSize)
		//fmt.Println("\t \t \t \tParentId: ", img.ParentID)

		// Gluster Image : "docker.io/gluster/gluster-centos"
		if img.RepoTags != nil {
			fmt.Println("Image: ", img.RepoTags)
		}

	}

	infoenv, _ := client.Info()
	if infoenv == nil {
		fmt.Println("Failed to Get Docker Info.. exiting")
		os.Exit(1)
	}
	fmt.Println(infoenv)

	// If we add support for specific OSs use below tag
	//fmt.Println(infoenv.Get("OperatingSystem"))

	docker_driver := infoenv.Get("Driver")
	if docker_driver != "devicemapper" {
		fmt.Println("The docker drivers other than the devicemapper is not supported at the moment ..exiting")
		os.Exit(1)
	}

	fmt.Println("List containers")

	containers, _ := client.ListContainers(docker.ListContainersOptions{All: false})

	for _, container := range containers {

		fmt.Println("\n Names:", container.Names)
		fmt.Println(" \t Container ID:", container.ID)
		if container.Image == "docker.io/gluster/gluster-centos" {
			fmt.Println("\n You already have a gluster Container running")

		}

	}

	fmt.Println("Create a Gluster Container")

	b := make(map[string][]docker.PortBinding)

	b["22/tcp"] = []docker.PortBinding{docker.PortBinding{HostPort: "22"}}

	hostConfig := &docker.HostConfig{}

	containerConfig := &docker.Config{
		Image:       "docker.io/gluster/gluster-centos",
		Cmd:         []string{"/sbin/init"},
		AttachStdin: true,
		Tty:         true,
		Entrypoint:  []string{"/sbin/init"},
		//ExposedPorts: map[string]struct{}{"22/tcp": {}

	}

	gluster_container, err := client.CreateContainer(docker.CreateContainerOptions{"gluster-centos", containerConfig, hostConfig})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(gluster_container.ID)

	gluster_container_id := gluster_container.ID

	new_container := client.StartContainer(gluster_container_id, hostConfig)

	if new_container != nil {
		fmt.Println("Gluster Container Started with ID:", new_container)
	}

	fmt.Println("End of Docker Handler")

	//fmt.Println(containers)
	/*
	   fmt.Printf("ID", containers.ID)
	*/
}

func Kubemode() {

	fmt.Println("Kubernetes ..Proceeding")

	config := client.Config{
		Host: kubeHost,
	}
	c, err := client.New(&config)

	if err != nil {
		log.Fatalln("Cant connect to Kubernetes API:", err)
	}

	s, err := c.Services(api.NamespaceDefault).Get(kubeService)

	if err != nil {
		log.Fatalln("Can't get service:", err)
	}

	fmt.Println("Name:", s.Name)

	fmt.Println(s.Spec)

	for p, _ := range s.Spec.Ports {
		fmt.Println("Port:", s.Spec.Ports[p].Port)
		fmt.Println("NodePort:", s.Spec.Ports[p].NodePort)
	}

	node := c.Nodes()

	fmt.Println("Nodes in your kubernetes Cluster")
	fmt.Println(node.List(k8api.ListOptions{}))
	fmt.Println("End of Kubernetes Handler")
	time.Sleep(1000 * time.Millisecond)

	/*
	   k8nodes,  := node.List(k8api.ListOptions{})
	   fmt.Println(k8nodes)
	   /*
	   for node := range k8nodes {
	       fmt.Println(k8node)
	   }*/

}

/*
func main() {
	go docker_mode()
	go kube()
	time.Sleep(1000 * time.Millisecond)

}

*/
