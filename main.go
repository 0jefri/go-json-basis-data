package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Order struct {
	ID        int       `json:"id"`
	Items     []string  `json:"items"`
	Status    string    `json:"status"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
}

var orders []Order
var idCounter int

func addOrder() {
	idCounter++
	var order Order
	order.ID = idCounter
	order.CreatedAt = time.Now()

	fmt.Println("Silahkan masukkan pesanan anda:")
	var itemsInput string
	fmt.Scanln(&itemsInput)

	order.Items = append(order.Items, itemsInput)
	order.Status = "Diproses"

	fmt.Println("Masukkan total pembayaran:")
	fmt.Scanln(&order.Total)

	orders = append(orders, order)

	orderJSON, _ := json.MarshalIndent(order, "", "  ")
	fmt.Println("Pesanan berhasil dibuat:")
	fmt.Println(string(orderJSON))
}

func editOrder() {
	// fmt.Println("Masukkan ID pesanan yang ingin diedit:")
	// var id int
	// fmt.Scanln(&id)

	// for i, order := range orders {
	// 	if order.ID == id {
	// 		fmt.Println("Pesanan ditemukan, masukkan item baru:")
	// 		var itemsInput string
	// 		fmt.Scanln(&itemsInput)
	// 		orders[i].Items = []string{itemsInput}

	// 		fmt.Println("Pesanan berhasil diupdate!")
	// 		return
	// 	}
	// }
	// fmt.Println("Pesanan tidak ditemukan!")
	fmt.Println("Masukkan ID pesanan yang ingin diedit:")
	var id int
	fmt.Scanln(&id)

	for i, order := range orders {
		if order.ID == id {
			fmt.Println("Pesanan ditemukan, item saat ini:", orders[i].Items)
			fmt.Println("Masukkan item baru untuk ditambahkan:")
			var itemsInput string
			fmt.Scanln(&itemsInput)

			newItems := []string{itemsInput}
			orders[i].Items = append(orders[i].Items, newItems...)

			fmt.Println("Item berhasil ditambahkan ke pesanan!")
			return
		}
	}
	fmt.Println("Pesanan tidak ditemukan!")
}

func viewOrders() {
	if len(orders) == 0 {
		fmt.Println("Tidak ada pesanan yang tersedia.")
		return
	}

	fmt.Println("Daftar Pesanan:")
	for _, order := range orders {
		orderJSON, _ := json.MarshalIndent(order, "", "  ")
		fmt.Println(string(orderJSON))
	}
}

func updateStatus() {
	fmt.Println("Masukkan ID pesanan yang ingin diubah statusnya:")
	var id int
	fmt.Scanln(&id)

	for i, order := range orders {
		if order.ID == id {
			fmt.Println("Status saat ini:", order.Status)
			fmt.Println("Masukkan status baru (Diproses/Diantar/Selesai):")
			fmt.Scanln(&orders[i].Status)

			fmt.Println("Status pesanan berhasil diperbarui!")
			return
		}
	}
	fmt.Println("Pesanan tidak ditemukan!")
}

func main() {
	for {
		fmt.Println("\n=== Aplikasi Reservasi Restoran ===")
		fmt.Println("1. Tambah Pesanan")
		fmt.Println("2. Edit Pesanan")
		fmt.Println("3. Lihat Pesanan")
		fmt.Println("4. Ubah Status Pesanan")
		fmt.Println("5. Keluar")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addOrder()
		case 2:
			editOrder()
		case 3:
			viewOrders()
		case 4:
			updateStatus()
		case 5:
			fmt.Println("Keluar dari aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}
