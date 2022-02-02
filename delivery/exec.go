package delivery

import "net/http"

type CoreGroup struct {
	core Core
}

func (c CoreGroup) Exec_AddDelivery(w http.ResponseWriter, r *http.Request) {
	c.core.AddDelivery("Dropship", 101, 147, "Adeniyi", "08031234567", "No. 1, Adeola Street, Lagos", "Payment", "Luzon", 1.0, 1.0, "")
}
