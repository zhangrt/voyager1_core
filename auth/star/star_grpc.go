package star

import (
	"fmt"
	"log"

	"github.com/zhangrt/voyager1_core/auth/grpc/pb"
	"github.com/zhangrt/voyager1_core/global"
	"google.golang.org/grpc"
)

func GetGrpcClient(opts ...grpc.DialOption) (*grpc.ClientConn, pb.AuthServiceClient) {
	conn, err := grpc.Dial(fmt.Sprintf(global.G_CONFIG.Grpc.Client.Host, ":", global.G_CONFIG.Grpc.Client.Port), opts)
	if err != nil {
		log.Fatal(err)
		return nil, nil
	}
	authClient := pb.NewAuthServiceClient(conn)

	return conn, authClient
}
func CloseConn(conn *grpc.ClientConn) {

	conn.Close()
}
