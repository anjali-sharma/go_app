// connect to DB, pop data from redis queue and insert to DB

package  main

import ( 
"fmt"
"github.com/xuyu/goredis"
"database/sql"
_ "github.com/go-sql-driver/mysql"
)

func fetchFromRedis(db *sql.DB, client *goredis.Redis) {

	reply, err := client.RPop("Rqueue")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(reply))

	storeinDB(db, string(reply))
}

func storeinDB(db *sql.DB, reply string) {
	stmtIns, err := db.Prepare("INSERT INTO device_info (payload, created_at) VALUES(?, CURRENT_TIMESTAMP)") // ? = placeholder

	if err != nil {
    fmt.Println() 
  }
  defer stmtIns.Close() 

  _, err2 := stmtIns.Exec(reply);

  if err2 != nil {
  	fmt.Println(err2)
  } else {
  	fmt.Println("Value inserted to DB...")
  }

}

func main() {

	db, err := sql.Open("mysql", "root:@/go_db")
	client, err2 := goredis.Dial(&goredis.DialConfig{Address: "127.0.0.1:6379"})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Database Connection established...")
	}

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Redis Connection established...")
	}

	fetchFromRedis(db, client)

	defer db.Close()
}