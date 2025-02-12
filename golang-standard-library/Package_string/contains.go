// Compare returns an integer comparing two strings lexicographically. The result will be 0 if a == b, -1 if a < b, and +1 if a > b.

// Compare is included only for symmetry with package bytes. It is usually clearer and always faster to use the built-in string comparison operators ==, <, >, and so on.



package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood","foo"))
	fmt.Println(strings.Contains("seafood","bar"))
	fmt.Println(strings.Contains("seafood",""))
	fmt.Println(strings.Contains("",""))

}


// Fungsi `strings.Contains` dalam paket `strings` di Go digunakan untuk memeriksa apakah sebuah string (`s`) mengandung sub-string tertentu (`substr`). Fungsi ini mengembalikan nilai `true` jika `substr` ada di dalam `s`, dan `false` jika tidak.

// Berikut adalah penjelasan mengapa hasil untuk kalimat 3 dan 4 dari kode Anda menghasilkan `true`:

// 1. `strings.Contains("seafood", "foo")` mengembalikan `true` karena string `"seafood"` memang mengandung sub-string `"foo"`.

// 2. `strings.Contains("seafood", "bar")` mengembalikan `false` karena string `"seafood"` tidak mengandung sub-string `"bar"`.

// 3. `strings.Contains("seafood", "")` mengembalikan `true`. Ini karena dalam konteks pencarian sub-string, string kosong (`""`) dianggap ada di setiap string, termasuk di awal, di akhir, dan di antara setiap karakter dari string `s`. Oleh karena itu, setiap string dianggap mengandung string kosong.

// 4. `strings.Contains("", "")` juga mengembalikan `true` karena alasan yang sama seperti poin 3. Bahkan jika string `s` itu sendiri kosong, string kosong dianggap ada di dalamnya.

// Ini adalah perilaku yang ditentukan oleh spesifikasi Go untuk fungsi `strings.Contains`. Jadi, meskipun mungkin terdengar kontra-intuitif pada awalnya, ini adalah bagian dari definisi fungsi tersebut.