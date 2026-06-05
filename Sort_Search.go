package main

import (
	"fmt"
	"strings"
)

func selectionAsc(T *Pengguna, max int, sortBy int) {
	/*{I.S.terdefinisi array T yang telah berisi max data Pengguna
	F.S.(menampilkan array T yang telah terurut // notes: mungkin di while loop main) 
	mengurutkan secara Ascending berdasarkan sortBy yang telah diketahui menggunakan prosedur Selection
	1 untuk sort berdasarkan prioritas, 2 berdasarkan durasi}*/
	var pass, k, acuan int
	var temp tugas

	for pass = 1; pass < max; pass++ {
		acuan = pass - 1
		
		for k = pass; k < max; k++ {
			var isLebihKecil bool
			if sortBy == 1 {
				isLebihKecil = T.daftar_tugas[k].Prioritas < T.daftar_tugas[acuan].Prioritas
			} else if sortBy == 2 {
				isLebihKecil = T.daftar_tugas[k].durasi_pengerjaan < T.daftar_tugas[acuan].durasi_pengerjaan
			}

			if isLebihKecil {
				acuan = k
			}
		}
		temp = T.daftar_tugas[pass-1]
		T.daftar_tugas[pass-1] = T.daftar_tugas[acuan]
		T.daftar_tugas[acuan] = temp
	}
}

func selectionDes(T *Pengguna, max int, sortBy int) {
	/*{I.S.terdefinisi array T yang telah berisi max data Pengguna
	F.S.(menampilkan array T yang telah terurut // notes: mungkin di while loop main)// 
	mengurutkan secara Descending berdasarkan sortBy yang telah diketahui menggunakan prosedur Selection}
	*/
	var pass, k, acuan int
	var temp tugas

	for pass = 1; pass < max; pass++ {
		acuan = pass - 1
		for k = pass; k < max; k++ {
			var isLebihBesar bool
			if sortBy == 1 {
				isLebihBesar = T.daftar_tugas[k].Prioritas > T.daftar_tugas[acuan].Prioritas
			} else if sortBy == 2 {
				isLebihBesar = T.daftar_tugas[k].durasi_pengerjaan > T.daftar_tugas[acuan].durasi_pengerjaan
			}

			if isLebihBesar {
				acuan = k
			}
		}
		temp = T.daftar_tugas[pass-1]
		T.daftar_tugas[pass-1] = T.daftar_tugas[acuan]
		T.daftar_tugas[acuan] = temp
	}
}

func insertionAsc(T *Pengguna, max int, sortBy int) {
	//{I.S.terdefinisi array T yang telah berisi max data pengguna
	//F.S.(menampilkan array T yang telah terurut // notes: mungkin di while loop main)
	// mengurutkan secara Ascending berdasarkan sortBy yang telah diketahui menggunakan prosedur Insertion}
	var k, pass int
	var temp tugas
	var isLebihKecil, stop bool
	
	for pass = 1; pass < max; pass++ {
		stop = false
		k = pass
		temp = T.daftar_tugas[k]
		for k > 0 && !stop{
			if sortBy == 1 {
				isLebihKecil = temp.Prioritas < T.daftar_tugas[k-1].Prioritas
			} else if sortBy == 2 {
				isLebihKecil = temp.durasi_pengerjaan < T.daftar_tugas[k-1].durasi_pengerjaan
			}
			if isLebihKecil {
				T.daftar_tugas[k] = T.daftar_tugas[k-1]
				k = k - 1
			}else{
				stop = true
			}
		}
		T.daftar_tugas[k] = temp
	}
}

func insertionDes(T *Pengguna, max int, sortBy int) {
	//{I.S.terdefinisi array T yang telah berisi max data pengguna
	//F.S.(menampilkan array T yang telah terurut // notes: mungkin di while loop main)
	// mengurutkan secara Descending berdasarkan sortBy yang telah diketahui menggunakan prosedur Insertion}
	var k, pass int
	var temp tugas
	var isLebihBesar, stop bool

	
	for pass = 1; pass < max; pass++ {
		stop = false
		k = pass
		temp = T.daftar_tugas[k]
		for k > 0 && !stop {
			if sortBy == 1 {
				isLebihBesar = temp.Prioritas > T.daftar_tugas[k-1].Prioritas
			} else if sortBy == 2 {
				isLebihBesar = temp.durasi_pengerjaan > T.daftar_tugas[k-1].durasi_pengerjaan
			}
			if isLebihBesar {
				T.daftar_tugas[k] = T.daftar_tugas[k-1]
				k = k - 1
			}else{
				stop = true
			}
		}
		T.daftar_tugas[k] = temp
	}
}

func totTDays(T *Pengguna) {
	//{I.S. terdefinisi array T yang merupakan daftar tugas pengguna
	//F.S. mengisi array T[idx].days dengan total hari berdasarkan perhitungan T[idx]day, T[idx]mon, dan T[idx]year}
	var n, i int
	n = T.T_tugas
	for i = 0; i < n; i++ {
		T.daftar_tugas[i].days = T.daftar_tugas[i].day + ((T.daftar_tugas[i].mon-1)*30) + ((T.daftar_tugas[i].year-1)*365)
	}
}

func binaryTugas(T Pengguna, N, d, m, y int) {
	//{I.S.terdefinisi array T, nilai d (tanggal), m (month), dan y (year)
	//F.S. menampilkan hasil pencarian secara binary berdasarkan jumlah hari dari waktu yang diketahui}
	
	var left, mid, right int
	var days, found int
	var done bool
	totTDays(&T)
	left = 0
	right = N
	days = d + ((m-1)*30) + ((y-1)*365)
	done = false

	//sort secara Asc	
	var pass, k, acuan int
	var temp tugas
	for pass = 1; pass < N-1; pass++ {
		acuan = pass - 1
		for k = pass; k < N-1; k++ {
			if T.daftar_tugas[acuan].days > T.daftar_tugas[k].days {
				acuan = k
			}
		}
		temp = T.daftar_tugas[acuan]
		T.daftar_tugas[acuan] = T.daftar_tugas[pass-1]
		T.daftar_tugas[pass-1] = temp
	}

	for left <= right && !done {
		mid = (left+right)/2
		if days < T.daftar_tugas[mid].days {
			right = mid - 1
		} else if days > T.daftar_tugas[mid].days {
			left = mid + 1
		} else {
			found = mid
			done = true
		}
	}
	if done {
		fmt.Println("Caught ya! Detective Millow, reporting for duty! ✨(🫵⌐■_■)")
		fmt.Println()
		fmt.Println("------------------------------------------------------------------------------")
		fmt.Printf("Tanggal            : %02d/%02d/%d\n", T.daftar_tugas[found].day, T.daftar_tugas[found].mon, T.daftar_tugas[found].year)
		fmt.Printf("Nama Tugas         : %s\n", T.daftar_tugas[found].nama_tugas)
		fmt.Printf("Durasi Pengerjaan  : %d\n", T.daftar_tugas[found].durasi_pengerjaan)
		fmt.Printf("Prioritas          : %d\n", T.daftar_tugas[found].Prioritas)
		fmt.Printf("Status Penyelesaian: %d\n", T.daftar_tugas[found].selesai)
		fmt.Println("------------------------------------------------------------------------------")
		fmt.Println()
	} else {
		fmt.Println("S-sorry...Detective Millow can't find it(⁠;⁠ŏ⁠﹏⁠ŏ⁠)")
	}
}

func totMDays(M *Pengguna) {
	//{I.S. terdefinisi array T yang merupakan daftar suasana hati pengguna
	//F.S. mengisi array T[idx].days dengan total hari berdasarkan perhitungan T[idx]day, T[idx]mon, dan T[idx]year}
	var n, i int
	n = M.T_suasana
	for i = 0; i < n; i++ {
		M.daftar_suasana[i].days = M.daftar_suasana[i].day + ((M.daftar_suasana[i].mon-1)*30) + ((M.daftar_suasana[i].year-1)*365)
	}
}

func binaryMood(M Pengguna, N, d, m, y int) {
	//{I.S.terdefinisi array M, nilai d (tanggal), m (month), dan y (year)
	//F.S. menampilkan hasil pencarian secara binary berdasarkan jumlah hari dari waktu yang diketahui}
	var left, mid, right int
	var days, found int
	var done bool
	totMDays(&M)
	left = 0
	right = N
	days = d + ((m-1)*30) + ((y-1)*365)
	done = false

	//sort secara Asc	
	var pass, k, acuan int
	var temp Suasana
	for pass = 1; pass < N-1; pass++ {
		acuan = pass - 1
		for k = pass; k < N-1; k++ {
			if M.daftar_suasana[acuan].days > M.daftar_suasana[k].days {
				acuan = k
			}
		}
		temp = M.daftar_suasana[acuan]
		M.daftar_suasana[acuan] = M.daftar_suasana[pass-1]
		M.daftar_suasana[pass-1] = temp
	}

	for left <= right && !done {
		mid = (left+right)/2
		if days < M.daftar_suasana[mid].days {
			right = mid - 1
		} else if days > M.daftar_suasana[mid].days {
			left = mid + 1
		} else {
			found = mid
			done = true
		}
	} 
	if done {
		fmt.Println("Caught ya! Detective Millow, reporting for duty! ✨(🫵⌐■_■)")
		fmt.Println()
		fmt.Println("------------------------------------------------------------------------------")
		fmt.Printf("Tanggal   : %02d/%02d/%d\n", M.daftar_suasana[found].day, M.daftar_suasana[found].mon, M.daftar_suasana[found].year)
		fmt.Printf("Skor Emosi: %d\n", M.daftar_suasana[found].skor_emosi)
		fmt.Printf("Deskripsi : %s\n", M.daftar_suasana[found].deskripsi_perasaan)
		fmt.Println("------------------------------------------------------------------------------")
		fmt.Println()
	} else {
		fmt.Println("S-sorry...Detective Millow can't find it(⁠;⁠ŏ⁠﹏⁠ŏ⁠)")
	}
}


func sequentialTugas(T Pengguna, cari string, max int) {
	//{I.S.terdefinisi array T yang telah berisi max data daftar_tuga dan kata kunci cari berisi nama_tugas
	//F.S.menampilkan data tugas berdasarkan kata kunci atau tanggal yang dicari}	
	var k int
	var found bool
	found = false
	k = 0
	for k <= max-1 && !found {
		found = cari == T.daftar_tugas[k].nama_tugas              
		k = k + 1
	}
	if found {
		// revisi layout
		k = k - 1
		fmt.Println("Caught ya! Detective Millow, reporting for duty! ✨(🫵⌐■_■)")
		fmt.Println("--------------------------------------------------")
		fmt.Printf("Tanggal            : %02d/%02d/%d\n", T.daftar_tugas[k].day, T.daftar_tugas[k].mon, T.daftar_tugas[k].year)
		fmt.Printf("Nama Tugas         : %s\n", T.daftar_tugas[k].nama_tugas)
		fmt.Printf("Durasi Pengerjaan  : %d\n", T.daftar_tugas[k].durasi_pengerjaan)
		fmt.Printf("Prioritas          : %d\n", T.daftar_tugas[k].Prioritas)
		fmt.Printf("Status Penyelesaian: %d\n", T.daftar_tugas[k].selesai)
		fmt.Println("--------------------------------------------------")
	}else{
		fmt.Println("S-sorry...Detective Millow can't find it(⁠;⁠ŏ⁠﹏⁠ŏ⁠)")
	}
}

func sequentialMood(M Pengguna, cari string, max int) {
	// {I.S. terdefinisi array M yang telah berisi max data daftar_suasana
	// F.S. menampilkan SEMUA catatan emosi berdasarkan kata kunci di deskripsi_perasaan}
	
	var k int
	var found bool = false
	
	for k = 0; k < max; k++ {
		kataKunci := strings.ToLower(cari)
		dataTugas := strings.ToLower(M.daftar_suasana[k].deskripsi_perasaan)
		if strings.Contains(dataTugas, kataKunci) {
			found = true 
			
			// revisi layout
			fmt.Println("Caught ya! Detective Millow, reporting for duty! ✨(🫵⌐■_■)")
			fmt.Println("--------------------------------------------------")
			fmt.Printf("Tanggal   : %02d/%02d/%d\n", M.daftar_suasana[k].day, M.daftar_suasana[k].mon, M.daftar_suasana[k].year)
			fmt.Printf("Skor Emosi: %d\n", M.daftar_suasana[k].skor_emosi)
			fmt.Printf("Deskripsi : %s\n", M.daftar_suasana[k].deskripsi_perasaan)
		}
	}
	
	if !found {
		fmt.Println("S-sorry...Detective Millow can't find it(⁠;⁠ŏ⁠﹏⁠ŏ⁠)")
	} else {
		fmt.Println("--------------------------------------------------")
	}
}
