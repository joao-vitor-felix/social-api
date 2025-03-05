package main

import "log"

func main() {
	app := &application{
		config: config{
			address: ":8080",
		},
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
