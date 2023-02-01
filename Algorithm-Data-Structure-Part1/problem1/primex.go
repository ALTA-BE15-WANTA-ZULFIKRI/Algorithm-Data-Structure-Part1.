package main 

import ( 
"fmt"
) 
// fungsi primeX untuk mencari bilangan prima berikutnya dari angka yang diinputkan
func primeX(number int) int{ 
	// inisialisasi variabel prime sebagai nilai awal
	prime := number
	// looping untuk mencari bilangan prima berikutnya
	for {
		// variabel prime ditambah 1 setiap looping
		prime++
		// inisialisasi variabel isPrime sebagai nilai awal
		isPrime := true
		// looping untuk mengecek apakah bilangan prima atau bukan
		for i := 2; i < prime; i++ {
			// jika hasil modulus dari prime dengan i sama dengan 0, maka variabel isPrime bernilai false
			if prime%i == 0 {
				isPrime = false
				break
			}
		}
		// jika variabel isPrime bernilai true, maka looping akan berhenti dan mengembalikan nilai prime
		if isPrime {
			return prime
		}
	}
} 

func main ()  { 
// mencetak nilai dari fungsi primeX dengan parameter angka yang berbeda-beda
fmt.Println(primeX(1))  //2
fmt.Println(primeX(5)) //11 
fmt.Println(primeX(8)) //19 
fmt.Println(primeX(9)) //23
fmt.Println(primeX(10)) //29
}