package main
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

type product struct{
    id int
    model string
    company string
    price int
}
 
func main() { 
 
    connStr := "host=localhost port=5432 user=gouser password=gopassword dbname=gopoductdb sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    } 
    defer db.Close()
     
    result, err := db.Exec("insert into Products (model, company, price) values ('Notebook', $1, $2)", 
        "Asus", 10000)
    if err != nil{
        panic(err)
    }

    fmt.Println(result.RowsAffected())  // количество добавленных строк

	rows, err := db.Query("select * from products")
    if err != nil {
        panic(err)
    }
    defer rows.Close()
    products := []product{}
     
    for rows.Next(){
        p := product{}
        err := rows.Scan(&p.id, &p.model, &p.company, &p.price)
        if err != nil{
            fmt.Println(err)
            continue
        }
        products = append(products, p)
    }
    for _, p := range products{
        fmt.Println(p.id, p.model, p.company, p.price)
    }

}