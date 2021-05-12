package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobaldia/encrypter-api/config"
	"io"
)

type Quote struct {
	Symbol string
	CurrentValue float64
	Variation float64
	PreviousClose float64
	Open float64
	MarketCapitalization float64
	Volume int
}

type EncryptedQuote struct {
	Value string
}

func EncryptQuote(ctx *gin.Context) {
	params := new(Quote)
	if ctx.BindJSON(params) != nil {
		ctx.String(400, "Bad Request")
		ctx.Abort()
		return
	}

	encryptedMessage, err := encryptQuote(*params)
	if err != nil {
		ctx.String(500, "Something went wrong")
		ctx.Abort()
		return
	}

	ctx.JSON(200, EncryptedQuote{
		Value: *encryptedMessage,
	})
}

func encryptQuote(quote Quote) (*string, error) {
	key := config.GetConfig().SecretKey
	if len(key) != 32 {
		return nil, errors.New("key length should be 32")
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	stringishQuote := []byte(fmt.Sprintf("%v", quote))
	cipherText := make([]byte, aes.BlockSize+len(stringishQuote))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], stringishQuote)

	encryptedMessage := base64.URLEncoding.EncodeToString(cipherText)

	return &encryptedMessage, nil
}