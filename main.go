package main

import ( "ServerGo/app")

func main() {
	// init server
	app.Init()

	// start server
	app.Server.Logger.Fatal(app.Server.Start(":1313"))

}