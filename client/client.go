package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
	//"math/rand"

	pb "JacobPaerre/Security-handin-2/proto"

	"google.golang.org/grpc"
)


type Patient struct {
	pb.UnimplementedShareSendingServiceServer
	id					int
    initialValue        int
	patientAddress		string
	localShare			int
	receivedShares		[]int
}

var (
    patientAddresses = map[int]string{
        0: "localhost:3001",
        1: "localhost:3002",
        2: "localhost:3003",
    }
)

func (p *Patient) generateShares() ([]int) {
    /*x := 10//rand.Intn(p.initialValue)
    y := 20//rand.Intn(p.initialValue - x)
    z := p.initialValue - x - y
    return []int{x, y, z}*/
	var shares []int
    switch p.id {
    case 0:
        shares = []int{30, 40, 30} // Sum = 100
    case 1:
        shares = []int{50, 70, 80} // Sum = 200
    case 2:
        shares = []int{150, 100, 50} // Sum = 300
    }
    return shares
}

func aggregateShares(receivedShares []int) int {
    sum := 0
    for _, share := range receivedShares {
        sum += share
    }
    return sum
}

func (p *Patient) SendShare(ctx context.Context, req *pb.Share) (*pb.Acknowledge, error) {
    log.Printf("Received share from %d: %d\n", req.SenderId, req.Share)
    p.receivedShares = append(p.receivedShares, int(req.Share))
    return &pb.Acknowledge{
        ReceiverId: int32(p.id),
        Message:    "Share received",
    }, nil
}

func sendSharesToOthers(peerAddress string, share int, senderId int) {
    // Connect to the peer via gRPC
    conn, err := grpc.Dial(peerAddress, grpc.WithInsecure()) // You can replace WithInsecure() with proper TLS credentials
    if err != nil {
        log.Fatalf("Failed to connect to peer: %v", err)
    }
    defer conn.Close()

    client := pb.NewShareSendingServiceClient(conn)
    _, err = client.SendShare(context.Background(), &pb.Share{
        SenderId: int32(senderId),
        Share:    int32(share),
    })
    if err != nil {
        log.Fatalf("Failed to send share: %v", err)
    }
}

func sendHospitalAggregation(hospitalAddress string, aggregation int, senderId int) {
    // Connect to the hospital via gRPC
    conn, err := grpc.Dial(hospitalAddress, grpc.WithInsecure()) // You can replace WithInsecure() with proper TLS credentials
    if err != nil {
        log.Fatalf("Failed to connect to hospital: %v", err)
    }
    defer conn.Close()

    client := pb.NewAggregationSendingServiceClient(conn)
    _, err = client.SendAggregation(context.Background(), &pb.Aggregation{
        SenderId:    int32(senderId),
        Aggregation: int32(aggregation),
    })
    if err != nil {
        log.Fatalf("Failed to send aggregation: %v", err)
    }
}

func runPatientServer(patient *Patient) {
    port := 3001 + patient.id
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
    if err != nil {
        log.Fatalf("Failed to listen on port %d: %v", port, err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterShareSendingServiceServer(grpcServer, patient)
    log.Printf("Starting gRPC server on port %d", port)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve gRPC server over %s: %v", patientAddresses[patient.id], err)
    }
}

func handleShares(patient *Patient, generatedShares []int) {
    // Each patient keeps its own share
    patient.receivedShares = append(patient.receivedShares, generatedShares[patient.id])
    time.Sleep(3 * time.Second)

    // Send shares to other patients
    for i, address := range patientAddresses {
        if i != patient.id {
            sendSharesToOthers(address, generatedShares[i], patient.id)
        }
    }
    time.Sleep(10 * time.Second)

    if len(patient.receivedShares) == 3 {
        aggregation := aggregateShares(patient.receivedShares)
        log.Printf("Patient %d aggregated value: %d\n", patient.id, aggregation)
        sendHospitalAggregation("localhost:3000", aggregation, patient.id)
    }
}

func main() {
	patientID := flag.Int("id", -1, "The patients ID")
    initialValue := 100 // flag.Int("init_val", -1, "The initial value")
    flag.Parse()

    port := 3001 + *patientID

    patient := &Patient{
        id:                 *patientID,
        initialValue:       initialValue,
        patientAddress:     fmt.Sprintf("localhost:%d", port),
        localShare:         0,
        receivedShares:     []int{},
    }

    generatedShares := patient.generateShares();

    go runPatientServer(patient)
    time.Sleep(10 * time.Second)

	handleShares(patient, generatedShares)
    time.Sleep(10 * time.Second)

}