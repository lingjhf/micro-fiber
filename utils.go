package microfiber

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Retry 重试函数
func Retry(fn func() error, n uint, d time.Duration) (err error) {
	var count uint
	for {
		err = fn()
		if err != nil {
			time.Sleep(d)
			count++
		} else {
			err = nil
			break
		}
		if count > n {
			break
		}
	}
	return
}

// GenerateCaptcha 随机生成4位数验证码
func GenerateCaptcha() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}

// CreateTablesIfNotExists 判断表如果不存在就创建表
func CreateTablesIfNotExists(db *gorm.DB, tables ...interface{}) error {
	for _, table := range tables {
		if !db.Migrator().HasTable(table) {
			if err := db.Migrator().CreateTable(table); err != nil {
				return err
			}
		}
	}
	return nil
}

func GeneratePasswordHash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func ComparePasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJwtWithKey(payload map[string]any, key string, expiration time.Duration) (string, error) {
	builder := jwt.NewBuilder().Expiration(time.Now().Add(expiration))
	for key, value := range payload {
		builder.Claim(key, value)
	}
	token, err := builder.Build()
	if err != nil {
		return "", err
	}
	jwtBytes, err := jwt.Sign(token, jwt.WithKey(jwa.HS256, []byte(key)))
	if err != nil {
		return "", err
	}
	return string(jwtBytes), nil
}

func ParseJwtWithKey(jwtString, key string) (jwt.Token, error) {
	return jwt.Parse([]byte(jwtString), jwt.WithKey(jwa.HS256, []byte(key)))
}

func GetOffsetLimit(page, limit uint) (uint, uint) {
	if limit <= 0 {
		limit = 10
	}
	if page > 1 {
		return (page - 1) * limit, limit
	}
	return 0, limit
}
