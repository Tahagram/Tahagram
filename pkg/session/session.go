package session

import (
	"Kavka/config"
	"Kavka/pkg/jwt_manager"
	"Kavka/utils/random"
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewSession(redisClient *redis.Client, authConfigs config.Auth) *Session {
	jwtManager := jwt_manager.NewJwtManager(authConfigs)
	return &Session{redisClient, authConfigs, jwtManager}
}

// "makeExpiration" returns the expiration time for a given token type.
func makeExpiration(tokenType string) time.Duration {
	var expiration time.Duration

	if tokenType == jwt_manager.RefreshToken {
		expiration = jwt_manager.RF_EXPIRE_DAY
	}

	if tokenType == jwt_manager.AccessToken {
		expiration = jwt_manager.AT_EXPIRE_DAY
	}

	return expiration
}

func (session *Session) saveToken(token string, tokenType string) error {
	payload := struct {
		TokenType string `json:"token_type"`
	}{
		TokenType: tokenType,
	}

	payloadJson, _ := json.Marshal(payload)

	expireTime := makeExpiration(tokenType)

	err := session.redisClient.Set(context.TODO(), token, payloadJson, expireTime).Err()
	if err != nil {
		return err
	}

	return nil
}

func (session *Session) Destroy(token string) error {
	err := session.redisClient.Del(context.Background(), token).Err()
	if err != nil {
		return err
	}

	return nil
}

func (session *Session) DestroyOTP(phone string) error {
	err := session.redisClient.Del(context.Background(), phone).Err()
	if err != nil {
		return err
	}

	return nil
}

// Login function is used to handle the login process for a user with the given phone number.
// It's just generate an OTP code then saves it in redis store with the key `phone`.
func (session *Session) Login(phone string) (int, error) {
	otp := random.GenerateOTP()
	payload, _ := json.Marshal(loginPayload{OTP: otp})
	expiration := session.authConfigs.OTP_EXPIRE_SECONDS * time.Second

	err := session.redisClient.Set(context.Background(), phone, payload, expiration).Err()
	if err != nil {
		return 0, err
	}

	return otp, nil
}

// Verify the generated otp code in Login method, if it was correct returns new tokens (Access, Refresh)
func (session *Session) VerifyOTP(phone string, otp int, staticID primitive.ObjectID) (LoginTokens, bool) {
	payload, getErr := session.redisClient.Get(context.Background(), phone).Result()
	if getErr != nil {
		return LoginTokens{}, false
	}

	var data loginPayload
	unmarshalErr := json.Unmarshal([]byte(payload), &data)
	if unmarshalErr != nil {
		return LoginTokens{}, false
	}

	if otp == data.OTP {
		accessToken, atOk := session.NewAccessToken(staticID)
		if !atOk {
			return LoginTokens{}, false
		}

		refreshToken, rfOk := session.NewRefreshToken(staticID)
		if !rfOk {
			return LoginTokens{}, false
		}

		session.DestroyOTP(phone)

		return LoginTokens{AccessToken: accessToken, RefreshToken: refreshToken}, true
	}

	return LoginTokens{}, false
}

func (session *Session) newToken(staticID primitive.ObjectID, tokenType string) (string, bool) {
	// Generate Token
	token, err := session.jwtManager.Generate(tokenType, staticID)

	if err != nil {
		return "", false
	}

	saveErr := session.saveToken(token, tokenType)

	if saveErr != nil {
		return "", false
	}

	return token, true
}

// Generates and stores a new access token with given phone
func (session *Session) NewAccessToken(staticID primitive.ObjectID) (string, bool) {
	return session.newToken(staticID, jwt_manager.AccessToken)
}

// Generates and stores a new refresh token with given phone
func (session *Session) NewRefreshToken(staticID primitive.ObjectID) (string, bool) {
	return session.newToken(staticID, jwt_manager.RefreshToken)
}

func (session *Session) DecodeToken(token string, tokenType string) (*jwt_manager.JwtClaims, error) {
	_, getErr := session.redisClient.Get(context.TODO(), token).Result()
	if getErr != nil {
		return nil, jwt_manager.ErrInvalidToken
	}

	claims, err := session.jwtManager.Verify(token, tokenType)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
