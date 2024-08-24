package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan total belanja seorang customer: Rp ")
	totalBelanjaStr, _ := reader.ReadString('\n')
	totalBelanjaStr = strings.TrimSpace(totalBelanjaStr)
	totalBelanja, err := strconv.ParseFloat(strings.Replace(totalBelanjaStr, ".", "", -1), 64)
	if err != nil {
		fmt.Println("Input total belanja tidak valid")
		return
	}

	fmt.Print("Pembeli membayar: Rp ")
	pembayaranStr, _ := reader.ReadString('\n')
	pembayaranStr = strings.TrimSpace(pembayaranStr)
	pembayaran, err := strconv.ParseFloat(strings.Replace(pembayaranStr, ".", "", -1), 64)
	if err != nil {
		fmt.Println("Input pembayaran tidak valid")
		return
	}

	kembalian, pecahanUang := HitungKembalian(totalBelanja, pembayaran)

	if kembalian == -1 {
		fmt.Println("Output: False, kurang bayar")
	} else {
		fmt.Printf("Kembalian yang harus diberikan kasir: %.3f, dibulatkan menjadi %.3f\n", kembalian, math.Floor(kembalian/100)*100)
		fmt.Println("Pecahan uang:")
		for _, pecahan := range pecahanUang {
			fmt.Println(pecahan)
		}
	}
}

func HitungKembalian(totalBelanja, pembayaran float64) (float64, []string) {
	if pembayaran < totalBelanja {
		return -1, nil
	}

	pecahan := []float64{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	kembalian := pembayaran - totalBelanja
	kembalian = math.Floor(kembalian/100) * 100
	pecahanUang := []string{}

	for _, p := range pecahan {
		if kembalian >= p {
			jumlah := int(kembalian / p)
			kembalian -= p * float64(jumlah)
			pecahanUang = append(pecahanUang, fmt.Sprintf("%d lembar %.f", jumlah, p))
		}
	}

	return pembayaran - totalBelanja, pecahanUang
}
