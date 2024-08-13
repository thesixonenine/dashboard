package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const WishHistoryFilePath = "C:\\Program Files\\Star Rail\\Game\\StarRail_Data\\webCaches\\2.27.0.0\\Cache\\Cache_Data\\data_2"
const JSONFilePath = "../../src/assets/data/star-rail-wish.json"

var re = regexp.MustCompile(`\p{C}`)

var gachaTypeMap = map[string]string{
	"11": "角色活动跃迁",
	"12": "光锥活动跃迁",
	"1":  "常驻跃迁",
	"2":  "新手跃迁",
}
var absParams = []string{"authkey_ver", "sign_type", "auth_appid", "lang", "authkey", "game_biz", "page", "size", "gacha_type"}

func main() {
	// 使用抽卡 URL 进行循环查询抽卡历史, 一但发现已经存在于历史 JSON 文件中, 则停止查询
	FetchWishes(FindURLFromLogFile(), LocalHistoryJSONFileToMap())
}

// FindURLFromLogFile 查询日志文件中的抽卡 URL
func FindURLFromLogFile() UrlParam {
	for _, p := range strings.Split(WishHistoryFilePath, "\\") {
		_, err := strconv.Atoi(strings.ReplaceAll(p, ".", ""))
		if err != nil {
			continue
		}
		log.Println("Version: " + p)
	}
	content, err := os.ReadFile(WishHistoryFilePath)
	if err != nil {
		log.Fatalf("日志文件[%s]未找到", WishHistoryFilePath)
	}
	lastUrl := ""
	nMap := map[string]string{}
	split := strings.Split(string(content), "1/0/")
	for i := len(split) - 1; i >= 0; i-- {
		s := split[i]
		t := re.ReplaceAllString(strings.TrimSpace(s), "")
		if !strings.HasPrefix(s, "http") || !strings.Contains(t, "getGachaLog") {
			continue
		}
		u, parseErr := url.Parse(t)
		if parseErr != nil {
			continue
		}
		queryMap := ParseQuery(u.RawQuery)
		queryMap["end_id"] = "0"
		if FetchData("https://api-takumi.mihoyo.com/common/gacha_record/api/getGachaLog?"+ParamMapToStr(queryMap)).Retcode != 0 {
			continue
		}

		for k, v := range queryMap {
			if slices.Contains(absParams, k) {
				nMap[k] = v
			}
		}
		lastUrl = u.Scheme + "://" + u.Host + u.Path
		log.Println("find the wish history url.")
		break
	}
	return UrlParam{lastUrl, nMap}
}

// ParseQuery 解析 URL 参数为 map
func ParseQuery(q string) map[string]string {
	m := map[string]string{}
	for _, s := range strings.Split(q, "&") {
		split := strings.Split(s, "=")
		if len(split) >= 2 {
			m[split[0]] = split[1]
		}
	}
	return m
}

// LocalHistoryJSONFileToMap 将本地抽卡历史 JSON 文件转为 map 对象, 如果文件不存在则创建
func LocalHistoryJSONFileToMap() map[string][]HKRPGWish {
	historyFile, err := os.OpenFile(JSONFilePath, syscall.O_RDWR|syscall.O_CREAT, os.ModePerm)
	if err != nil {
		log.Fatalf("打开文件[%s]异常,err[%s]\n", JSONFilePath, err.Error())
	}
	defer func() {
		if err := historyFile.Close(); err != nil {
			log.Fatalf("关闭文件[%s]异常,err[%s]\n", JSONFilePath, err.Error())
		}
	}()
	historyFileContent, err := io.ReadAll(historyFile)
	if err != nil {
		log.Fatalf("读取文件[%s]异常,err[%s]\n", JSONFilePath, err.Error())
	}
	history := map[string][]HKRPGWish{}
	_ = json.Unmarshal(historyFileContent, &history)
	return history
}

// FetchWishes 从指定 URL 及参数中拉取抽卡参数, 并追加到 Map 中
func FetchWishes(urlParam UrlParam, localHistoryMap map[string][]HKRPGWish) {
	if urlParam.BaseUrl == "" || len(urlParam.ParamMap) == 0 {
		log.Println("cannot find the wish history url.")
		return
	}
	paramMap := urlParam.ParamMap
	// 循环抽卡类型
	for k, v := range gachaTypeMap {
		fmt.Printf("开始获取[%s]\n", v)
		localIdList := MapToId(localHistoryMap[k])
		// 如果是新手跃迁且已经抽了50抽则直接跳过
		if k == "2" && len(localIdList) == 50 {
			continue
		}
		page := 1
		size := 5
		paramMap["gacha_type"] = k
		paramMap["page"] = strconv.Itoa(page)
		paramMap["end_id"] = "0"
		paramMap["size"] = strconv.Itoa(size)
		for {
			wishList := FetchData(urlParam.BaseUrl + "?" + ParamMapToStr(paramMap)).Data.List
			if wishList == nil {
				fmt.Println("未获取到数据")
				break
			}
			isContains := false
			for _, wish := range wishList {
				if slices.Contains(localIdList, wish.Id) {
					isContains = true
					continue
				}
				localHistoryMap[k] = append(localHistoryMap[k], wish)
				fmt.Println(wish.String())
			}
			if isContains {
				break
			}
			dataLen := len(wishList)
			if dataLen == 0 {
				break
			}
			paramMap["end_id"] = wishList[dataLen-1].Id
			if dataLen < size {
				break
			}
			page++
		}
	}
	StoreWishes(localHistoryMap)
}

// MapToId 将抽卡对象列表中的Id转成切片
func MapToId(wishes []HKRPGWish) []string {
	var idList []string
	for i := range wishes {
		idList = append(idList, wishes[i].Id)
	}
	return idList
}

// StoreWishes 存储抽卡历史
func StoreWishes(wishMap map[string][]HKRPGWish) {
	wishMap = SortWishMap(wishMap)
	marshal, err := json.Marshal(wishMap)
	if err != nil {
		log.Fatalf("JSON序列化异常[%s]\n", err.Error())
		return
	}
	WriteToFile(JSONIndent(marshal))
}

// JSONIndent 进行 JSON 格式化
func JSONIndent(marshal []byte) []byte {
	var out bytes.Buffer
	err := json.Indent(&out, marshal, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	return out.Bytes()
}

// WriteToFile 写入文件
func WriteToFile(marshal []byte) {
	err := os.WriteFile(JSONFilePath, marshal, syscall.O_RDWR|syscall.O_CREAT)
	if err != nil {
		log.Fatalf("写入文件异常[%s]\n", err.Error())
	}
}

// SortWishMap 根据抽卡 ID (时间)进行排序
func SortWishMap(ist map[string][]HKRPGWish) map[string][]HKRPGWish {
	list := map[string][]HKRPGWish{}
	for k, v := range ist {
		sort.Slice(v, func(i, j int) bool {
			return v[i].Id < v[j].Id
		})
		list[k] = v
	}
	return list
}

// ParamMapToStr 将参数 Map 使用 & 连接, 转成字符串
func ParamMapToStr(paramMap map[string]string) string {
	var paramStr []string
	for k, v := range paramMap {
		paramStr = append(paramStr, k+"="+v)
	}
	return strings.Join(paramStr, "&")
}

// FetchData 从指定 URL 获取抽卡历史并转成分页对象
func FetchData(link string) Page[HKRPGWish] {
	time.Sleep(5 * time.Second)
	resp, err := http.Get(link)
	if err != nil {
		log.Fatalf("HTTP请求异常,err[%s]", err.Error())
	}
	body := resp.Body
	defer func() {
		if err := body.Close(); err != nil {
			log.Fatalf("关闭HTTP请求异常,err[%s]", err.Error())
		}
	}()
	bodyByte, httpReadErr := io.ReadAll(resp.Body)
	if httpReadErr != nil {
		log.Fatalf("读取HTTP Body异常,err[%s]", err.Error())
	}
	p := Page[HKRPGWish]{}
	if resp.StatusCode != 200 {
		p.Retcode = -1
		return p
	}
	_ = json.Unmarshal(bodyByte, &p)
	return p
}

type HKRPGWish struct {
	Uid       string `json:"uid"`
	GachaId   string `json:"gacha_id"`
	GachaType string `json:"gacha_type"`
	ItemId    string `json:"item_id"`
	Count     string `json:"count"`
	Time      string `json:"time"`
	Name      string `json:"name"`
	Lang      string `json:"lang"`
	ItemType  string `json:"item_type"`
	RankType  string `json:"rank_type"`
	Id        string `json:"id"`
}

func (wish HKRPGWish) String() string {
	return fmt.Sprintf("%s %s %s", wish.Time, gachaTypeMap[wish.GachaType], wish.Name)
}

type Page[T any] struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		Page           int    `json:"page"`
		Size           int    `json:"size"`
		List           []T    `json:"list"`
		Region         string `json:"region"`
		RegionTimeZone int    `json:"region_time_zone"`
	} `json:"data"`
}

type UrlParam struct {
	BaseUrl  string
	ParamMap map[string]string
}
