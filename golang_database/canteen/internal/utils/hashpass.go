package utils

import (
	// "testing"

	"golang.org/x/crypto/bcrypt"
)

//hash password
func HashPass(password string) (string, error)  {
	hashed, err:= bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
	
}


//if passhash == compatibles return true

func CheckPassHash(password, hash string) bool  {
	err:=bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err==nil
}

// func TestHashPassword(t*testing.T)  {
// 	kata:="farrel123"
// 	hashedpass, err:=HashPass(kata)
// 	if err != nil {
// 		t.Errorf("HashPassword failed: %v", err)
// 		return 
// 	}
// 	if hashedpass==""{
// 		t.Errorf("hashpassword returned an empty hash")
// 		return
// 	}
// 	t.Logf("HashPassword succeeded. Hashed password: %s", hashedpass)
// }

// func TestCheckPasswordHash(t *testing.T) {
// 	password := "jonathan"
// 	hashedPassword, _ := HashPass(password)

// 	// Test for correct password
// 	if !CheckPassHash(password, hashedPassword) {
// 		t.Errorf("CheckPasswordHash returned false for correct password")
// 	}

	// Test for incorrect password
// 	wrongPassword := "jonathan2"
// 	if CheckPassHash(wrongPassword, hashedPassword) {
// 		t.Errorf("CheckPasswordHash returned true for incorrect password")
// 	}
// }


	// //test for correct password
	// if !CheckPassHash(kata, hashedpass) {
	// 	t.Errorf("Checkpassword hashed returned true for incorect password")
	// 	fmt.Println("hashed password: ", hashedpass)
	// }

	// //test wrong pass
	// wrongpass:="jonathan"
	// if CheckPassHash(wrongpass,hashedpass) {
	// 	t.Errorf("CheckPasswordHash returned true for incorrect password")
	// }




