package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	text := `Rau diếp cá VietGAP 100g Trứng gà Tafa size 60g hộp 6 quả Mầm hướng dương hộp 200g Xà Lách Sồi Xanh 250G Đậu bắp 300g Tía Tô Lotte 100G 
	 Khoai tây vàng Đà Lạt Choice L 400g Hỗn hợp nấm tươi Trâm Anh khay 300g Thanh long loại 1 1kg-1.1kg (trái trên 300G) 
	Ổi ruột đỏ 1kg - 1.1kg (trái từ 200g trở lên) Chuối cau 800g - 1kg Hành tây Đà Lạt (6-9 củ/kg) 900g-1kg Mầm củ cải trắng hộp 200g 
	Ba chỉ bò Mỹ hotpot Thảo Tiến 300g Trái Ôliu Hỗn Hợp (Tím/Đen/Xanh) Nguyên Trái Epoch 370G Lốc 2 gói nấm kim châm Hàn Quốc gói 150g 
	Đậu Nành Nhật Bản Đông Lạnh DaLat Agri Food Gói 400G Bánh Tiểu Long Nhân Súp Kiểu Thượng Hải CJ Cầu Tre Nhân Súp Tôm Và Thịt 300G 
	Bánh Bao Xíu Mại Trứng Muối Choice L Gói 300G`
	sum := 0
	// gr
	pat := regexp.MustCompile(`[0-9]+[g|G] `)
	matches := pat.FindAllStringSubmatch(text, -1) // matches is [][]string

	for _, match := range matches {
		fmt.Printf(" value=<%s>\n", match[0])
		i, err := strconv.Atoi(strings.TrimSuffix(match[0], "g "))
		if err == nil {
			fmt.Printf("i=%d\n", i)
		} else {
			i, err = strconv.Atoi(strings.TrimSuffix(match[0], "G "))
			fmt.Printf("i=%d\n", i)
		}
		
		sum += i
	}
	fmt.Printf("==========TOTAL = %d (gr) = %d (kg)\n", sum, sum/1000)
	
	// kg
	pat = regexp.MustCompile(`[0-9]+kg\s+[^-]+`)
	matches = pat.FindAllStringSubmatch(text, -1) // matches is [][]string

	for _, match := range matches {
		fmt.Printf(" value=<%s> ===== <%s>\n", match[0], match[0][:1])
		
		i, _ := strconv.Atoi(match[0][:1])
		
		fmt.Printf("i=%d\n", i)
		sum += i*1000
	}
	fmt.Printf("TOTAL = %d (gr) = %d (kg)\n", sum, sum/1000)
}
