package services

import (
	"engine/lib/services/events"
	pb "engine/lib/structs"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

//db-interface
type ICurrencyRepository interface {
	CreateCurrency(query *pb.Currency) (*pb.Currency, error)
	GetAllCurrency() ([]*pb.Currency, error)
	FindCurrency(req *pb.Query_Currency) (*pb.Currency, error)
	UpdateCurrency(obj *pb.Currency) (*pb.Currency, error)
	NewCurrInc() uint32
}

//service
type ServiceCurrency struct {
	bus  *events.Bus
	repo ICurrencyRepository
}

//constructor
func NewCurrencyService(repo ICurrencyRepository) *ServiceCurrency {
	return &ServiceCurrency{
		repo: repo,
	}
}

func (s *ServiceCurrency) AddEventBus(bus *events.Bus) {
	s.bus = bus
}

func (s *ServiceCurrency) Notify(event *pb.Event) {
	s.bus.NewEvent(event)
}

// create object
func (s *ServiceCurrency) CreateCurrency(ctx context.Context, req *pb.Query_CreateCurrency) (*pb.Response_CurrencyQuery, error) {
	// Save our consignment

	if len(req.Items) > 0 {
		var currencies []*pb.Currency

		for _, v := range req.Items {
			curr, err := s.repo.CreateCurrency(v)
			if err != nil {
				return nil, fmt.Errorf("CreateCurrency: %s ", err)
			}

			s.Notify(&pb.Event{Type: &pb.Event_NewCurrency{
				NewCurrency: &pb.EventNewCurrency{Currency: curr},
			}})

			log.Printf("CreateCurrency: %s (%s) ", curr.Id, curr.Symbol)
			currencies = append(currencies, curr)
		}

		// Return matching the `Response` message we created in our
		// protobuf definition.
		return &pb.Response_CurrencyQuery{
			Created:     true,
			Items:       currencies,
			ItemsCount:  uint32(len(currencies)),
			QueryStatus: pb.QueryStatus_Query_Success,
		}, nil
	} else if req.Object != nil {
		curr, err := s.repo.CreateCurrency(req.Object)
		if err != nil {
			return nil, fmt.Errorf("CreateCurrency: %s ", err)
		}

		s.Notify(&pb.Event{Type: &pb.Event_NewCurrency{
			NewCurrency: &pb.EventNewCurrency{Currency: curr},
		}})

		return &pb.Response_CurrencyQuery{
			Created:     true,
			Object:      curr,
			QueryStatus: pb.QueryStatus_Query_Success,
		}, nil
	}
	return nil, fmt.Errorf("CreateCurrency: Incorrect query ")
}

// view
func (s *ServiceCurrency) GetCurrency(ctx context.Context, req *pb.Query_Currency) (*pb.Response_CurrencyQuery, error) {
	resp := &pb.Response_CurrencyQuery{QueryStatus: pb.QueryStatus_Query_Success}

	c, err := s.repo.FindCurrency(req)
	if err != nil {
		return nil, fmt.Errorf("GetCurrency: %s ", err)
	}

	if c != nil {
		resp.Object = c
		return resp, nil
	}

	cs, err := s.repo.GetAllCurrency()
	if err != nil {
		return nil, fmt.Errorf("GetCurrency: %s ", err)
	}
	resp.Items = cs
	resp.ItemsCount = uint32(len(cs))

	return resp, nil
}

func (s *ServiceCurrency) Deactivate(ctx context.Context, req *pb.Query_Currency) (*pb.Currency, error) {
	if aid := ctx.Value("admin-id"); aid == nil {
		return nil, status.Error(codes.PermissionDenied, "ServiceCurrency Deactivate: not allowed")
	}
	curr, err := s.repo.FindCurrency(req)
	if err != nil {
		return nil, fmt.Errorf("ServiceCurrency Deactivate, %s", err)
	}

	curr.Active = false
	deactivated, err := s.repo.UpdateCurrency(curr)
	if err == nil {
		s.Notify(&pb.Event{Type: &pb.Event_CurrencyDeactivated{
			CurrencyDeactivated: &pb.EventCurrencyDeactivated{Currency: deactivated},
		}})
		return deactivated, nil
	}
	return nil, fmt.Errorf("ServiceCurrency Deactivate, %s", err)
}

func (s *ServiceCurrency) Activate(ctx context.Context, req *pb.Query_Currency) (*pb.Currency, error) {
	if aid := ctx.Value("admin-id"); aid == nil {
		return nil, status.Error(codes.PermissionDenied, "ServiceCurrency Activate: not allowed")
	}
	curr, err := s.repo.FindCurrency(req)
	if err != nil {
		return nil, fmt.Errorf("ServiceCurrency Activate, %s", err)
	}

	curr.Active = true

	activated, err := s.repo.UpdateCurrency(curr)
	if err == nil {
		s.Notify(&pb.Event{Type: &pb.Event_CurrencyActivated{
			CurrencyActivated: &pb.EventCurrencyActivated{Currency: activated},
		}})
		return activated, nil
	}
	return nil, fmt.Errorf("ServiceCurrency Activate, %s", err)

}
