package client

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// @TODO will enahance this going forward
type GRPCHandler interface {
}

type GRPCClient struct {
	Host       string
	ClientConn *grpc.ClientConn
}

func NewGRPCClient(ctx context.Context) (*GRPCClient, error) {

	host := os.Getenv("AVIATRIX_GRPC_HOST")
	// Connect to the gRPC server
	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Failed to connect gRPC host: %v, error: %v", host, err))
		return nil, err
	}

	client := &GRPCClient{
		Host:       host,
		ClientConn: conn,
	}

	return client, err
}
