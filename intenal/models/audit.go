package models

import (
	"fmt"
	"strings"
)

type AuditStatusChange struct {
	ID  string
	Old string
	New string
}

type AuditRequest struct {
	Method string
	URL    string
	Query  string
	Body   string
}

type AuditResponse struct {
	Code string
	Body string
}

func (a AuditStatusChange) String() string {
	var fields []string

	if a.ID != "" {
		fields = append(fields, fmt.Sprintf("ID: %s", a.ID))
	}
	if a.Old != "" {
		fields = append(fields, fmt.Sprintf("Old: %s", a.Old))
	}
	if a.New != "" {
		fields = append(fields, fmt.Sprintf("New: %s", a.New))
	}

	return fmt.Sprintf("[STATUS CHANGE] %s", strings.Join(fields, ", "))
}

func (a AuditRequest) String() string {
	var fields []string

	if a.Method != "" {
		fields = append(fields, fmt.Sprintf("Method: %s", a.Method))
	}
	if a.URL != "" {
		fields = append(fields, fmt.Sprintf("URL: %s", a.URL))
	}
	if a.Query != "" {
		fields = append(fields, fmt.Sprintf("Query: %s", a.Query))
	}
	if a.Body != "" {
		fields = append(fields, fmt.Sprintf("Body: %s", a.Body))
	}

	return fmt.Sprintf("[REQUEST] %s", strings.Join(fields, ", "))
}

func (a AuditResponse) String() string {
	var fields []string

	if a.Code != "" {
		fields = append(fields, fmt.Sprintf("Code: %s", a.Code))
	}
	if a.Body != "" {
		fields = append(fields, fmt.Sprintf("Body: %s", a.Body))
	}

	return fmt.Sprintf("[RESPONSE] %s", strings.Join(fields, ", "))
}
