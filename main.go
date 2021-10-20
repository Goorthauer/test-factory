package main

import (
	"factory/Fight"
	"factory/MonsterFactory/Factory"
	"factory/MonsterFactory/Monster/Orc"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	app := &cli.App{
		Name:  "factory",
		Usage: "check factory",
		Action: func(c *cli.Context) error {
			fmt.Println("You need more parameters")
			return nil
		},
		Commands: []cli.Command{
			{
				Name:  "create-orc",
				Usage: "create orc x(count orcs)",
				Action: func(c *cli.Context) error {
					argOne, _ := strconv.Atoi(c.Args().First())
					for i := 1; i <= argOne; i++ {
						fmt.Printf("-----------------i'm %v orc-------------------\n", i)
						fmt.Printf("Создаю орка.\n")
						orc := Orc.CreateSimple()
						orc.Say()
						fmt.Printf("-----------------i'm %v orc-------------------\n\n", i)

					}
					return nil
				},
			},
			{
				Name:  "fight-basic",
				Usage: "create basic fight",
				Action: func(c *cli.Context) error {
					Fight.Fight()
					return nil
				},
			},
			{
				Name:  "create-basic",
				Usage: "create basic mode",
				Action: func(c *cli.Context) error {
					basic := new(Factory.Basic)
					orc := basic.CreateOrc()
					zombie := basic.CreateZombie()
					orc.Say()
					zombie.Say()
					return nil
				},
			},
			{
				Name:  "create-hard",
				Usage: "create hard mode",
				Action: func(c *cli.Context) error {
					basic := new(Factory.Hard)
					orc := basic.CreateOrc()
					zombie := basic.CreateZombie()
					orc.Say()
					zombie.Say()
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
