package main

const arrmax = 999

type Suasana struct{
	/*struktur yang berisikan catatan 
	suasan masing masing pasien*/
	skor_emosi int
	deskripsi_perasaan string
	day, mon, year, days int
}

type tugas struct{
	/*struktur yang berisikan catatan 
	daftar tugas masing masing pasien*/
	nama_tugas string //kata kunci
	durasi_pengerjaan, Prioritas int
	selesai int // 1 untuk selesai, 0 untuk belum
	day, mon, year, days int
}

type Pengguna struct{
	/*struktur yang berisikan data pengguna*/
	Nama string 
	Status_Penguna string 
	T_suasana, T_tugas int
	daftar_suasana [arrmax]Suasana
	daftar_tugas [arrmax]tugas
}

type arrTemp [arrmax]float64

type arrTemp2 [arrmax]string

type Tanggal[arrmax] int
/*riwayat percakapan not found*/
