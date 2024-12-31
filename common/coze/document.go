package coze

import (
	"Ai-HireSphere/utils"
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
