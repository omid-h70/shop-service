package domain

import (
	"context"
	"github.com/pkg/errors"
)

var (
	ErrInvalidMinTransactionAmount = errors.New("invalid Min Transaction Amount")
	ErrInvalidMaxTransactionAmount = errors.New("invalid Max Transaction Amount")
	ErrInvalidAccountBalance       = errors.New("invalid Account Balance")
)

type (
	VendorEntity struct {
		id          string
		name        string
		phoneNumber string
		status      int
	}

	VendorReportRequest struct {
		id string
	}

	VendorReportResponse struct {
		ID        string
		DelayTime string
	}

	VendorRepository interface {
		GetAllVendorsDelayReports(ctx context.Context, input VendorReportRequest) ([]VendorReportResponse, error)
	}
)
