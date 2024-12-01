package cli

import (
	"github.com/dangquangdon/gpm/config"
	"github.com/urfave/cli/v2"
)

func InitCli(cfg config.Config) *cli.App {
	return &cli.App{
		Name:  "gpm",
		Usage: "CLI to make the workflow smoother when working with PostgREST",
		Commands: []*cli.Command{
			{
				Name:  "migrate-up",
				Usage: "Run all `.up.sql` migration files in the migration directory",
				Action: func(ctx *cli.Context) error {
					return MigrateUp(cfg)
				},
			},
			{
				Name:  "migrate-down",
				Usage: "Run all `.down.sql` migration files in the migration directory. Effectively revert all the changes from the `up` files",
				Action: func(ctx *cli.Context) error {
					return MigrateDown(cfg)
				},
			},
			{
				Name:  "migrate-to",
				Usage: "Migrate to a specific version. Revert all changes from the migration after the specified version",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "version", Aliases: []string{"v"}},
				},
				Action: func(ctx *cli.Context) error {
					return MigrateTo(cfg, ctx.String("version"))
				},
			},
			{
				Name:  "make-migrations",
				Usage: "Create new migration files",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
				},
				Action: func(ctx *cli.Context) error {
					name := ctx.String("name")
					return MakeMigration(cfg, name)
				},
			},
			{
				Name:  "check-db",
				Usage: "Check the database config",
				Action: func(ctx *cli.Context) error {
					return CheckDbConfig(cfg)
				},
			},
			{
				Name:  "migrate-force",
				Usage: "Force the migration to a specific version when it's dirty",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "version", Aliases: []string{"v"}},
				},
				Action: func(ctx *cli.Context) error {
					return MigrateForce(cfg, ctx.String("version"))
				},
			},
			{
				Name:  "run-pg-rest",
				Usage: "Run PostgREST API server",
				Action: func(ctx *cli.Context) error {
					return RunPGRest(cfg)
				},
			},
		},
	}

}
