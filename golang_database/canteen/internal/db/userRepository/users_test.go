package userRepository_test

import (
	"context"
	"fmt"
	"kantin2/canteen/internal/db/driver"
	"kantin2/canteen/internal/utils"

	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

// represents the user model

type User struct{
	id_pembeli string
	nama_users string
	email string
	password string
	img_user *[]byte
}

//update profile & delete profile picture if needed (database transaction)
func UpdateProfile(tx pgx.Tx, user *User, updateImg bool, deleteImg bool) error  {

	//prepare statements
	stmUpdate := "UPDATE users SET nama_users=$1, email=$2, password=$3, img_user=$4 WHERE id_pembeli=$5"
	stmtDeleteImage := "UPDATE users SET img_user = NULL WHERE id_pembeli = $1"
	stmtUpdateImage := "UPDATE users SET img_user = $1 WHERE id_pembeli = $2"

	//hash password if it is not empty
	var hashedPassword string
	var err error
	if user.password != "" {
		hashedPassword, err = utils.HashPass(user.password)
		if err != nil {
			return err
		}
	} else {
		hashedPassword = user.password
	}

	// update profile
	_, err = tx.Exec(context.Background(), stmUpdate, user.nama_users, user.email, hashedPassword, user.img_user, user.id_pembeli)
	if err != nil {
		return err
	}

	//delete profile pic
	if deleteImg {
		_, err = tx.Exec(context.Background(), stmtDeleteImage, user.id_pembeli)
		if err != nil {
			return err
		}
	} else if updateImg && user.img_user != nil {
		_, err = tx.Exec(context.Background(), stmtUpdateImage, *user.img_user, user.id_pembeli)
		if err != nil {
			return err
		}
	}

	return nil
}


func TestUpdateProfile(t *testing.T) {
	driver.InitDBPool()
	defer driver.CloseDBPool()

	db := driver.GetDbpool()
	
	//DB ACID TRANSACTION

	//mulai transaksi
	tx, err:= db.Begin(context.Background())
	require.NoError(t, err)
	defer tx.Rollback(context.Background()) //rollback jika terjadi kesalahan

	id_pembeli := "9c8e980c-4310-48ea-b5a9-c6aa02e9d712"
	newPass := "terbarujo2"

	//buat objek user untuk pengujian
	user:=&User{
		id_pembeli: id_pembeli,
		nama_users: "Jonathan123",
		email: "joanathan123@gmail.com",
		password: newPass, // Password yang akan di-hash
		img_user: nil,
	}

	//panggil fungsi updateprofile dalam transaksi
	err=UpdateProfile(tx,user, false,  false)
	require.NoError(t, err)

	//verifikasi pembaruan dalam transaksi
	var nama_users, email, passwordHash string
	err=tx.QueryRow(context.Background(), "SELECT nama_users, email, password FROM users WHERE id_pembeli=$1", id_pembeli).Scan(&nama_users, &email, &passwordHash)
	require.NoError(t, err)

	//commit transaksi jika semua operasi berhasil
	err=tx.Commit(context.Background())
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}

	fmt.Printf("Update user Details: \n")
	fmt.Printf("Nama User: %s \n", nama_users)
	fmt.Printf("Email: %s \n", email)
	fmt.Printf("Password: %s \n", passwordHash)
	require.NoError(t,err)

	// Verifikasi hasil
	require.Equal(t, "Jonathan123", nama_users)
	require.Equal(t, "joanathan123@gmail.com", email)
	// Optionally, you can verify that the password hash is valid without checking the exact value

	fmt.Printf("Update user Details: \n")
	fmt.Printf("Nama User: %s \n", nama_users)
	fmt.Printf("Email: %s \n", email)
	fmt.Printf("Password: %s \n", passwordHash)
	require.NoError(t, err)

	// require.Equal(t, "Jonathan123", nama_users)
	// require.Equal(t, "joanathan123@gmail.com", email)
	// err = utils.CheckPassHash(newPass,passwordHash)
	// require.NoError(t, err)

}


