package main

import "fmt"

type Komentar struct {
	ID       int
	Teks     string
	Sentimen string
	Panjang  int
}

type DataKomentar struct {
	List   [100]Komentar
	Jumlah int
	NextID int
}

var kataPositif = [4]string{"bagus", "baik", "hebat", "luar biasa"}
var kataNegatif = [4]string{"buruk", "jelek", "parah", "menyedihkan"}
var kataNetral = [4]string{"normal", "standar", "baik-baik saja", "biasa"}

func HitungPanjang(teks string) int {
	p := 0
	for i := 0; i < len(teks); i++ {
		p++
	}
	return p
}

func Mengandung(teks string, kata string) bool {
	n := len(teks)
	m := len(kata)
	for i := 0; i <= n-m; i++ {
		sama := true
		for j := 0; j < m; j++ {
			if teks[i+j] != kata[j] {
				sama = false
				break
			}
		}
		if sama {
			return true
		}
	}
	return false
}

func AnalisisSentimen(teks string) string {
	skor := 0
	for i := 0; i < 4; i++ {
		if Mengandung(teks, kataPositif[i]) {
			skor++
		}
	}
	for i := 0; i < 4; i++ {
		if Mengandung(teks, kataNegatif[i]) {
			skor--
		}
	}
	for i := 0; i < 4; i++ {
		if Mengandung(teks, kataNetral[i]) {
		}
	}

	if skor > 0 {
		return "positif"
	}
	return "negatif"
}

func TambahKomentar(data *DataKomentar) {
	var teks string
	
	if data.Jumlah >= 100 {
		fmt.Println("Data penuh.")
		return
	}
	fmt.Print("Masukkan komentar: ")
	fmt.Scanln(&teks)
	s := AnalisisSentimen(teks)
	p := HitungPanjang(teks)

	data.List[data.Jumlah] = Komentar{data.NextID, teks, s, p}
	data.Jumlah++
	data.NextID++
}

func TampilkanSemua(data *DataKomentar) {
	for i := 0; i < data.Jumlah; i++ {
		k := data.List[i]
		fmt.Printf("ID: %d | %s | %s | %d karakter\n", k.ID, k.Sentimen, k.Teks, k.Panjang)
	}
}

func EditKomentar(data *DataKomentar) {
	var id int
	fmt.Print("Masukkan ID yang ingin diubah: ")
	fmt.Scanln(&id)
	for i := 0; i < data.Jumlah; i++ {
		if data.List[i].ID == id {
			var teks string
			fmt.Print("Masukkan komentar baru: ")
			fmt.Scanln(&teks)
			s := AnalisisSentimen(teks)
			p := HitungPanjang(teks)
			data.List[i].Teks = teks
			data.List[i].Sentimen = s
			data.List[i].Panjang = p
			fmt.Println("Komentar diubah.")
			return
		}
	}
	fmt.Println("ID tidak ditemukan.")
}

func HapusKomentar(data *DataKomentar) {
	var id int
	fmt.Print("Masukkan ID yang ingin dihapus: ")
	fmt.Scanln(&id)
	for i := 0; i < data.Jumlah; i++ {
		if data.List[i].ID == id {
			for j := i; j < data.Jumlah-1; j++ {
				data.List[j] = data.List[j+1]
			}
			data.Jumlah--
			fmt.Println("Komentar dihapus.")
			return
		}
	}
	fmt.Println("ID tidak ditemukan.")
}

func CariKomentarSequential(data *DataKomentar) {
	var kata string
	fmt.Print("Masukkan kata kunci: ")
	fmt.Scanln(&kata)
	found := false
	for i := 0; i < data.Jumlah; i++ {
		if Mengandung(data.List[i].Teks, kata) {
			k := data.List[i]
			fmt.Printf("ID: %d | %s | %s\n", k.ID, k.Sentimen, k.Teks)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ditemukan.")
	}
}

func CariKomentarBinary(data *DataKomentar) {
	var kata string
	fmt.Print("Masukkan kata kunci: ")
	fmt.Scanln(&kata)

	low := 0
	high := data.Jumlah - 1
	for low <= high {
		mid := (low + high) / 2
		teks := data.List[mid].Teks
		if teks == kata {
			k := data.List[mid]
			fmt.Printf("ID: %d | %s | %s\n", k.ID, k.Sentimen, k.Teks)
			return
		} else if teks < kata {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Tidak ditemukan.")
}

func UrutkanSelection(data *DataKomentar) {
	var pilihan int
	fmt.Println("Urut berdasarkan:")
	fmt.Println("1. Panjang")
	fmt.Println("2. Sentimen (positif ke negatif)")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&pilihan)

	for i := 0; i < data.Jumlah-1; i++ {
		idx := i
		for j := i + 1; j < data.Jumlah; j++ {
			if pilihan == 1 {
				if data.List[j].Panjang < data.List[idx].Panjang {
					idx = j
				}
			} else if pilihan == 2 {
				if data.List[j].Sentimen < data.List[idx].Sentimen {
					idx = j
				}
			}
		}
		tmp := data.List[i]
		data.List[i] = data.List[idx]
		data.List[idx] = tmp
	}
	fmt.Println("Data diurutkan.")
}

func UrutkanInsertion(data *DataKomentar) {
	var pilihan int
	fmt.Println("Urut berdasarkan:")
	fmt.Println("1. Panjang")
	fmt.Println("2. Sentimen (positif ke negatif)")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&pilihan)

	for i := 1; i < data.Jumlah; i++ {
		tmp := data.List[i]
		j := i - 1
		for j >= 0 && ((pilihan == 1 && data.List[j].Panjang > tmp.Panjang) ||
			(pilihan == 2 && data.List[j].Sentimen > tmp.Sentimen)) {
			data.List[j+1] = data.List[j]
			j--
		}
		data.List[j+1] = tmp
	}
	fmt.Println("Data diurutkan.")
}

func Statistik(data *DataKomentar) {
	var pos, neg int
	for i := 0; i < data.Jumlah; i++ {
		if data.List[i].Sentimen == "positif" {
			pos++
		} else if data.List[i].Sentimen == "negatif" {
			neg++
		}
	}
	fmt.Println("Statistik:")
	fmt.Println("Positif:", pos)
	fmt.Println("Negatif:", neg)
}

func Menu() {
	var komentar DataKomentar
	komentar.NextID = 1
	var pilihan int

	for {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1. Tambah Komentar")
		fmt.Println("2. Lihat Semua Komentar")
		fmt.Println("3. Ubah Komentar")
		fmt.Println("4. Hapus Komentar")
		fmt.Println("5. Cari Komentar (Sequential)")
		fmt.Println("6. Cari Komentar (Binary)")
		fmt.Println("7. Urutkan (Selection)")
		fmt.Println("8. Urutkan (Insertion)")
		fmt.Println("9. Statistik Sentimen")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scanln(&pilihan)

		if pilihan == 0 {
			fmt.Println("Keluar dari program.")
			break
		} else if pilihan == 1 {
			fmt.Println("Silakan masukkan komentar baru.")
			TambahKomentar(&komentar)
		} else if pilihan == 2 {
			fmt.Println("Menampilkan semua komentar.")
			TampilkanSemua(&komentar)
		} else if pilihan == 3 {
			fmt.Println("Silakan masukkan ID komentar yang ingin diedit.")
			EditKomentar(&komentar)
		} else if pilihan == 4 {
			fmt.Println("Silakan masukkan ID komentar yang ingin dihapus.")
			HapusKomentar(&komentar)
		} else if pilihan == 5 {
			fmt.Println("Silakan masukkan teks komentar yang ingin dicari (Sequential Search).")
			CariKomentarSequential(&komentar)
		} else if pilihan == 6 {
			fmt.Println("Silakan masukkan teks komentar yang ingin dicari (Binary Search).")
			CariKomentarBinary(&komentar)
		} else if pilihan == 7 {
			fmt.Println("Silakan pilih kriteria pengurutan (1 = Panjang, 2 = Sentimen) - Selection Sort.")
			UrutkanSelection(&komentar)
		} else if pilihan == 8 {
			fmt.Println("Silakan pilih kriteria pengurutan (1 = Panjang, 2 = Sentimen) - Insertion Sort.")
			UrutkanInsertion(&komentar)
		} else if pilihan == 9 {
			fmt.Println("Menampilkan statistik sentimen komentar.")
			Statistik(&komentar)
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func main() {
	Menu()
}
