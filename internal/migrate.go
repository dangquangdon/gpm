package internal

import (
	"os"
	"os/exec"

	"github.com/dangquangdon/gpm/config"
)

func ExecMigrate(cfg config.Config, args []string) error {
	dbConnStr := GetDbConnStr(cfg)
	cmdArgs := []string{
		"-path",
		cfg.MigrationDir,
		"-database",
		dbConnStr,
	}
	cmdArgs = append(cmdArgs, args...)
	cmd := exec.Command("migrate", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func ExecBaseMigrate(args []string) error {
	cmd := exec.Command("migrate", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
