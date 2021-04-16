package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"myProject/dao"
	"myProject/demo"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wCRC16Table = []int{
	0x0000, 0xC0C1, 0xC181, 0x0140, 0xC301, 0x03C0, 0x0280, 0xC241,
	0xC601, 0x06C0, 0x0780, 0xC741, 0x0500, 0xC5C1, 0xC481, 0x0440,
	0xCC01, 0x0CC0, 0x0D80, 0xCD41, 0x0F00, 0xCFC1, 0xCE81, 0x0E40,
	0x0A00, 0xCAC1, 0xCB81, 0x0B40, 0xC901, 0x09C0, 0x0880, 0xC841,
	0xD801, 0x18C0, 0x1980, 0xD941, 0x1B00, 0xDBC1, 0xDA81, 0x1A40,
	0x1E00, 0xDEC1, 0xDF81, 0x1F40, 0xDD01, 0x1DC0, 0x1C80, 0xDC41,
	0x1400, 0xD4C1, 0xD581, 0x1540, 0xD701, 0x17C0, 0x1680, 0xD641,
	0xD201, 0x12C0, 0x1380, 0xD341, 0x1100, 0xD1C1, 0xD081, 0x1040,
	0xF001, 0x30C0, 0x3180, 0xF141, 0x3300, 0xF3C1, 0xF281, 0x3240,
	0x3600, 0xF6C1, 0xF781, 0x3740, 0xF501, 0x35C0, 0x3480, 0xF441,
	0x3C00, 0xFCC1, 0xFD81, 0x3D40, 0xFF01, 0x3FC0, 0x3E80, 0xFE41,
	0xFA01, 0x3AC0, 0x3B80, 0xFB41, 0x3900, 0xF9C1, 0xF881, 0x3840,
	0x2800, 0xE8C1, 0xE981, 0x2940, 0xEB01, 0x2BC0, 0x2A80, 0xEA41,
	0xEE01, 0x2EC0, 0x2F80, 0xEF41, 0x2D00, 0xEDC1, 0xEC81, 0x2C40,
	0xE401, 0x24C0, 0x2580, 0xE541, 0x2700, 0xE7C1, 0xE681, 0x2640,
	0x2200, 0xE2C1, 0xE381, 0x2340, 0xE101, 0x21C0, 0x2080, 0xE041,
	0xA001, 0x60C0, 0x6180, 0xA141, 0x6300, 0xA3C1, 0xA281, 0x6240,
	0x6600, 0xA6C1, 0xA781, 0x6740, 0xA501, 0x65C0, 0x6480, 0xA441,
	0x6C00, 0xACC1, 0xAD81, 0x6D40, 0xAF01, 0x6FC0, 0x6E80, 0xAE41,
	0xAA01, 0x6AC0, 0x6B80, 0xAB41, 0x6900, 0xA9C1, 0xA881, 0x6840,
	0x7800, 0xB8C1, 0xB981, 0x7940, 0xBB01, 0x7BC0, 0x7A80, 0xBA41,
	0xBE01, 0x7EC0, 0x7F80, 0xBF41, 0x7D00, 0xBDC1, 0xBC81, 0x7C40,
	0xB401, 0x74C0, 0x7580, 0xB541, 0x7700, 0xB7C1, 0xB681, 0x7640,
	0x7200, 0xB2C1, 0xB381, 0x7340, 0xB101, 0x71C0, 0x7080, 0xB041,
	0x5000, 0x90C1, 0x9181, 0x5140, 0x9301, 0x53C0, 0x5280, 0x9241,
	0x9601, 0x56C0, 0x5780, 0x9741, 0x5500, 0x95C1, 0x9481, 0x5440,
	0x9C01, 0x5CC0, 0x5D80, 0x9D41, 0x5F00, 0x9FC1, 0x9E81, 0x5E40,
	0x5A00, 0x9AC1, 0x9B81, 0x5B40, 0x9901, 0x59C0, 0x5880, 0x9841,
	0x8801, 0x48C0, 0x4980, 0x8941, 0x4B00, 0x8BC1, 0x8A81, 0x4A40,
	0x4E00, 0x8EC1, 0x8F81, 0x4F40, 0x8D01, 0x4DC0, 0x4C80, 0x8C41,
	0x4400, 0x84C1, 0x8581, 0x4540, 0x8701, 0x47C0, 0x4680, 0x8641,
	0x8201, 0x42C0, 0x4380, 0x8341, 0x4100, 0x81C1, 0x8081, 0x4040,
}

func Crc16(pDataIn []byte, iLenIn int) int {
	wResult := 0
	wTableNo := 0
	for i := 0; i < iLenIn; i++ {
		wTableNo = (wResult & 0xff) ^ (int(pDataIn[i]) & 0xff)
		i2 := (wResult >> 8) & 0xff
		wResult = i2 ^ wCRC16Table[wTableNo]
	}
	return wResult
}

func requestProcess(wg *sync.WaitGroup) {
	data, _ := dao.HttpGet()
	account := demo.ReviewData{}

	err := json.Unmarshal([]byte(data), &account)
	if err != nil {
		log.Fatalln(err)
	}

	/* vehicle */
	for index, a := range account.Data.List {
		fmt.Println(index)

		var info demo.AlgTaskInfoReqPkg

		info.EncodeKey = "0"
		info.InSystemTime = time.Now().Format("2006-01-02 15:04:05")

		info.PicProcessTime = "4"
		info.ResponseChan = "alg_response_queue"
		info.SessionID = ""
		info.TimeoutSecond = "2400"

		hiddenDangerId := a.HiddenDangerID
		println(" id = " + hiddenDangerId)

		for _, b := range a.PicList {
			var item demo.AlgTaskListReqAttr
			item.AlgIDs = append(item.AlgIDs, 7002)
			item.AlgTaskKeyCode = hiddenDangerId
			item.Other.InSystemTime = time.Now().Format("2006-01-02 15:04:05")
			item.Other.ReponseChan = "alg_response_queue"
			item.Other.TimeoutSecond = "2400"
			item.EncodeKey = strconv.Itoa(Crc16([]byte(item.AlgTaskKeyCode), len(item.AlgTaskKeyCode)))
			item.Photo.MainPhoto = b
			item.Photo.MainPhotoPath = b

			//res, err := dao.QueryDb("select zhaoPianXiaZaiDiZhi from photo_info where taskId = '5ffd0902a787864299050821'")
			res, err := dao.QueryDb("select zhaoPianXiaZaiDiZhi from photo_info where taskId = '" + hiddenDangerId + "'")
			if err != nil || len(res) == 0 {
				continue
			}

			for _, item1 := range res {
				item.Photo.SubPhoto = item1
				item.Photo.SubPhotoPath = item1
				info.AlgTasks = append(info.AlgTasks, &item)

				//println("---------------")
				//b, _ := json.Marshal(info)
				//var out bytes.Buffer
				//err = json.Indent(&out, b, "", "\t")
				//println(out.String())
				//println("----------------")

				time.Sleep(time.Second * time.Duration(2))
				_ = demo.SendPkgToRedis(info)
			}
		}
	}

	wg.Done()
}

func getRedisDataProcess(wg *sync.WaitGroup) {
	for {
		time.Sleep(time.Second * time.Duration(1))
		if v, ok := demo.GetPkgFormRedis(); ok == nil {

			res := demo.AlgTaskInfoRespPkg{}

			//item := bytes.TrimPrefix([]byte(v[1]), []byte("\xef\xbb\xbf"))
			println("---------------")
			println(v[1])
			println("+++++++++++++++")

			err := json.Unmarshal([]byte(v[1]), &res)
			if err != nil {
				log.Printf("error decoding sakura response: %v", err)
				if e, ok := err.(*json.SyntaxError); ok {
					log.Printf("syntax error at byte offset %d", e.Offset)
				}
				log.Printf("sakura response: %q", v[1])
			}

			println("000000000000   ===  " + res.AlgTaskKeyCode)
			var desc string
			var result string

			mainPhotoUrl := res.PhotoSavePath.MainPhoto[27 : len(res.PhotoSavePath.MainPhoto)]
			subPhotoUrl := res.PhotoSavePath.SubPhoto[27 : len(res.PhotoSavePath.SubPhoto)]

			if res.ReturnResultMemberList[0].Num7002[0].Value == "true" {
				result = "5"
				desc = "不通过，是同一张照片"
			} else {

				//https://lmd-internal.oss-cn-hangzhou.aliyuncs.com/xuhuipatrol/review/5ff268b1a787867c6095982e/1609722005309.jpg
				sql := `select shuoMing from reviewDataPhoto where zhaoPianXiaZaiDiZhi = '` + mainPhotoUrl + `'`
				println("sql is  == " + sql)
				res, err := dao.QueryDb(sql)
				if err != nil || len(res) == 0 {
					println("**************************** no result !")

					result = "5"
					desc = "没有查到对比照片结果"
				} else {
					if res[0] != "[通过]" {
						result = "5"
						desc = "不是同一张照片，但是复核照片不通过(" + res[0] + ")"
					} else {
						result = "1"
						desc = "不是同一张照片且复核照片结果通过"
					}
				}
			}

			/* 数据库插入 */
			sql := "INSERT INTO reviewDataComparePhoto (taskId, mainPicUrl, subPicUrl, jieGuo, shuoMing) VALUES('" + res.AlgTaskKeyCode + "','" +
				mainPhotoUrl + "','" + subPhotoUrl + "'," + result + ",'" + desc + "')"

			println("sql is " + sql)
			//_ = dao.DbInsert(sql)
		}
	}
	//wg.Done()
}

func client() {
	conn, err := net.Dial("tcp", "127.0.0.1:20001")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer conn.Close() // 关闭连接
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		_, err = conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

func main() {
	//obj := dao.Person1{}
	//
	//fmt.Print("%#v\n", obj)

	//var wg sync.WaitGroup
	//count := 0
	//wg.Add(2)
	//go requestProcess(&wg)
	//go getRedisDataProcess(&wg)
	//
	//wg.Wait()
	//fmt.Println("count 的值为：", count)

	client()
}
