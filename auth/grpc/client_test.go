package grpc_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	service "github.com/zhangrt/voyager1_core/auth/grpc/pb"

	"google.golang.org/grpc"
)

func TestClient(t *testing.T) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	authClient := service.NewAuthServiceClient(conn)
	res, err := authClient.GetUser(context.Background(), &service.Token{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZjBjY2M1ZGEtMzA5NC00MWJkLWJmN2UtNzE1MDZjMTdjNDQ3IiwiSUQiOjc5NDI4MDIxMTY5MjgxNDMzNywiQWNjb3VudCI6InRlc3QiLCJOYW1lIjoiQklHIE1vbnN0ZXIiLCJBdXRob3JpdHlJZCI6Ijk1MjgiLCJBdXRob3JpdHkiOnsiQ3JlYXRlZEF0IjoiMjAyMi0wOS0wNlQxOTo1ODowMy40MTM1MDgrMDg6MDAiLCJVcGRhdGVkQXQiOiIyMDIyLTA5LTA2VDE5OjU4OjA0LjY1NDI4MSswODowMCIsIkRlbGV0ZWRBdCI6bnVsbCwiYXV0aG9yaXR5SWQiOiI5NTI4IiwiYXV0aG9yaXR5TmFtZSI6Iua1i-ivleinkuiJsiIsInBhcmVudElkIjoiMCIsImRlZmF1bHRSb3V0ZXIiOiI0MDQifSwiQXV0aG9yaXRpZXMiOm51bGwsIkRlcGFydE1lbnRJZCI6IiIsIkRlcGFydE1lbnROYW1lIjoiIiwiVW5pdElkIjoiIiwiVW5pdE5hbWUiOiIiLCJCdWZmZXJUaW1lIjoxMjAsImV4cCI6MTY2MzkwODYxMiwiaXNzIjoiZ3NhZmV0eSIsIm5iZiI6MTY2MzkwNzQzMn0.q3r3QwpLGcAq45OHinhB1wncEbATCjXwKdbMApgXLVM",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
