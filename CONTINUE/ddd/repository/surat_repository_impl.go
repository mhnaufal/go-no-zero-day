package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"no-zero-day/CONTINUE/ddd/entity"
	"strconv"
)

type suratRepositoryImpl struct {
	DB *sql.DB
}

func NewSuratRepository(db *sql.DB) SuratRepository {
	return &suratRepositoryImpl{DB: db}
}

func (repository *suratRepositoryImpl) Insert(ctx context.Context, surat entity.Surat) (entity.Surat, error) {
	queryInsert := "INSERT INTO surat(nama) VALUES(?);"
	result, err := repository.DB.ExecContext(ctx, queryInsert, surat.Nama)
	if err != nil {
		log.Fatalln("Gagal memasukkan data surat baru")
		return surat, err
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		log.Fatalln("Gagal mendapatkan data surat baru")
		return surat, err
	}
	surat.Id = int32(insertedId)

	log.Println("Berhasil menambahkan data surat baru dengan ID: ", insertedId)
	return surat, nil
}

func (repository *suratRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Surat, error) {
	querySelect := "SELECT id, nama FROM surat where id = ? LIMIT 1;"
	surat := entity.Surat{}
	result, err := repository.DB.QueryContext(ctx, querySelect, id)
	if err != nil {
		log.Fatalln("Gagal mencari surat dengan ID: ", id)
		return surat, err
	}
	defer result.Close()

	if result.Next() {
		err = result.Scan(&surat.Id, &surat.Nama)
		if err != nil {
			return surat, nil
		}
		return surat, nil
	} else {
		return surat, errors.New("Surat ID: " + strconv.Itoa(int(id)) + " tidak ada")
	}
}

func (repository *suratRepositoryImpl) FindAll(ctx context.Context) ([]entity.Surat, error) {
	querySelect := "SELECT id, nama FROM surat;"
	result, err := repository.DB.QueryContext(ctx, querySelect)
	if err != nil {
		log.Fatalln("Gagal mengambil semua data surat")
		return nil, err
	}
	defer result.Close()

	var surats []entity.Surat
	for result.Next() {
		surat := entity.Surat{}
		result.Scan(&surat.Id, &surat.Nama)
		surats = append(surats, surat)
	}

	return surats, nil
}
