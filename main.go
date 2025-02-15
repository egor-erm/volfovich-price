package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"price/data"

	"github.com/gin-gonic/gin"
)

var price string

func main() {
	go update()
	run()
}

func run() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	router := gin.New()

	router.GET("/", Handle)

	fmt.Println("Run!")
	if err := router.Run(":88"); err != nil {
		log.Println("error running web:", err)
	}
}

func Handle(c *gin.Context) {
	c.String(http.StatusOK, price)
}

func update() {
	// Формирование URL
	url := fmt.Sprintf("https://api.dexscreener.com/latest/dex/pairs/%s/%s", "solana", "7efsbtclckyxje67rbwqazxcfsb6vbyr8cns5hvpscx4")

	// Создание нового HTTP-клиента
	client := &http.Client{}

	for {
		// Выполнение GET-запроса
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Установка заголовков
		req.Header.Set("Accept", "*/*")

		// Отправка запроса
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Чтение ответа
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error read:", err)
			return
		}

		var root data.Root
		err = json.Unmarshal(body, &root)
		if err != nil {
			fmt.Println("Error json:", err)
			return
		}

		price = root.Pairs[0].PriceUSD

		resp.Body.Close()
		time.Sleep(30 * time.Second)
	}
}
