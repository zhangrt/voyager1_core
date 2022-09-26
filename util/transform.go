package util

import (
	"reflect"
	"unsafe"

	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/zhangrt/voyager1_core/auth/luna"
	pb "github.com/zhangrt/voyager1_core/zinx/pb"
)

func ProtoResultTransformClaims(result *pb.Result) *luna.CustomClaims {

	return ProtoClaimsTransformClaims(result.Claims)
}

func ProtoUserTransformClaims(result *pb.User) *luna.CustomClaims {

	return ProtoClaimsTransformClaims(result.Claims)
}

func ProtoClaimsTransformClaims(result *pb.CustomClaims) *luna.CustomClaims {
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
