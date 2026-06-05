package main 

import "fmt"

func main(){
	var oHome, oEdit, oInput, oSearch, oSort, oKey, oDate, oSel, oIns, oTren int
	var d, m, y int
	var run, edit, input, search, sSort, iSort bool
	var keyword, date, sort, tren bool
	var sortBy int
	var tgl string
	var cari string
	var data Pengguna
	var Temp arrTemp
	var Temp2 arrTemp2

	loadData(&data)
	run = true
	for run {
		//fungsi print tabel
		cls()
		showHeader()
		showTables(data)
		fmt.Println("Hi-ya, user!")
		fmt.Println("Millow's here! ✧⁠◝⁠(⁠⁰⁠▿⁠⁰⁠)⁠◜⁠✧")
		fmt.Println("What are we doing today? ♪⁠ヽ⁠(⁠･⁠ˇ⁠∀⁠ˇ⁠･⁠ゞ⁠)")
		fmt.Println("⪩ (0)Show conversation")
		fmt.Println("⪩ (1)Edit data")
		fmt.Println("⪩ (2)Search data")
		fmt.Println("⪩ (3)Sort task data")
		fmt.Println("⪩ (4)Trend statistics")
		fmt.Println("⪩ (5)Exit")
		fmt.Println()
		fmt.Print("⟢ ")
		fmt.Scan(&oHome)
		switch oHome {
		case 0: //log percakapan
			cls()
			openLog()
			fmt.Print("\n\n(5)Close log")
			fmt.Println()
			fmt.Print("⟢ ")
			fmt.Scan(&oHome)
		case 1: //edit data
			edit = true
			for edit {
				cls()
				showTables(data)
				fmt.Println( )
				fmt.Println("⪩ (1)Input data")
				fmt.Println("⪩ (2)Replace data")
				fmt.Println("⪩ (3)Delete data")
				fmt.Println("⪨ (4)Back ←⁠(⁠*⁠꒪⁠ヮ⁠꒪⁠*⁠)")
				fmt.Println()
				fmt.Print("⟢ ")
				fmt.Scan(&oEdit)
				switch oEdit {
				case 1: //input data
					input = true
					for input {
						cls()
						showTables(data)
						fmt.Println("⪩ (1)Mood")
						fmt.Println("⪩ (2)Task")
						fmt.Println("⪨ (3)Back ←⁠(⁠*⁠꒪⁠ヮ⁠꒪⁠*⁠)")
						fmt.Println()
						fmt.Print("⟢ ")
						fmt.Scan(&oInput)
						switch oInput {
						case 1: //input mood
							cls()
							showTables(data)
							fmt.Println("Millow: Alrighty! what do you feel today? ⋋⁠✿⁠ ⁠⁰⁠ ⁠o⁠ ⁠⁰⁠ ⁠✿⁠⋌")
							fmt.Println("--------------------------------------------------------")
							fmt.Println()
							fmt.Println("Masukkan skor emosi, tanggal, dan deskripsi perasaan secara berurutan! (❁´◡❁)`")
							fmt.Println("--------------------------------------------------------------------------------")
							fmt.Println("Note!")
							fmt.Println("Masukkan skor emosi dengan skala 1-100")
							fmt.Println("Contoh format tanggal: 13/06/2027")
							fmt.Println("Hanya gunakan titik (.) untuk mengakhiri input deskripsi perasaan")
							fmt.Println("Akhiri input dengan '-1'")
							fmt.Println("Format keseluruhan: 50 13/06/2027 Hari ini aku mau makan nasi padang. -1")
							fmt.Println()
							inputDataSuasana(&data, &Temp2)
						case 2: //input task
							cls()
							showTables(data)
							fmt.Println("Millow: Alrighty! what task do we have today? ⋋⁠✿⁠ ⁠⁰⁠ ⁠o⁠ ⁠⁰⁠ ⁠✿⁠⋌")
							fmt.Println("-----------------------------------------------------------")
							fmt.Println()
							fmt.Println("Masukkan data nama tugas, tanggal(), durasi, skala prioritas, status selesai (1 = selesai, 0 = belum) tugas secara berurutan")
							fmt.Println("--------------------------------------------------------------------------------------------------------------------------")
							fmt.Println("Note!")
							fmt.Println("Input nama tugas menggunakan underscore (_) sebagai pengganti spasi")
							fmt.Println("Contoh format tanggal: 13/06/2027")
							fmt.Println("Masukkan durasi pengerjaan dalam menit")
							fmt.Println("Skala prioritas berupa nomor")
							fmt.Println("Status selesai berupa angka 1(selesai) atau 0(belum)")
							fmt.Println("Akhiri input dengan 'END'")
							fmt.Println("Format keseluruhan: Tugas_Alpro_Sorting 13/06/2027 30 2 1 END")
							fmt.Println()
							inputDataTugas(&data)
						case 3: // back
							input = false
						}
					}
				case 2: //ini replace
					cls()
					showTables(data)
					fmt.Println("Millow: Hmm...What do ya want to change, besto friendo? (⁠｡⁠•̀⁠ᴗ⁠-⁠)⁠✧")
					UbahData(&data, &Temp2)
				case 3://ini delete
					cls()
					showTables(data)
					fmt.Println("Millow: Bye-bye memoriee~ (⁠╥⁠﹏⁠╥⁠)")
					deleteData(&data)
				case 4://back
					edit = false
				}
			}
		case 2: // ini search
			search = true
			for search {
				cls()
				showTables(data)
				fmt.Println("⪩ (1)Search by keyword")
				fmt.Println("⪩ (2)Search by date")
				fmt.Println("⪨ (3)Back ←⁠(⁠*⁠꒪⁠ヮ⁠꒪⁠*⁠)")
				fmt.Println()
				fmt.Print("⟢ ")
				fmt.Scan(&oSearch)
				switch oSearch {
				case 1: //by keyword
					keyword = true 
					cls()
					showTables(data)
					for keyword {
						fmt.Println("⪩ (1)Mood")
						fmt.Println("⪩ (2)Task")
						fmt.Println("⪨ (3)Back ←⁠(⁠*⁠꒪⁠ヮ⁠꒪⁠*⁠)")
						fmt.Println()
						fmt.Print("⟢ ")
						fmt.Scan(&oKey)
						switch oKey {
						case 1://mood
							cls()
							showTables(data)
							fmt.Println("Millow: Detective Millow Holmes, ready for duty! (⌐■_■)ノ🔎")
							fmt.Println("Masukkan 1 kata untuk mencari data suasana hati")
							fmt.Print("⟢ ")
							fmt.Scan(&cari)
							sequentialMood(data, cari, data.T_suasana)
						case 2://task
							cls()
							showTables(data)
							fmt.Println("Millow: Detective Millow Holmes, ready for duty! (⌐■_■)ノ🔎")
							fmt.Println("Masukkan nama tugas secara lengkap untuk mencari data")
							fmt.Print("⟢ ")
							fmt.Scan(&cari)
							sequentialTugas(data, cari, data.T_tugas)
						case 3: //back
							keyword = false
						}
					}
				case 2: //date
					date = true 
					cls()
					showTables(data)
					for date {
						fmt.Println("⪩ (1)Mood")
						fmt.Println("⪩ (2)Task")
						fmt.Println("⪨ (3)Back ←⁠(⁠*⁠꒪⁠ヮ⁠꒪⁠*⁠)")
						fmt.Println()
						fmt.Print("⟢ ")
						fmt.Scan(&oDate)
						switch oDate {
						case 1: //mood
							cls()
							showTables(data)
							fmt.Println("Millow: Detective Millow Holmes, ready for duty! (⌐■_■)ノ🔎")
							fmt.Println("Masukkan tanggal (01 02 2019)")
							fmt.Println()
							fmt.Print("⟢ ")
							fmt.Scan(&d, &m, &y)
							binaryMood(data, data.T_suasana, d, m, y)
						case 2: //task
							cls()
							showTables(data)
							fmt.Println("Millow: Detective Millow Holmes, ready for duty! (⌐■_■)ノ🔎")
							fmt.Println("Masukkan tanggal (01 02 2019)")
							fmt.Println()
							fmt.Print("⟢ ")
							fmt.Scan(&d, &m, &y)
							binaryTugas(data, data.T_tugas, d, m, y)
						case 3: //back
							date = false
						}
					}
				case 3: //back
					search = false
				}
			}
		case 3://ini sorting
		sort = true
			for sort {
				cls()
				showTables(data)
				fmt.Println("⪩ (1)Selection Sort")
				fmt.Println("⪩ (2)Insertion Sort")
				fmt.Println("⪨ (3)Back ←⁠(⁠*⁠꒪⁠ヮ⁠꒪⁠*⁠)")
				fmt.Println()
				fmt.Print("⟢ ")
				fmt.Scan(&oSort)
				switch oSort {
				case 1://selection sort
					sSort = true
					for sSort {
						cls()
						showTables(data)
						fmt.Println("⪩ (1)Ascending")
						fmt.Println("⪩ (2)Descending")
						fmt.Println("⪨ (3)Back ←⁠(⁠*⁠꒪⁠ヮ⁠꒪⁠*⁠)")
						fmt.Println()
						fmt.Print("⟢ ")
						fmt.Scan(&oSel)
						switch oSel {
						case 1://accending
							cls()
							showTables(data)
							fmt.Println("Millow: (⊃＞ ⌂ ＜)⊃━✿✿✿ Magiiiiiic Sort Maniaaaaa!")
							fmt.Println("⪩ (1)By priority")
							fmt.Println("⪩ (2)By duration")
							fmt.Scan(&sortBy)
							selectionAsc(&data, data.T_tugas, sortBy)
							//show data
						case 2://decending
							cls()
							showTables(data)
							fmt.Println("Millow: (⊃＞ ⌂ ＜)⊃━✿✿✿ Magiiiiiic Sort Maniaaaaa!")
							fmt.Println("⪩ (1)By priority")
							fmt.Println("⪩ (2)By duration")
							fmt.Scan(&sortBy)
							selectionDes(&data, data.T_tugas, sortBy)
							//show data
						case 3://back
							sSort = false
						} 
					}
				case 2: //insertion sort
					iSort = true
					for iSort {
						cls()
						showTables(data)
						fmt.Println("⪩ (1)Ascending")
						fmt.Println("⪩ (2)Descending")
						fmt.Println("⪨ (3)Back ←⁠(⁠*⁠꒪⁠ヮ⁠꒪⁠*⁠)")
						fmt.Println()
						fmt.Print("⟢ ")
						fmt.Scan(&oIns)
						switch oIns {
						case 1: //ascending
							cls()
							showTables(data)
							fmt.Println("Millow: (⊃＞ ⌂ ＜)⊃━✿✿✿ Magiiiiiic Sort Maniaaaaa!")
							fmt.Println("⪩ (1)By priority")
							fmt.Println("⪩ (2)By duration")
							fmt.Scan(&sortBy)
							insertionAsc(&data, data.T_tugas, sortBy)
						case 2: //descending
							cls()
							showTables(data)
							fmt.Println("Millow: (⊃＞ ⌂ ＜)⊃━✿✿✿ Magiiiiiic Sort Maniaaaaa!")
							fmt.Println("⪩ (1)By priority")
							fmt.Println("⪩ (2)By duration")
							fmt.Scan(&sortBy)
							insertionDes(&data, data.T_tugas, sortBy)
							//show data
						case 3: //back
							iSort = false
						}
					}
				case 3: //back
					sort = false
				}
			}
		case 4://statistik tren
			tren = true
			cls()
			showTables(data)
			for tren {
				fmt.Println("⪩ (1)Weekly Mood")
				fmt.Println("⪩ (2)Daily Tasks")
				fmt.Println("⪨ (3)Back ←⁠(⁠*⁠꒪⁠ヮ⁠꒪⁠*⁠)")
				fmt.Println()
				fmt.Print("⟢ ")
				fmt.Scan(&oTren)
				switch oTren {
				case 1://mood
					cls()
					showTables(data)
					fmt.Println("Millow: Lets see hows your mood this week~ ( ✌︎'ω')✌︎")
					presentaseSuasanaMingguan(&data, &Temp)
				case 2://task
					cls()
					showTables(data)
					var hari, bulan, tahun int
					fmt.Println("Millow: Ooh, ooh! Time to reveal your daily stats! ＼( ° ∇ ° )／")
					fmt.Println("Masukkan tanggal (01/02/2019)")
					fmt.Println()
					fmt.Scan(&tgl)
					fmt.Sscanf(tgl, "%d/%d/%d\n", &hari, &bulan, &tahun)
					fmt.Printf("Presentase tugas selesai: %.2f%%\n\n\n", presentaseTugasHarian(&data, hari, bulan, tahun))
				case 3://back
					tren = false
				}
			}
		case 5://out atau close
			fmt.Println("Millow: Buh-bye! See ya later! (⁠~⁠‾⁠▿⁠‾⁠)⁠~")
			run = false
		}
	}
}

