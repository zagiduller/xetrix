package simple

import (
	"engine/lib/helper"
	pb "engine/lib/structs"
	"fmt"
	"github.com/google/uuid"
	"log"
)

func (repo *Repository) NewUserInc() uint32 {
	return uint32(len(repo.users))
}

func (repo *Repository) CreateUser(u *pb.User) (*pb.User, error) {
	if len(u.Id) == 0 {
		u.Id = uuid.New().String()
		u.Inc = repo.NewUserInc()
	}
	u.CreatedAt = helper.CurrentTimestamp()
	updated := append(repo.users, u)
	repo.users = updated
	log.Printf("Repo-CreateUser: user appended")
	return u, nil
}

func (repo *Repository) GetAllUsers() ([]*pb.User, error) {
	return repo.users, nil
}

func (repo *Repository) FindUser(id string) (*pb.User, error) {
	for _, p := range repo.users {
		if p.Id == id {
			return p, nil
		}
	}
	return nil, nil
}

func (repo *Repository) FindByLogin(name string) (*pb.User, error) {
	for _, p := range repo.users {
		if name == p.Name {
			return p, nil
		}

	}
	return nil, fmt.Errorf("Storage-FindByNamePassword: Incorrect login or password")
}

func (repo *Repository) FindByNamePassword(name, password string) (*pb.User, error) {
	for _, p := range repo.users {
		log.Printf("FindByNamePassword: %s @ %s == %s @ %s ", p.Name, p.Password, name, password)
		if name == p.Name && p.Password == password {
			return p, nil
		}
	}
	return nil, fmt.Errorf("Storage-FindByNamePassword: Incorrect login or password")
}
