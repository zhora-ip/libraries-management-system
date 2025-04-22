package physbookservice

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *PhysBookService) FindByBookID(ctx context.Context, ID int64) (*svc.FindPBookByBookIDResponse, error) {

	var (
		data = make(map[int64]*svc.BookAvailability)
		resp = &svc.FindPBookByBookIDResponse{}
	)

	if err := s.txManager.RunSerializable(ctx, func(ctxTx context.Context) error {
		books, err := s.pbRepo.FindByBookID(ctxTx, ID)

		switch {
		case err != nil:
			return err
		case len(books) == 0:
			return models.ErrObjectNotFound
		}

		shouldReturn, err1 := s.processLib(ctxTx, books, data)
		if shouldReturn {
			return err1
		}

		for _, v := range data {
			resp.Data = append(resp.Data, v)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *PhysBookService) processLib(ctxTx context.Context, books []*models.PhysBook, data map[int64]*svc.BookAvailability) (bool, error) {
	for _, b := range books {
		lib, err := s.lRepo.FindByID(ctxTx, b.LibraryID)
		if err != nil {
			return true, err
		}
		ba, ok := data[lib.ID]
		if !ok {
			data[lib.ID] = &svc.BookAvailability{
				Library:     lib,
				PhysBookIDs: []int64{b.ID},
				Amount:      1,
			}
			continue
		}
		ba.PhysBookIDs = append(ba.PhysBookIDs, b.ID)
		ba.Amount++
	}
	return false, nil
}
