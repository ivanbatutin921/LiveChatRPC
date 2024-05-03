package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	pb "github.com/ivanbatutin921/LiveChatGrpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	
)

type Server struct {
	app   *fiber.App
	proto pb.LiveChatClient
	port  string
}

func (s *Server) runGrpcServer() {
	conn, err := grpc.Dial("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("не получилось соединиться: %v", err)
	}
	s.proto = pb.NewLiveChatClient(conn)
}

func (s *Server) Run() {
	s.app = fiber.New()
	s.runGrpcServer()
	log.Fatal(s.app.Listen(":" + s.port))
}

func main() {
	s := &Server{port:os.Getenv("PORT")}
	s.Run()
}
