package main

import (
  "fmt"
  "os"

  "github.com/learnsecurely/ansikube"
  "github.com/caarlos0/spin"
  "github.com/urfave/cli"
)

var version = "master"

func main() {
  app := cli.NewApp()
  app.Name = "example"
  app.Version = version
  app.Author = "Jeremy Pruitt (jeremy@learnsecurely.com)"
  app.Usage = "This is an example app written in Go"
  app.Action = func(c *cli.Context) error {
    spin := spin.New("\033[36m %s Working...\033[m")
    spin.Start()
    err := ansikube.Foo()
    spin.Stop()
    if err != nil {
      return err
    }
    fmt.Println("Done!")
    return nil
  }
  _ = app.Run(os.Args)
}
