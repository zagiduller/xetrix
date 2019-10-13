package services

import (
	"engine/lib/services/events"
	pb "engine/lib/structs"
	"fmt"
	"golang.org/x/net/context"
	"log"
)

type ITransactionRepository interface {
	SaveTx(tx *pb.Tx) (*pb.Tx, error)
	UpdateTx(tx *pb.Tx) (*pb.Tx, error)
	GetTx(req *pb.Query_Tx) (*pb.Tx, error)
	GetTxs(req *pb.Query_Tx) ([]*pb.Tx, error)
	GetAllTxs() ([]*pb.Tx, error)
	ConfirmTx(tx *pb.Tx) (bool, error)
	UpdateContractAvailable(c *pb.Contract) (*pb.Contract, error)
	//UpdateContractStatus(c *pb.Contract) (*pb.Contract, error)
}

func NewTransactionService(repo ITransactionRepository, bus *events.Bus) *ServiceTransaction {
	return &ServiceTransaction{
		repo: repo,
		bus:  bus,
	}
}

type ServiceTransaction struct {
	repo ITransactionRepository
	bus  *events.Bus
}

func (s *ServiceTransaction) AddEventBus(bus *events.Bus) {
	s.bus = bus
	//bus.Subscribe(s,
	//	&pb.Event{Type: &pb.Event_PaySystemRawTx{}},
	//)
}

func (s *ServiceTransaction) Update(event *pb.Event) {

}

func (s *ServiceTransaction) Notify(event *pb.Event) {
	s.bus.NewEvent(event)
}

/**
Проверки перенести в отдельные вальдаторы:
Проверка аккаунтов
*/
func (s *ServiceTransaction) CreateTx(ctx context.Context, _tx *pb.Tx) (*pb.Response_Tx, error) {
	if _tx.Reason == nil || _tx.Reason.Status == pb.TxReason_UNREASON_TX {
		return nil, fmt.Errorf("CreateTx: Unreason tx ")
	}

	// Обеспечение уникальности транзакции
	if len(_tx.Reason.InPStxId) > 0 {
		if find, _ := s.GetTx(ctx, &pb.Query_Tx{InPStxId: _tx.Reason.InPStxId}); find != nil {
			return nil, fmt.Errorf("CreateTx: PaymentSystemID not unique")
		}
	}

	_tx.Status = pb.TxStatus_UNCONFIRMED
	tx, err := s.repo.SaveTx(_tx)
	if err != nil {
		return nil, fmt.Errorf("CreateTx: %s ", err)
	}

	for _, rtx := range _tx.Related {
		resp, err := s.CreateTx(ctx, rtx)
		if err != nil {
			log.Printf("CreateTx: Related tx: %s ", err)
		}
		rtx.Id = resp.Object.Id
	}

	if _, err := s.repo.UpdateTx(tx); err != nil {
		log.Printf("CreateTx: %s", err)
	}

	log.Printf("CreateTx: id: %s from: %s to: %s amount: %f reason %s\n",
		tx.Id,
		_tx.FromAddress,
		_tx.ToAddress,
		_tx.Amount,
		_tx.Reason.Status)

	s.Notify(&pb.Event{Type: &pb.Event_NewTransaction{NewTransaction: &pb.EventNewTx{Tx: _tx}}})

	return &pb.Response_Tx{
		Created:     true,
		Object:      tx,
		QueryStatus: pb.QueryStatus_Query_Success,
	}, nil

}

func (s *ServiceTransaction) GetTx(ctx context.Context, query *pb.Query_Tx) (*pb.Response_Tx, error) {
	txs, err := s.repo.GetTxs(query)
	if err != nil {
		return nil, fmt.Errorf("GetTx: %s ", err)
	}
	return &pb.Response_Tx{
		Items:       txs,
		ItemsCount:  uint32(len(txs)),
		QueryStatus: pb.QueryStatus_Query_Success,
	}, nil
}

func (s *ServiceTransaction) GetAllTxs(ctx context.Context, q *pb.Empty) (*pb.Response_Tx, error) {
	aid := ctx.Value("admin-id")
	if aid != nil {
		txs, err := s.repo.GetAllTxs()
		if err != nil {
			return nil, fmt.Errorf("GetAllTxs: %s ", err)
		}
		return &pb.Response_Tx{
			Items:       txs,
			ItemsCount:  uint32(len(txs)),
			QueryStatus: pb.QueryStatus_Query_Success,
		}, nil
	}
	return nil, fmt.Errorf("GetAllTxs: Not Allowed")
}
