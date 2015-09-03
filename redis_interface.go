package  main

import ( 
	"fmt"
	"github.com/xuyu/goredis"
	"time"
	)

func main() {
	client, err := goredis.Dial(&goredis.DialConfig{Address: "127.0.0.1:6379"})

	for {

		fmt.Println("\033[H\033[2J")
		fmt.Println("============= Data in Redis Key Value Pair ==========================\n")

	    keys, err2 := client.ExecuteCommand("KEYS", "*")

	    length := len(keys.Multi)
	    //fmt.Println(keys.)

	    for i := 0; i < length; i++ {
	    	key := string(keys.Multi[i].Bulk)
	    	//fmt.Println(key)
		     value , err3 := client.Get(key)
		    fmt.Println("Key: ", string(key), "\tValue: ", string(value))
		    Printerror(err)
		    Printerror(err2)
		    Printerror(err3)
		}
	    time.Sleep(25 * time.Millisecond)
	}
}

func Printerror(err error){
   return
}
