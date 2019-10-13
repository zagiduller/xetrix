package main

import (
	"engine/lib/gateway"
	"engine/lib/payments"
	"engine/lib/services"
	"engine/lib/services/events"
	"engine/lib/storage"
	"engine/lib/structs"
	"fmt"
	"github.com/olebedev/config"
	"google.golang.org/grpc"
	"log"
)

func main() {
	//appCtx, cancel := context.WithCancel(context.Background())
	//csig := make(chan os.Signal, 1)
	//signal.Notify(csig, os.Interrupt)
	//go func(){
	//	for sig := range csig {
	//		// sig is a ^C, handle it
	//		log.Println("Signal handle. Stop app: ", sig)
	//		cancel()
	//		os.Exit(0)
	//	}
	//}()

	cfg, err := config.ParseYamlFile("./config/app.yml")
	if err != nil {
		panic(err)
	}

	var repo services.IRepo

	cfgstorage, _ := cfg.Get("engine.database")
	switch cfgstorage.UString("type") {
	case "leveldb":
		repo = storage.InitLevelDBRepo(cfgstorage.UString("dbpath"))
		break
	default:
		repo = storage.InitSimpleRepo()
	}

	bus := events.NewBus()

	// Инициализация регистра платежных систем
	addrMap := make(map[string][]*structs.Account)
	if accs, err := repo.GetAllAccount(); err == nil {
		addrMap = addressesMap(accs)
	}
	currencies, _ := repo.GetAllCurrency()

	cfgpayments, _ := cfg.Get("engine.payments")
	preg := payments.InitRegistry(cfgpayments, currencies, addrMap)
	go preg.Run()
	bus.InitSubscribePublishers(preg)
	/////

	go services.Run(cfg, repo, bus, preg)

	rpchost, _ := cfg.String("engine.host")
	rpcport, _ := cfg.String("engine.port")

	grpcConn, err := grpc.Dial(fmt.Sprintf("%s:%s", rpchost, rpcport), grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed connect to grpc", err)
	}
	defer grpcConn.Close()

	go gateway.Run(grpcConn, cfg, bus, preg)

	//go websocket.Init(grpcConn, bus).Run(":50000")

	fmt.Scanln()
}

func addressesMap(accs []*structs.Account) map[string][]*structs.Account {
	addrMap := make(map[string][]*structs.Account)
	for _, a := range accs {
		addrMap[a.Currency.Symbol] = append(addrMap[a.Currency.Symbol], a)
	}
	return addrMap
}
