package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	token         = ""
	commandPrefix = "!"
)

func init() {
	token = os.Getenv("DISC_MSG_EMBEDDER_TOKEN")
	if token != "" {
		return
	}

	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
	if token != "" {
		return
	}

	log.Fatal("no token provided")
}

func main() {
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error creating the bot", err.Error())
		return
	}

	bot.AddHandler(checkMessage)

	err = bot.Open()
	if err != nil {
		log.Fatal("Error opening the websocket connection", err.Error())
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()
}

func checkMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, commandPrefix) {
		return
	}

	command := strings.Fields(m.Content[1:])

	if command[0] != "embed" {
		return
	}

	log.Println("received valid command", command)
	if len(command) == 1 {
		printTemplate(s, m.ChannelID)
		return
	}
	if len(command) > 2 {
		sendEmbed(s, command[1], strings.Join(command[2:], " "), m)
	}
}

func printTemplate(s *discordgo.Session, chanelID string) {
	log.Println("printing template")
	file, err := ioutil.ReadFile("embed_example.json")
	if err != nil {
		log.Println(err)
		return
	}
	s.ChannelMessageSend(chanelID, string(file))
}

func sendEmbed(s *discordgo.Session, chanelID, jsonString string, m *discordgo.MessageCreate) {
	log.Println("sending embedded message")
	log.Println("json string", jsonString)
	msg := new(discordgo.MessageSend)
	embed := new(discordgo.MessageEmbed)
	err := json.Unmarshal([]byte(jsonString), embed)
	if err != nil {
		log.Println(err)
		return
	}
	msg.Content = "Message from " + m.Author.Mention()
	msg.Embed = embed
	_, err = s.ChannelMessageSendComplex(chanelID, msg)
	if err != nil {
		log.Println(err)
	}
}
