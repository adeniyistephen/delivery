package delivery

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func Test(log *zap.SugaredLogger, db *sqlx.DB) *mux.Router {
	cg := CoreGroup{
		core: NewCore(log, db),
	}

	r := mux.NewRouter()

	r.HandleFunc("/", cg.Exec_AddDelivery).Methods("GET")
	return r
}
