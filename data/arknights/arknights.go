package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const JSONFilePath = "../../src/assets/data/arknights.json"

// 收集每次从官网获取的抽卡数据, 然后合并到本地JSON文件中.
func main() {
	token := ""
	csrf := ""

	history := LocalHistory()
	var maxTs int64 = 0
	var maxWish = Wish{}
	for _, wish := range history {
		if wish.Ts > maxTs {
			maxTs = wish.Ts
			maxWish = wish
		}
	}
	log.Printf("本地最新抽卡记录: %s\n", maxWish.String())
	page := 1
	for {
		pageStr := "第" + strconv.Itoa(page) + "页"
		curGacha := FetchGacha(page, token, csrf)
		if curGacha.Code != 0 {
			log.Printf(pageStr+"查询失败,code=%d\n", curGacha.Code)
			break
		}
		listLen := len(curGacha.Data.List)
		if listLen == 0 {
			log.Println(pageStr + "无数据,结束查询")
			break
		} else {
			log.Printf(pageStr+"共%d条数据\n", listLen)
		}
		hasNew := false
		for _, wish := range curGacha.Data.List {
			if wish.Ts > maxTs {
				log.Printf("查询到新的抽卡记录,%s\n", wish.String())
				hasNew = true
				history = append(history, wish)
			}
		}
		if !hasNew {
			log.Println("已获取所有最新的抽卡记录")
			break
		}
		if listLen < 10 {
			break
		}
		time.Sleep(1 * time.Second)
		page++
	}
	StoreWishes(history)
}

func LocalHistory() []Wish {
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
	var history []Wish
	_ = json.Unmarshal(historyFileContent, &history)
	return history
}

func FetchGacha(page int, token string, csrf string) Gacha {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://ak.hypergryph.com/user/api/inquiry/gacha?page="+strconv.Itoa(page)+"&channelId=1&token="+token, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("x-csrf-token", csrf)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	history := Gacha{}
	_ = json.Unmarshal(bodyText, &history)
	return history
}

// StoreWishes 存储抽卡历史
func StoreWishes(wishMap []Wish) {
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

// SortWishMap 根据 ts (时间)进行排序
func SortWishMap(ist []Wish) []Wish {
	sort.Slice(ist, func(i, j int) bool {
		return ist[i].Ts < ist[j].Ts
	})
	return ist
}

type Gacha struct {
	Code int64 `json:"code"`
	Data struct {
		List       []Wish `json:"list"`
		Pagination struct {
			Current int64 `json:"current"`
			Total   int64 `json:"total"`
		} `json:"pagination"`
	} `json:"data"`
	Msg string `json:"msg"`
}
type Wish struct {
	Chars []Char `json:"chars"`
	Pool  string `json:"pool"`
	Ts    int64  `json:"ts"`
}
type Char struct {
	IsNew  bool   `json:"isNew"`
	Name   string `json:"name"`
	Rarity int    `json:"rarity"`
}

func (c Char) String() string {
	n := ""
	if c.IsNew {
		n = "(新)"
	}
	return fmt.Sprintf("%s-%s"+n, c.Name, strconv.Itoa(c.Rarity+1)+"星")
}

func (w Wish) String() string {
	if w.Ts == 0 {
		return ""
	}
	tm := TsToTime(w.Ts)
	chars := make([]string, 0)
	for _, char := range w.Chars {
		chars = append(chars, char.String())
	}
	wishType := "抽卡"
	if len(w.Chars) == 1 {
		wishType = "单抽"
	} else if len(w.Chars) == 10 {
		wishType = "十连"
	}
	return fmt.Sprintf("[%s]Time[%s]Pool[%s]Chars[%s]", wishType, tm, w.Pool, strings.Join(chars, ","))
}

func TsToTime(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02 03:04:05")
}
