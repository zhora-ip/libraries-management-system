package auditlogs

import "github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/db"

type AuditLogsRepo struct {
	db db.DB
}

func NewAuditLogs(database db.DB) *AuditLogsRepo {
	return &AuditLogsRepo{db: database}
}
