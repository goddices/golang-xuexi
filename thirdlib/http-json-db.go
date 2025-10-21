package thirdlib

import (
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/go-resty/resty/v2"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ResponseData struct {
	Areas []Area `json:"areas"`
}

type Area struct {
	AreaID   string  `json:"area_id" gorm:"size:50;primaryKey"`
	AreaName string  `json:"area_name" gorm:"size:50;`
	CenterXy string  `json:"center_xy" gorm:"size:50;`
	Rate     float32 `json:"rate" gorm:"size:50;`
}

func RestyTest() {
	// 全局 Resty 客户端（建议初始化一次）
	var client = resty.New()
	resp, err := client.R().
		SetAuthToken("eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50VHlwZSI6ImNhb2hlamluZyIsImF1ZCI6WyI0ZUZmUUtoeXZ1SHJpQUdFb2w0ZVIxIl0sImVuY2wiOlsiIl0sImV4cCI6MTc2MzYyMDk1NCwiaXNzIjoibWlyYWNsZSIsImp0aSI6InFBdlRvUFo3Iiwic3ViIjoic2IudXNlciIsInR5cCI6ImFjY2VzcyJ9.9iRaytXbzstMOowSyN3iReP9Wq035MpaYAYVfRvBdeJl4dXr2QrdlJqQGhlulgV9EZi0NrmsS1jD3wNz2NrcVw").
		Get("https://test.sheepwall.com/query/digiflow/api/v1/commercial-wow?enclosure_id=LinGangA7ZCrBAsIcWeqUN&week_time=1748793600")
	if err != nil {
		panic(err)
	}
	fmt.Println("Response Info:")
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Body       :", resp.String())

	// fmt.Println(len(resp.Body()))

	var responseData ResponseData
	err = sonic.Unmarshal(resp.Body(), &responseData)
	if err != nil {
		panic(err)
	}
	fmt.Println("Parsed Response Data:")
	for _, area := range responseData.Areas {
		fmt.Printf("  Area ID: %s, Area Name: %s, Center XY: %s, Rate: %.2f\n",
			area.AreaID, area.AreaName, area.CenterXy, area.Rate)
	}

	dsn := "host=10.8.0.1 user=postgres password=pinefield port=25432 dbname=space_mcp sslmode=disable"
	// 通过 GORM 驱动连接 PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接 PostgreSQL 失败")
	}
	db.AutoMigrate(&Area{})
	fmt.Println("成功连接到 PostgreSQL 数据库并自动迁移 Area 结构体")
	// 后续通过 db 执行 ORM 操作（如创建表、查询数据等）

	for _, area := range responseData.Areas {
		var existingArea Area
		result := db.First(&existingArea, "area_id = ?", area.AreaID)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				// 记录不存在，插入新记录
				db.Create(&area)
				fmt.Printf("插入新区域：%s\n", area.AreaName)
			} else {
				// 其他错误
				fmt.Printf("查询区域失败：%v\n", result.Error)
			}
		} else {
			// 记录已存在，更新记录
			db.Model(&existingArea).Updates(area)
			fmt.Printf("更新区域：%s\n", area.AreaName)
		}
	}
}
