package main

import (
	"context"
	"fmt"

	"github.com/apache/arrow-adbc/go/adbc"
	"github.com/apache/arrow-adbc/go/adbc/drivermgr"
)

func main() {
	ctx := context.Background()

	if err := Run(ctx); err != nil {
		fmt.Println(err)
		return
	}
}

func Run(ctx context.Context) error {
	var drv drivermgr.Driver
	db, err := drv.NewDatabase(map[string]string{
		"driver":     "libduckdb.dylib",
		"entrypoint": "duckdb_adbc_init",
		"path":       ":memory:",
	})
	if err != nil {
		return err
	}

	cnxn, err := db.Open(ctx)
	if err != nil {
		return err
	}
	defer cnxn.Close()

	if err := countRowsInTable(ctx, cnxn); err != nil {
		return err
	}

	if err := selectAll(ctx, cnxn); err != nil {
		return err
	}

	return nil
}

func selectAll(ctx context.Context, cnxn adbc.Connection) error {
	stmt, err := cnxn.NewStatement()
	if err != nil {
		return err
	}

	if err := stmt.SetSqlQuery(`SELECT os_name FROM 'report_data.parquet'`); err != nil {
		return err
	}
	defer stmt.Close()

	rdr, _, err := stmt.ExecuteQuery(ctx)
	if err != nil {
		return err
	}
	defer rdr.Release()

	for rdr.Next() {
		rec := rdr.Record()
		rec.Release()
	}

	if err := rdr.Err(); err != nil {
		return err
	}

	return nil
}

func countRowsInTable(ctx context.Context, cnxn adbc.Connection) error {
	stmt, err := cnxn.NewStatement()
	if err != nil {
		return err
	}
	defer stmt.Close()

	if err := stmt.SetSqlQuery(`SELECT count(*) FROM 'report_data.parquet'`); err != nil {
		return err
	}

	rdr, _, err := stmt.ExecuteQuery(ctx)
	if err != nil {
		return err
	}
	defer rdr.Release()

	for rdr.Next() {
		rec := rdr.Record()

		// fmt.Println(rec.Column(0).String())

		rec.Release()
	}

	if err := rdr.Err(); err != nil {
		return err
	}

	return nil
}
