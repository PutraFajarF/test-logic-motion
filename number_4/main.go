package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan jumlah cuti bersama: ")
	totalLeaveStr, _ := reader.ReadString('\n')
	totalLeaveStr = strings.TrimSpace(totalLeaveStr)
	var totalLeave int
	fmt.Sscanf(totalLeaveStr, "%d", &totalLeave)

	fmt.Print("Masukkan tanggal join karyawan (yyyy-mm-dd): ")
	joinDate, _ := reader.ReadString('\n')
	joinDate = strings.TrimSpace(joinDate)

	fmt.Print("Masukkan tanggal rencana cuti (yyyy-mm-dd): ")
	planDateLeave, _ := reader.ReadString('\n')
	planDateLeave = strings.TrimSpace(planDateLeave)

	fmt.Print("Masukkan durasi cuti (hari): ")
	leaveDurationStr, _ := reader.ReadString('\n')
	leaveDurationStr = strings.TrimSpace(leaveDurationStr)
	var leaveDuration int
	fmt.Sscanf(leaveDurationStr, "%d", &leaveDuration)

	hasil, alasan := IsAvailableLeave(totalLeave, joinDate, planDateLeave, leaveDuration)

	if hasil {
		fmt.Printf("Output: %t\n", hasil)
	} else {
		fmt.Printf("Output: %t\nAlasan: %s\n", hasil, alasan)
	}
}

func IsAvailableLeave(totalLeave int, joinDate, planDateLeave string, leaveDuration int) (bool, string) {
	layout := "2006-01-02"
	joinDateTime, _ := time.Parse(layout, joinDate)
	planDateLeaveTime, _ := time.Parse(layout, planDateLeave)

	threshold := joinDateTime.AddDate(0, 0, 180)

	if planDateLeaveTime.Before(threshold) {
		return false, "Karena belum 180 hari sejak tanggal join karyawan"
	}

	endOfYear := time.Date(joinDateTime.Year(), 12, 31, 0, 0, 0, 0, time.UTC)
	remainDay := endOfYear.Sub(threshold).Hours() / 24

	leaveQuota := int(math.Floor(remainDay / 365 * float64(14-totalLeave)))

	if leaveDuration > leaveQuota {
		return false, fmt.Sprintf("Karena hanya boleh mengambil %d hari cuti", leaveQuota)
	}

	return true, ""
}
