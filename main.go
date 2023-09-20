package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

/**
Made by your ultimate engineer, with love!
*/
func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	} 
}

func main() {
	// These tokens are disbled
	os.Setenv("SLACK_BOT_TOKEN","xoxb-5905014337303-5932238586081-xCLQXUzYmsTM30o4B1ge0WLb")
	os.Setenv("SLACK_APP_TOKEN","xapp-1-A05T403LRN0-5919461077571-31c075bd25bc5ad3552d16c0b5c1da4cf109de50bd19e827a55fd3d47a516637")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"),os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())
	
	bot.Command("My yob is <year>", &slacker.CommandDefinition {
		Description: "yob calculator",
		Examples: []string{"my yob is 2022"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2023 - yob
			r := fmt.Sprintf("age is %d ", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

} 
