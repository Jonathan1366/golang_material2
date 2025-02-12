package utils

import (
	"crypto/rand"
	"fmt"
	"regexp"
	"testing"
)

func GenerateUserId() (string, error) {
	var uuid [16]byte
	_, err := rand.Read(uuid[:])
	if err != nil {
		return "", err
	}
	//set uuid version and variants acording to RFC 4122
	uuid[6]=(uuid[6] & 0x0f)|0x40 //set version to 4
	uuid[8]=(uuid[8] & 0x3f)|0x80 //set version to 4

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
	
}

func TestUuid(t*testing.T)  {
	id_pembeli, err:=GenerateUserId()
	if err != nil {
		t.Fatalf("GenerateUserId failed: %v", err)
	}
	if id_pembeli=="" {
		t.Log("Succeded\n")		
	}

	// Define a regular expression for UUID v4
	// UUID v4 format: xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx
	// where y is one of [8, 9, A, B]

	re:=regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
	if !re.MatchString(id_pembeli){
		t.Errorf("GenerateUserId returned an invalid UUID: %s", id_pembeli)
	} else{
		// Log the result if it's valid
		t.Logf("GenerateUserId succeeded. Generated UUID: %s", id_pembeli)
	}
}
