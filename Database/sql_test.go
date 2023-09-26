package Database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

// function untuk mengirimkan perintah SQL ke database menggunakan function DB ExecContex(contex, sql, params)
// untuk query yang mengembalikan data perlu dengan cara lain
func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	//direkomendasikan menggunakan ExecContex
	query := "INSERT INTO customers(id, name) values ('rizal', 'Rizal')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

// untuk query yang yang memerlukan return maka perlu menggunakan QueryContex
func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customers"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	// hasil query adalah sebuah data structs
	// rows digunakan untuk melakukan iterasi terhadap hasil dari query
	// gunakan next untuk iterasi terhadap hasil query
	// untuk membaca tiap data gunakan Scan
	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}

	//jika sudah dieksusi rows perlu diclose
	defer rows.Close()
}

// tipe data column
// untuk menangkap data dengan tipe column selain varchar/string
// maping tipe data
// varchar, char => string
// int, bigint => int32, int64
// float, double => float32, float64
// boolean => bool
// date, datetime, time, timestamp => time.Time
func TestQueryColumn(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	//disaranklan disebutkan semua nama kolom agar kita tau posisi kolom saat mengambil data
	query := "SELECT id, name, email, balance, rating, birth_date, merried, created_at FROM customers"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var createdAt time.Time
		var birthDate sql.NullTime
		var merried bool
		// tambahkan parseTime = true pada dns agar golang bisa convert date, datatime, time,timestamp menjadi time

		var err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &merried, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("======")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate.Time)
		}
		fmt.Println("Married:", merried)
		fmt.Println("Created At:", createdAt)
	}

	defer rows.Close()
}

// nullable type
// jika kolom bisa null gunakan tipe data seperti berikut
// tipe data nullable
// string => database/sql.NullString
// bool => database/sql.NullBool
// float64 => database/sql.NullFloat64
// int32 => database/sql.NullInt32
// time.Time => database/sql.NullTime
// hasilnya akan struct

// sql injection
// sql injection teknik menyalahkan sebuah celah kemananan
// berjalan dengan cara mengirim input user dengan perintah salah sehingga hasil sql menjadi tidak valid
// jangan membuat SQL secara manual dengan menggabungkan string
// gunakan function Execute atau Query dengan parameter
func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	//username := "admin'; #"
	username := "admin"
	password := "admin"

	query := "SELECT username FROM users WHERE username = '" + username + "' AND password= '" + password + "' LIMIT 1"
	fmt.Println(query)
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

// sql dengan parameter
// untuk mengatasi SQL Injection
// untuk menandai sql perlu parameter maka perlu tambahkan ?
func TestSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin' ; #"
	password := "admin"

	query := "SELECT username FROM users WHERE username = ? AND password = ? LIMIT 1"
	// posisi params sesuai urutan query
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "bima"
	password := "bima"

	//direkomendasikan menggunakan ExecContex
	query := "INSERT INTO users(username, password) values (?, ?)"
	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}

// auto increment
// mendapatkan id setelah insert
func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "bima@gmail.com"
	comment := "test comment"

	//direkomendasikan menggunakan ExecContex
	query := "INSERT INTO comments(email, comment) values (?, ?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new comment with id", insertId)
}

// prepared statement
// statement disiapakan kemudian diisi parameter
// jika kita ingin melakukan hal yang sama sekaligus hanya berbeda parameter maka kita bisa menggunakan prepared statement
// kita bisa membuat manual tanpa menggunakan query atau exec
// saat menggunakan function query atau exec bisa jadi koneksinya berbeda sehingga kita akan terus2an meminta ke pull, lebih mudah menggunakan satu pool kemudian run querynya
// saat membuat prepare statement secara otomattis akan mengenali koneksi database yang digunakan, sehingga akan menggunakan koneksi yang sama akibatnya akan lebih cepat
// harus di close ketika sudah tidak digunakan
func TestPreparedStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	//prepare statement
	script := "INSERT INTO comments(email, comment) values (?, ?)"
	stmt, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	//cara menggunakannya
	for i := 0; i < 10; i++ {
		email := "bima" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment ke " + strconv.Itoa(i)
		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		lastInsertId, err := result.LastInsertId()
		fmt.Println("Comment id", lastInsertId)
	}
}

// database transaction
// setelah selesai transaksi kita bisa melakukan commit atau rollback
func TestTreansactionn(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email, comment) values (?, ?)"
	for i := 0; i < 10; i++ {
		email := "bima" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment ke " + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}
		lastInsertId, err := result.LastInsertId()
		fmt.Println("Comment id", lastInsertId)
	}

	//err = tx.Commit()
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
