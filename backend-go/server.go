package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 6699
// Hash 7489
// Hash 7455
// Hash 6045
// Hash 6505
// Hash 2266
// Hash 1416
// Hash 4727
// Hash 8568
// Hash 4773
// Hash 9077
// Hash 9692
// Hash 5597
// Hash 8671
// Hash 7062
// Hash 5628
// Hash 5034
// Hash 6288
// Hash 1125
// Hash 4932
// Hash 5195
// Hash 6633
// Hash 9243
// Hash 2140
// Hash 7399
// Hash 8959
// Hash 7046
// Hash 1399
// Hash 1966