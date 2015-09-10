package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"time"

	pb "github.com/harlow/go-micro-services/service.auth/proto"
	trace "github.com/harlow/go-micro-services/trace"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	port       = flag.Int("port", 10001, "The server port")
	jsonDBFile = flag.String("json_db_file", "data/customers.json", "A json file containing a list of customers")
	serverName = "service.auth"
)

// newServer creates a new authServer and loads the customers from
// JSON file into customers map
func newServer() *authServer {
	s := new(authServer)
	s.loadCustomers(*jsonDBFile)
	return s
}

// authServer struct w/ customers map
type authServer struct {
	customers map[string]*pb.Customer
}

// VerifyToken finds a customer by authentication token.
func (s *authServer) VerifyToken(ctx context.Context, args *pb.Args) (*pb.Customer, error) {
	md, _ := metadata.FromContext(ctx)

	t := trace.Tracer{TraceID: md["traceID"]}
	t.In(serverName, args.From)
	defer t.Out(args.From, serverName, time.Now())

	customer := s.customers[args.AuthToken]
	if customer == nil {
		return &pb.Customer{}, errors.New("Invalid Token")
	}

	return customer, nil
}

// loadCustomers loads customers from a JSON file.
func (s *authServer) loadCustomers(filePath string) {
	// open data file
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to load file: %v", err)
	}

	// unmarshal JSON
	customers := []*pb.Customer{}
	if err := json.Unmarshal(file, &customers); err != nil {
		log.Fatalf("Failed to unmarshal json: %v", err)
	}

	// create customer lookup map
	s.customers = make(map[string]*pb.Customer)
	for _, c := range customers {
		s.customers[c.AuthToken] = c
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
