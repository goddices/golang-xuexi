package thirdlib

import (
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/go-resty/resty/v2"
)

type ResponseData struct {
	Areas []Area `json:"areas"`
}

type Area struct {
	AreaID   string  `json:"area_id"`
	AreaName string  `json:"area_name"`
	CenterXy string  `json:"center_xy"`
	Rate     float32 `json:"rate"`
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

	fmt.Println(len(resp.Body()))

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
}
