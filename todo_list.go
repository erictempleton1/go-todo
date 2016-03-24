package main

import (
    "fmt"
    "log"
    "time"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

const (
    Host = "localhost"
    Database = "test_go"
)

type Todos struct {
    Name string
    Created time.Time
    Done bool
}

func main() {
    session, err := mgo.Dial(Host)
    if err != nil {
        panic(err)
    }

    defer session.Close()

    c := session.DB(Database).C("test_collection")
    err = c.Insert(&Todos{"Take out trash", time.Now(), false},
                   &Todos{"Clean up room", time.Now(), false})

    if err != nil {
        log.Fatal(err)
    }

    result := Todos{}
    err = c.Find(bson.M{}).One(&result)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Name:", result)
}
