package mongo

import (
	"engine/lib/helper"
	pb "engine/lib/structs"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func (repo *Repository) CreateUser(u *pb.User) (*pb.User, error) {
	if len(u.Id) == 0 {
		u.Id = uuid.New().String()
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Storage-CreateParticipant: %s ", err)
	}
	u.Password = string(pwd)
	u.CreatedAt = helper.CurrentTimestamp()
	if err := repo.users.Insert(&u); err != nil {
		return nil, fmt.Errorf("Storage-CreateParticipant: %s ", err)
	}
	log.Printf("Storage-User added %+v", u)

	return u, nil
}

func (repo *Repository) GetAllUsers() ([]*pb.User, error) {
	var res []*pb.User
	if err := repo.users.Find(nil).All(&res); err != nil {
		return nil, fmt.Errorf("Storage-GetAllUsers: %s ", err)
	}

	return res, nil
}

func (repo *Repository) FindUser(id string) (*pb.User, error) {
	var u *pb.User
	if err := repo.users.Find(bson.M{"id": id}).One(u); err != nil {
		return nil, err
	}

	return u, nil
}

//FindByNamePassword TODO: Не безопасная операция!
//В связке логин-пароль, логик не подразумевается как уникальный
func (repo *Repository) FindByNamePassword(name, pass string) (*pb.User, error) {

	var us []*pb.User
	if err := repo.users.Find(bson.M{"name": name}).All(&us); err != nil {
		return nil, fmt.Errorf("Storage-FindByNamePassword: %s ", err)
	}
	if us == nil {
		return nil, fmt.Errorf("Storage-FindByNamePassword: Incorrect login or password")
	}

	for _, p := range us {
		if err := bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(pass)); err != nil {
			continue
		} else {
			return p, nil
		}
	}

	return nil, fmt.Errorf("Storage-FindByNamePassword: Incorrect login or password")
}
