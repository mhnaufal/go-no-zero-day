package test

import (
	"context"
	"fmt"
	database "no-zero-day/CONTINUE/chapter11"
	"no-zero-day/CONTINUE/ddd/entity"
	"no-zero-day/CONTINUE/ddd/repository"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsertSurat(t *testing.T) {
	suratRepository := repository.NewSuratRepository(database.GetConnection())

	ctx := context.Background()
	surat := entity.Surat{
		Nama: "Surat Pendamping",
	}

	result, err := suratRepository.Insert(ctx, surat)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindByIdSurat(t *testing.T) {
	suratRepository := repository.NewSuratRepository(database.GetConnection())

	ctx := context.Background()
	suratId := 1

	result, err := suratRepository.FindById(ctx, int32(suratId))
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindAllSurat(t *testing.T) {
	suratRepository := repository.NewSuratRepository(database.GetConnection())

	ctx := context.Background()

	results, err := suratRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
