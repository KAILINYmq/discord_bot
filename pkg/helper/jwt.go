package helper

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtVerify interface {
	GenSignature() (string, error)
	VerifySignature(sign string) bool
}

const (
	privateKey = "xxxxxx#!dw1dfwwqreq21"
	payIssuer  = "aaaaaworld"
)

type payClaims struct {
	User      string
	TradeType string
	TradeItem string
	Amount    string
	TradeNo   string
	jwt.StandardClaims
}

type bonusClaims struct {
	GameId      string
	BonusId     string
	BonusAmount string
	Count       int64
	jwt.StandardClaims
}

func (b *bonusClaims) GenSignature() (string, error) {
	nowTime := time.Now()

	expireTime := nowTime.Add(time.Minute * 5)

	b.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Issuer:    payIssuer,
	}
	//jwt.SigningMethodES256
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, b).SignedString([]byte(privateKey))
	return token, err
}

func (b *bonusClaims) VerifySignature(sign string) bool {
	cc, err := bonusParsePayToken(sign)
	if err != nil {
		return false
	}

	if cc.ExpiresAt <= time.Now().Unix() {
		return false
	}

	if cc.GameId != b.GameId || cc.Count != b.Count || cc.BonusAmount != b.BonusAmount || cc.BonusId != b.BonusId {
		return false
	}
	return true
}

func bonusParsePayToken(sign string) (*bonusClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(sign, &bonusClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(privateKey), nil
	})

	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*bonusClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("parse error")
}

func (p *payClaims) GenSignature() (string, error) {
	nowTime := time.Now()

	expireTime := nowTime.Add(time.Minute * 5)

	p.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Issuer:    payIssuer,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, p).SignedString([]byte(privateKey))
	return token, err
}

func (p *payClaims) VerifySignature(sign string) bool {
	cc, err := payParsePayToken(sign)
	if err != nil {
		return false
	}

	if cc.ExpiresAt <= time.Now().Unix() {
		return false
	}

	if cc.User != p.User ||
		cc.TradeNo != p.TradeNo ||
		cc.Amount != p.Amount ||
		cc.TradeType != p.TradeType ||
		cc.TradeItem != p.TradeItem {
		return false
	}
	return true
}

func payParsePayToken(sign string) (*payClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(sign, &payClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(privateKey), nil
	})

	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*payClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("parse error")
}

func NewPayJwt(user, tradeType, tradeItem, amount, tradeNo string) JwtVerify {
	return &payClaims{
		User:      user,
		TradeType: tradeType,
		TradeItem: tradeItem,
		Amount:    amount,
		TradeNo:   tradeNo,
	}
}

func NewBonusJwt(gameId, bonusId, bonusAmount string, count int64) JwtVerify {
	return &bonusClaims{
		GameId:      gameId,
		BonusId:     bonusId,
		BonusAmount: bonusAmount,
		Count:       count,
	}
}
