# fileStoreService
A go  microservice for storing files to a file store with a command line application for performing different commands on the store
=============PART1==================================================================
Steps to run the application:
Install go in the environment/host from https://golang.org/doc/install
Setup the Path variables GOPATH and GOROOT
git clone repo https://github.com/Kartikdubey/fileStoreService.git
Execute go mod tidy
Start the server go run server/server.go
Option one---
From another terminal ensure mycommandLineApp command works by typing mycommandLineApp storeLs .IF some error is there copy the binary from mycommandLineApp folder to
envPath/bin folder
   RUN the following commands as per requiements to interact with server
   mycommandLineApp storeAdd pqr.txt      (to add a file,can even add multiple files)
   mycommandLineApp storeLs               (to list all the files)
   mycommandLineApp storeRm abc.txt       (to delete files)
   mycommandLineApp storeUpdate pqr.txt   (to update a file)
   mycommandLineApp storeWc               (to get total word counts from all the files)
   mycommandLineApp storefreqWords        (to get less or more frequent words)
Option two---
We can also send command  to server using one other option provided in server/cmd/cli.go  go run server/cmd/cli.go store wc ---This will return count of all the words
=====================================================================================



====================FOR BUILDING DOCKER IMAGE==========================================================
To build the image use below command from the same dir where this repo has been cloned: sudo docker build . -t go-docker

For running the container: sudo docker run --detach --name az go-docker:latest

For entering in the container environment sudo docker exec -it az /bin/bash

For creating kubernetes pod from the image kubectl create -f kb8s-server-api.yaml

======================TESTING IN PROCESS FOR PART 2=====================================================
