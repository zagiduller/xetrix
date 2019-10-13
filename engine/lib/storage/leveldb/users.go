package leveldb

import (
	"bytes"
	"encoding/binary"
	"engine/lib/helper"
	pb "engine/lib/structs"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
)

func (repo *Repository) NewUserInc() uint32 {
	var inc uint32
	v, err := repo.users.Get([]byte("inc-"), nil)
	if err != nil {
		inc = 0
	} else {
		inc = binary.LittleEndian.Uint32(v)
		inc += 1
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, inc)
	repo.users.Put([]byte("inc-"), buf.Bytes(), nil)

	return inc
}

func (repo *Repository) CreateUser(u *pb.User) (*pb.User, error) {
	if exist, _ := repo.FindByLogin(u.Name); exist != nil {
		return nil, fmt.Errorf("Repo-CreateUser: User with login %s already exist", u.Name)
	}

	if len(u.Id) == 0 {
		u.Id = uuid.New().String()
		u.Inc = repo.NewUserInc()
	}
	u.CreatedAt = helper.CurrentTimestamp()

	byted, err := proto.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("Repo-CreateUser: %s", err)
	}
	batch := new(leveldb.Batch)

	batch.Put([]byte("object-"+u.Id), byted)
	batch.Put([]byte("name-"+u.Name), []byte(u.Id))
	if err := repo.users.Write(batch, nil); err != nil {
		return nil, fmt.Errorf("Repo-CreateUser: %s", err)
	}
	log.Printf("Repo-CreateUser: user appended")
	return u, nil
}

func (repo *Repository) GetAllUsers() ([]*pb.User, error) {
	var result []*pb.User
	iter := repo.users.NewIterator(util.BytesPrefix([]byte("object-")), nil)
	for iter.Next() {
		obj := new(pb.User)
		if err := proto.Unmarshal(iter.Value(), obj); err == nil {
			result = append(result, obj)
		} else {
			fmt.Printf("Repo-GetAllUsers: Error. %s", err)
		}
	}

	return result, nil
}

func (repo *Repository) FindUser(id string) (*pb.User, error) {
	if len(id) > 0 {
		user := new(pb.User)
		byted, err := repo.users.Get([]byte("object-"+id), nil)
		if err != nil {
			return nil, fmt.Errorf("Repo-FindUser: %s", err)
		}
		if err := proto.Unmarshal(byted, user); err != nil {
			return nil, fmt.Errorf("Repo-FindUser: %s", err)
		}
		return user, nil
	}
	return nil, nil
}

func (repo *Repository) FindByLogin(name string) (*pb.User, error) {
	if len(name) > 0 {
		user := new(pb.User)

		byteId, err := repo.users.Get([]byte("name-"+name), nil)
		if err != nil {
			return nil, fmt.Errorf("Repo-FindByLogin: %s", err)
		}
		byted, err := repo.users.Get([]byte("object-"+string(byteId)), nil)
		if err != nil {
			return nil, fmt.Errorf("Repo-FindByLogin: %s", err)
		}
		if err := proto.Unmarshal(byted, user); err != nil {
			return nil, fmt.Errorf("Repo-FindByLogin: %s", err)
		}
		return user, nil
	}
	return nil, fmt.Errorf("Repo-FindByLogin: Incorrect login or password")
}
