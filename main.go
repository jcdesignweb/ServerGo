package main

import (
	"ServerGo/app"
	"ServerGo/boot"
)

/**
 * @author Juan Andres Carmena
 */
func main() {
	// init server
	app.Init()

	// start server
	app.Server.Logger.Fatal(app.Server.Start(boot.App.Port))

}