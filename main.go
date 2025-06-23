package main

import (
	// "fmt"

	"github.com/Meh-Mehul/mailing-service/utils"
	// "github.com/joho/godotenv"
	"gofr.dev/pkg/gofr"
)

func main() {
    // err := godotenv.Load()
    // if err != nil {
    //     fmt.Println("Error loading .env file")
    //     return
    // }
	rdb := utils.InitRedis();
	go utils.StartWorker(rdb);
	app := gofr.New();
	app.POST("/send", utils.HandleSend(rdb));
	app.Run();

}