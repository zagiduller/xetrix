package services

import (
	"context"
	"engine/lib/structs"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

//var (
//	sessions = make(map[string]*structs.Session)
//)

type ServiceAuth struct {
	signKey []byte
	su      *ServiceUser
	// Добавить защиту в виде User-Agent
	Sessions map[string]*structs.User
}

func NewAuthService(secret []byte, su *ServiceUser) *ServiceAuth {
	s := &ServiceAuth{
		signKey:  secret,
		su:       su,
		Sessions: make(map[string]*structs.User),
	}
	return s
}

func (s *ServiceAuth) SignUp(ctx context.Context, auth *structs.Authenticate) (*structs.Session, error) {
	r, err := s.su.CreateUser(ctx, &structs.User{
		Name:     auth.Name,
		Password: auth.Password,
		Email:    auth.Email,
	})
	if err != nil {
		return nil, err
	}
	sessid := uuid.New().String()
	t := s.createToken(sessid)
	ss := &structs.Session{
		SessionId: sessid,
		Token:     t,
	}
	s.Sessions[ss.SessionId] = r.Object

	return ss, nil
}

func (s *ServiceAuth) StartSession(ctx context.Context, auth *structs.Authenticate) (*structs.Session, error) {
	ru, err := s.su.FindByNamePassword(ctx, auth)
	if err != nil {
		return nil, fmt.Errorf("StartSession: ", err)
	}
	sess := structs.Session{}

	sess.SessionId = uuid.New().String()
	sess.Token = s.createToken(sess.SessionId)
	s.Sessions[sess.SessionId] = ru.Object

	return &sess, nil
}

func (s *ServiceAuth) CloseSession(ctx context.Context, empty *structs.Empty) (*structs.Bool, error) {
	return nil, nil
}

func (s *ServiceAuth) DecodeSession(ctx context.Context, session *structs.Session) (*structs.User, error) {
	token, err := jwt.Parse(session.Token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("DecodeSession: Unexpected signing method: %v", token.Header["alg"])
		}
		return s.signKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("DecodeSession: %s", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sessid, ok := claims["sessid"]
		if !ok {
			return nil, fmt.Errorf("DecodeSession: Token structure error - 'sessid' not find")
		}
		user, ok := s.Sessions[sessid.(string)]
		if !ok {
			return nil, fmt.Errorf("DecodeSession: Session not exist")
		}

		return &structs.User{
			Id:        user.Id,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			Status:    user.Status,
		}, nil

	}
	return nil, fmt.Errorf("DecodeSession: Invalid token")
}

func (s *ServiceAuth) createToken(sessid string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sessid": sessid,
		//"sub":  u.Id,
		//"name": u.Name,
	})

	// Подписываем токен нашим секретным ключем
	tokenString, err := token.SignedString(s.signKey)
	if err != nil {
		log.Printf("createToken: %s", err)
	}

	return tokenString
}

// Устанавливает pid в ctx. Серверная часть
func PidInteceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Способы получения токена:
	// Token, metadata
	if ctx.Value("pid") == nil && ctx.Value("admin-id") == nil {
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if len(md["token"]) > 0 {
				ctx = context.WithValue(ctx, "token", md["token"][0])
			}
			if len(md["pid"]) > 0 {
				ctx = context.WithValue(ctx, "pid", md["pid"][0])
			}
			if len(md["admin-id"]) > 0 {
				ctx = context.WithValue(ctx, "admin-id", md["admin-id"][0])
			}
		}
	}
	return handler(ctx, req)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
