package main

import "fmt"
import "gopkg.in/mgo.v2"
import "os"
import "net/http"
import "html/template"

type mahasiswa struct {
	Id    int     `bson:"nim"`
	Nama  string  `bson:"nama"`
	Grade float64 `bson:"ipk"`
}

func connect() *mgo.Session {
	var session, err = mgo.Dial("localhost")
	if err != nil {
		os.Exit(0)
	}
	return session
}

func tampil(res http.ResponseWriter, req *http.Request) {
	session := connect()
	defer session.Close()

	halaman, _ := template.New("h1").Parse(indeks)

	collection := session.DB("ka").C("mahasiswa")
	var hasil []mahasiswa
	var err = collection.Find(nil).All(&hasil)
	if err != nil {
		fmt.Println("gagal mengambil")
		os.Exit(0)

	}

	for _, data_mahasiswa := range hasil {
		fmt.Println("nim :", data_mahasiswa.Id)
		fmt.Println("nama :", data_mahasiswa.Nama)
		fmt.Println("ipk :", data_mahasiswa.Grade)
		fmt.Println("")
	}
	halaman.Execute(res, hasil)
}

const indeks = `<table width='100%' border='10'>
				<tr><td width='20%'>{{.nim}}</td><td width='60%'>{{.nama}}</td><td width='20%'>{{.ipk}}</td></tr>
				</table>`

func main() {

	http.HandleFunc("/", tampil)

	fmt.Println("localhost:8080 running now....")
	http.ListenAndServe(":8080", nil)

}
