package main

import (
	"github.com/OpenCal-FYDP/EmailNotification/internal/service"
	"github.com/OpenCal-FYDP/EmailNotification/rpc"
	"log"
	"net/http"
)

func main() {
	service := service.New()
	server := rpc.NewEmailNotificationServiceServer(service)
	log.Fatal(http.ListenAndServe(":8080", server))
}
