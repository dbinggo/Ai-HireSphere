package coze

import (
	"Ai-HireSphere/common/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const createDocUrl = "https://api.coze.cn/open_api/knowledge/document/create"

type CozeDocApi struct {
	Token     string `json:"access_token"`
	DataSetID string `json:"dataset_id"`
}

type docHeader struct {
	Authorization string `json:"Authorization"`
	ContentType   string `json:"Content-Type"`
	AgwJsConv     string `json:"Agw-Js-Conv"`
}

type createDocBody struct {
	DatasetID     string                 `json:"dataset_id"`
	DocumentBases []createDocBases       `json:"document_bases"`
	ChunkStrategy createDocChunkStrategy `json:"chunk_strategy"`
}

type createDocBases struct {
	Name       string              `json:"name"`
	SourceInfo createDocSourceInfo `json:"source_info"`
}

type createDocSourceInfo struct {
	DocumentSource int    `json:"document_source"`
	FileType       string `json:"file_type"`
	FileBase64     string `json:"file_base64"`
}

type createDocChunkStrategy struct {
	ChunkType int `json:"chunk_type"`
}

func NewCozeDocApi(token, datasetID string) *CozeDocApi {
	return &CozeDocApi{
		Token:     token,
		DataSetID: datasetID,
	}
}

func (api *CozeDocApi) CreateDoc(files []utils.FileBase) error {
	var body createDocBody
	var header docHeader

	header = getDocApiHeader(api.Token)
	body.DatasetID = api.DataSetID
	// 这里要写死0
	body.ChunkStrategy.ChunkType = 0
	body.DocumentBases = NewDocBases(files)

	headerByte, _ := json.Marshal(header)
	headerMap := make(map[string]string)
	err := json.Unmarshal(headerByte, &headerMap)
	if err != nil {
		return err
	}

	bodyByte, _ := json.Marshal(body)

	request, err := http.NewRequest("POST", createDocUrl, strings.NewReader(string(bodyByte)))
	if err != nil {
		return err
	}

	for key, value := range headerMap {
		request.Header.Set(key, value)
	}

	resp, err := (&http.Client{}).Do(request)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	respByte, _ := io.ReadAll(resp.Body)
	fmt.Println(string(respByte))

	return nil
}

func getDocApiHeader(token string) docHeader {
	var header docHeader
	header.Authorization = "Bearer " + token
	header.ContentType = "application/json"
	header.AgwJsConv = "str"

	return header
}

func NewDocBases(bases []utils.FileBase) []createDocBases {
	var docBases []createDocBases
	for i := 0; i < len(bases); i++ {
		docBase := createDocBases{
			Name: bases[i].Filename,
			SourceInfo: createDocSourceInfo{
				// 这里要写死0
				DocumentSource: 0,
				FileType:       bases[i].Ext,
				FileBase64:     base64UrlEncode(bases[i].FileByte),
			},
		}
		docBases = append(docBases, docBase)
	}
	return docBases
}

//{"code":0,
//	"document_infos":[{
//	"char_count":0,
//	"chunk_strategy":{},
//	"create_time":0,
//	"document_id":"7454933945252216868",
//	"format_type":0,
//	"hit_count":0,
//	"name":"test01",
//	"size":0,
//	"slice_count":0,
//	"source_type":0,
//	"status":0,
//	"type":"",
//	"update_interval":0,
//	"update_time":0,
//	"update_type":0
//	}],
//	"msg":""
//}
