package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	//input------------------------
    // pl("name : ")
	// reader := bufio.NewReader(os.Stdin)
	// name, err := reader.ReadString('\n')
	// if err == nil {
	// 	pl("hai bang ",name ,"tes")
	// } else {
	// 	log.Fatal(err)
	// }

	//variabel------------------------
	// var name string = "udin"
	// pl(name)
	// var v1, v2 = 2.1, 5.6
	// pl(v1+v2)
	// v3 := 5.2
	// pl(v3)
	// pl(reflect.TypeOf(6.7))

	// //casting string to ascii------------------------
	// v4 := int(5)
	// pl(v4)
	// v5 := "50000"
	// v6, err := strconv.Atoi(v5)
	// pl(v6, err, reflect.TypeOf(v6))

	// //casting int to ascii------------------------
	// v7 := 7000
	// v8 := strconv.Itoa(v7)
	// pl(v8, reflect.TypeOf(v8))

	// //casting string to float------------------------
	// v9 := "8.2"
	// if v10, err := strconv.ParseFloat(v9, 64); err == nil {
	// 	pl(v10)
	// }

	// //casting float to string------------------------
	// v12 := 8.2
	// v11 := fmt.Sprintf("%f", v12)
	// io.WriteString(os.Stdout,v11)
	
	// //if else------------------------
	// iAge := 55
	// if (iAge >=1) && (iAge <= 21) {
	// 	pl("Penting")
	// } else if (iAge == 25) || (iAge == 50) || (iAge == 65) {
	// 	pl("Penting")
	// } else {
	// 	pl("Tak Penting")
	// }

	// //String------------------------
	// sV1 := "A Word"
	// replacer := strings.NewReplacer("A", "Another")
	// sV2 := replacer.Replace(sV1)
	// pl(sV2)
	// pl("Panjang : ", len(sV2))
	// pl("Mengandung Another : ", strings.Contains(sV2, "Another"))
	// pl("Index o ada di : ", strings.Index(sV2,"o"))
	// pl("Hasil replace : ", strings.Replace(sV2,"o","0", -1))
	// sV3 := "\n        Amazing Yeah\n"
	// sV3 = strings.TrimSpace(sV3)
	// pl(sV3)
	// pl("Hasil Split : ", strings.Split("a-b-c-d", "-"))
	// pl("Hasil Lower : ", strings.ToLower("Hai"))
	// pl("Hasil Uppercase : ", strings.ToUpper("Hai"))
	// pl("Apakah punya prefix : ", strings.HasPrefix("School", "Sc"))
	// pl("Apakah punya prefix : ", strings.HasSuffix("School", "ol"))

	// //character = rune------------------------
	// rStr := "abcdefg"
	// pl("Jumlah Rune : ", utf8.RuneCountInString(rStr))
	// for i, runeVal := range rStr {
	// 	fmt.Printf("%d : %#U : %c\n",i,runeVal,runeVal)
	// }

	// //time------------------------
	// now := time.Now()
	// pl(now.Year(), now.Month(), now.Day())
	// pl(now.Hour(), now.Minute(), now.Second())

	// //math------------------------
	// pl("5 + 4 = ", 5+4)
	// pl("5 % 2 = ", 5%2)
	// x := 2
	// x++
	// pl(x)

	// // random number
	// seedSecs := time.Now().Unix()
	// rand.Seed(seedSecs)
	// randNum := rand.Intn(50) + 1 //random number from 1-49 + 1
	// pl(randNum)

	// //mat function
	// pl("Math Abs : ",  math.Abs(-10))
	// pl("Math Pow : ", math.Pow(2,3))
	// r90 := 90 * math.Pi / 180 // derajat
	// pl("Sin (90) : ", math.Sin(r90))
	
	// //printf
	// /*
	// 	%d = integer
	// 	%c = character
	// 	%f = float
	// 	%t = boolean
	// 	%s = string
	// 	%o = base 8
	// 	%x = base 16
	// */
	// fmt.Printf("%9f\n", 3.14)
	// fmt.Printf("%.2f\n", 3.14232)
	// fmt.Printf("%9.f\n", 3.14232) //menghapus semua desimal
	// sp1 := fmt.Sprintf("%9.f", 3.14232)
	// pl(sp1)

	// //for loop------------------------
	// //for initialization; condition; postStatement {BODY}
	// for i:= 0; i<=5; i++ {
	// 	pl(i)
	// }
	// fx :=0
	// for fx <= 5 {
	// 	pl(fx)
	// 	fx++
	// }

	// //mini project, tebak angka random
	// seedSecs := time.Now().Unix()
	// rand.Seed(seedSecs)
	// randNum := rand.Intn(50) + 1 //random number from 1-49 + 1
	// pl(randNum)
	
	// for true {
	// fmt.Print("Tebak angkanya : ")
	// reader := bufio.NewReader(os.Stdin)
	// input, err := reader.ReadString('\n')
	// if err!=nil {
	// 	log.Fatal(err)
	// }
	// guess := strings.TrimSpace(input)
	// angkaInput, err := strconv.Atoi(guess)
	// if err!=nil {
	// 	log.Fatal(err)
	// }
	// if angkaInput > randNum{
	// 	pl("Pilih lebih kecil!")
	// } else if angkaInput < randNum {
	// 	pl("Pilih lebih besar")
	// } else if angkaInput == randNum{
	// 	pl("Selamat!!")
	// 	break
	// }
	// }
	
	arr := []int{1,2,3}
	for _, num := range arr {
		pl(num)
	}
}