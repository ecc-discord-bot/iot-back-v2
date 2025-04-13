package services

import (
	"auth/models"
	"auth/utils"

	"github.com/golang-jwt/jwt/v5"
)

type SessionArgs struct {
	UserID    string // ユーザーID
	RemoteIP  string // リモートIP
	UserAgent string // ユーザーエージェント
}

func GenSessionToken(SessionID string) (string, error) {
	// トークンを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"SessionID": SessionID,
	})

	// トークンに署名
	signedToken, err := token.SignedString([]byte(TokenSecret))

	// エラー処理
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// トークンを検証
func ValidateSessionToken(tokenString string) (string, error) {
	// トークンを検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(TokenSecret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS512.Alg()}))

	// エラー処理
	if err != nil {
		return "", err
	}
	
	// トークンを検証
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims["SessionID"].(string), nil
	}

	return "", nil
}

// セッションを作成してトークンを返す
func NewSession(args SessionArgs) (string, error) {
	// ユーザーIDを取得
	user, err := models.GetUser(args.UserID)

	// エラー処理
	if err != nil {
		return "", err
	}

	// セッションIDを生成
	SessionID := utils.GenID()

	// セッションを作成
	session := models.Session{
		SessionID: SessionID,
		UserID:    args.UserID,
		RemoteIP:  args.RemoteIP,
		UserAgent: args.UserAgent,
	}

	// セッションを追加
	if err := user.NewSession(&session); err != nil {
		return "", err
	}

	// トークンを生成
	token, err := GenSessionToken(SessionID)

	return token, err
}

func GetSession(tokenString string) (*models.Session, error) {
	// トークンを検証
	SessionID, err := ValidateSessionToken(tokenString)

	// エラー処理
	if err != nil {
		return nil, err
	}

	// セッションを取得
	return models.GetSession(SessionID)
}