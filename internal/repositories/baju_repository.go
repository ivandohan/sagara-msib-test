package repositories

import (
	"database/sql"
	"sagara-msib-test/internal/entities"
)

type BajuRepository interface {
	// POST
	Create(baju entities.Baju) error

	// GET
	GetByID(id int) (entities.Baju, error)
	GetAll() ([]entities.Baju, error)
	GetBajuOrderByEmptyStok() ([]entities.Baju, error)
	GetBajuOrderByStok(stok int, kondisi string) (bajuList []entities.Baju, err error)

	// PUT
	Update(baju entities.Baju) error

	// DELETE
	Delete(id int) error
}

type bajuRepository struct {
	db *sql.DB
}

func NewInventoryBajuRepository(db *sql.DB) (ibr BajuRepository) {
	ibr = &bajuRepository{
		db: db,
	}
	return ibr
}

func (br *bajuRepository) Create(baju entities.Baju) error {
	_, err := br.db.Exec(`INSERT INTO baju (warna, ukuran, harga, stok, nama, brand) VALUES ($1, $2, $3, $4, $5, $6)`,
		baju.Warna, baju.Ukuran, baju.Harga, baju.Stok, baju.Nama, baju.Brand)
	return err
}

func (br *bajuRepository) GetByID(id int) (entities.Baju, error) {
	row := br.db.QueryRow(`SELECT id, warna, ukuran, harga, stok, nama, brand FROM baju WHERE id = $1`, id)
	var baju entities.Baju
	err := row.Scan(&baju.ID, &baju.Warna, &baju.Ukuran, &baju.Harga, &baju.Stok, &baju.Nama, &baju.Brand)
	if err != nil {
		return entities.Baju{}, err
	}
	return baju, nil
}

func (br *bajuRepository) GetAll() (bajuList []entities.Baju, err error) {
	rows, err := br.db.Query(`SELECT id, warna, ukuran, harga, stok, nama, brand FROM baju`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var baju entities.Baju
		if err := rows.Scan(&baju.ID, &baju.Warna, &baju.Ukuran, &baju.Harga, &baju.Stok, &baju.Nama, &baju.Brand); err != nil {
			return nil, err
		}
		bajuList = append(bajuList, baju)
	}
	return bajuList, nil
}

func (br *bajuRepository) GetBajuOrderByEmptyStok() (bajuList []entities.Baju, err error) {
	rows, err := br.db.Query(`SELECT id, warna, ukuran, harga, stok, nama, brand FROM baju WHERE stok=0`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var baju entities.Baju
		if err := rows.Scan(&baju.ID, &baju.Warna, &baju.Ukuran, &baju.Harga, &baju.Stok, &baju.Nama, &baju.Brand); err != nil {
			return nil, err
		}
		bajuList = append(bajuList, baju)
	}
	return bajuList, nil
}

func (br *bajuRepository) GetBajuOrderByStok(stok int, kondisi string) (bajuList []entities.Baju, err error) {
	var (
		rows *sql.Rows
	)

	if kondisi == ">" {
		rows, err = br.db.Query(`SELECT id, warna, ukuran, harga, stok, nama, brand FROM baju WHERE stok>$1`, stok)
	} else if kondisi == "<" {
		rows, err = br.db.Query(`SELECT id, warna, ukuran, harga, stok, nama, brand FROM baju WHERE stok<$1`, stok)
	} else if kondisi == "=" {
		rows, err = br.db.Query(`SELECT id, warna, ukuran, harga, stok, nama, brand FROM baju WHERE stok=$1`, stok)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var baju entities.Baju
		if err := rows.Scan(&baju.ID, &baju.Warna, &baju.Ukuran, &baju.Harga, &baju.Stok, &baju.Nama, &baju.Brand); err != nil {
			return nil, err
		}
		bajuList = append(bajuList, baju)
	}
	return bajuList, nil
}

func (br *bajuRepository) Update(baju entities.Baju) (err error) {
	_, err = br.db.Exec(`UPDATE baju SET warna = $1, ukuran = $2, harga = $3, stok = $4, nama = $5, brand = $6 WHERE id = $7`,
		baju.Warna, baju.Ukuran, baju.Harga, baju.Stok, baju.Nama, baju.Brand, baju.ID)
	return err
}

func (br *bajuRepository) Delete(id int) (err error) {
	_, err = br.db.Exec(`DELETE FROM baju WHERE id = $1`, id)
	return err
}
