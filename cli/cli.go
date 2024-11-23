package cli

import (
	"github.com/dangquangdon/gpm/config"
	"github.com/urfave/cli/v2"
)

func InitCli(cfg config.Config) *cli.App {
	return &cli.App{
		Name: "gpm",
		Commands: []*cli.Command{
			{
				Name: "migrate-up",
				Action: func(ctx *cli.Context) error {
					return MigrateUp(cfg)
				},
			},
			{
				Name: "migrate-down",
				Action: func(ctx *cli.Context) error {
					return MigrateDown(cfg)
				},
			},
			{
				Name: "migrate-to",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "version", Aliases: []string{"v"}},
				},
				Action: func(ctx *cli.Context) error {
					return MigrateTo(cfg, ctx.String("version"))
				},
			},
			{
				Name: "make-migrations",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
				},
				Action: func(ctx *cli.Context) error {
					name := ctx.String("name")
					return MakeMigration(cfg, name)
				},
			},
			{
				Name: "check-db",
				Action: func(ctx *cli.Context) error {
					return CheckDbConfig(cfg)
				},
			},
			{
				Name: "migrate-force",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "version", Aliases: []string{"v"}},
				},
				Action: func(ctx *cli.Context) error {
					return MigrateForce(cfg, ctx.String("version"))
				},
			},
			{
				Name: "run-pg-rest",
				Action: func(ctx *cli.Context) error {
					return RunPGRest(cfg)
				},
			},
		},
	}

}
