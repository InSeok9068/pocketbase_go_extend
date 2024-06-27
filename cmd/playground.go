package main

import (
	"fmt"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"log"
)

func main() {
	app := pocketbase.New()

	app.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {
		user, err := app.Dao().FindAuthRecordByEmail("users", "dlstjr9068@gmail.com")
		fmt.Print(user)
		fmt.Print(err)
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
