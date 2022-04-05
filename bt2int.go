package main

import "fmt"

type tongluong interface { //khai bao 1 interface
	luong() int
}

type chinhthuc struct { //khai bao 1 du lieu thuoc kieu cau truc
	bh         int
	luongthang int
	thuong     int
}

type hopdong struct { //khai bao cau truc
	luongthang int
	bh         int
}

func (c chinhthuc) luong() int {
	return c.luongthang + c.thuong
}

func (h hopdong) luong() int {
	return h.luongthang
}

func total(t []tongluong) {
	tongchiphi := 0
	for _, v := range t {
		tongchiphi = tongchiphi + v.luong()
	}
	fmt.Printf("Tong chi phi cho luong nv: $%d", tongchiphi)
}
func main() {
	emp1 := chinhthuc{1, 1000, 100}
	emp2 := chinhthuc{2, 1500, 50}
	emp3 := hopdong{3, 200}
	employee := []tongluong{emp1, emp2, emp3}
	total(employee)
}
