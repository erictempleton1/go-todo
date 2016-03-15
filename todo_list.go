package main

import (
    "os"
    "time"
    "github.com/codegangsta/cli"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Todos struct {
    Name string
    Created time.Time
    Done bool
}


func main() {
    app := cli.NewApp()
    app.Name = "boom"
    app.Usage = "make an explosive entrance"
    app.Action = func(c *cli.Context) {
        println("Hello", c.Args()[0])
    }

    app.Run(os.Args)
}