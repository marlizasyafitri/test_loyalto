package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// input jumlah pemain dan jumlah dadu
	var jml_pemain, jml_dadu int
	fmt.Println("masukkan jumlah pemain : ")
	fmt.Scan(&jml_pemain)
	fmt.Println("masukkan jumlah dadu : ")
	fmt.Scan(&jml_dadu)

	fmt.Println("Jumlah Pemain : ", jml_pemain, ", Jumlah Dadu : ", jml_dadu)

	var daduPemain = make([]int, jml_pemain)  // jumlah dadu masing-masing pemain
	var angkaDadu = make(map[int][]int)       // angka yang di keluarkan oleh tiap-tiap dadu
	var tempDadu = make(map[int][]int)        // temporary variable utk angka yang di keluarkan oleh tiap-tiap dadu
	var pointPemain = make([]int, jml_pemain) // point yang dimiliki masing-masing pemain

	// inisialisasi awal jumlah dadu yang dimiliki pemain
	for a := 0; a < jml_pemain; a++ {
		daduPemain[a] = jml_dadu
	}

	// Process dadu
	for b := 1; b > 0; b++ {
		fmt.Println("Lemparan Dadu ke-", b)

		for c := 0; c < jml_pemain; c++ {
			var dadu []int

			for d := 0; d < daduPemain[c]; d++ {
				randInt := rand.Intn(6) + 1 // maximal 6, minimal 1
				dadu = append(dadu, randInt)
			}
			angkaDadu[c] = dadu
			tempDadu[c] = angkaDadu[c]
			fmt.Println("dadu Pemain ke ", c, " : ", angkaDadu[c], ", Point = ", pointPemain[c])
		}
		// Evaluasi hasil lemparan dadu
		fmt.Println("#evaluasi")

		for z := 0; z < jml_pemain; z++ {
			var hasilDadu = make(map[int][]int)

			for l, y := range angkaDadu[z] {
				if y == 6 {
					pointPemain[z]++
				} else if y == 1 && l < len(angkaDadu[z])-1 {
					if z == jml_pemain-1 {
						angkaDadu[0] = append(angkaDadu[0], y)
					} else {
						angkaDadu[z+1] = append(angkaDadu[z+1], y)
					}
				} else if y == 1 && l == len(tempDadu[z])-1 {
					angkaDadu[z+1] = append(angkaDadu[z+1], y)
				} else {
					hasilDadu[z] = append(hasilDadu[z], y)
				}
			}
			angkaDadu[z] = hasilDadu[z]
		}

		sisaPemain := jml_pemain
		for i := 0; i < jml_pemain; i++ {

			fmt.Println("dadu Pemain ke ", i, " : ", angkaDadu[i], ", Point = ", pointPemain[i])
			daduPemain[i] = len(angkaDadu[i])

			if len(angkaDadu[i]) == 0 {
				sisaPemain--
			}
		}

		if sisaPemain <= 1 {
			break
		}
		fmt.Println()
	}

}
