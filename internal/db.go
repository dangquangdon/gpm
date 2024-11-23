package internal

import (
	"fmt"

	"github.com/dangquangdon/gpm/config"
)

func GetDbConnStr(cfg config.Config) string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DbUser,
		cfg.DbPass,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
	)
}
