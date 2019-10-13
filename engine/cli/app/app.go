package app

import (
	pb "engine/lib/structs"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func AppInit(rpc string) *App {
	if len(rpc) > 0 {
		conn, connErr := grpc.Dial(rpc, grpc.WithInsecure())
		if connErr != nil {
			log.Fatalf("Did not connect: %v", connErr)
		}
		log.Printf("Connected to: %s ", rpc)

		app := App{
			ctx:       context.Background(),
			srvc_auth: pb.NewServiceAuthClient(conn),
			srvc_c:    pb.NewServiceCurrencyClient(conn),
			srvc_p:    pb.NewServiceUserClient(conn),
			srvc_a:    pb.NewServiceAccountClient(conn),
			srvc_ab:   pb.NewServiceAccountBalanceClient(conn),
			srvc_t:    pb.NewServiceTransactionClient(conn),
			srvc_o:    pb.NewServiceOrderClient(conn),
			srvc_tp:   pb.NewServiceTransactionProcessingClient(conn),
			srvc_cm:   pb.NewServiceCommissionClient(conn),
		}

		app.Conn = conn

		return &app
	}
	panic("rpc not set")
}

type App struct {
	ctx       context.Context
	srvc_auth pb.ServiceAuthClient
	srvc_c    pb.ServiceCurrencyClient
	srvc_p    pb.ServiceUserClient
	srvc_a    pb.ServiceAccountClient
	srvc_ab   pb.ServiceAccountBalanceClient
	srvc_t    pb.ServiceTransactionClient
	srvc_tp   pb.ServiceTransactionProcessingClient
	srvc_o    pb.ServiceOrderClient
	srvc_cm   pb.ServiceCommissionClient
	U         *pb.User
	Session   *pb.Session
	Conn      *grpc.ClientConn
	Header    grpc.CallOption
}
