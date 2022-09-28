package db

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/build-chain/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/build-chain/pkg/db/ent"

	_ "github.com/mattn/go-sqlite3"

	// ent policy runtime
	_ "github.com/NpoolPlatform/build-chain/pkg/db/ent/runtime"
)

func client() (*ent.Client, error) {
	dataDir := config.GetENV().DataDir
	dbFile := "buildchain.sqlite.db"
	dbFileURL := fmt.Sprintf("%v/%v?cache=shared&_fk=1", dataDir, dbFile)
	client, err := ent.Open("sqlite3", dbFileURL)
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}
	return client, nil
}

func Init() error {
	cli, err := client()
	if err != nil {
		return err
	}
	defer cli.Close()
	return cli.Schema.Create(context.Background())
}

var entclient *ent.Client

func Client() (*ent.Client, error) {
	var err error
	if entclient != nil {
		return entclient, nil
	}
	entclient, err = client()
	return entclient, err
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
