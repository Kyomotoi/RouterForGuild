package global

import "fmt"

type Type string

const (
	TypeBot    Type = "Bot"
	TypeNormal Type = "Bearer"
)

type Token struct {
	AppID       uint64
	AccessToken string
	Type        Type
}

// New new一个 token
func New(tokenType Type) *Token {
	return &Token{
		Type: tokenType,
	}
}

// UserToken 获取用户 token
func UserToken(appID uint64, accessToken string) *Token {
	return &Token{
		AppID:       appID,
		AccessToken: accessToken,
		Type:        TypeNormal,
	}
}

// GetAuthToken 获取验证 token
func (t *Token) GetAuthToken() string {
	if t.Type == TypeNormal {
		return t.AccessToken
	}
	return fmt.Sprintf("%v.%s", t.AppID, t.AccessToken)
}

// init 初始化所需参数
func (t *Token) init() {
	t.AppID = Conf.AppID
	t.AccessToken = Conf.Token
}
