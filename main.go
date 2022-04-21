package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Commands events")
		fmt.Println(event.Timestamp)
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-2918691710243-3416454394949-PqKT3dG86B6OST8bpYJDHmYI")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03CPRTC4LR-3419269522163-b3d0ac84b60468237a4585ed57353f7aa63229a639f8ad3c980da8a3d44b0221")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Example:     "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("error")
			}
			age := 2022 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	// need to know it better way
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
