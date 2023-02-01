package main

import "fmt"

func PrimaSegiEmpat(wide, high, start int) string {
	var result string                //variabel result digunakan untuk menyimpan hasil dari fungsi
	for i := 0; i < high; i++ {      // looping untuk mencetak baris
		for j := 0; j < wide; j++ {  // looping untuk mencetak kolom
			if isPrime(start) {      // memanggil fungsi isPrime untuk mengecek apakah angka tersebut prima atau bukan
				result += fmt.Sprintf("%d\t", start) // jika prima, maka angka tersebut akan ditambahkan ke variabel result
			}
			start++ // menambah nilai start
		}
		result += "\n" // menambahkan newline setelah setiap baris
	}

    result += fmt.Sprintf("%d", start-1)  // menambahkan nilai terakhir dari start ke variabel result
	return result   // mengembalikan nilai dari variabel result
}

func isPrime(num int) bool {    // isPrime adalah fungsi yang digunakan untuk mengecek apakah angka tersebut prima atau bukan
	if num <= 1 {     // jika angka kurang dari sama dengan 1, maka bukan prima
		return false
	}

	for i := 2; i < num; i++ { // looping untuk mengecek apakah angka tersebut habis dibagi angka lain
		if num%i == 0 {        // jika habis dibagi, maka bukan prima
			return false
		}
	}

	return true  // jika tidak habis dibagi, maka angka tersebut prima
}
}

func main() {
	fmt.Println(PrimaSegiEmpat(2, 3, 13,))  // memanggil fungsi PrimaSegiEmpat dengan parameter wide = 2, high = 3, dan start = 13
	fmt.Println(PrimaSegiEmpat(5, 2, 1,))   // memanggil fungsi PrimaSegiEmpat dengan parameter wide = 5, high = 2, dan start = 1
	// Output:
	// 2      3       5
	// 7      11      13
	// 17     19      23
    // 23
}