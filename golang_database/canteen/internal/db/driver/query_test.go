package driver

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestQueryScan(t *testing.T) {
	InitDBPool()
	// defer CloseDBPool()
	
	query:=`SELECT id_pembeli, nama_users, email, password, role, img_user, saldo FROM "users"`
	ctx, cancel:= context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err:=dbPool.Query(ctx, query)
	if err!=nil {
		t.Fatalf("Query failed: %v\n", err)
	}
	defer rows.Close()

	var rowCount int
	for rows.Next(){
		var id_pembeli, nama_users, email, password, role string 
		var img_user[] byte
		var saldo string

		// scan
		err:=rows.Scan(&id_pembeli, &nama_users, &email, &password, &role, &img_user, &saldo)
		if err!=nil {
			t.Fatalf("Row Scan failed: %v\n", err)
		}
		
		// Convert saldo to float64
		var saldoFloat float64
		fmt.Scanf(saldo, "%f",saldoFloat)

		//output the result for debugging or verification
		fmt.Printf("ID Pembeli: %s, Nama: %s, Email: %s, Password: %s, Role: %s, Saldo: %.2f\n",id_pembeli, nama_users, email, password, role, saldoFloat)
			rowCount++
	}

	//ensure there were no err during iteration
	if rows.Err() != nil {
		t.Fatalf("Rows iteration failed: %v\n", rows.Err())
	}

	fmt.Printf("Total rows retrieved: %d\n", rowCount)

}


func TestQuerysql(t*testing.T)  {
	InitDBPool()
	// defer CloseDBPool()

	query:=`SELECT id_pembeli, nama_users FROM "users"`

	ctx, cancel:= context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err:=dbPool.Query(ctx, query)
	if err!=nil {
		t.Fatalf("Query failed: %v\n", err)
	}
	defer rows.Close()

	var rowCount int
	for rows.Next(){
		var id_pembeli string 
		var nama_users string 

		err:=rows.Scan(&id_pembeli, &nama_users)
		if err!=nil {
			t.Fatalf("Row Scan failed: %v\n", err)
		}
		fmt.Printf("ID Pembeli: %s\n, Nama users: %s\n ", id_pembeli, nama_users)
		rowCount++

}
	if rows.Err() != nil {
		t.Fatalf("Rows iteration failed: %v\n", rows.Err())
	}

	fmt.Printf("Total rows retrieved: %d\n", rowCount)
}

