package cli

import (
	"fmt"

	"github.com/dangquangdon/gpm/config"
	"github.com/dangquangdon/gpm/internal"
)

func MigrateUp(cfg config.Config) error {
	return internal.ExecMigrate(cfg, []string{
		"-verbose",
		"up",
	})
}

func MakeMigration(cfg config.Config, name string) error {
	return internal.ExecBaseMigrate([]string{
		"create",
		"-seq",
		"-ext",
		"sql",
		"-dir",
		cfg.MigrationDir,
		name,
	})
}

func CheckDbConfig(cfg config.Config) error {
	dbConfig := internal.GetDbConnStr(cfg)
	fmt.Println(dbConfig)
	return nil
}

func MigrateForce(cfg config.Config, version string) error {
	return internal.ExecMigrate(cfg, []string{
		"-verbose",
		"force",
		version,
	})
}

func MigrateDown(cfg config.Config) error {
	return internal.ExecMigrate(cfg, []string{
		"--verbose",
		"down",
	})
}

func MigrateTo(cfg config.Config, version string) error {
	return internal.ExecMigrate(cfg, []string{
		"-verbose",
		"goto",
		version,
	})
}

func RunPGRest(cfg config.Config) error {
	return internal.ExecPGRest(cfg)
}
