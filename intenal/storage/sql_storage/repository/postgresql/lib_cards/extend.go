package libcards

import (
	"context"
	"time"
)

func (s *LibCardsRepo) Extend(ctx context.Context, ID int64, extTime time.Time) error {
	_, err := s.db.Exec(ctx, `UPDATE lib_cards SET expires_at = $1 WHERE id = $2`, extTime, ID)
	return err
}
