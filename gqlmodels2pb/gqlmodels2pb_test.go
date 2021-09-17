package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	ProjectPath = "C:/Users/Administrator/GoProjects/src/solarland/backendv2/"
)

var transMessage2Object = map[string]string{
	"ListPlayingUsers":                   "PlayingGameUserConnection",
	"MarketListCount":                    "Count",
	"AssetInfo":                          "Asset",
	"CreateAssetWorld":                   "UpdateAssetWorldPayload",
	"UserAssetCount":                     "Count",
	"SwapBanner":                         "",
	"AssetRecommendRule":                 "RecommendRule",
	"BmsMarketAssetList":                 "MarketAssetListPayload",
	"GetAllDevices":                      "DevicePayload",
	"EditRecommendGame":                  "RecommendGame",
	"GetdevActivityParticipants":         "DevActivityParticipantsResult",
	"GetGameAuditDetail":                 "DisplayGameAuditInfo",
	"GenerateVivoxSquadChannelJoinToken": "VivoxSquadChannelJoinToken",
	"GenerateAgoraJoinToken":             "AgoraJoinToken",
	"GenerateVivoxGameChannelJoinToken":  "VivoxGameChannelJoinToken",
	"GenerateVivoxLoginToken":            "VivoxLoginToken",
	"DeveloperGameDataView":              "DeveloperGameDisplayItemV2",
	"QueryUserInfo":                      "DisplayUserInfo",
	"HottestGameList":                    "OnlineGameListPayload",
	//"SwapBanner": "",
}

var ssMap = map[string]string{
	"DeveloperGamesDisplayInfo":     "",
	"AddIrisStrategy":               "",
	"HTTPCallback":                  "",
	"GetIMFriendFlagsMap":           "",
	"QueryUserLimitAvatarList":      "",
	"QueryUserAvatar":               "",
	"SetUserAvatar":                 "",
	"LoginByLilithSDK":              "",
	"LoginByGooglePlay":             "",
	"RegisterByBms":                 "",
	"Teleport2Homestead":            "",
	"TeleportPlayersToNewServer":    "",
	"TeleportToServer":              "",
	"TeleportPlayers":               "",
	"SyncAssetStoreInfosToAssetsvr": "",
}

//	//.BatchGetUserProfile:"UserProfile" map
var transMessageListObject = map[string]string{
	"BannerList":                 "Banner",
	"AuditAssetList":             "Asset",
	"UserMarketAssetList":        "Asset",
	"MarketAssetRecommendList":   "Asset",
	"MarketSearchAssetList":      "Asset",
	"AssetClassList":             "AssetClass",
	"AssetCategoryList":          "AssetCategory",
	"AssetLevelList":             "AssetLevel",
	"AssetTypeList":              "AssetType",
	"UserImportedList":           "Asset",
	"UserHistoryAssetList":       "Asset",
	"UserAssetStorageFolderList": "AssetMyStorage",
	"UserAssetStorageList":       "AssetMyStorage",
	"UserAssetList":              "Asset",
	"BmsBannerList":              "Banner",
	"BatchCreateHubAccount":      "CreateHubUserResult",
	"HubUserList":                "DisplayHubUser",
	"GameVersionList":            "GameVersionSummary",
	"GetPublishedGameList":       "DisplayGameAuditSummary",
	"GetGameListByAuditStatus":   "DisplayGameAuditSummary",
	"GetGameAuditDetail":         "DisplayGameAuditInfo",
	"SearchPublishGameList":      "Game",
	"UserGameList":               "Game",
	"UserGame":                   "Game",
	"GameList":                   "Game",
	"GameByIDs":                  "Game",
	"UserPlayedGameList":         "PlayedGame",
	"UserInfos":                  "DisplayUserInfo",
	"UserListInGame":             "DisplayUserInfo",
	"SearchUserList":             "DisplayUserInfo",
	"QueryIMFriends":             "DisplayUserInfo",
	"QueryFriends":               "DisplayUserInfo",
}

var notSameObjectMap = map[string]struct{}{
	"AssetFilter":                   {},
	"BannerLink":                    {},
	"HomesteadItem":                 {},
	"Status":                        {},
	"Homestead":                     {},
	"AssetWorldListInput":           {},
	"GenGameTinyCodeReply":          {},
	"User":                          {},
	"DisplayDevActivitySummary":     {},
	"Banner":                        {},
	"VersionNoticePayload":          {},
	"VersionNoticeInput":            {},
	"DeletePostInput":               {},
	"DisplayDevActivityDetail":      {},
	"Post":                          {},
	"PublishUndoGameInput":          {},
	"Asset":                         {},
	"Game":                          {},
	"RoomListInput":                 {},
	"HubBanner":                     {},
	"PlayerHubGameCommentListInput": {},
	"GenGameTinyCodeInput":          {},
	"HubBannerLink":                 {},
	"Masterpiece":                   {},
}

type MessageInfo struct {
	Name           string
	RowsNum        int
	BeginPos       int
	JsonBeginIndex int
	EndPos         int
}

func GetMessageInfos(content string) []MessageInfo {
	leftIdx := 0
	retMessageInfos := make([]MessageInfo, 0)
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

		retMessageInfos = append(retMessageInfos, MessageInfo{
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

func TestFiles(t *testing.T) {
	allObjectMap, allAliasMap := tRender()
	prefix := fmt.Sprintf("%v%v", ProjectPath, "deploy/app/base/proto/avatar")
	files, err := ioutil.ReadDir(prefix)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	for _, f := range files {
		backPrefixPath := fmt.Sprintf("%v/%v/", prefix, f.Name())
		protoFileName := "grpc.proto"
		protoFilePath := fmt.Sprintf("%v%v", backPrefixPath, protoFileName)
		if !exists(protoFilePath) {
			continue
		}
		copyBackup(backPrefixPath, protoFileName)
		backupFile := fmt.Sprintf("%v%v.back.txt", backPrefixPath, protoFileName)
		bs, err := ioutil.ReadFile(backupFile)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		messageInfos := GetMessageInfos(string(bs))
		RewriteData2ProtoFile(string(bs), protoFilePath, messageInfos, allObjectMap, allAliasMap)
	}
}

func GetMessageGenString(messageName string, allObjectMap AllObjectMap, aliasMap VariableAliasMap, extraRows int) (string, map[string]struct{}) {
	flag := ""
	origMessageName := messageName
	dependObjectMap := make(map[string]struct{})
	if strings.Contains(messageName, "Response") {
		flag = "Response"
		messageName = strings.Replace(messageName, "Response", "", 1)
	} else if strings.Contains(messageName, "Request") {
		flag = "Request"
		messageName = strings.Replace(messageName, "Request", "", 1)
	}
	if flag == "" {
		panic("no Response message " + messageName)
	}
	if val, ok := transMessage2Object[messageName]; ok {
		messageName = val
	}
	// TODO list 类型 待检测
	if val, ok := transMessageListObject[messageName]; ok {
		return fmt.Sprintf("  string deprecated_json = %v;\n  repeated %v list = %v ;\n", 1+extraRows, val, 2+extraRows), map[string]struct{}{val: {}}
	}
	// TODO 检测用的名字不对
	if _, ok := notSameObjectMap[messageName]; ok {
		//fmt.Println("in notSameObjectMap ", messageName)
	}
	for _, objectMap := range allObjectMap {
		var val *MyObject
		if flag == "Response" {
			val = getObjectFromMapRes(messageName, objectMap)
		} else if flag == "Request" {
			val = getObjectFromMapIntp(messageName, objectMap)
		}
		if val != nil {
			ret := fmt.Sprintf("  string deprecated_json = %v;\n", 1+extraRows)
			for i, fil := range val.FieldArr {
				ret += fmt.Sprintf("  %v %v = %v;\n", transType(aliasMap, fil), ns(fil.Name), i+2+extraRows)
				if !fil.IsBaseType {
					dependObjectMap[fil.Type] = struct{}{}
				}
			}
			return ret, dependObjectMap
		}
	}
	fmt.Println("message cant find target object messageName ", messageName, " origMessageName ", origMessageName)
	return fmt.Sprintf("  string deprecated_json = %v;\n", extraRows+1), dependObjectMap
}

func getObjectFromMapRes(messageName string, objectMap ObjectMap) *MyObject {
	if val, ok := objectMap[messageName]; ok {
		return val
	}
	if val, ok := objectMap[messageName+"Payload"]; ok {
		return val
	}
	if val, ok := objectMap[messageName+"Response"]; ok {
		return val
	}
	if val, ok := objectMap[messageName+"Result"]; ok {
		return val
	}
	if val, ok := objectMap[messageName+"Reply"]; ok {
		return val
	}
	//if val, ok := objectMap["rpcmsg."+messageName]; ok {
	//	return val
	//}
	return nil
}

func getObjectFromMapIntp(messageName string, objectMap ObjectMap) *MyObject {
	if val, ok := objectMap[messageName]; ok {
		return val
	}
	if val, ok := objectMap[messageName+"Input"]; ok {
		return val
	}
	if val, ok := objectMap[messageName+"Request"]; ok {
		return val
	}
	//if val, ok := objectMap["rpcmsg."+messageName]; ok {
	//	return val
	//}
	return nil
}

func RewriteData2ProtoFile(content string, targetFile string, messageInfos []MessageInfo, allObjectMap AllObjectMap, aliasMap VariableAliasMap) {
	newString := content
	lth := len(messageInfos)
	allDependObjectMap := make(map[string]struct{})
	for i := lth - 1; i >= 0; i-- {
		jsonBeginIdx := messageInfos[i].JsonBeginIndex
		jsonEndIdx := messageInfos[i].EndPos
		genStr, dependObjectMap := GetMessageGenString(messageInfos[i].Name, allObjectMap, aliasMap, messageInfos[i].RowsNum)
		if genStr == "+" {
			//fmt.Println("in targetFile ", targetFile)
		}
		newString = newString[:jsonBeginIdx] + genStr + newString[jsonEndIdx:]
		for k := range dependObjectMap {
			allDependObjectMap[k] = struct{}{}
		}
	}
	dependStr := ""
	for _, objsMap := range allObjectMap {
		for _, v := range objsMap {
			if _, ok := allDependObjectMap[v.Name]; ok {
				dependStr += getMessGenStr(v, aliasMap)
			}
		}
	}
	newString = newString + dependStr
	err := ioutil.WriteFile(targetFile, []byte(newString), 0600)
	if err != nil {
		panic(err.Error())
	}
}

func copyBackup(prefixPath string, fileName string) {
	backupFileName := prefixPath + fileName + ".back.txt"
	if !exists(backupFileName) || AlwaysCopy {
		bs, err := ioutil.ReadFile(prefixPath + fileName)
		if err != nil {
			panic(err.Error())
		}
		err = ioutil.WriteFile(backupFileName, bs, 0600)
		if err != nil {
			panic(err.Error())
		}
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

type ConfPath struct {
	Name       string //业务的名字
	Prefix     string
	GoFileName string
}

// 配置
const (
	ShowAst       = false
	ShowRender    = false
	RenderObjects = true
	AlwaysCopy    = false
)

func tRender() (AllObjectMap, VariableAliasMap) {
	prefixPaths := []ConfPath{
		{

			Name:   "gate",
			Prefix: fmt.Sprintf("%v%v", ProjectPath, "cluster/gate/gate/gen/"),
		},
		{
			Name:   "asset",
			Prefix: fmt.Sprintf("%v%v", ProjectPath, "cluster/asset/graphql/gen/"),
		},
		{
			Name:   "bms",
			Prefix: fmt.Sprintf("%v%v", ProjectPath, "cluster/gate/bms/gen/"),
		},
		{
			Name:       "ex",
			Prefix:     "./gen/",
			GoFileName: "example.go",
		},
	}
	allObjectMap := make(AllObjectMap)
	allAliasMap := make(VariableAliasMap)
	for _, v := range prefixPaths {
		var tempAliasMap VariableAliasMap
		allObjectMap[v.Name], tempAliasMap = RenderFile(v)
		for k, v := range tempAliasMap {
			allAliasMap[k] = v
		}
	}
	//notSameObjectMap := printRepetitionObjects(allObjectMap)
	//for k := range notSameObjectMap {
	//	fmt.Println("不相同的结构 ", k)
	//}
	return allObjectMap, allAliasMap
}

func TestRender(t *testing.T) {
	//t.Skip("用于在本地生成pb文件")
	_, _ = tRender()
}

func printRepetitionObjects(allObjectMap AllObjectMap) map[string]struct{} {
	repetitionObjectsMap := make(map[string]map[string]*MyObject)
	for busName, objectMap := range allObjectMap {
		for k, v := range objectMap {
			if _, ok := repetitionObjectsMap[k]; !ok {
				repetitionObjectsMap[k] = make(map[string]*MyObject)
			}
			repetitionObjectsMap[k][busName] = v
		}
	}
	notSameObjectMap := make(map[string]struct{})
	for k, repObjectMap := range repetitionObjectsMap {
		if len(repObjectMap) > 1 {
			for busName := range repObjectMap {
				fmt.Println("存在重复的结构:業務", busName, k)
			}
			retObjName := comprOjects(repObjectMap)
			if retObjName != "" {
				notSameObjectMap[retObjName] = struct{}{}
			}
			fmt.Println()
		}
	}
	return notSameObjectMap
}

func comprOjects(objectMap map[string]*MyObject) string {
	a := objectMap["bms"]
	b := objectMap["gate"]
	if a == nil || b == nil {
		return ""
	}
	same := true
	for k, v := range a.FieldMap {
		if _, ok := b.FieldMap[k]; !ok {
			fmt.Printf("gate 少的字段 %v 具体:%+v\n", k, v)
			same = false
		}
	}
	for k, v := range b.FieldMap {
		if _, ok := a.FieldMap[k]; !ok {
			fmt.Printf("bms 少的字段 %v 具体:%+v\n", k, v)
			same = false
		}
	}
	if !same {
		return a.Name
	}
	return ""
}

type ObjectMap map[string]*MyObject
type VariableAliasMap map[string]string
type SpeStructMap map[string]struct{}
type AllObjectMap map[string]ObjectMap // asset/bms/gate

func RenderFile(conf ConfPath) (ObjectMap, VariableAliasMap) {
	subfix := "models_gen.go"
	if conf.GoFileName != "" {
		subfix = conf.GoFileName
	}
	abspath, err := filepath.Abs(conf.Prefix + subfix)
	if err != nil {
		panic(abspath)
	}
	// 创建用于解析源文件的对象
	fset := token.NewFileSet()
	// 解析源文件，返回ast.File原始文档类型的结构体。
	f, err := parser.ParseFile(fset, abspath, nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}
	if ShowAst {
		ast.Print(fset, f)
	}
	var objectMap = make(ObjectMap)
	var aliasMap = make(VariableAliasMap)
	speStructMap := make(SpeStructMap)
	VisitStruct(aliasMap, objectMap, speStructMap, f)

	if ShowRender {
		fmt.Println(renderObjects(speStructMap, aliasMap, objectMap))
	}
	if RenderObjects {
		content := []byte(renderObjects(speStructMap, aliasMap, objectMap))
		err = ioutil.WriteFile(conf.Prefix+"models.proto", content, 0600)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(fmt.Sprintf("./tosearch/%v.models.proto", conf.Name), content, 0600)
		if err != nil {
			panic(err)
		}
	}
	return objectMap, aliasMap
}

func renderObjects(speStructMap SpeStructMap, aliasMap VariableAliasMap, objectMap ObjectMap) string {
	retStr := "//<框出特殊的结构 待处理>\n"
	for k := range speStructMap {
		retStr += fmt.Sprintf("//%v\n", k)
	}
	retStr += "//<前面的是特殊的结构>\n\n\n\n\n"
	for _, v := range objectMap {
		retStr += getMessGenStr(v, aliasMap)
	}
	return retStr
}

func getMessGenStr(v *MyObject, aliasMap VariableAliasMap) string {
	retStr := ""
	objStr := "message " + v.Name + " {\n"
	for i, fil := range v.FieldArr {
		objStr += fmt.Sprintf("  %v %v = %v;\n", transType(aliasMap, fil), ns(fil.Name), i+2)
	}
	retStr += objStr + "}\n\n"
	return retStr
}

//func RenderObject(v MyObject) string{
//	objStr := "message " + v.Name + " {\n"
//	for i, fil := range v.FieldArr {
//		objStr += fmt.Sprintf("  %v %v = %v;\n", transType(aliasMap, fil), ns(fil.Name), i+2)
//	}
//	retStr += objStr + "}\n\n"
//}

func ns(s string) string {
	lth := len(s)
	i := lth - 1
	t := false
	for ; i > 1; i-- {
		c := s[i]
		if !t && c >= 'A' && c <= 'Z' {
			t = true
		}
		if t && 'a' <= c && c <= 'z' {
			s = s[:i+1] + "_" + s[i+1:]
			t = false
		}
	}
	ind := strings.Index(s, "ID")
	lth = len(s)
	if ind > 0 && ind+2 < lth && s[ind+2] >= 'A' && s[ind+2] <= 'Z' {
		s = s[:ind+2] + "_" + s[ind+2:]
	}
	return strings.ToLower(s)
}

func transType(aliasMap VariableAliasMap, field *MyField) string {
	ret := ""
	if field.IsMap {
		return fmt.Sprintf("map<%v,%v>", field.MKeyType, transType(aliasMap, &MyField{
			Name:       field.Name,
			Type:       field.Type,
			IsArr:      field.IsArr,
			IsSelector: field.IsSelector,
			IsStar:     field.IsStar,
		}))
	}
	if field.IsArr {
		ret += "repeated "
	}
	filType := toBaseType(aliasMap, field.Type)
	star := ""
	if field.IsStar {
		star = "*"
	}
	switch star + filType {
	case "*int":
		ret += "google.protobuf.Int64Value"
	case "*int64":
		ret += "google.protobuf.Int64Value"
	case "*string":
		ret += "google.protobuf.StringValue"
	case "*bool":
		ret += "google.protobuf.BoolValue"
	default:
		ret += filType
	}
	return ret
}

func toBaseType(aliasMap VariableAliasMap, filType string) string {
	val := aliasMap[filType]
	if val != "" {
		filType = val
	}
	switch filType {
	case "int":
		return "int64"
	case "UploadFile":
		return "string"
	}
	return filType
}

func VisitStruct(aliasMap VariableAliasMap, objectMap ObjectMap, speStructMap SpeStructMap, root ast.Node) {
	ast.Inspect(root, func(n ast.Node) bool {
		if file, ok := n.(*ast.File); ok {
			for _, val := range file.Decls {
				VisitStruct(aliasMap, objectMap, speStructMap, val)
			}
		}
		if vnode, ok := n.(*ast.GenDecl); ok {
			for _, val := range vnode.Specs {
				VisitStruct(aliasMap, objectMap, speStructMap, val)
			}
		}
		if vnode, ok := n.(*ast.TypeSpec); ok {
			if _, ok2 := vnode.Type.(*ast.StructType); ok2 {
				solveStruct(aliasMap, objectMap, speStructMap, vnode)
			}
			if vtai, ok2 := vnode.Type.(*ast.Ident); ok2 {
				aliasMap[vnode.Name.Name] = vtai.Name
			}
		}
		return true
	})
}

func solveStruct(aliasMap VariableAliasMap, objectMap ObjectMap, speStructMap SpeStructMap, node *ast.TypeSpec) {
	structName := node.Name.Name
	sType, ok := node.Type.(*ast.StructType)
	if !ok {
		panic("node.Type.(*ast.StructType)")
	}
	fieldList := sType.Fields.List
	retFieldList := make([]*MyField, len(fieldList))
	retFieldMap := make(map[string]*MyField, len(fieldList))
	for i, field := range fieldList {
		fieldName := field.Names[0].Name
		myField := &MyField{
			IsBaseType: false,
		}
		solveField(aliasMap, myField, speStructMap, field)

		retFieldList[i] = myField
		retFieldMap[fieldName] = myField
	}
	objectMap[structName] = &MyObject{
		Name:     structName,
		FieldArr: retFieldList,
		FieldMap: retFieldMap,
	}
}

func solveField(aliasMap VariableAliasMap, myField *MyField, speStructMap SpeStructMap, node ast.Node) interface{} {
	switch node.(type) {
	// 单值类型
	case *ast.Field:
		field := node.(*ast.Field)
		myField.Name = field.Names[0].Name
		return solveField(aliasMap, myField, speStructMap, field.Type)
		// 数组类型
	case *ast.ArrayType:
		myField.IsArr = true
		arrType := node.(*ast.ArrayType)
		return solveField(aliasMap, myField, speStructMap, arrType.Elt)
		// 指针类型
	case *ast.StarExpr:
		myField.IsStar = true
		nextNode := node.(*ast.StarExpr)
		return solveField(aliasMap, myField, speStructMap, nextNode.X)
		// 值类型
	case *ast.Ident:
		nextNode := node.(*ast.Ident)
		myField.Type = nextNode.Name
		myField.IsBaseType = isBaseType(aliasMap, myField.Type)
		return nextNode.Name
		// 表达式类型
	case *ast.SelectorExpr:
		myField.IsSelector = true
		nextNode := node.(*ast.SelectorExpr)
		myField.Type = nextNode.X.(*ast.Ident).Name + "." + nextNode.Sel.Name
		speStructMap[myField.Type] = struct{}{}
		return myField.Type
		// map类型
	case *ast.MapType:
		myField.IsMap = true
		nextNode := node.(*ast.MapType)
		myField.MKeyType = solveField(aliasMap, myField, speStructMap, nextNode.Key).(string)
		myField.MValueType = solveField(aliasMap, myField, speStructMap, nextNode.Value).(string)
		// 接口类型
	case *ast.InterfaceType:
		myField.Type = "interface{}"
		return "interface{}"
	default:
		panic("unknown type")
	}
	return ""
}

func isBaseType(aliasMap VariableAliasMap, s string) bool {
	s = toBaseType(aliasMap, s)
	switch s {
	case "int", "string", "int32", "int64", "bool":
		return true
	case "*int", "*string", "*int32", "*int64", "*bool":
		return true
	}
	return false
}

type MyObject struct {
	Name     string
	FieldArr []*MyField
	FieldMap map[string]*MyField
}

type MyField struct {
	Name       string
	Type       string
	IsBaseType bool
	IsArr      bool
	IsSelector bool
	IsStar     bool
	IsMap      bool
	MKeyType   string
	MValueType string
}
