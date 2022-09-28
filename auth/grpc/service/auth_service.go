package service

import (
	pb "github.com/zhangrt/voyager1_core/auth/grpc/pb"
	luna "github.com/zhangrt/voyager1_core/auth/luna"

	"context"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
}

func (auth *AuthService) ReadAuthentication(c context.Context, p *pb.Token) (*pb.Result, error) {
	result := new(pb.Result)

	return result, nil
}

func (auth *AuthService) GrantedAuthority(c context.Context, p *pb.Policy) (*pb.Result, error) {
	result := new(pb.Result)

	return result, nil

}
func (auth *AuthService) GetUser(c context.Context, p *pb.Token) (*pb.User, error) {
	user := new(pb.User)
	claims, _ := luna.GetUser(p.Token)
	if claims != nil {
		user.Claims = protoTransformClaims(claims)
		user.UserID = int64(claims.ID)
		user.UUID = claims.UUID.String()
		user.AuthorityId = claims.AuthorityId
	}
	return user, nil
}

// 转换
func protoTransformClaims(c *luna.CustomClaims) *pb.CustomClaims {
	p := pb.CustomClaims{
		Claims: &pb.BaseClaims{
			UserID:      int64(c.BaseClaims.ID),
			UUID:        c.BaseClaims.UUID.String(),
			AuthorityId: c.BaseClaims.AuthorityId,
			Account:     c.BaseClaims.Account,
			Name:        c.BaseClaims.Name,
		},
		BufferTime: c.BufferTime,
		Standard: &pb.StandardClaims{
			Audience:  c.StandardClaims.Audience,
			ExpiresAt: c.StandardClaims.ExpiresAt,
			Id:        c.StandardClaims.Id,
			IssuedAt:  c.StandardClaims.IssuedAt,
			Issuer:    c.StandardClaims.Issuer,
			NotBefore: c.StandardClaims.NotBefore,
			Subject:   c.StandardClaims.Subject,
		},
	}
	return &p
}
