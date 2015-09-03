// web server, render HTML data from DB
package  main

import ( 
"fmt"
//"html/template"
"os"
"database/sql"
_ "github.com/go-sql-driver/mysql"
)

func write2html(db *sql.DB) {
	file, _ := os.Create("index.html")

	file.Write([]byte(`
<!DOCTYPE html>
	<html>
	<head>
	<style>
		table, th, td {
    	border: 1px solid black;
    	border-collapse: collapse;
		}
		th, td {
    	padding: 5px;
		}
	</style>
	</head>

	<body>
	`))

  file.Write([]byte(`
  <table style="width:100%">
  <tr>
  `))

  rows, _ := db.Query("SELECT * FROM device_info")
  columns, _ := rows.Columns()
  values := make([]sql.RawBytes, len(columns))
  scanArgs := make([]interface{}, len(values))
  for i := range values {
  	scanArgs[i] = &values[i]
  }

  var str string

  for i := range values {
  	str = "<th>" + columns[i] + "</th>" 
  	file.Write([]byte(str))
  }

  file.Write([]byte(`</tr>
  `))
  
  for rows.Next() {
        
  	err := rows.Scan(scanArgs...)
  	if err != nil {
      fmt.Println(err) 
    }
    
    var s string
    s = "<tr>"

    var value string
    for i, col := range values {
            
      if col == nil {
        value = "NULL"
      } else {
        value = string(col)
      }
      s += "<td>" + value +"</td>"
      fmt.Print(columns[i], ": ", value, "\t")
    }

    s+= "</tr>"
    file.Write([]byte(s))
          
    fmt.Println("\n-----------------------------------")
  }

  if err := rows.Err(); err != nil {
    fmt.Println(err) 
  }

	file.Write([]byte(`
	</table>
	</body>
	</html>`))
}

func main() {
	
	db, _ := sql.Open("mysql", "root:@/go_db")

  write2html(db)

}
