package internal

import (
	"os"
	"os/exec"

	"github.com/dangquangdon/gpm/config"
)

func setRestEnvVar(cfg config.Config) error {
	dbConnStr := GetDbConnStr(cfg)
	restEnvVarMap := map[string]string{
		"PGRST_DB_URI":         dbConnStr,
		"PGRST_DB_SCHEMAS":     cfg.RestDbSchema,
		"PGRST_DB_ANON_ROLE":   cfg.RestDbAnonRole,
		"PGRST_JWT_SECRET":     cfg.RestJwtSecret,
		"PGRST_SERVER_PORT":    cfg.RestServerPort,
		"PGRST_DB_PRE_REQUEST": cfg.RestPreRequest,
	}

	for key, value := range restEnvVarMap {
		err := os.Setenv(key, value)
		if err != nil {
			return err
		}
	}

	return nil
}

func ExecPGRest(cfg config.Config) error {
	err := setRestEnvVar(cfg)
	if err != nil {
		return err
	}

	cmd := exec.Command("postgrest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
