package services

import (
	"engine/lib/helper"
	"engine/lib/services/events"
	pb "engine/lib/structs"
	"fmt"
	"github.com/olebedev/config"
	"golang.org/x/net/context"
	"log"
)

type IUserRepository interface {
	CreateUser(u *pb.User) (*pb.User, error)
	GetAllUsers() ([]*pb.User, error)
	FindUser(id string) (*pb.User, error)
	FindByLogin(name string) (*pb.User, error)
	NewUserInc() uint32
}

func NewUserService(repo IUserRepository) *ServiceUser {
	return &ServiceUser{
		repo: repo,
	}
}

/* */

type ServiceUser struct {
	repo IUserRepository
	bus  *events.Bus
}

func (s *ServiceUser) AddEventBus(bus *events.Bus) {
	s.bus = bus
}

func (s *ServiceUser) Notify(event *pb.Event) {
	s.bus.NewEvent(event)
}

func (s *ServiceUser) CreateUser(ctx context.Context, u *pb.User) (*pb.Response_User, error) {
	s.GetUser(ctx, &pb.Query_User{})
	//ID назначается системой
	if len(u.Password) < 6 {
		return nil, fmt.Errorf("CreateUser: Short password")
	}
	var p *pb.User
	var err error
	u.Password, _ = hashPassword(u.Password)
	p, err = s.repo.CreateUser(u)
	if err != nil {
		return nil, fmt.Errorf("CreateUser: %s", err)
	}

	log.Printf("CreateUser: created: %s - %s ", p.Id, p.Name)

	s.Notify(
		&pb.Event{
			Type: &pb.Event_NewUser{NewUser: &pb.EventNewUser{User: u}},
		},
	)

	return &pb.Response_User{Created: true, Object: p, QueryStatus: pb.QueryStatus_Query_Success}, nil
}

func (s *ServiceUser) GetUser(ctx context.Context, req *pb.Query_User) (*pb.Response_User, error) {
	participants, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("GetUser: %s ", err)
	}
	if len(req.Id) > 0 {
		for _, p := range participants {
			if p.Id == req.Id {
				return &pb.Response_User{Object: p, QueryStatus: pb.QueryStatus_Query_Success}, nil
			}
		}
	}
	return nil, fmt.Errorf("GetUser: Not find user")
}

func (s *ServiceUser) GetInfo(ctx context.Context, empty *pb.Empty) (*pb.User, error) {
	pid := ctx.Value("pid")
	if pid != nil {
		fmt.Printf("GetInfo pid: %T %#v \n", pid, pid)
		resp, err := s.GetUser(ctx, &pb.Query_User{Id: pid.(string)})
		if err != nil {
			return nil, fmt.Errorf("GetInfo: %s", err)
		}
		if resp != nil {
			return &pb.User{
				Id:        resp.Object.Id,
				Name:      resp.Object.Name,
				Email:     resp.Object.Email,
				CreatedAt: resp.Object.CreatedAt,
			}, nil
		}
	}
	return nil, fmt.Errorf("GetInfo: Not find user")
}

func (s *ServiceUser) FindByNamePassword(ctx context.Context, auth *pb.Authenticate) (*pb.Response_User, error) {
	u, err := s.repo.FindByLogin(auth.Name)
	if err != nil {
		return nil, err
	}
	if helper.CheckPasswordHash(auth.Password, u.Password) {
		return &pb.Response_User{Object: u, QueryStatus: pb.QueryStatus_Query_Success}, nil
	}
	return nil, fmt.Errorf("FindByNamePassword: User not find")
}

func (s *ServiceUser) GetAllUsers(ctx context.Context, req *pb.Empty) (*pb.Response_User, error) {
	aid := ctx.Value("admin-id")
	if aid != nil {
		participants, err := s.repo.GetAllUsers()
		if err != nil {
			return nil, fmt.Errorf("GetAllUsers: %s ", err)
		}

		return &pb.Response_User{Items: participants, QueryStatus: pb.QueryStatus_Query_Success}, nil

	}

	return nil, fmt.Errorf("GetUser: Not accepted")
}

func (s *ServiceUser) InitAdminUser(cfgadmin *config.Config) *pb.User {
	//существует ли администратор
	if adm, err := s.repo.FindByLogin("Administrator"); err == nil {
		return adm
	}

	name, pass := cfgadmin.UString("name"), cfgadmin.UString("password")

	if len(name) > 3 && len(pass) > 4 {
		uadmin, _ := s.CreateUser(context.Background(), &pb.User{
			Name:     name,
			Email:    "admin@admin",
			Password: pass,
			Status:   pb.UserStatus_ADMINISTRATOR,
		})

		fmt.Printf("Administrator created. UserID: %s  \n Email: %s \t Name: %s \t Password: %s \n",
			uadmin.Object.Id,
			uadmin.Object.Email,
			uadmin.Object.Name,
			uadmin.Object.Password)

		return uadmin.Object
	}

	panic("Admin name or password not setup in config")
}
