package main

import (
	"fmt"
)

func selectionAsc(T *Pengguna, max int, sortBy int) {
	/*{I.S.terdefinisi array T yang telah berisi max data Pengguna
	F.S.(menampilkan array T yang telah terurut // notes: mungkin di while loop main) 
	Deskripsi: Prosedur ini berfungsi untuk mengurutkan array T secara Ascending berdasarkan sortBy 
	yang telah diketahui menggunakan prosedur Selection, pengguna diminta untuk menginput nomor 1 
	untuk sort berdasarkan prioritas, dan 2 berdasarkan durasi*/
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
	Deskripsi: Prosedur ini berfungsi untuk mengurutkan array T secara Descending berdasarkan 
	sortBy yang telah diketahui menggunakan prosedur Selection, pengguna diminta untuk menginput 
	nomor 1 untuk sort berdasarkan prioritas, dan 2 berdasarkan durasi
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
	// Deskripsi: Prosedur ini berfungsi untuk mengurutkan array T secara Ascending berdasarkan sortBy yang 
	// telah diketahui menggunakan prosedur Insertion, pengguna diminta untuk menginput nomor 1 
	// untuk sort berdasarkan prioritas, dan 2 berdasarkan durasi
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
	// Deskripsi: Prosedur ini berfungsi untuk mengurutkan array T secara Descending berdasarkan 
	// sortBy yang telah diketahui menggunakan prosedur Insertion, pengguna diminta untuk 
	// menginput nomor 1 untuk sort berdasarkan prioritas, dan 2 berdasarkan durasi.
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
	//F.S. mengisi array T.daftar_tugas[idx].days dengan total hari berdasarkan perhitungan T.daftar_tugas[idx]day, 
	//T.daftar_tugas[idx]mon, dan T.daftar_tugas[idx]year}
	//Deskripsi: Prosedur ini berfungsi untuk menghitung jumlah hari berdasarkan hari, bulan,
	//dan tahun yang telah diketahui pada array pada masing-masing index lalu menyimpannya ke dalam 
	//array T.daftar_tugas[idx].days sesuai indexnya.
	var n, i int
	n = T.T_tugas
	for i = 0; i < n; i++ {
		T.daftar_tugas[i].days = T.daftar_tugas[i].day + ((T.daftar_tugas[i].mon-1)*30) + ((T.daftar_tugas[i].year-1)*365)
	}
}

func binaryTugas(T Pengguna, N, d, m, y int) {
	//{I.S.terdefinisi array T, nilai d (tanggal), m (month), dan y (year)
	//F.S. menampilkan hasil pencarian secara binary berdasarkan jumlah hari dari waktu yang diketahui}
	//Deskripsi: Prosedur ini berfungsi untuk melakukan pencarian data tertentu pada array T dengan mengurutkan array T 
	//terlebih dahulu secara ascending menggunakan metode selection lalu membandingkan jumlah hari setiap data dengan 
	//jumlah hari dari input waktu yang dicari.Prosedur ini memanggil prosedur TotTDays untuk menghitung 
	//jumlah hari berdasarkan hari, bulan, dan tahun dari masing-masing index data tugas.
	var left, mid, right, i int
	var days, found, awal, akhir int
	var done bool
	totTDays(&T)
	left = 0
	right = N-1
	days = d + ((m-1)*30) + ((y-1)*365)
	done = false

	//sort secara Asc	
	var pass, k, acuan int
	var temp tugas
	for pass = 1; pass < N; pass++ {
		acuan = pass - 1
		for k = pass; k < N; k++ {
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
		awal = found
		for awal > 0 && days == T.daftar_tugas[awal-1].days {
			awal--
		}
		akhir = found
		for akhir < N-1 && days == T.daftar_tugas[akhir+1].days {
			akhir++
		}
		for i = awal; i <= akhir; i++ {
			fmt.Println("Caught ya! Detective Millow, reporting for duty! ✨(🫵⌐■_■)")
			fmt.Println()
			fmt.Println("------------------------------------------------------------------------------")
			fmt.Printf("Tanggal            : %02d/%02d/%d\n", T.daftar_tugas[i].day, T.daftar_tugas[i].mon, T.daftar_tugas[i].year)
			fmt.Printf("Nama Tugas         : %s\n", T.daftar_tugas[i].nama_tugas)
			fmt.Printf("Durasi Pengerjaan  : %d\n", T.daftar_tugas[i].durasi_pengerjaan)
			fmt.Printf("Prioritas          : %d\n", T.daftar_tugas[i].Prioritas)
			fmt.Printf("Status Penyelesaian: %d\n", T.daftar_tugas[i].selesai)
			fmt.Println("------------------------------------------------------------------------------")
			fmt.Println()
		}
	} else {
		fmt.Println("S-sorry...Detective Millow can't find it(⁠;⁠ŏ⁠﹏⁠ŏ⁠)")
	}
}

func totMDays(M *Pengguna) {
	//{I.S. terdefinisi array M yang merupakan daftar suasana hati pengguna
	//F.S. mengisi array M.daftar_suasana[idx].days dengan total hari berdasarkan perhitungan M.daftar_suasana[idx]day, 
	//M.daftar_suasana[idx]mon, dan M.daftar_suasana[idx]year}
	//Deskripsi: Prosedur ini berfungsi untuk menghitung jumlah hari berdasarkan hari, bulan,
	//dan tahun yang telah diketahui pada array pada masing-masing index lalu menyimpannya ke dalam 
	//array M.daftar_suasana[idx].days sesuai indexnya.
	var n, i int
	n = M.T_suasana
	for i = 0; i < n; i++ {
		M.daftar_suasana[i].days = M.daftar_suasana[i].day + ((M.daftar_suasana[i].mon-1)*30) + ((M.daftar_suasana[i].year-1)*365)
	}
}

func binaryMood(M Pengguna, N, d, m, y int) {
	//{I.S.terdefinisi array M, nilai d (tanggal), m (month), dan y (year)
	//F.S. menampilkan hasil pencarian secara binary berdasarkan jumlah hari dari waktu yang diketahui}
	//Deskripsi: Prosedur ini berfungsi untuk melakukan pencarian data tertentu pada array M dengan mengurutkan array M 
	//terlebih dahulu secara ascending menggunakan metode selection lalu membandingkan jumlah hari setiap data dengan 
	//jumlah hari dari input waktu yang dicari.Prosedur ini memanggil prosedur TotMDays untuk menghitung 
	//jumlah hari berdasarkan hari, bulan, dan tahun dari masing-masing index data suasana hati.
	var left, mid, right, i int
	var days, found, awal, akhir int
	var done bool
	totMDays(&M)
	left = 0
	right = N-1
	days = d + ((m-1)*30) + ((y-1)*365)
	done = false

	//sort secara Asc	
	var pass, k, acuan int
	var temp Suasana
	for pass = 1; pass < N; pass++ {
		acuan = pass - 1
		for k = pass; k < N; k++ {
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
		awal = found
		for awal > 0 && days == M.daftar_suasana[awal-1].days {
			awal--
		}
		akhir = found
		for akhir < N-1 && days == M.daftar_suasana[akhir+1].days {
			akhir++
		}
		for i = awal; i <= akhir; i++ {
			fmt.Println("Caught ya! Detective Millow, reporting for duty! ✨(🫵⌐■_■)")
			fmt.Println()
			fmt.Println("------------------------------------------------------------------------------")
			fmt.Printf("Tanggal   : %02d/%02d/%d\n", M.daftar_suasana[i].day, M.daftar_suasana[i].mon, M.daftar_suasana[i].year)
			fmt.Printf("Skor Emosi: %d\n", M.daftar_suasana[i].skor_emosi)
			fmt.Printf("Deskripsi : %s\n", M.daftar_suasana[i].deskripsi_perasaan)
			fmt.Println("------------------------------------------------------------------------------")
			fmt.Println()
		}
	} else {
		fmt.Println("S-sorry...Detective Millow can't find it(⁠;⁠ŏ⁠﹏⁠ŏ⁠)")
	}
}


func sequentialTugas(T Pengguna, cari string, max int) {
	//{I.S.terdefinisi array T yang telah berisi max data daftar_tuga dan kata kunci cari berisi nama_tugas
	//F.S.menampilkan data tugas berdasarkan kata kunci atau tanggal yang dicari}	
	//Deskripsi: Prosedur ini berfungsi untuk mencari data tertentu pada array T dengan cara membandingkan
	//nama tugas yang dicari dengan setiap nama tugas pada array T.
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
	//Deskripsi: Prosedur ini berfungsi untuk mencari data tertentu pada array M dengan cara mengecek
	//setiap data pada array M apabila terdapat data yang deskripsi perasaannya mengandung kata yang dicari.
	var k, i, j int
	var found, isSama, cocok bool
	var deskripsi string
	found = false
	
	for k = 0; k < max; k++ {
		deskripsi = M.daftar_suasana[k].deskripsi_perasaan
		isSama = false 
		
		if len(cari) > 0 && len(cari) <= len(deskripsi) {
			
			for i = 0; i <= len(deskripsi)-len(cari) && !isSama; i++ {
				cocok = true
				
				
				for j = 0; j < len(cari) && cocok; j++ {
					if deskripsi[i+j] != cari[j] {
						cocok = false
					}
				}	
				
				if cocok {
					isSama = true 
				}
			}
		}
		
		if isSama {
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
