package app

import (
	pb "engine/lib/structs"
	"fmt"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

func (app *App) SignUp(name, pass string) error {
	log.Printf("Регистрируется новый участник: Name: %s, Pass: %s \n", name, pass)

	//p, err := app.srvc_tp.CreateUserAndAccountGenerate(app.ctx, &pb.User{Name: name, Password: pass})
	sess, err := app.srvc_auth.SignUp(app.ctx, &pb.Authenticate{Name: name, Password: pass, Email: name + "@" + name})
	if err != nil {
		return fmt.Errorf("SignUp: %s ", err)
	}
	decoded, err := app.srvc_auth.DecodeSession(app.ctx, sess)
	if err != nil {
		return fmt.Errorf("SignUp: %s ", err)
	}

	app.ctx = metadata.AppendToOutgoingContext(app.ctx, "pid", decoded.Id)

	app.U = decoded

	log.Printf("Session token: %s ", sess.Token)

	time.Sleep(5 * time.Second)
	//if err = app.StartSession(name, pass); err != nil {
	//	return fmt.Errorf("SignUp: %s ", err)
	//}
	return nil
}

//func (app *App) Migrate(p *pb.Participant) error {
//
//}
