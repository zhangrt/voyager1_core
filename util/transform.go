package util

import (
	"reflect"
	"unsafe"

	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	grpc_pb "github.com/zhangrt/voyager1_core/auth/grpc/pb"
	"github.com/zhangrt/voyager1_core/auth/luna"
	zinx_pb "github.com/zhangrt/voyager1_core/zinx/pb"
)

func ZinxProtoResult2Claims(result *zinx_pb.Result) *luna.CustomClaims {

	return ZinxProtoClaimsTransformClaims(result.Claims)
}

func GrpcProtoResult2Claims(result *grpc_pb.Result) *luna.CustomClaims {
	return GrpcProtoClaimsTransformClaims(result.Claims)
}

func GrpcProtoUser2Claims(result *grpc_pb.User) *luna.CustomClaims {
	return GrpcProtoClaimsTransformClaims(result.Claims)
}

func ZinxProtoUser2Claims(result *zinx_pb.User) *luna.CustomClaims {

	return ZinxProtoClaimsTransformClaims(result.Claims)
}

func GrpcProtoClaimsTransformClaims(result *grpc_pb.CustomClaims) *luna.CustomClaims {

	s := result.Claims.UUID

	bys := String2BytesSlicePlus(s)

	claims := &luna.CustomClaims{
		BaseClaims: luna.BaseClaims{
			ID:          uint(result.Claims.UserID),
			UUID:        uuid.FromBytesOrNil(bys),
			Account:     result.Claims.Account,
			Name:        result.Claims.Name,
			AuthorityId: result.Claims.AuthorityId,
		},
		BufferTime: result.BufferTime,
		StandardClaims: jwt.StandardClaims{
			Audience:  result.Standard.Audience,
			ExpiresAt: result.Standard.ExpiresAt,
			Id:        result.Standard.Id,
			IssuedAt:  result.Standard.IssuedAt,
			Issuer:    result.Standard.Issuer,
			NotBefore: result.Standard.NotBefore,
			Subject:   result.Standard.Subject,
		},
	}

	return claims
}

func ZinxProtoClaimsTransformClaims(result *zinx_pb.CustomClaims) *luna.CustomClaims {

	s := result.Claims.UUID

	bys := String2BytesSlicePlus(s)

	claims := &luna.CustomClaims{
		BaseClaims: luna.BaseClaims{
			ID:          uint(result.Claims.UserID),
			UUID:        uuid.FromBytesOrNil(bys),
			Account:     result.Claims.Account,
			Name:        result.Claims.Name,
			AuthorityId: result.Claims.AuthorityId,
		},
		BufferTime: result.BufferTime,
		StandardClaims: jwt.StandardClaims{
			Audience:  result.Standard.Audience,
			ExpiresAt: result.Standard.ExpiresAt,
			Id:        result.Standard.Id,
			IssuedAt:  result.Standard.IssuedAt,
			Issuer:    result.Standard.Issuer,
			NotBefore: result.Standard.NotBefore,
			Subject:   result.Standard.Subject,
		},
	}

	return claims
}

func String2BytesSlicePlus(str string) []byte {
	bytesSlice := []byte{}                                                                                            //此处定义了一个空切片
	stringData := &(*(*reflect.StringHeader)(unsafe.Pointer(&str))).Data                                              //取得StringHeader的Data地址
	byteSliceData := &(*(*reflect.SliceHeader)(unsafe.Pointer(&bytesSlice))).Data                                     //取得SliceHeader的Data地址
	*byteSliceData = *stringData                                                                                      //将StringHeader.Data的值赋给SliceHeader.Data
	(*(*reflect.SliceHeader)(unsafe.Pointer(&bytesSlice))).Len = (*(*reflect.StringHeader)(unsafe.Pointer(&str))).Len //设置长度

	return bytesSlice
}
