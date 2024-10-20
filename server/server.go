package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"

	pb "JacobPaerre/Security-handin-2/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)


type Hospital struct {
	pb.UnimplementedAggregationSendingServiceServer
	id              int
	hospitalAddress string
	receivedShares  []int
}

func hospitalAggregateShares(receivedShares []int) int {
	sum := 0
	for _, share := range receivedShares {
		sum += share
	}
	return sum
}

func (h *Hospital) SendAggregation(ctx context.Context, req *pb.Aggregation) (*pb.Acknowledge, error) {
	log.Printf("Share from %d: %d\n", req.SenderId, req.Aggregation)
	h.receivedShares = append(h.receivedShares, int(req.Aggregation))

	if len(h.receivedShares) % 3 == 0 {
		aggregatedShares := hospitalAggregateShares(h.receivedShares)
		log.Printf("Total share sum: %d\n", aggregatedShares)
	}

	return &pb.Acknowledge{
		ReceiverId: int32(h.id),
		Message:    "Share received",
	}, nil
}

func runHospitalServer(hospital *Hospital) {
	// Set up a listener on the hospital's address
	lis, err := net.Listen("tcp", hospital.hospitalAddress)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", hospital.hospitalAddress, err)
	}

	// Load certificates
	tlsCredentials, err := loadTLSCredentials("server-cert.pem", "server-key.pem")
	if err != nil {
		log.Fatalf("Failed to load TLS credentials: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(tlsCredentials),
	)
	pb.RegisterAggregationSendingServiceServer(grpcServer, hospital)
	log.Printf("Hospital server running at %s", hospital.hospitalAddress)

	// Start serving incoming gRPC requests
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over %s: %v", hospital.hospitalAddress, err)
	}
}

func loadTLSCredentials(certPath, keyPath string) (credentials.TransportCredentials, error) {
    // Load server's certificate and private key
    serverCert, err := tls.LoadX509KeyPair(certPath, keyPath)
    if err != nil {
        return nil, err
    }

    // Create the credentials and return it
    config := &tls.Config{
        Certificates: []tls.Certificate{serverCert},
        ClientAuth:   tls.NoClientCert,
    }

    return credentials.NewTLS(config), nil
}

func main() {
	hospital := &Hospital{
		id:              3,
		hospitalAddress: "localhost:3000",
		receivedShares:  []int{},
	}

	runHospitalServer(hospital)
}
