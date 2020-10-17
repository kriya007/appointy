package main

import (
	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

const (
	url = "localhost"
)

func main(){

	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	log.Infof("Connected to MongoDb", url)
}


