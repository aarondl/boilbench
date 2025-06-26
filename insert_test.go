package main

import (
	"context"
	"database/sql"
	"github.com/gobuffalo/pop/v6"
	"testing"

	"github.com/aarondl/boilbench/gorms"
	"github.com/aarondl/boilbench/gorps"
	"github.com/aarondl/boilbench/mimic"
	"github.com/aarondl/boilbench/models"
	"github.com/aarondl/boilbench/xorms"
	"github.com/aarondl/sqlboiler/v4/boil"
	"gopkg.in/gorp.v1"
	"gorm.io/gorm"
	"xorm.io/xorm"
)

func BenchmarkGORMInsert(b *testing.B) {
	store := gorms.Jet{
		ID: 1,
	}

	exec := jetQueryInsert()
	exec.NumInput = -1
	mimic.NewResult(exec)

	gormdb, err := gorm.Open(gormMimicDialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	b.Run("gorm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := gormdb.Create(&store).Error
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkGORPInsert(b *testing.B) {
	store := gorps.Jet{
		ID: 1,
	}

	query := jetQueryInsert()
	query.NumInput = -1
	mimic.NewQuery(query)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	gorpdb := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	if err != nil {
		panic(err)
	}

	gorpdb.AddTable(gorps.Jet{}).SetKeys(true, "ID")

	b.Run("gorp", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := gorpdb.Insert(&store)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkXORMInsert(b *testing.B) {
	store := xorms.Jet{
		Id: 1,
	}

	exec := jetExec()
	exec.NumInput = -1
	mimic.NewResult(exec)

	xormdb, err := xorm.NewEngine("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("xorm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := xormdb.Insert(&store)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkBoilInsert(b *testing.B) {
	store := models.Jet{
		ID: 1,
	}

	exec := jetExec()
	exec.NumInput = -1
	mimic.NewResult(exec)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("boil", func(b *testing.B) {
		ctx := context.Background()

		for i := 0; i < b.N; i++ {
			err := store.Insert(ctx, db, boil.Infer())
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkPOPInsert(b *testing.B) {
	dsn := "postgres://BenchmarkPOPInsert"
	store := models.Jet{
		ID: 1,
	}

	exec := jetExec()
	exec.NumInput = -1
	mimic.NewResultDSN(dsn, exec)

	popdb, err := pop.NewConnection(&pop.ConnectionDetails{Driver: "mimic", Dialect: "postgres", URL: dsn})
	if err != nil {
		panic(err)
	}

	err = popdb.Open()
	if err != nil {
		panic(err)
	}

	b.Run("pop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := popdb.Save(&store)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
