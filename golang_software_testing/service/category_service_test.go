package service

import (
	"fmt"
	"golang_software_testing/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var CategoryRepository = &repository.CategoryRepositoryMock{Mock:mock.Mock{}}
var categoryService = CategoryService{Repository: CategoryRepository}
func TestCategoryService_Get(t*testing.T){
// program mock ny
CategoryRepository.Mock.On("FindById","1").Return (nil)
category, err:=categoryService.Get("1")
assert.Nil(t, category)
assert.NotNil(t, err)

}

//BENCHMARK
// func BenchmarkHelloWorld(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		fmt.Println("hai Jonathan")
// 	}
	
// }

// func BenchmarkHelloWorld2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		fmt.Println("hai Christopher")
// 	}
	
// }

//SUB BENCHMARK
// func BenchmarkSub(b *testing.B) {
// 	b.Run("Jonathan", func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 					fmt.Println("hai Jonathan")
// 				}
// 	})
// 	b.Run("Christoper", func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 					fmt.Println("hai Christopher")
// 				}
// 	},
// )}
func BenchmarkSub(b *testing.B) {
	b.Run("Jonathan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fmt.Println("hai Jonathan")
		}
	})
	b.Run("Christoper", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fmt.Println("hai Christopher")
		}
	})
	
}

// TABLE TESTING BENCHMARK
// func BenchmarkTable(b *testing.B){
// 	benchmark:= []struct{
// 		name string
// 		request string
// 	}{
// 		name:"Jonathan Farrel Emanuel",
// 		request: "coca-cola",
// 	};
// 	{
// 		name:="Jonathan Farrel Emanuel";
// 		request:= "coca-cola",
// 	}
// 	for _, bm := range benchmark{
// 		b.Run(bm.name, 	func (b *testing.B)  {
// 			for i := 0; i < count; i++ {
// 				BenchmarkHelloWorld(benchmark.request)
// 			}
// 		})
// 	}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Jonathan Farrel Emanuel",
			request: "coca-cola",
		},
		{
			name:    "Farrel",
			request: "Sprite",
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				 BenchmarkHelloWorld(bm.request)
			}
		})
	}
}


func Benchmark123 (b*testing.B){
	result:=[] struct{
		id int
		names string
	}{
		{id: 32220020, 
		names:"Jonathan Farrel Emanuel",
	},
		{id: 32220021, 
			names:"Budi Pekerti",
		},
	}
	for _, score:=range result{
		b.Run(score.id, func (b *testing.B)  {
			for i := 0; i < b.N; i++ {
				Benchmark123(score.names)
			}
		})
	}

}


	




	
	

