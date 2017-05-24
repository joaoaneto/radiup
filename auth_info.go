package main

import (
	//"github.com/joaoaneto/radiup/streamer"
	"github.com/joaoaneto/radiup/streamer/repository"
	"fmt"
)

func main() {

	//authInfo1 := streamer.OAuthInfo{"asdasd-2112-kdasd", "kfewe-3123-dasda"}
	//authInfo2 := streamer.OAuthInfo{"asidsadjasid", "12312dasd313"}
	//authInfo3 := streamer.OAuthInfo{"dasjd71en1nd", "dkasdas-123-daisnd"}

	//register test
	persistor := repository.NewPersistor()
	//persistor.RegisterOAuthInfo(authInfo1)
	//persistor.RegisterOAuthInfo(authInfo2)
	//persistor.RegisterOAuthInfo(authInfo3)

	//remove test
	//persistor.RemoveOAuthInfo("asdasd-2112-kdasd")

	//search test
	authInfoReceive := persistor.SearchOAuthInfo("asidsadjasid")
	fmt.Println("O que foi recebido: ", authInfoReceive.SecretKey)

	//update test
	//persistor.UpdateOAuthInfo("dasjd71en1nd", "zezinchigaboa")
	//fmt.Println("Atualizou!")

}