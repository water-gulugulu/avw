package request

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// Custom claims structure
type CustomClaims struct {
	UUID        uuid.UUID // 唯一标示
	ID          uint      // 用户ID
	Username    string    // 用户名
	NickName    string    // 用户名
	AuthorityId string    // 权限ID
	BufferTime  int64
	jwt.StandardClaims
}
