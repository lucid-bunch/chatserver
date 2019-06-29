package app

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"

	pb "chatserver/chatpb"

	"google.golang.org/grpc"
)

// App type
type App struct {
	Server   *grpc.Server
	Messages []*pb.Message
}

// NewApp constructs a new server app
func NewApp() *App {
	messages := make([]*pb.Message, 0, 100)
	return &App{Messages: messages}
}

// Listen binds and listens for connections on specified port
func (app *App) Listen(port int) {
	strport := strconv.Itoa(port)
	lis, err := net.Listen("tcp", ":"+strport)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	app.Server = grpc.NewServer()
	pb.RegisterChatServerServer(app.Server, app)

	log.Printf("Running gRPC server on port %s", strport)
	if err := app.Server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Close closes server
func (app *App) Close(sig os.Signal) {
	log.Printf("Caught signal (%s). Gracefully shutting down...", sig)
	app.Server.GracefulStop()
}

// SendMessage sends message
func (app *App) SendMessage(ctx context.Context, in *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	log.Printf("Message received: %s", in.GetMessage())
	app.Messages = append(app.Messages, in.GetMessage())
	return &pb.SendMessageResponse{Messages: app.Messages}, nil
}

// ReceiveMessages sends message
func (app *App) ReceiveMessages(ctx context.Context, in *pb.ReceiveMessagesRequest) (*pb.ReceiveMessagesResponse, error) {
	return &pb.ReceiveMessagesResponse{Messages: app.Messages}, nil
}
