package services

import (
	"encoding/base64"
	"fmt"
	"regexp"

	"github.com/dgrijalva/jwt-go"
	"github.com/estromenko/yoof-api/internal/models"
	"github.com/estromenko/yoof-api/internal/repo"
	"golang.org/x/crypto/argon2"
)

type UserServiceConfig struct {
	JWTSecret string `json:"jwt_secret" yaml:"jwt_secret"`
	Salt      string `json:"salt" yaml:"salt"`
}

type UserService struct {
	Config *UserServiceConfig
	repo   *repo.UserRepo
}

// NewUserService ...
func NewUserService(repo *repo.UserRepo, config *UserServiceConfig) *UserService {
	return &UserService{
		repo:   repo,
		Config: config,
	}
}

func (s *UserService) Repo() *repo.UserRepo {
	return s.repo
}

func (u *UserService) hashPassword(password string) string {
	// Helper struct for password hashing
	type passwordConfig struct {
		time    uint32
		memory  uint32
		threads uint8
		keyLen  uint32
	}

	c := &passwordConfig{
		time:    1,
		memory:  64 * 1024,
		threads: 4,
		keyLen:  32,
	}

	hash := argon2.IDKey(
		[]byte(password),
		[]byte(u.Config.Salt),
		c.time,
		c.memory,
		c.threads,
		c.keyLen,
	)

	b64Salt := base64.RawStdEncoding.EncodeToString([]byte(u.Config.Salt))
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
	return fmt.Sprintf(format, argon2.Version, c.memory, c.time, c.threads, b64Salt, b64Hash)
}

// ComparePasswords ...
func (u *UserService) ComparePasswords(user *models.User, pass string) bool {
	return u.hashPassword(pass) == user.Password
}

func (u *UserService) validate(user *models.User) string {
	errors := ""

	if ok, _ := regexp.MatchString(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`, user.Email); !ok {
		errors += "Email is not valid. "
	}
	if ok, _ := regexp.MatchString(`[a-zA-Z0-9]{3,255}`, user.Username); !ok {
		errors += "Username is not valid. "
	}
	if ok, _ := regexp.MatchString(`.{6,}`, user.Password); !ok {
		errors += "Password is not valid. "
	}

	return errors
}

// Create ...
func (u *UserService) Create(user *models.User) error {

	// Validation
	if message := u.validate(user); message != "" {
		return fmt.Errorf(message)
	}

	user.Password = u.hashPassword(user.Password)

	if err := u.repo.Create(user); err != nil {
		return err
	}

	return nil
}

// GenerateToken ...
func (u *UserService) GenerateToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email

	return token.SignedString([]byte(u.Config.JWTSecret))
}

func (u *UserService) Login(email string, password string) (*models.User, string, error) {
	user, err := u.repo.FindByEmail(email)
	if err != nil {
		return nil, "", err
	}

	if !u.ComparePasswords(user, password) {
		return nil, "", fmt.Errorf("Wrong email or password")
	}

	token, err := u.GenerateToken(user)
	return user, token, err
}
