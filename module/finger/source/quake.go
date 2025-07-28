package source

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/color"
)

type QuakeServiceInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		IP   string `json:"ip"`
		Port int    `json:"port"`
		ID   string `json:"id"`
	} `json:"data"`
	Meta struct {
		Pagination struct {
			Count     int `json:"count"`
			PageIndex int `json:"page_index"`
			PageSize  int `json:"page_size"`
			Total     int `json:"total"`
		} `json:"pagination"`
	} `json:"meta"`
}

func GetQuakeConfig() Config {
	//创建一个空的结构体,将本地文件读取的信息放入
	c := &Config{}
	//创建一个结构体变量的反射
	cr := reflect.ValueOf(c).Elem()
	//打开文件io流
	f, err := os.Open("config.ini")
	if err != nil {
		log.Fatal(err)
		color.RGBStyleFromString("237,64,35").Println("[Error] Quake configuration file error!!!")
		os.Exit(1)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	//我们要逐行读取文件内容
	s := bufio.NewScanner(f)

	for s.Scan() {
		//以=分割,前面为key,后面为value
		var str = s.Text()
		var index = strings.Index(str, "=")
		var key = strings.TrimSpace(str[0:index])
		var value = strings.TrimSpace(str[index+1:])
		//通过反射将字段设置进去
		cr.FieldByName(key).Set(reflect.ValueOf(value))
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	//返回Config结构体变量
	return *c
}

func quake_api(keyword string, apiKey string, pageSize int) (string, []byte) {
	url := "https://quake.360.net/api/v3/search/quake_service"
	
	data := make(map[string]interface{})
	data["query"] = keyword
	data["start"] = "0"
	data["size"] = strconv.Itoa(pageSize)
	data["include"] = []string{"ip", "port"}
	
	jsonData, _ := json.Marshal(data)
	
	return url, jsonData
}

// 请求api
func quakeHttp(url string, jsonData []byte, apiKey string, timeout string) *QuakeServiceInfo {
	var itime, err = strconv.Atoi(timeout)
	if err != nil {
		log.Println("quake超时参数错误: ", err)
	}
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{
		Timeout:   time.Duration(itime) * time.Second,
		Transport: transport,
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonData)))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("X-QuakeToken", apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*;q=0.8")
	req.Header.Set("Connection", "close")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2227.0 Safari/537.36")
	
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	
	res := &QuakeServiceInfo{}
	json.Unmarshal(result, &res)
	
	return res
}

// 从ID中提取域名
func extractDomainFromID(id string) string {
	if strings.Contains(id, "_") {
		parts := strings.Split(id, "_")
		if len(parts) >= 2 {
			return parts[0]
		}
	}
	return ""
}

func Quakeip(ips string) (urls []string) {
	color.RGBStyleFromString("244,211,49").Println("请耐心等待quake搜索......")
	quake := GetQuakeConfig()
	
	// 如果输入已经包含ip:前缀，直接使用；否则添加ip=前缀
	keyword := ips
	if !strings.Contains(ips, "ip:") && !strings.Contains(ips, "ip=") {
		keyword = `ip="` + ips + `"`
	}
	
	// 确保不会超速
	time.Sleep(time.Second * 2)
	
	url, jsonData := quake_api(keyword, quake.Quake_key, 100)
	res := quakeHttp(url, jsonData, quake.Quake_key, quake.Fofa_timeout)
	
	if res.Data != nil {
		for _, result := range res.Data {
			// 从ID中提取域名
			domain := extractDomainFromID(result.ID)
			
			// 如果没有提取到域名，使用IP
			if domain == "" {
				domain = result.IP
			}
			
			// 根据端口判断协议 - 只要端口号包含443就使用https
			var protocol string
			if strings.Contains(strconv.Itoa(result.Port), "443") {
				protocol = "https"
			} else {
				protocol = "http"
			}
			
			// 构建URL
			url := fmt.Sprintf("%s://%s:%d", protocol, domain, result.Port)
			urls = append(urls, url)
		}
	}
	
	return urls
}

func Quakeall(keyword string) (urls []string) {
	color.RGBStyleFromString("244,211,49").Println("请耐心等待quake搜索......")
	quake := GetQuakeConfig()
	
	// 确保不会超速
	time.Sleep(time.Second * 2)
	
	url, jsonData := quake_api(keyword, quake.Quake_key, 100)
	res := quakeHttp(url, jsonData, quake.Quake_key, quake.Fofa_timeout)
	
	if res.Data != nil {
		for _, result := range res.Data {
			// 从ID中提取域名
			domain := extractDomainFromID(result.ID)
			
			// 如果没有提取到域名，使用IP
			if domain == "" {
				domain = result.IP
			}
			
			// 根据端口判断协议 - 只要端口号包含443就使用https
			var protocol string
			if strings.Contains(strconv.Itoa(result.Port), "443") {
				protocol = "https"
			} else {
				protocol = "http"
			}
			
			// 构建URL
			url := fmt.Sprintf("%s://%s:%d", protocol, domain, result.Port)
			urls = append(urls, url)
		}
	}
	
	return urls
}

func Quakeall_out(keyword string) (result [][]string) {
	quake := GetQuakeConfig()
	
	// 确保不会超速
	time.Sleep(time.Second * 2)
	
	url, jsonData := quake_api(keyword, quake.Quake_key, 100)
	res := quakeHttp(url, jsonData, quake.Quake_key, quake.Fofa_timeout)
	
	if res.Data != nil {
		for _, d := range res.Data {
			var row []string
			row = append(row, d.IP)
			
			// 从ID中提取域名
			domain := extractDomainFromID(d.ID)
			
			// 如果没有提取到域名，使用IP
			if domain == "" {
				domain = d.IP
			}
			
			// 根据端口判断协议 - 只要端口号包含443就使用https
			var protocol string
			if strings.Contains(strconv.Itoa(d.Port), "443") {
				protocol = "https"
			} else {
				protocol = "http"
			}
			
			// 构建URL
			url := fmt.Sprintf("%s://%s:%d", protocol, domain, d.Port)
			row = append(row, url)
			row = append(row, "") // 标题为空
			row = append(row, strconv.Itoa(d.Port))
			row = append(row, "http") // 协议
			result = append(result, row)
		}
	}
	
	return result
}

func Quakeips_out(filename string) [][]string {
	fmt.Println("开始使用quake批量搜索ip,请耐心等待....")
	keys := keyword_ips(filename)
	var results [][]string
	for _, x := range keys {
		result := Quakeall_out(x)
		results = append(results, result...)
	}
	fmt.Println("共收集" + strconv.Itoa(len(results)) + "条数据，已保存！！")
	return results
}
