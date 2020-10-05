package IPWhitelist

import (
	"context"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func RegisterWhitelistServiceHandler(ctx context.Context, mux *mux.Router, conn *grpc.ClientConn) error {
	return nil
}


