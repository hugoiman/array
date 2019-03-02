package admin

func checkErr(err error){
  if err != nil {
    panic(err.Error())
  }
}
