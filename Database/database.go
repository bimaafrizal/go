package Database

import (
	"database/sql"
	"time"
)

// manajemen koneksi = database pooling
// manajemen koneksi sudah otomatis
// kita bisa menentukan jumlah minimal dan maksimal koneksi yang dibuat golang
// ketika melebihi batas maka request menunggu terlebih dahulu
// SetMaxIdleConns mengatur jumlah koneksi minimal
// SetMaxOpenConns mengatur jumlah koneksi maksimal
// SetConnMaxIdleTime mengatur berapa lama koneksi yang sudah tidak digunakan akan dihapus, koneksi akan dihapus sampai jumlah koneksi minimal
// SetConnMaxLifetime berapa lama koneksi boleh digunakan, jika koneksinya minimal maka akan tetap dihapus tapi akan dibuat kembali
// closenya di unit test
func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
