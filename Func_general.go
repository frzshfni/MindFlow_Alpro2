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

func cls() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func showHeader() {
	pad := ""
	fmt.Printf("%-45s╔═══════════════════════════════════════════════════════════════════════════════════╗\n", pad)
	fmt.Printf("%-45s║                                                                                   ║\n", pad)
	fmt.Printf("%-45s║                                                                                   ║\n", pad)
	fmt.Printf("%-45s║                                    🔮MINDFLOW☁️                                   ║\n", pad)
	fmt.Printf("%-45s║                  Your daily companion for a calmer and clearer day!               ║\n", pad)
	fmt.Printf("%-45s║                                                                                   ║\n", pad)
	fmt.Printf("%-45s║                                                                                   ║\n", pad)
	fmt.Printf("%-45s╚═══════════════════════════════════════════════════════════════════════════════════╝\n", pad)
}

func showTables(T Pengguna) {
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
func PrintTabel(data Pengguna){
	var maxRows int
	maxRows = data.T_suasana
	if data.T_tugas > maxRows {
		maxRows = data.T_tugas
	}

	if maxRows == 0 {
		fmt.Printf("║ %-3s╎ %-8s╎ %-54s╎ %-16s║%s║ %-3s╎ %-8s╎ %-31s╎ %-18s╎ %-10s╎ %-7s║\n",
			"", "", "", "", "", "   ", "", "", "", "", "", "")
	}

	// 3. Looping data
	for i := 0; i < maxRows; i++ {
		// DATA SUASANA (KIRI) ---
		tglS, deskripsi, skor, sNo := "", "", "", ""
		if i < data.T_suasana {
			s := data.daftar_suasana[i]
			sNo = fmt.Sprintf("%d", i+1)
			tglS = fmt.Sprintf("%02d/%02d/%02d", s.day, s.mon, s.year%100) 
			
			deskripsi = s.deskripsi_perasaan
			skor = fmt.Sprintf("%d", s.skor_emosi)
		}

		// DATA TUGAS (KANAN) ---
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

		// 4. Cetak satu baris untuk kedua tabel (kiri dan kanan)
		fmt.Printf("║ %-3s╎ %-8s╎ %-54s╎ %-16s║%s║ %-3s╎ %-8s╎ %-31s╎ %-18s╎ %-10s╎ %-7s║\n",
			sNo, tglS, deskripsi, skor, "   ", tNo, tglT, namaT, durasi, prioritas, status)
	}

}

func inputDataTugas(data *Pengguna) {
	/* I.S. terdefinisi array T yang kosong atau terisi data minimal 1
	F.S. array data terisi data baru ke n*/
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
	F.S. data suasana untuk masing-masing pengguna terisi sebanyak k dan disimpan di dalam array data T_suasana pengguna */
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
//
	data.T_suasana = k
}

// Input kata dengan spasi dan diakhiri dengan titik (.)
func isiDeskripsi(P *Pengguna, kataPart *arrTemp2, i int) {
	var j int = 0

	fmt.Scan(&kataPart[j])
	for !strings.Contains(kataPart[j], ".") {
		j++
		fmt.Scan(&kataPart[j])
	}
	panjangKata := j + 1
	P.daftar_suasana[i].deskripsi_perasaan = gabungDeskripsi(kataPart, panjangKata)
}

func gabungDeskripsi(kataPart *arrTemp2, panjang int) string {
	var j int
	var hasil string
	
	for j = 0; j < panjang; j++ {
		if j > 0 {
			hasil += " "
		}
		hasil += kataPart[j]
	}
	
	return hasil
}

// fungsi delete//
func deleteData(data *Pengguna)bool{
	/*I.S. terdefinisi array data > 0
	F.S. data pengguna terdelete sesuai dengan input dari pengguna	
	*/
	var jenis string
	var Nomer int
	var status bool
	fmt.Print("Nama tabel dari data yang akan diubah (Tugas/Suasana): ")
	fmt.Scan(&jenis)
	fmt.Print("Data yang akan dihapus nomor:")
	fmt.Scan(&Nomer)
	if jenis == "Tugas"{
		for i:= Nomer; i < data.T_suasana; i++{
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
		fmt.Scanf("%d/%d/%d", &data.daftar_tugas[Nomer].day, &data.daftar_tugas[Nomer].mon, &data.daftar_tugas[Nomer].year)
		fmt.Scan(&data.daftar_tugas[Nomer-1].durasi_pengerjaan)
		fmt.Scan(&data.daftar_tugas[Nomer-1].Prioritas)
		fmt.Scan(&data.daftar_tugas[Nomer-1].selesai)
		fmt.Println("Note!")
	}else if Jenis == "Suasana"{
		fmt.Println("Masukkan skor emosi, tanggal, dan deskripsi perasaan secara berurutan! (❁´◡❁)`")
		fmt.Println("-Contoh format: 50 13/06/2027 aku mau makan.-")
		fmt.Scan(&data.daftar_suasana[Nomer-1].skor_emosi)
		fmt.Scanf("%d/%d/%d", &data.daftar_suasana[Nomer-1].day, &data.daftar_suasana[Nomer].mon, &data.daftar_suasana[Nomer].year)
		isiDeskripsi(data, kataPart, Nomer-1)
		fmt.Println("Note!")
	}

}
// fungsi presentase tugas
func presentaseTugasHarian(data *Pengguna, hari int, bulan int, tahun int)float64{
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

func loadData(data *Pengguna) {
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
	//LOG file aja
func openLog(){
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

