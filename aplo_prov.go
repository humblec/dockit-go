
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
    "fmt"
    "log"
    "os"
    "flag"
   // One other dockerclient
    "github.com/fsouza/go-dockerclient"
)

var (
    APLO_VERSION = "1.0"
   // configfile string
    image string
    showVersion bool

)

func init() {
    flag.StringVar(&image, "image", "", "Docker Image Name")
   // flag.StringVar(&configfile, "config", "", "Configuration file")
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

    fmt.Printf("Docker Image :%s", image)

    endpoint := "unix:///var/run/docker.sock"
   
    fmt.Println("\n Aplo Provisioner Connected to the Docker Deamon")
    
    client, _ := docker.NewClient(endpoint)
    
    //fmt.Println("Client is :", client)
    
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
        //fmt.Println("\t \t Created: ", img.Created)
        //fmt.Println("\t \t Size: ", img.Size)
        //fmt.Println("\t \t \tVirtualSize: ", img.VirtualSize)
        //fmt.Println("\t \t \t \tParentId: ", img.ParentID)

        //"docker.io/gluster/gluster-centos" 
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

    //fmt.Println(infoenv.Get("OperatingSystem"))
    
    docker_driver := infoenv.Get("Driver")
    if docker_driver != "devicemapper" {
        fmt.Println("The docker drivers other than the devicemapper is not supported at the moment ..exiting")
        os.Exit(1  )
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
    
    hostConfig := &docker.HostConfig {}
    
    containerConfig  := &docker.Config {
        Image:        "docker.io/gluster/gluster-centos",
        Cmd:   []string{"/sbin/init"},
        AttachStdin: true,
        Tty:          true,
        Entrypoint:   []string{"/sbin/init"},
        //ExposedPorts: map[string]struct{}{"22/tcp": {}

    }
   
    //client.CreateContainer(docker.CreateContainerOptions{"gluster-centos",containerConfig, hostConfig})
    gluster_container, err := client.CreateContainer(docker.CreateContainerOptions{"gluster-centos",containerConfig, hostConfig})
    //client.CreateContainer()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(gluster_container.ID)

    gluster_container_id := gluster_container.ID

    new_container := client.StartContainer(gluster_container_id, hostConfig)

    if new_container != nil {
        fmt.Println("Gluster Container Started with ID:",new_container)
    }
    //fmt.Println(containers)
    /*
    fmt.Printf("ID", containers.ID)
    */
    
    

}