package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func init() {
	RootCmd.AddCommand(ServerCmd)
}

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start turtl server",
	Long:  `Start turtl server`,
	Run: func(cmd *cobra.Command, args []string) {
		err := validateServerArgs(args)
		if err != nil {
			log.Fatal(err)
		} else {
			var wg sync.WaitGroup
			config := newServerConfig()
			wg.Add(1)
			go startServer(config, &wg)
			wg.Wait()
		}
	},
}

func validateServerArgs(args []string) error {
	return nil
}

type serverConfig struct {
	port uint
}

func newServerConfig() *serverConfig {
	return &serverConfig{
		port: 50010,
	}
}

func startServer(config *serverConfig, wg *sync.WaitGroup) {
	defer wg.Done()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	RegisterTurtlServer(grpcServer, newTurtlServer())
	log.Println("started turtl RPC server")
	grpcServer.Serve(lis)
}

type turtlServer struct {
}

func (gs *turtlServer) Command(ctx context.Context, command *Request) (*Response, error) {
	log.Printf("received command: \"%s\"", command.Command)

	response := Response{}

	return &response, nil
}

func newTurtlServer() *turtlServer {
	s := new(turtlServer)
	return s
}
