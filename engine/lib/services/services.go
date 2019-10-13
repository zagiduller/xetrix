package services

import (
	"engine/lib/payments"
	"engine/lib/services/events"
	"engine/lib/structs"
	"fmt"
	"github.com/olebedev/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sync"
)

var (
	Mu = &sync.Mutex{}
)

type IRepo interface {
	ICurrencyRepository
	IUserRepository
	IAccountRepository
	IOrderRepository
	ITransactionRepository
	ICommissionRepository
}

func Run(c *config.Config, repo IRepo, bus *events.Bus, payreg *payments.Registry) {
	port, _ := c.String("engine.port")

	fmt.Printf("GRPC services Host: :%s \n", port)

	server := grpc.NewServer(grpc.UnaryInterceptor(PidInteceptor))

	// Инициализация сервисов
	sc := NewCurrencyService(repo)
	su := NewUserService(repo)

	// Сервис Авторизации
	secretKey, _ := c.String("engine.services.auth.secretKey")
	sau := NewAuthService([]byte(secretKey), su)
	st := NewTransactionService(repo, bus)
	sa := NewAccountService(repo, sc, su, st, payreg)

	sab := NewAccountBalanceService(repo, sa, st, bus)

	commissionRule, _ := c.Float64("engine.services.commission.rule")
	scm := NewCommissionService(repo, sa, sab, "", commissionRule)
	so := NewOrderService(repo, sc, su, st, sa, sab, scm)

	stp := NewTransactionProcessingService(repo, st, sa, so)

	// Шина событий
	bus.InitSubscribePublishers(sc, su, sa, stp, so, sab)

	// Получение или создание администратора
	adminuser, _ := c.Get("engine.admin")
	uadmin := su.InitAdminUser(adminuser)

	//Установка юзера который будет принимать комиссии
	scm.SetSystemUserId(uadmin.Id)

	// Регистрация сервисов в grpc
	structs.RegisterServiceCurrencyServer(server, sc)
	structs.RegisterServiceUserServer(server, su)
	structs.RegisterServiceAuthServer(server, sau)
	structs.RegisterServiceAccountServer(server, sa)
	structs.RegisterServiceAccountBalanceServer(server, sab)
	structs.RegisterServiceTransactionServer(server, st)
	structs.RegisterServiceTransactionProcessingServer(server, stp)
	structs.RegisterServiceOrderServer(server, so)

	reflection.Register(server)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
