package testing

import (
	"encoding/json"
	"github.com/aaronchen2k/openstc-common/src/domain"
	"github.com/aaronchen2k/openstc-common/src/libs/file"
	"github.com/aaronchen2k/openstc-common/src/libs/http"
	"log"
	"testing"
)

func TestUpdateResult(t *testing.T) {
	url := httpUtils.GenUrl("http://localhost:8080", "build/uploadResult")
	log.Println(url)

	files := []string{"/Users/aaron/openstc/testResult.zip"}
	extraParams := map[string]string{}

	result := domain.RpcResult{}
	result.Code = 1
	result.Msg = "success"

	json, _ := json.Marshal(result)
	extraParams["json"] = string(json)

	fileUtils.Upload(url, files, extraParams)
}
