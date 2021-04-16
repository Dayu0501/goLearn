package dao

import "testing"

func TestQueryDb(t *testing.T) {
	sql := `select shuoMing from reviewDataPhoto where zhaoPianXiaZaiDiZhi like '%5ffba873a78786429905031a/1610328128666.jpg%'`
	//sql := "select jieGuo from reviewDataPhoto where taskId = '5ff268b1a787867c6095982e'"
	res, err := QueryDb(sql)
	if err != nil || len(res) == 0 {
		println("no result !")
	}

	println(res[0])

	//for _, item1 := range res {
	//	println(item1)
	//	println("----------")
	//}
}
