package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "idg",
		Usage: "id gen",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "uuidv4",
				Usage:   "uuidv1 or uuidv4 or ulid",
			},
		},
		Action: func(c *cli.Context) error {
			format := c.String("format")
			g, err := GetGenerator(format)
			if err != nil {
				return err
			}
			id, err := g.Generate()
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", id)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
