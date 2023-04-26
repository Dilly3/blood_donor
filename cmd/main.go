package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/dilly3/blood-donor/database/mongodb"
	"github.com/dilly3/blood-donor/internal"
	"github.com/dilly3/blood-donor/web"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	mongoUrl := os.Getenv("MONGO_URL")
	mongoTimeout, err := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
	if err != nil {
		fmt.Println("strconv error ====> ", err)
	}
	fmt.Println("welcome to Blood donor App")
	mongoDb := mongodb.NewMongoDb(mongoUrl, mongoTimeout)
	impl := internal.NewServImpl(mongoDb)
	config := web.NewConfig(impl)

	gin := config.Serve()
	serv := &http.Server{
		Handler: gin,
		Addr:    ":8000",
	}
	c := make(chan os.Signal, 1)
	go func() {
		fmt.Println("running server")
		err := serv.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		signal.Notify(c, syscall.SIGINT)
	}()

	fmt.Println("shutting down server ===>", <-c)
}
