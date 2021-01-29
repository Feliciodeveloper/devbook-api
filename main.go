package main

import (
	"api/src/config"
	"api/src/orm"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)
/*func init(){
	key := make([]byte,64)
	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}
	str64 := base64.StdEncoding.EncodeToString(key)
	fmt.Println(str64)
}*/
func main() {
	config.Loading()
	orm.AutoMigration()
	r := router.Generate()
	fmt.Printf("Rodando a API na porta %d\n",config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",config.Port),r))
}
