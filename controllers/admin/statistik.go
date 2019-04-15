package admin

import (
  // "fmt"
  "net/http"
  "array/models/admin"
  // "array/models/member"
  "array/structs"
  "github.com/labstack/echo"
)

var total_pemasukan, average_pemasukan, total_pengeluaran, average_pengeluaran, total_laba, average_laba int

func HitungPemasukan(c echo.Context) error {
  pemasukan := admin.GetPemasukan()
  total_pemasukan = 0
  for i := 0; i < len(pemasukan.PemasukanBulanan); i++ {
    total_pemasukan += pemasukan.PemasukanBulanan[i].JumlahPemasukan
  }

  average_pemasukan = total_pemasukan/len(pemasukan.PemasukanBulanan)
  // persentase_pemasukan := fmt.Sprintf("%.2f", float64(average_pemasukan)/float64(total_pemasukan)*100)

  data := struct {
    Pemasukan           structs.PemasukanBulanan
    TotalPemasukan      int
    AveragePemasukan    int
  } {
    pemasukan,
    total_pemasukan,
    average_pemasukan,
  }

  return c.JSON(http.StatusOK, data)
}

func HitungPengeluaran(c echo.Context) error {
  pengeluaran := admin.GetPengeluaran()
  total_pengeluaran = 0
  for i := 0; i < len(pengeluaran.Pengeluaran); i++ {
    total_pengeluaran += pengeluaran.Pengeluaran[i].JumlahPengeluaran
  }

  average_pengeluaran = total_pengeluaran/len(pengeluaran.Pengeluaran)

  data := struct {
    Pengeluaran           structs.Pengeluaran
    TotalPengeluaran      int
    AveragePengeluaran    int
  } {
    pengeluaran,
    total_pengeluaran,
    average_pengeluaran,
  }

  return c.JSON(http.StatusOK, data)
}

func HitungLaba(c echo.Context) error {
  // Hitung Pemasukan
  pemasukan := admin.GetPemasukan()
  total_pemasukan = 0
  for i := 0; i < len(pemasukan.PemasukanBulanan); i++ {
    total_pemasukan += pemasukan.PemasukanBulanan[i].JumlahPemasukan
  }

  // Hitung Pengeluaran
  pengeluaran := admin.GetPengeluaran()
  total_pengeluaran = 0
  for i := 0; i < len(pengeluaran.Pengeluaran); i++ {
    total_pengeluaran += pengeluaran.Pengeluaran[i].JumlahPengeluaran
  }

  total_laba   = total_pemasukan - total_pengeluaran
  average_laba = total_laba/len(pemasukan.PemasukanBulanan)

  data := struct {
    TotaLaba        int
    AverageLaba     int
  } {
    total_laba,
    average_laba,
  }
  // fmt.Printf("%+v\n",pemasukan)

  return c.JSON(http.StatusOK, data)
}
