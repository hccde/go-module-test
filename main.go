package main
import (
	"net/http"
	"fmt"
	"io/ioutil"
)


func main(){
	res,err := http.Get("http://www.baidu.com/")
	if err != nil {
		fmt.Println(11111)
		fmt.Println(err.Error())
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
    fmt.Println(string(body))
}

// main()
