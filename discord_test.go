package main

import (
	"DiscordRolesBot/internal/Scheduled_work"
	"DiscordRolesBot/internal/async_work"
	"DiscordRolesBot/pkg/discord"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func init() {
	initConfig()
	initLogger()
	initMongo()
	initDB()
	initRedis()
	Scheduled_work.ScheduledWork()
}

func TestDiscord(t *testing.T) {
	if err := discord.NewDiscord(); err != nil {
		t.Error(err)
	}

	user, err := discord.Discord.User("@me")
	if err != nil {
		t.Fatal(err.Error())
	}

	st, err := discord.Discord.UserGuilds(10, "", "")
	t.Log(len(st))
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("Guild:")
	for _, v := range st {
		t.Log(v.ID)
		t.Log(v.Name)
	}
	t.Log("User:")
	t.Log(user.ID)
	t.Log(user.Username)

	err = async_work.AddRoles("address_test:XXXXXXXXXXXXXXXXXXXXXXX")
	if err != nil {
		fmt.Println(err.Error())
	}
	t.Log("success")
	defer discord.Discord.Close()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stop
	log.Println("server shutdown......")
}
