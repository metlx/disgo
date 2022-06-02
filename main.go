package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!m ︻デ-═一" {
		// because
		s.ChannelMessageSend(m.ChannelID, "that's dad")
	}

	if m.Content == "!spectrum" {
		// tribute to a good person
		s.ChannelMessageSend(m.ChannelID, "blazing fast!!!!")
	}

	if m.Content == "!editor" {
		// link to my editor
		s.ChannelMessageSend(m.ChannelID, "https://github.com/hlissner/doom-emacs")
	}

	if m.Content == "!meme" {
		// random memes
		s.ChannelMessageSend(m.ChannelID, rUrl())
	}

	if m.Content == "!gn" {
		// goodnight meme
		s.ChannelMessageSend(m.ChannelID, "https://memesbams.com/wp-content/uploads/2017/09/best-funny-goodnight-memes.jpg")
	}

	if m.Content == "!what?" {
		// rick roll
		s.ChannelMessageSend(m.ChannelID, "(https://www.tomorrowtides.com/secret887.html)")
	}

	if m.Content == "!secret" {
		// rick roll again
		s.ChannelMessageSend(m.ChannelID, "||omfg look at this https://www.tomorrowtides.com/secret887.html||")
	}

	if m.Content == "!︻デ-═一" {
		// sniper
		s.ChannelMessageSend(m.ChannelID, "pew")
		s.ChannelMessageSend(m.ChannelID, "pew")
		s.ChannelMessageSend(m.ChannelID, "pew")
	}

	if m.Content == "!commands" {
		// list all commands
		s.ChannelMessageSend(
			m.ChannelID,
			fmt.Sprint(
				"!what: what is this? \n",
				"!secret: tells you a secret. \n",
				"!︻デ-═一: shoots everything. \n",
				"!editor: shows m's editor. \n",
				"!gn: goodnight meme. \n",
				"!meme: random meme. \n",
				"!spectrum: the OG. \n",
			),
		)
	}
}

func rNum() int {
	rand.Seed(time.Now().UnixNano())
	rN := rand.Intn(13)
	return rN
}

func rUrl() string {
	urls := []string{
		"https://memecomplete.com/edit/https://api.memegen.link/images/gandalf/_/what_did_i_just_say~q.jpg",
		"https://memecomplete.com/share/https://api.memegen.link/images/ggg.jpg",
		"https://memecomplete.com/share/https://api.memegen.link/images/mb/'member/star_wars~q.jpg",
		"https://memecomplete.com/share/https://api.memegen.link/images/noidea/i_have_no_idea/what_i'm_doing.jpg",
		"https://memecomplete.com/share/https://api.memegen.link/images/rollsafe/can't_get_fired/if_you_don't_have_a_job.jpg",
		"https://memecomplete.com/share/https://api.memegen.link/images/aag/_/aliens.jpg",
		"https://memecomplete.com/share/https://api.memegen.link/images/sohappy/if_i_could_use_this_meme/i_would_be_so_happy.jpg",
		"https://memecomplete.com/share/https://api.memegen.link/images/ll/_/hhhehehe.jpg",
		"https://memecomplete.com/share/https://api.memegen.link/images/mmm.jpg",
		"https://memecomplete.com/share/https://api.memegen.link/images/keanu.jpg",
		"https://memecomplete.com/share/https://api.memegen.link/images/leo.jpg",
		"https://memecomplete.com/share/https://api.memegen.link/images/interesting.jpg",
		"https://memecomplete.com/share/https://api.memegen.link/images/philosoraptor.jpg",
	}
	return urls[rNum()]
}
