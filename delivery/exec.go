package delivery

import (
	"context"
	"log"
	"net/http"
)

type CoreGroup struct {
	core Core
}

func (c CoreGroup) Exec_AddDelivery(w http.ResponseWriter, r *http.Request) {
	//c.core.AddDelivery("Parcel", 101, 147, "Adeniyi", "08031234567", "No. 1, Adeola Street, Lagos", "Payment", "Luzon", 1.0, 1.0, "I ordered for 4 PPAR")
	err := c.core.AddDelivery(context.Background())
	if err != nil {
		log.Println("error: %w", err)
	}
}