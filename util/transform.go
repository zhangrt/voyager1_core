package util

// 数据结构体之间转化的工具方法

import (
	"reflect"
	"unsafe"

	"github.com/golang-jwt/jwt/v4"
	grpc_pb "github.com/zhangrt/voyager1_core/auth/grpc/pb"
	"github.com/zhangrt/voyager1_core/auth/luna"
	grpc_pbJ "github.com/zhangrt/voyager1_core/com/gs/voyager1_core/auth/grpc/pb"
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

func GrpcProtoUser2ClaimsJ(result *grpc_pbJ.User) *luna.CustomClaims {
	return GrpcProtoClaimsTransformClaimsJ(result.Claims)
}

func ZinxProtoUser2Claims(result *zinx_pb.User) *luna.CustomClaims {

	return ZinxProtoClaimsTransformClaims(result.Claims)
}

func GrpcLunaClaimsTransformProtoClaims(claims *luna.CustomClaims) *grpc_pb.CustomClaims {
	if claims == nil {
		return nil
	}
	result := &grpc_pb.CustomClaims{
		Claims: &grpc_pb.BaseClaims{
			ID:      claims.Id,
			Account: claims.BaseClaims.Account,
			Name:    claims.BaseClaims.Name,
			RoleIds: claims.BaseClaims.RoleIds,
		},
		BufferTime: claims.BufferTime,
		Standard: &grpc_pb.StandardClaims{
			Audience:  claims.StandardClaims.Audience,
			ExpiresAt: claims.StandardClaims.ExpiresAt,
			Id:        claims.StandardClaims.Id,
			IssuedAt:  claims.StandardClaims.IssuedAt,
			Issuer:    claims.StandardClaims.Issuer,
			NotBefore: claims.StandardClaims.NotBefore,
			Subject:   claims.StandardClaims.Subject,
		},
	}
	return result
}

func GrpcLunaClaimsTransformProtoClaimsJ(claims *luna.CustomClaims) *grpc_pbJ.CustomClaims {
	if claims == nil {
		return nil
	}
	result := &grpc_pbJ.CustomClaims{
		Claims: &grpc_pbJ.BaseClaims{
			ID:      claims.Id,
			Account: claims.BaseClaims.Account,
			Name:    claims.BaseClaims.Name,
			RoleIds: claims.BaseClaims.RoleIds,
		},
		BufferTime: claims.BufferTime,
		Standard: &grpc_pbJ.StandardClaims{
			Audience:  claims.StandardClaims.Audience,
			ExpiresAt: claims.StandardClaims.ExpiresAt,
			Id:        claims.StandardClaims.Id,
			IssuedAt:  claims.StandardClaims.IssuedAt,
			Issuer:    claims.StandardClaims.Issuer,
			NotBefore: claims.StandardClaims.NotBefore,
			Subject:   claims.StandardClaims.Subject,
		},
	}
	return result
}

func GrpcProtoClaimsTransformClaims(result *grpc_pb.CustomClaims) *luna.CustomClaims {
	if result == nil {
		return nil
	}
	s := result.Claims.ID

	// bys := String2BytesSlicePlus(s)

	claims := &luna.CustomClaims{
		BaseClaims: luna.BaseClaims{
			// ID:      uuid.FromBytesOrNil(bys),
			ID:      s,
			Account: result.Claims.Account,
			Name:    result.Claims.Name,
			RoleIds: result.Claims.RoleIds,
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

func GrpcProtoClaimsTransformClaimsJ(result *grpc_pbJ.CustomClaims) *luna.CustomClaims {
	if result == nil {
		return nil
	}
	s := result.Claims.ID

	// bys := String2BytesSlicePlus(s)

	claims := &luna.CustomClaims{
		BaseClaims: luna.BaseClaims{
			// ID:      uuid.FromBytesOrNil(bys),
			ID:      s,
			Account: result.Claims.Account,
			Name:    result.Claims.Name,
			RoleIds: result.Claims.RoleIds,
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
	if result == nil {
		return nil
	}
	s := result.Claims.ID

	claims := &luna.CustomClaims{
		BaseClaims: luna.BaseClaims{
			ID:      s,
			Account: result.Claims.Account,
			Name:    result.Claims.Name,
			RoleIds: result.Claims.RoleIds,
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
