package qtoken

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	// 密钥
	tokenSecretKey = "#qf&frame@123qwer456thik789bvcx8529kjio63741!%.#"
)

type Content struct {
	Id      uint64                 `json:"i"`
	Roles   []string               `json:"r"`
	Customs map[string]interface{} `json:"c"`
}

// GenerateToken
//
//	@Description: 生成Token
//	@param content token的内容
//	@param ttl 生存时间 单位小时
//	@return string
//	@return error
func GenerateToken(content Content, ttl int64) (string, error) {
	// 创建一个新的令牌声明
	claims := jwt.MapClaims{}

	now := time.Now()
	claims["i"] = content.Id
	if content.Roles != nil {
		claims["r"] = content.Roles
	}
	claims["e"] = jwt.NewNumericDate(now.Add(time.Duration(ttl) * time.Hour))
	if content.Customs != nil {
		js, _ := json.Marshal(content.Customs)
		claims["c"] = string(js)
	}

	// 使用HS256算法创建一个新的令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名令牌
	signedToken, err := token.SignedString([]byte(tokenSecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// AnalyseToken
//
//	@Description: 解析Token
//	@param tokenString
//	@return *T
//	@return error
func AnalyseToken(tokenString string) (*Content, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	// 打印解析的令牌信息
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		model := &Content{}
		js, _ := json.Marshal(claims)
		_ = json.Unmarshal(js, model)
		// 判断是否超过有效值
		if int64(claims["e"].(float64)) < time.Now().Unix() {
			return model, errors.New("token已过期")
		}
		return model, nil
	}
	return nil, errors.New("token无效")
}

// CheckToken
//
//	@Description: 验证token是否有效，当token无效或过期时，返回false
//	@param tokenString
//	@return bool
func CheckToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecretKey), nil
	})
	if err != nil {
		return false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if int64(claims["e"].(float64)) > time.Now().Unix() {
			return true
		}
	}
	return false
}
