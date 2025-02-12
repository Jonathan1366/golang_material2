// package yours

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// func TestSomething(test *testing.T) {
//     // assert == fail()

//     result := assert.Equal(test, 123, 123, "they should be equal")
//     fmt.Printf("Result of first assertion: %v\n", result)

//     // check if result is true or false
//     if result {
//         test.Log("First assertion successfully")
//     } else {
//         test.Error("First assertion failed")
//     }

//     // test demo with unequal values
//     result = assert.Equal(test, 123, 456, "they should not be equal")
//     fmt.Printf("Result of second assertion: %v\n", result)

//     if !result {
//         test.Error("Second assertion failed as expected: 123 != 456")
//     }

//     fmt.Println("Test with assert done")

//     //require == failNow()

//     nama:="Jonathan"

//     require.Equal(test, "Jonathan", nama, "result must be 'Jonathan'")

//     if test.Failed() {
//         test.Error("require failed")
//     }else{
//         test.Log("passed")
//     }

// }

package yours

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	// "github.com/stretchr/testify/require"
)

func TestSomething(test *testing.T) {
	// assert == fail()

	result := assert.Equal(test, 123, 123, "they should be equal")
	fmt.Printf("Result of first assertion: %v\n", result)

	// check if result is true or false
	if result {
		test.Log("First assertion successfully")
	} else {
		test.Error("First assertion failed")
	}

	// test demo with unequal values
	result = assert.Equal(test, 123, 456, "they should not be equal")
	fmt.Printf("Result of second assertion: %v\n", result)

	if !result {
		test.Error("Second assertion failed as expected: 123 != 456")
	}

	fmt.Println("Test with assert done")

	//require == failNow()

	nama := "Jonathan"

	require.Equal(test, "jonathan",nama, "result must be 'Jonathan'")

	// Handle assertion result
	if test.Failed() {
		test.Error("Require assertion failed: expected 'Jonathan' but got something else")
	} else {
		test.Log("Require assertion passed: 'Jonathan' is equal to 'Jonathan'")
	}


}


// Kadang dalam keadaan tertentu, kita ingin membatalkan eksekusi unit test
// Di Go-Lang juga kita bisa membatalkan eksekusi unit test jika kita mau
// Untuk membatalkan unit test kita bisa menggunakan function Skip() 

func TestSkip(test2 *testing.T){
    if runtime.GOOS=="darwin" {
        test2.Skip("Skipping test on macOS")
    } else if runtime.GOOS=="windows" {
        test2.Skip("skipping test on windows")

    
    }
    fmt.Println("Running test on",runtime.GOOS)
}

func TestMain(m *testing.M){
    //before
    fmt.Println("before unit test")

    m.Run()

    //after
    fmt.Print("after unit test\n")
}

func TestSubTest(t *testing.T){
    t.Run("Jonathan_Farrel",func (t*testing.T)  {

        // result,  expected result, desc 

        assert.Equal(t,"Jonatha","Jonathan Farrel", "expected 'Jonathan Farrel'")

    })
    t.Run("Farrel", func (t*testing.T)  {
        assert.Equal(t,"Farrel","Farrel","expected 'Farrel'")
    })
}

func TestHelloWorldTable(t*testing.T){
	test:= []struct {
	name string
	request string
	expected string 
	}{
	{	
		name:"Jonathan Farrel Emanuel",
		request:"Jonathan",
		expected:"Emanuel",
	},
	
}
	for _, test:= range test{
		t.Run(test.name, func (t*testing.T)  {
			result:=test.request
			assert.Equal(t, test.expected,result)
	},
)
}
}