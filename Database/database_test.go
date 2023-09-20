package Database

// harus menggunakan _ agar kita bisa menggunakan fungsi init tanpa membutuhkan data lainnya
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

// cara memanggil database berbeda beda
// jika sudah tidak digunakan database perlu ditutup
func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang")
	if err != nil {
		panic(err)
	}

	//close db
	defer db.Close()

	// gunakan db

}
