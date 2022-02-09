package delivery

import (
	"context"

	"github.com/adeniyistephen/delivery/db"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Core struct {
	delivery db.Delivery
}

func NewCore(log *zap.SugaredLogger, sqlxDB *sqlx.DB) Core {
	return Core{
		delivery: db.NewDelivery(log, sqlxDB),
	}
}

func (c Core) AddDelivery(ctx context.Context) error {
	// Transaction starts here
	tran := func(tx sqlx.ExtContext) error {
		// Insert delivery
		c.delivery.Tran(tx).AddDelivery("Parcel", 101, 147, "Adeniyi", "08031234567", "No. 1, Adeola Street, Lagos", "Payment", "Luzon", 1.0, 1.0, "I ordered for 4 PPAR")
		return nil
	}
	
	if err := c.delivery.WithinTran(ctx, tran); err != nil {
		return err
	}

	return nil
}