package main

import (
	"fmt"
)

func main() {
	var c = cron.New()

	defer c.Stop()

	c.AddFunc("* * * * * *",
		func() {
			fmt.Println("Me muestro cada segundo")
		},
	)

	c.Start()

}
