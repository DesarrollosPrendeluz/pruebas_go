package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Skus struct {
	ID          int64
	Sku_type_id int64
	Item_id     int64
	Code        string
	Created_at  []uint8
	Updated_at  []uint8
}

func main() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:30306",
		DBName: "prend_items_luz",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	//fmt.Println("Connected!")

	// results, err := getSkus()
	// log.Println(results)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Lista de Skus: %v\n", results)

	router := gin.Default()
	router.GET("/test", getSkusApi)
	router.Run("localhost:8080")
}

func getSkus() ([]Skus, error) {
	var skusList []Skus
	rows, err := db.Query("SELECT * from skus")

	if err != nil {
		return nil, fmt.Errorf("skus : %v", err)

	}

	defer rows.Close()

	for rows.Next() {
		var sku Skus
		if err := rows.Scan(&sku.ID, &sku.Item_id, &sku.Sku_type_id, &sku.Code, &sku.Created_at, &sku.Updated_at); err != nil {
			return nil, fmt.Errorf("error iterando %v", err)
		}
		skusList = append(skusList, sku)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist: %v", err)
	}

	return skusList, nil
}

func getSkusApi(c *gin.Context) {
	results, err := getSkus()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Nothing found"})
	}
	c.IndentedJSON(http.StatusOK, results)
}
