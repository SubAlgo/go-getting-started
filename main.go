package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

/*
require (
	github.com/gin-gonic/gin v1.7.4
	github.com/heroku/x v0.0.33
)


 channelSecret: "af42dc438bbcf308b5b0d274b4e1846e"
  channelAccessToken: "CX7kjGwq6ASjy3wd2SRihDD4XhlEzVKbTQ07JIUqGhNhXHuQwJ1L9NdP80uvSpqFz7qpmsdSQO0r9HmvEITCUGoy4j/zJWxwx09+5P8Mklzbo1H2FBnrrPXYx3iFhl+iZU74LMu0q8HEpQCj/vk1DgdB04t89/1O/w1cDnyilFU="
  yourLineID: "Ue56654f7c4a75297a0ecd19695fef261"
*/
func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK1",
		})
	})

	router.POST("/webhok", func(ctx *gin.Context) {
		bot, err := linebot.New("af42dc438bbcf308b5b0d274b4e1846e", "CX7kjGwq6ASjy3wd2SRihDD4XhlEzVKbTQ07JIUqGhNhXHuQwJ1L9NdP80uvSpqFz7qpmsdSQO0r9HmvEITCUGoy4j/zJWxwx09+5P8Mklzbo1H2FBnrrPXYx3iFhl+iZU74LMu0q8HEpQCj/vk1DgdB04t89/1O/w1cDnyilFU=")
		if err != nil {
			log.Println(err)
		}
		events, err := bot.ParseRequest(ctx.Request)
		if err != nil {
			log.Println(err)
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				log.Println("userID: ", event.Source.UserID)
				log.Println("groupID: ", event.Source.GroupID)
				log.Println("RoomID: ", event.Source.RoomID)
				log.Println("ReplyToken: ", event.ReplyToken)
				//lineMsg := event.Message
				log.Println("line_message:", event.Message)

				msg := linebot.NewTextMessage("test_message: ")
				_, err := bot.PushMessage(event.Source.GroupID, msg).Do()
				if err != nil {
					log.Println(err)
				}
			}
		}
	})

	router.Run(":" + port)
}
