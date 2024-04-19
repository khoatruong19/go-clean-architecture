package initializers

import (
	"database/sql"

	"github.com/khoatruong19/go-clean-architecture/config"
	db "github.com/khoatruong19/go-clean-architecture/db/sqlc"
	_ "github.com/lib/pq"
)

func ConnectToDB(cfg *config.Config) *db.Store {
	conn, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		panic(err)
	}

	store := db.NewStore(conn)

	return store
}
