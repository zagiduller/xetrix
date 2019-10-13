package email

import (
	"github.com/olebedev/config"
	"log"
	"testing"
)

func TestSender_Send(t *testing.T) {
	cfg, _ := config.ParseJson("{}")
	log.Printf("%s", cfg.Set("host", "smtp.gmail.com"))
	cfg.Set("port", ":587")
	cfg.Set("from", "rimastery@gmail.com")
	cfg.Set("username", "example@example.com")
	cfg.Set("password", "aaa")

	sender := NewSender(cfg)
	if err := sender.Send("design.mgn@gmail.com"); err != nil {
		t.Error(err)
	}

}
