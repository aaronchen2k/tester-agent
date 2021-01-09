package test

import (
	"encoding/json"
	_domain "github.com/aaronchen2k/tester/internal/pkg/domain"
	_fileUtils "github.com/aaronchen2k/tester/internal/pkg/libs/file"
	_httpUtils "github.com/aaronchen2k/tester/internal/pkg/libs/http"
	"log"
	"testing"
)

func TestUpdateResult(t *testing.T) {
	url := _httpUtils.GenUrl("http://localhost:8080", "build/uploadResult")
	log.Println(url)

	files := []string{"/Users/aaron/tester/testResult.zip"}
	extraParams := map[string]string{}

	result := _domain.RpcResult{}
	result.Code = 1
	result.Msg = "success"

	json, _ := json.Marshal(result)
	extraParams["json"] = string(json)

	_fileUtils.Upload(url, files, extraParams)
}
