package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

// định nghĩa một kiểu dữ liệu có các trường thông tin cơ bản liên quan đến một mục công việc,
// và sử dụng các tag json để chỉ định tên của các trường khi chuyển đổi thành dữ liệu JSON.
type TodoItems struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

func main() {
	//ket noi voi csdl
	dsn := os.Getenv("DB_CONN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(db)
	now := time.Now().UTC()
	item := TodoItems{
		Id:          1,
		Title:       "This is item 1",
		Description: "description 1",
		Status:      "Doing",
		CreatedAt:   &now,
		UpdatedAt:   nil,
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": item,
		})
	})
	r.Run(":3000")

}
