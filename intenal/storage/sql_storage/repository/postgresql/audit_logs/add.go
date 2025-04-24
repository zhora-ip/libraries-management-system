package auditlogs

import (
	"context"
	"time"
)

func (r *AuditLogsRepo) Add(ctx context.Context, log string) (int, error) {
	var ID int

	err := r.db.ExecQueryRow(
		ctx,
		`INSERT INTO
			audit_logs (
				created_at,
				log
			)
			VALUES ($1,$2)
			RETURNING id;`,
		time.Now(),
		log,
	).Scan(&ID)

	return ID, err
}
