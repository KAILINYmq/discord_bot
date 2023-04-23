package main

import (
	"DiscordRolesBot/global"
	"DiscordRolesBot/internal/Scheduled_work"
	"DiscordRolesBot/internal/async_work"
	"DiscordRolesBot/internal/routers"
	"DiscordRolesBot/pkg/discord"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	initConfig()
	initLogger()
	initMongo()
	initDB()
	initRedis()
	async_work.Jobs()
	Scheduled_work.ScheduledWork()
}

func main() {
	if err := discord.NewDiscord(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("discord is running......")
	}
	defer discord.Discord.Close()

	router := routers.NewRouter()
	srv := &http.Server{
		Addr:    global.Config.Base.Port,
		Handler: router,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			global.Logger.Error("Gin server start error:", err.Error())
			panic(err.Error())
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stop
	log.Println("server shutdown......")
}
