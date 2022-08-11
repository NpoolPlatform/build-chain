package db

import (
	"context"
	"fmt"
	"log"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/build-chain/pkg/db/ent"

	_ "github.com/mattn/go-sqlite3"

	// ent policy runtime
	_ "github.com/NpoolPlatform/build-chain/pkg/db/ent/runtime"
)

func client() (*ent.Client, error) {
	client, err := ent.Open("sqlite3", "buildchain.sqlite.db?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client, nil
}

func Init() error {
	cli, err := client()
	if err != nil {
		return err
	}
	return cli.Schema.Create(context.Background())
}

func Client() (*ent.Client, error) {
	return client()
}

func WithTx(ctx context.Context, fn func(ctx context.Context, tx *ent.Tx) error) error {
	cli, err := Client()
	if err != nil {
		return err
	}

	tx, err := cli.Tx(ctx)
	if err != nil {
		return fmt.Errorf("fail get client transaction: %v", err)
	}

	succ := false
	defer func() {
		if !succ {
			err := tx.Rollback()
			if err != nil {
				logger.Sugar().Errorf("fail rollback: %v", err)
				return
			}
		}
	}()

	if err := fn(ctx, tx); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %v", err)
	}

	succ = true
	return nil
}

func WithClient(ctx context.Context, fn func(ctx context.Context, cli *ent.Client) error) error {
	cli, err := Client()
	if err != nil {
		return fmt.Errorf("fail get db client: %v", err)
	}

	if err := fn(ctx, cli); err != nil {
		return err
	}
	return nil
}
