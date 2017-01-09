//author aamsur

package siswa

type Siswa struct {
	Nama   string
	Alamat string
	Hobby  string
	Umur   int
}

type Kelas struct {
	Siswa []Siswa
}

func (k *Kelas) Daftar(sw Siswa) {
	k.Siswa = append(k.Siswa, sw)
}

