package domain

import (
	"github.com/google/uuid"
	"time"
)

type Wallet struct {
	Uuid       *uuid.UUID
	WalletName string

	Income float64
	Spends float64

	Created *time.Time
}
