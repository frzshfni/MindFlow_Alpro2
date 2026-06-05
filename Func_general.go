package main 

import (
	"fmt"
	"math"
	"os"
	"bufio"
	"os/exec"
	"runtime"
	"strings"
)

//fungsi clear terminal
func cls() {
	/*{menghapus atau melakukan clear pada terminal sehingga hasil berikutnya terlihat lebih rapi}
	//Deskripsi: Prosedur yang berfungsi untuk melakukan clear pada terminal setiap kali ia dipanggil.
	*/
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

//prosedur untuk menampilkan template header
func showHeader() {
	//{menampilkan template header atau judul saat program dijalankan}
	//Deskripsi: Prosedur yang berfungsi untuk menampilkan template header yang telah dibuat secara otomatis saat dipanggil tanpa membutuhkan parameter apapun.
	var pad string = ""
	fmt.Printf("%-45s╔═══════════════════════════════════════════════════════════════════════════════════╗\n", pad)
	fmt.Printf("%-45s║                                                                                   ║\n", pad)
	fmt.Printf("%-45s║                                                                                   ║\n", pad)
	fmt.Printf("%-45s║                                    🔮MINDFLOW☁️                                   ║\n", pad)
	fmt.Printf("%-45s║                  Your daily companion for a calmer and clearer day!               ║\n", pad)
	fmt.Printf("%-45s║                                                                                   ║\n", pad)
	fmt.Printf("%-45s║                                                                                   ║\n", pad)
	fmt.Printf("%-45s╚═══════════════════════════════════════════════════════════════════════════════════╝\n", pad)
}

//prosedur untuk menampilkan template tabel data
func showTables(T Pengguna) {
	//{menampilkan template judul tabel data suasana hati dan tabel data tugas}
	//Deskripsi: prosedur untuk menampilkan template tabel suasana dan tabel tugas yang telah dibuat secara bersampingan dari isi array dengan parameter array pengguna.
	var gap string = "   " 

	fmt.Printf("╔════════════════════════════════════════════════════════════════════════════════════════╗%s╔════════════════════════════════════════════════════════════════════════════════════════╗\n", gap)
	fmt.Printf("║                                                                                        ║%s║                                                                                        ║\n", gap)
	fmt.Printf("║                                    Tabel Suasana Hati                                  ║%s║                                      Tabel Tugas                                       ║\n", gap)
	fmt.Printf("║                                     ✨Moodie-Mood!✨                                   ║%s║                                     🚀Doie-Do!🚀                                       ║\n", gap)
	fmt.Printf("║                                                                                        ║%s║                                                                                        ║\n", gap)
	fmt.Printf("╠════╤═════════╤═══════════════════════════════════════════════════════╤═════════════════╣%s╠════╤═════════╤════════════════════════════════╤═══════════════════╤═══════════╤════════╣\n", gap)
	fmt.Printf("║ No ╎ Tanggal ╎                   Deskripsi Perasaan                  ╎    Skor Emosi   ║%s║ No ╎ Tanggal ╎           Nama Tugas           ╎ Durasi Pengerjaan ╎ Prioritas ╎ Status ║\n", gap)
	fmt.Printf("╠════╪═════════╪═══════════════════════════════════════════════════════╪═════════════════╣%s╠════╪═════════╪════════════════════════════════╪═══════════════════╪═══════════╪════════╣\n", gap)
	
	PrintTabel(T)
	
	fmt.Printf("╚════╧═════════╧═══════════════════════════════════════════════════════╧═════════════════╝%s╚════╧═════════╧════════════════════════════════╧═══════════════════╧═══════════╧════════╝\n", gap)
}

//prosedur untuk menampilkan isi tabel data
func PrintTabel(data Pengguna){
	//{I.S.terdefinisi array data pengguna
	//F.S.menampilkan array data sesuai dengan template susunan pada tabel}
	//Deskripsi: Prosedur yang berfungsi untuk mengisi data pada tabel sesuai dengan format yang telah dirancang dengan array data yang ingin ditampilkan.
	var maxRows int
	maxRows = data.T_suasana
	if data.T_tugas > maxRows {
		maxRows = data.T_tugas
	}

	if maxRows == 0 {
		fmt.Printf("║ %-3s╎ %-8s╎ %-54s╎ %-16s║%s║ %-3s╎ %-8s╎ %-31s╎ %-18s╎ %-10s╎ %-7s║\n",
			"", "", "", "", "", "   ", "", "", "", "", "", "")
	}
	
	for i := 0; i < maxRows; i++ {
		// DATA SUASANA (KIRI)
		tglS, deskripsi, skor, sNo := "", "", "", ""
		if i < data.T_suasana {
			s := data.daftar_suasana[i]
			sNo = fmt.Sprintf("%d", i+1)
			tglS = fmt.Sprintf("%02d/%02d/%02d", s.day, s.mon, s.year%100) 
			
			deskripsi = s.deskripsi_perasaan
			skor = fmt.Sprintf("%d", s.skor_emosi)
		}

		// DATA TUGAS (KANAN)
		tglT, namaT, durasi, prioritas, status, tNo := "", "", "", "", "", ""
		if i < data.T_tugas {
			t := data.daftar_tugas[i]
			tNo = fmt.Sprintf("%d", i+1)
			tglT = fmt.Sprintf("%02d/%02d/%02d", t.day, t.mon, t.year%100)
			
			namaT = t.nama_tugas
			
			durasi = fmt.Sprintf("%d Menit", t.durasi_pengerjaan) 
			prioritas = fmt.Sprintf("%d", t.Prioritas)
			
			if t.selesai == 1 {
				status = "Selesai"
			} else {
				status = "Belum"
			}
		}

		// Cetak satu baris untuk kedua tabel (kiri dan kanan)
		fmt.Printf("║ %-3s╎ %-8s╎ %-54s╎ %-16s║%s║ %-3s╎ %-8s╎ %-31s╎ %-18s╎ %-10s╎ %-7s║\n",
			sNo, tglS, deskripsi, skor, "   ", tNo, tglT, namaT, durasi, prioritas, status)
	}

}

func inputDataTugas(data *Pengguna) {
	/* I.S. terdefinisi array T yang kosong atau terisi data minimal 1
	F.S. array data terisi data baru ke n*/
	//Deskripsi: Prosedur yang berfungsi untuk melakukan input dari user untuk mengisi array daftar_tugas yang 
	//merupakan bagian dari array Pengguna. User diminta untuk melakukan input temp yang merupakan nama tugas, 
	//lalu diikuti dengan input tanggal, durasi pengerjaan, prioritas, dan status penyelesaian.
	var temp string
	var i int
	fmt.Println()
	fmt.Print("⟢ ")
	fmt.Scan(&temp)	
	i = data.T_tugas
	for temp != "END" && i < arrmax {
		data.daftar_tugas[i].nama_tugas = temp
		fmt.Scanf("%d/%d/%d", &data.daftar_tugas[i].day, &data.daftar_tugas[i].mon, &data.daftar_tugas[i].year)
		fmt.Scan(&data.daftar_tugas[i].durasi_pengerjaan, &data.daftar_tugas[i].Prioritas, &data.daftar_tugas[i].selesai)
		fmt.Scan(&temp)
		i++
	}
	data.T_tugas = i
}
func inputDataSuasana(data *Pengguna, kataPart *arrTemp2){
	/*I.S. terdefinisi array data suasana pengguna kosong atau ada minimal 1.
	Proses : data suasana untuk pengguna secara berurutan diinput dan Input untuk seorang pengguna berhenti jika tanggal diisi "0 / stop"
	F.S. data suasana untuk masing-masing pengguna terisi sebanyak k dan disimpan di dalam array data T_suasana pengguna 
	Deskripsi: prosedur yang berfungsi untuk menerima input dari user untuk mengisi array daftar_suasana yang merupakan bagian dari array Pengguna.
	User diminta untuk melakukan input temp yang merupakan skor emosi, lalu diikuti dengan input tanggal dan deskripsi perasaaan, prosedur ini memanggil prosedur lain
	yaitu isiDeskripsi untuk menyempurnakan hasil input deskripsi perasaan.*/
	var temp int
	var k int 
	
	fmt.Println()
	fmt.Print("⟢ ")
	fmt.Scan(&temp)
	fmt.Println()
	k = data.T_suasana
	for temp != -1 && k < arrmax{
		data.daftar_suasana[k].skor_emosi = temp
		fmt.Scanf("%d/%d/%d", &data.daftar_suasana[k].day, &data.daftar_suasana[k].mon, &data.daftar_suasana[k].year)
		isiDeskripsi(data, kataPart, k)
		fmt.Scan(&temp)
		k++
	}
	data.T_suasana = k
}

// Input kata dengan spasi dan diakhiri dengan titik (.)
func isiDeskripsi(P *Pengguna, kataPart *arrTemp2, i int) {
	//{I.S.terdefinisi array P, array kataPart, serta i berupa index array
	//F.S.menggabungkan seluruh array kataPart menjadi sebuah string menggunakan prosedur gabungDeskripsi dan menyimpannya di array P}
	//Deskripsi: prosedur yang berfungsi untuk menerima string deskrisi perasaan perkata kata berdasarkan yang telah diketahui dari prosedur inputDataSuasana
	//ke dalam array p index ke i. Prosedur ini memanggil prosedur gabungDeskripsi untuk membantu proses penggabungan array.
	var j int = 0

	fmt.Scan(&kataPart[j])
	for !strings.Contains(kataPart[j], ".") {
		j++
		fmt.Scan(&kataPart[j])
	}
	panjangKata := j + 1
	P.daftar_suasana[i].deskripsi_perasaan = gabungDeskripsi(kataPart, panjangKata)
}

//fungsi untuk menggabung kata kata deskripsi perasaan
func gabungDeskripsi(kataPart *arrTemp2, panjang int) string {
	//{mengembalikan hasil berupa gabungan array kataPart menjadi sebuah string}
	//Deskripsi: Fungsi yang befungsi untuk menggabungkan kata kata dari setiap isi dari 
	//index array kataPart yang telah diketahui dari prosedur isiDeskripsi menjadi sebuah
	//string baru berupa kalimat dengan spasi
	var j int
	var hasil string
	
	for j = 0; j < panjang; j++ {
		if j > 0 {
			hasil += " "p
		}
		hasil += kataPart[j]
	}
	
	return hasil
}

// fungsi delete//
func deleteData(data *Pengguna)bool{
	/*I.S. terdefinisi array data > 0
	F.S. data pengguna terdelete sesuai dengan input dari pengguna	
	Deskripsi: prosedur ini berfungsi untuk menghapus data array daftar_suasana atau daftar_tugas 
	dari array Pengguna pada index tertentu dengan meminta pengguna untuk menginput nama dari
	tabel beserta nomor dari data yang ingin dihapus. Setelah data dihapus, data dari index setelahnya akan naik 
	untuk menggantikan sehingga jumlah index array yang terisi berkurang sebesar 1.
	*/
	var jenis string
	var Nomer int
	var status bool
	fmt.Print("Nama tabel dari data yang akan diubah (Tugas/Suasana): ")
	fmt.Scan(&jenis)
	fmt.Print("Data yang akan dihapus nomor:")
	fmt.Scan(&Nomer)
	if jenis == "Tugas"{
		for i:= Nomer; i < data.T_tugas; i++{
			data.daftar_tugas[i-1] = data.daftar_tugas[i]
		}
		data.T_tugas--
		status = true  
	}else if jenis == "Suasana"{
		for i:= Nomer; i < data.T_suasana; i++{
			data.daftar_suasana[i-1] = data.daftar_suasana[i]
		}
		data.T_suasana--
		status = true
	}else{
		status = false
	}
	return status
}

//fungsi mengubah data
func UbahData(data *Pengguna, kataPart *arrTemp2){
	/*I.S. terdefini array data > 0
	F.S. mengubah data i yang berada di dalam array
	Deskripsi: Prosedur yang berfungsi untuk mengubah data yang berada pada array index tertentu
	dengan data baru yang sesuai dengan input data yang baru. Pengguna diminta menginput nama dari
	tabel beserta nomor dari data yang ingin diganti.
	*/
	var Jenis string
	var Nomer int
	fmt.Print("Nama tabel dari data yang akan diubah (Tugas/Suasana): ")
	fmt.Scan(&Jenis)
	fmt.Print("Data yang akan diubah nomor:")
	fmt.Scan(&Nomer)
	if Jenis == "Tugas"{
		fmt.Println("Masukkan data nama tugas, tanggal(), durasi, skala prioritas, status selesai (1 = selesai, 0 = belum) tugas secara berurutan")
		fmt.Println("-Contoh format: Tugas_Alpro_Sorting 13/06/2027 30 2 1-")
		fmt.Scan(&data.daftar_tugas[Nomer-1].nama_tugas)
		fmt.Scanf("%d/%d/%d", &data.daftar_tugas[Nomer-1].day, &data.daftar_tugas[Nomer-1].mon, &data.daftar_tugas[Nomer-1].year)
		fmt.Scan(&data.daftar_tugas[Nomer-1].durasi_pengerjaan)
		fmt.Scan(&data.daftar_tugas[Nomer-1].Prioritas)
		fmt.Scan(&data.daftar_tugas[Nomer-1].selesai)
		fmt.Println("Note!")
	}else if Jenis == "Suasana"{
		fmt.Println("Masukkan skor emosi, tanggal, dan deskripsi perasaan secara berurutan! (❁´◡❁)`")
		fmt.Println("-Contoh format: 50 13/06/2027 aku mau makan.-")
		fmt.Scan(&data.daftar_suasana[Nomer-1].skor_emosi)
		fmt.Scanf("%d/%d/%d", &data.daftar_suasana[Nomer-1].day, &data.daftar_suasana[Nomer-1].mon, &data.daftar_suasana[Nomer-1].year)
		isiDeskripsi(data, kataPart, Nomer-1)
		fmt.Println("Note!")
	}

}
// fungsi presentase tugas
func presentaseTugasHarian(data *Pengguna, hari int, bulan int, tahun int)float64{
	/*I.S. terdefinisi array data, hari, bulan dan tahun dalam integer
	F.S. mengembalikan presentase tugas yang selesai pada hari bulan dan tahun tersebut 
	Deskripsi: Fungsi ini berfungsi untuk menampilkan presentase tugas pada hari, bulan dan tahun tertentu.
	Fungsi ini memanggil prosedur totTDays untuk menghitung jumlah hari berdasarkan hari, bulan, dan tahun 
	dari masing-masing index data tugas.
	*/
	var tHari, done, tTugas int
	done = 0
	tTugas = 0
	tHari = hari + ((bulan-1)*30) + ((tahun-1)*365)
	totTDays(data)
	for i:= 0; i < data.T_tugas; i++{
		if data.daftar_tugas[i].days == tHari{
			tTugas++
			if data.daftar_tugas[i].selesai == 1{
				done++
			}
		}
	}
	if tTugas == 0{
		return 0
	}else{
		return float64(done)/float64(tTugas) * 100
	}
}
//fungsi presentase suasana
func presentaseSuasanaMingguan(data *Pengguna, Temp *arrTemp){
	/*I.S. terdefinisi array data dan array Temp
	F.S. menampilkan grafik batang horizontal dari rentang waktu kurang lebih 7 hari/ 1 minggu 
	Deskripsi: prosedur ini berfungsi untuk menampilkan presentase suasana pada hari, bulan dan tahun tertentu dalam rentang kurang lebih 7 hari pada isi dari array daftar_suasana 
	pada array Pengguna. prosedur ini memanggil prosedur GrafikTrenMood untuk menampilkan presentase dalam grafik batang horizontal.
	Prosedur ini juga memanggil totMDays untuk menghitung jumlah hari berdasarkan hari, bulan, dan tahun 
	dari masing-masing index data suasana hati.
	*/
	var h1, b1, t1, h2, b2, t2 int
	var tglAwal, tglAkhir string

	fmt.Println("Format rentang tanggal: DD/MM/YYYY DD/MM/YYYY (pisahkan dengan spasi)")
	fmt.Print("Masukkan rentang: ")
	fmt.Scan(&tglAwal, &tglAkhir)

	fmt.Sscanf(tglAwal, "%d/%d/%d", &h1, &b1, &t1)
	fmt.Sscanf(tglAkhir, "%d/%d/%d", &h2, &b2, &t2)
	
	var Tmin, Tmaks, n int
	Tmin = h1 + ((b1-1)*30) + ((t1-1)*365)
	Tmaks = h2 + ((b2-1)*30) + ((t2-1)*365)
	
	totMDays(data)
	
	n = 0
	for i:= 0; i < data.T_suasana; i++{
		if data.daftar_suasana[i].days >= Tmin && data.daftar_suasana[i].days <= Tmaks{
			Temp[n] = float64(data.daftar_suasana[i].skor_emosi)
			n++
		}
	}
	GrafikTrenMood(Temp, n)

}

// fungsi tren grafik
func GrafikTrenMood(data *arrTemp, n int) {
	/* I.S. terdefinisi data Array yang sudah terisi.
	F.S. : Menampilkan grafik batang horizontal 
	Deskripsi: procedur ini menampilkan tren grafik yang sudah di adjust ke dalam skala 1 - 10 dari 
	1 - 100 dengan menggunakan linear scaling atau normalisasi min-max
	untuk menampilkan grafik batang secara horizontal langsung di terminal.
	*/

	var i int
	var nilai float64
	var nilaiBaru float64
	var nilaiBulat int
	var barGrafik string

	if n == 0 {
		fmt.Println("Data kosong!")
	}else{

		fmt.Printf("\n\n====== Grafik Tren Mood =====\n")
		fmt.Println("Data Asli | Skala (1-10) | Grafik Tren")
		fmt.Println("--------------------------------------------------")

		for i = 0; i < n; i++ {
			nilai = data[i]

			//rumus linear scaling atau rumus normalisasi min - max
			nilaiBaru = ((nilai - 1.0) / (100.0 - 1.0)) * (10.0 - 1.0) + 1.0 
			
			if nilaiBaru > 10.0 {
				nilaiBaru = 10.0
			} else if nilaiBaru < 1.0 {
				nilaiBaru = 1.0
			}

			nilaiBulat = int(math.Round(nilaiBaru))
			barGrafik = strings.Repeat("█", nilaiBulat)

			fmt.Printf("%9.0f | %12.1f | %s\n\n", nilai, nilaiBaru, barGrafik)
		}
		fmt.Println() 
	}
}
// fungsi load data array dummy
func loadData(data *Pengguna) {
	/*I.S. terdefinisi array data kosong
	F.S Array data terisi dengan data yang di load dari file sData.txt
	Deskripsi: procedur ini membuka dan melakukan scan pada file sData.txt untuk dimuat ke dalam array data pengguna 
	*/
	file, err := os.Open("sData.txt")
	if err != nil {
		fmt.Println("Millow: Dummy data tidak ditemukan. Mulai dengan tabel kosong! ☁️")
	}else{
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			p := strings.Split(scanner.Text(), "|") 

			if p[0] == "TUGAS" && len(p) == 6 && data.T_tugas < arrmax {
				t := &data.daftar_tugas[data.T_tugas]
				
				t.nama_tugas = p[1]
				fmt.Sscanf(p[2], "%d/%d/%d", &t.day, &t.mon, &t.year)
				fmt.Sscanf(p[3], "%d", &t.durasi_pengerjaan)
				fmt.Sscanf(p[4], "%d", &t.Prioritas)
				fmt.Sscanf(p[5], "%d", &t.selesai)
				
				data.T_tugas++
				
			} else if p[0] == "SUASANA" && len(p) == 4 && data.T_suasana < arrmax {
				s := &data.daftar_suasana[data.T_suasana] 
				
				fmt.Sscanf(p[1], "%d", &s.skor_emosi)
				fmt.Sscanf(p[2], "%d/%d/%d", &s.day, &s.mon, &s.year)
				s.deskripsi_perasaan = p[3]
				
				data.T_suasana++
			}
		}
	}
}

// fungsi percakapan dibawah 
	//membaca log data
func openLog(){
	// Deskripsi: Prosedur ini berfungsi untuk menampilkan log percakapan yang sudah 
	// disiapkan dalam file percakapan.log.
	file, err := os.Open("percakapan.log")
	if err != nil {
		fmt.Println("Millow: Wah, belum ada riwayat percakapan nih! ☁️")
	}else{
	defer file.Close()
		var gap string = "   "
		fmt.Printf("%-45s╔══════════════════════════════════════════════════════╗\n", gap)
		fmt.Printf("%-45s║              📜 Riwayat Percakapan Millow            ║\n", gap)
		fmt.Printf("%-45s╚══════════════════════════════════════════════════════╝\n\n", gap)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}

