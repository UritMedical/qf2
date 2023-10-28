package qtoken

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	// 密钥
	tokenSecretKey = "#qf&frame@123456789852963741!%.#"
)

// GenerateToken
//
//	@Description: 生成Token
//	@param content token的内容
//	@param ttl 生存时间 单位小时
//	@return string
//	@return error
func GenerateToken(content map[string]interface{}, ttl int64) (string, error) {
	// 创建一个新的令牌声明
	claims := jwt.MapClaims{}

	now := time.Now()
	for k, v := range content {
		claims[k] = v
	}
	claims["IssuedAt"] = jwt.NewNumericDate(now)
	claims["ExpiresAt"] = jwt.NewNumericDate(now.Add(time.Duration(ttl) * time.Hour))

	// 使用HS256算法创建一个新的令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名令牌
	signedToken, err := token.SignedString([]byte(tokenSecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// VerifyToken
//
//	@Description: 验证Token
//	@param tokenString
//	@return *T
//	@return error
func VerifyToken[T any](tokenString string) (*T, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	// 打印解析的令牌信息
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		model := new(T)
		js, _ := json.Marshal(claims)
		_ = json.Unmarshal(js, model)
		// 判断是否超过有效值
		if int64(claims["ExpiresAt"].(float64)) < time.Now().Unix() {
			return model, errors.New("token已过期")
		}
		return model, nil
	}
	return nil, errors.New("token无效")
}
