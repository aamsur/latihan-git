package siswa

import (
	"net/http"
	"encoding/json"
	"os"
)

type(
	HandleRequest struct {
		Nama   string `json:"nama"`
		Alamat string `json:"alamat"`
		Hobby  string `json:"hobby"`
		Umur   int `json:"umur"`
	}
)

func SiswaHandler(w http.ResponseWriter, r *http.Request) {
	var rq HandleRequest
	var sw Siswa
	var k Kelas

	var decoder = json.NewDecoder(r.Body)

	e := decoder.Decode(&rq)
	if e != nil {
		panic("error bro")
		os.Exit(0)
	}

	sw.Nama = rq.Nama
	sw.Alamat = rq.Alamat
	sw.Hobby = rq.Hobby
	sw.Umur = rq.Umur

	k.Daftar(sw)

	var rt, _ = json.Marshal(k.Siswa)
	w.Header().Set("Content-Type", "application/json")
	w.Write(rt)
}