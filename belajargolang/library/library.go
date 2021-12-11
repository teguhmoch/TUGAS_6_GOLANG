package library

import "fmt"

func MahasiswaPublic(){
  var s1 = mahasiswa{"asep",20}
  var s2 = mahasiswa{"yusuf",21}
  s1.datasaya()
  s2.datasaya()
}

type mahasiswa struct {
  nama string
  umur int
}

func (s mahasiswa)datasaya(){
  fmt.Println("Nama Mahasiswa :",s.nama)
  fmt.Println("Umur           :",s.umur)
}
