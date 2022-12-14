package luna

import (
	"strings"
	"testing"

	"github.com/zhangrt/voyager1_core/constant"
	"github.com/zhangrt/voyager1_core/global"

	pb "github.com/zhangrt/voyager1_core/zinx/pb"
)

var s Star = Star{}

func pre() {
	global.G_CONFIG.JWT.SigningKey = "gsafety"
	global.G_CONFIG.AUTHKey.RefreshToken = "new-token"
}

func TestCheck(t *testing.T) {
	pre()
	s := Star{}
	s.CheckToken(&pb.Authentication{
		Key:   "test",
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZjBjY2M1ZGEtMzA5NC00MWJkLWJmN2UtNzE1MDZjMTdjNDQ3IiwiSUQiOjc5NDI4MDIxMTY5MjgxNDMzNywiQWNjb3VudCI6InRlc3QiLCJOYW1lIjoiQklHIE1vbnN0ZXIiLCJBdXRob3JpdHlJZCI6Ijk1MjgiLCJBdXRob3JpdHkiOnsiQ3JlYXRlZEF0IjoiMjAyMi0wOS0wNlQxOTo1ODowMy40MTM1MDgrMDg6MDAiLCJVcGRhdGVkQXQiOiIyMDIyLTA5LTA2VDE5OjU4OjA0LjY1NDI4MSswODowMCIsIkRlbGV0ZWRBdCI6bnVsbCwiYXV0aG9yaXR5SWQiOiI5NTI4IiwiYXV0aG9yaXR5TmFtZSI6Iua1i-ivleinkuiJsiIsInBhcmVudElkIjoiMCIsImRlZmF1bHRSb3V0ZXIiOiI0MDQifSwiQXV0aG9yaXRpZXMiOm51bGwsIkRlcGFydE1lbnRJZCI6IiIsIkRlcGFydE1lbnROYW1lIjoiIiwiVW5pdElkIjoiIiwiVW5pdE5hbWUiOiIiLCJCdWZmZXJUaW1lIjoxMjAsImV4cCI6MTY2MzgzNjUxMCwiaXNzIjoiZ3NhZmV0eSIsIm5iZiI6MTY2MzgzNTMzMH0.akA6tQJPoNbJ-HF6-s-EG9PnKQ4T5WRT6sjKLXf3984",
	})
}

func TestString(t *testing.T) {
	pre()
	m := "new-token" + constant.MARKER + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZjBjY2M1ZGEtMzA5NC00MWJkLW"
	s := strings.Split(m, constant.MARKER)
	println(len(s))
	println(s[0])
	println(s[1])
}

func TestGetUserInfo(t *testing.T) {
	pre()
	s.GetUserInfo(&pb.Authentication{
		Key:   "test",
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZjBjY2M1ZGEtMzA5NC00MWJkLWJmN2UtNzE1MDZjMTdjNDQ3IiwiSUQiOjc5NDI4MDIxMTY5MjgxNDMzNywiQWNjb3VudCI6InRlc3QiLCJOYW1lIjoiQklHIE1vbnN0ZXIiLCJBdXRob3JpdHlJZCI6Ijk1MjgiLCJBdXRob3JpdHkiOnsiQ3JlYXRlZEF0IjoiMjAyMi0wOS0wNlQxOTo1ODowMy40MTM1MDgrMDg6MDAiLCJVcGRhdGVkQXQiOiIyMDIyLTA5LTA2VDE5OjU4OjA0LjY1NDI4MSswODowMCIsIkRlbGV0ZWRBdCI6bnVsbCwiYXV0aG9yaXR5SWQiOiI5NTI4IiwiYXV0aG9yaXR5TmFtZSI6Iua1i-ivleinkuiJsiIsInBhcmVudElkIjoiMCIsImRlZmF1bHRSb3V0ZXIiOiI0MDQifSwiQXV0aG9yaXRpZXMiOm51bGwsIkRlcGFydE1lbnRJZCI6IiIsIkRlcGFydE1lbnROYW1lIjoiIiwiVW5pdElkIjoiIiwiVW5pdE5hbWUiOiIiLCJCdWZmZXJUaW1lIjoxMjAsImV4cCI6MTY2MzgzNjUxMCwiaXNzIjoiZ3NhZmV0eSIsIm5iZiI6MTY2MzgzNTMzMH0.akA6tQJPoNbJ-HF6-s-EG9PnKQ4T5WRT6sjKLXf3984",
	})
}
