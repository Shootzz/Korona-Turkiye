package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"os"
)

func main() {
	// curl the API content
	url := "https://covid19.saglik.gov.tr/covid19api?getir=sondurum"
        req, _ := http.NewRequest("GET", url, nil)
        res, _ := http.DefaultClient.Do(req)
        body, _ := ioutil.ReadAll(res.Body)
        defer res.Body.Close()
        result := string(body)

	// trim the result get the variables
	r_trim := strings.Split(result, "\"")
	vaka := r_trim[11]
	test_sayisi := r_trim[7]
	iyilesen := r_trim[19]
	tarih := r_trim[3]

	// check the file or created and write it
	if _, err := os.Stat("vakalar.txt"); err == nil {
                b, err := ioutil.ReadFile("vakalar.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
                if strings.Contains(string(b), tarih) {
                        fmt.Println("Tarih           :"+tarih+"\n"+"Vaka Sayısı     :"+vaka+"\n"+"Test Sayısı     :"+test_sayisi+"\n"+"İyileşen Sayısı :"+iyilesen+"\n")
			fmt.Println("Yeni vakalar henüz açıklanmadı..")
                }else{
			f, err := os.OpenFile("vakalar.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}
			f.WriteString("\nTarih           :"+tarih+"\n"+"Vaka Sayısı     :"+vaka+"\n"+"Test Sayısı     :"+test_sayisi+"\n"+"İyileşen Sayısı :"+iyilesen+"\n")
                	fmt.Println("Tarih           :"+tarih+"\n"+"Vaka Sayısı     :"+vaka+"\n"+"Test Sayısı     :"+test_sayisi+"\n"+"İyileşen Sayısı :"+iyilesen+"\n")
			fmt.Println("Yeni veriler vakalar dosyasına yazıldı..")
		}
        } else {
                f, err := os.Create("vakalar.txt")
                if err != nil {
                        fmt.Println(err)
                        return
                }
                f.WriteString("\nTarih           :"+tarih+"\n"+"Vaka Sayısı     :"+vaka+"\n"+"Test Sayısı     :"+test_sayisi+"\n"+"İyileşen Sayısı :"+iyilesen+"\n")
		fmt.Println("Tarih           :"+tarih+"\n"+"Vaka Sayısı     :"+vaka+"\n"+"Test Sayısı     :"+test_sayisi+"\n"+"İyileşen Sayısı :"+iyilesen+"\n")
		fmt.Println("Vakalar.txt dosyası oluşturuldu ve veriler yazıldı..")
        }
}
