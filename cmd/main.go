package main

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	test := flag.Bool("test", false, "Test mode")
	testEnvFile := flag.String("testEnvFile", "test.env", ".env file with testing bindings")

	flag.Parse()

	if *test {
		err := godotenv.Load(*testEnvFile)
		if err != nil {
			log.Fatalf("Error loading %s file\n", *testEnvFile)
		}
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Start()
}
