package main

import ( "ServerGo/app")

/**
 * @author Juan Andres Carmena
 */
func main() {
	// init server
	app.Init()

	// start server
	app.Server.Logger.Fatal(app.Server.Start(":1313"))

}