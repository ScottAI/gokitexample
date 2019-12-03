package grpc

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	vault "gokitexample"
	"gokitexample/pb"
	"google.golang.org/grpc"
)

// New makes a new gokitexample.Service client.
func New(conn *grpc.ClientConn) vault.Service {
	var hashEndpoint = grpctransport.NewClient(
		conn, "Vault", "Hash",
		vault.EncodeGRPCHashRequest,
		vault.DecodeGRPCHashResponse,
		pb.HashResponse{},
	).Endpoint()
	var validateEndpoint = grpctransport.NewClient(
		conn, "Vault", "Validate",
		vault.EncodeGRPCValidateRequest,
		vault.DecodeGRPCValidateResponse,
		pb.ValidateResponse{},
	).Endpoint()
	return vault.Endpoints{
		HashEndpoint:     hashEndpoint,
		ValidateEndpoint: validateEndpoint,
	}
}
