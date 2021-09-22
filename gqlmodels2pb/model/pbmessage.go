package model

import "strings"

// pb结构的信息
type PbMessageStructData struct {
	Name           string //pb结构的名字
	RowsNum        int    //字段的数量
	BeginPos       int    //开始的位置
	JsonBeginIndex int    //json的开始的位置
	EndPos         int    //结束的位置
}

// 从文件中解析pb结构的信息
func GetMessageInfos(content string) []PbMessageStructData {
	leftIdx := 0
	retMessageInfos := make([]PbMessageStructData, 0)
	for {
		messageIdx := strings.Index(content[leftIdx:], "message")
		//fmt.Println("messageIdx ", messageIdx)
		if messageIdx < 0 {
			break
		}
		messageIdx += leftIdx
		idx1 := strings.Index(content[messageIdx:], "{")
		//fmt.Println("idx1 ", idx1)
		if idx1 < 0 {
			break
		}
		idx1 += messageIdx
		//fmt.Println("content ", content[messageIdx:idx1])
		messageName := getMessageName(content[messageIdx:idx1])
		//fmt.Println("messageName ", messageName)
		idx2 := strings.Index(content[idx1:], "}")
		//fmt.Println("idx2 ", idx2)
		if idx2 < 0 {
			break
		}
		idx2 += idx1
		deprecatedIdx := strings.Index(content[idx1:], "deprecated")
		jsonIdx := strings.Index(content[idx1:], "_json")
		//fmt.Println("jsonIdx ", jsonIdx)
		if jsonIdx < 0 {
			break
		}
		jsonIdx += idx1
		jsonBeginIndex := strings.LastIndex(content[:jsonIdx], "\n")
		jsonBeginIndex += 1
		if jsonIdx > idx2 {
			leftIdx = idx2 + 2
			continue
		}
		rowsNum := strings.Count(content[idx1:idx2], "\n")
		//fmt.Println("rowsNum:", rowsNum, " data:", content[idx1:idx2])
		// json被禁止
		if deprecatedIdx > 0 && jsonIdx > deprecatedIdx+idx1 {
			leftIdx = idx2 + 2
			continue
		}
		//fmt.Println("messageIdx ", messageIdx)
		//fmt.Println("idx1 ", idx1)
		//fmt.Println("content ", content[messageIdx:idx1])
		//fmt.Println("messageName ", messageName)
		//fmt.Println("idx2 ", idx2)
		//fmt.Println("jsonIdx ", jsonIdx)

		retMessageInfos = append(retMessageInfos, PbMessageStructData{
			Name:           messageName,
			RowsNum:        rowsNum - 2,
			BeginPos:       idx1,
			JsonBeginIndex: jsonBeginIndex,
			EndPos:         idx2,
		})
		leftIdx = idx2 + 2
		//fmt.Println()
	}
	return retMessageInfos
}

func getMessageName(content string) string {
	s1 := strings.Replace(content, "message", "", 1)
	s2 := strings.Replace(s1, "{", "", 1)
	return strings.Trim(s2, " ")
}
