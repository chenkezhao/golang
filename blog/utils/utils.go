package utils

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	// "time"
)

type IPInfo struct {
	Code int           `json:"code"`
	Data *IPDetailInfo `json:"data"`
}

type IPDetailInfo struct {
	Country    string `json:"country"`    // "中国"
	Country_id string `json:"country_id"` // "CN"
	Area       string `json:"area"`       // "华北"
	Area_id    string `json:"area_id"`    //0000"
	Region     string `json:"region"`     //"北京市"
	Region_id  string `json:"region_id"`  //"0000"
	City       string `json:"city"`       //"北京市"
	City_id    string `json:"city_id"`    //"0100"
	County     string `json:"county"`     // ""
	County_id  string `json:"county_id"`  // "-1"
	Isp        string `json:"isp"`        //天信科技"
	Isp_id     string `json:"isp_id"`     //0157"
	Ip         string `json:"ip"`         //3.141"
}

//使用淘宝接口，http://ip.taobao.com/service/getIpInfo.php?ip=119.57.173.141
func GetAddressByIP(ip string) string {
	var addressInfo = "空白"

	req := httplib.Get("http://ip.taobao.com/service/getIpInfo.php?ip=" + ip)
	// req.SetTimeout(100*time.Second, 30*time.Second)

	_, err := req.Response()
	if err != nil {
		fmt.Println("------------------err:", err)
		addressInfo = "40行-err"
	}
	// fmt.Println("------------------resp:", resp)

	var r IPInfo
	err = req.ToJSON(&r)
	if err != nil {
		fmt.Println("------------------err：", err)
		addressInfo = "54行-err"
	}
	// fmt.Println("--------------"+ip+"----->>>>>>>>>>>>>>>>>>>>>>>>>>>", result.Data)
	d := r.Data
	// fmt.Println("--------------"+ip+"----->>>>>>>>>>>>>>>>>>>>>>>>>>>", d)
	if d != nil {
		addressInfo = fmt.Sprintf("%s-%s-%s-%s-%s", d.Country, d.Area, d.City, d.County, d.Isp)
	}
	return addressInfo
}
