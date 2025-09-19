package grpc_clients

import (
	"os"
	pb "ride-sharing/shared/proto/trip"

	"google.golang.org/grpc"
)

type tripServiceClient struct {
	client pb.TripServiceClient
	conn   *grpc.ClientConn
}

func NewTripServiceClient() (*tripServiceClient, error) {
	tripServiceURL := os.Getenv("TRIP_SERVICE_URL")
	if tripServiceURL == "" {
		tripServiceURL = "localhost:50053"
	}
	conn, err := grpc.NewClient(tripServiceURL)
	if err != nil {
		return nil, err
	}
	client := pb.NewTripServiceClient(conn)
	return &tripServiceClient{client: client, conn: conn}, nil
}
func (c *tripServiceClient) Close() {
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return
		}
	}
}
