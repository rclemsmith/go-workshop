package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"github.com/buger/jsonparser"
)

type ApiResponse struct {
	kind       string
	apiVersion string
	metadata   map[string]string
	items      []map[string]map[string]interface{}
}

func getDeploymentInfo(wg *sync.WaitGroup,namespace string,deployment string){
	wg.Add(1)
	client := http.Client{}
	url := fmt.Sprintf("http://localhost:8080/apis/apps/v1/namespaces/%s/deployments/%s",namespace,deployment)
	// fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6Imd1Q3FoRXZvUEhRcHp6ZTRzU2JlLWF6Y085Mzd3Zl9oRmlPTEVKYU5KQ3MifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJjbGVhbmVyIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6ImNsZWFuZXIiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiY2xlYW5lciIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6ImM0N2I2YjcwLTJjYWYtNGNhMS1iMGY4LTRlNTQ0NzJkN2UxZCIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpjbGVhbmVyOmNsZWFuZXIifQ.ESeIMkUMq_hokC41iNbW0xBw1m_pU-5fAkSAvsNUpk35W_sKxa9oOgjMrLpOGZ9rOWvgxBPy0swyeWwGq7qkhuOkLBfznpsFeAb8as-oCP0q_FxPC_g3N1TIkxX56nObBAJZgh75YTIkhP_uMiTbi69jfOwizA1wnYs8-QWYY07nqD7dxsqkQ-MubdPbypQ-9BCydP-Z5wRepq3PS-fOMQvVXZ3XarWxAYBeK_UWZlisRcJksz2GXWtEU8saI-Noa0YaLh_yK0Xx3MoyqRZooxyh_YiG72HoPkP09XLEEt1kHhDm7y8lXKE-OMZ2o7NEVDnOvP4x-P4ZMb_-VbkH3Q")
	res,err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Body", jsonparser.GET(body,"kind"))
	wg.Done()
}

func deleteDeployments(wg *sync.WaitGroup,channel <-chan map[string]string) {
	for deploys := range channel {
		fmt.Println(deploys)
		go getDeploymentInfo(wg,deploys["namespace"],deploys["deploymentName"])
	}
	wg.Done()
}


func main() {

	channel := make(chan map[string]string)
	wg := &sync.WaitGroup{}

	go deleteDeployments(wg,channel)
	go deleteDeployments(wg,channel)
	go deleteDeployments(wg,channel)
	go deleteDeployments(wg,channel)

	client := http.Client{}
	// url := "https://kubernetes.default.svc"
	url := "http://localhost:8080"
	req, _ := http.NewRequest("GET", url+"/api/v1/namespaces", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6Imd1Q3FoRXZvUEhRcHp6ZTRzU2JlLWF6Y085Mzd3Zl9oRmlPTEVKYU5KQ3MifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJjbGVhbmVyIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6ImNsZWFuZXIiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiY2xlYW5lciIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6ImM0N2I2YjcwLTJjYWYtNGNhMS1iMGY4LTRlNTQ0NzJkN2UxZCIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpjbGVhbmVyOmNsZWFuZXIifQ.ESeIMkUMq_hokC41iNbW0xBw1m_pU-5fAkSAvsNUpk35W_sKxa9oOgjMrLpOGZ9rOWvgxBPy0swyeWwGq7qkhuOkLBfznpsFeAb8as-oCP0q_FxPC_g3N1TIkxX56nObBAJZgh75YTIkhP_uMiTbi69jfOwizA1wnYs8-QWYY07nqD7dxsqkQ-MubdPbypQ-9BCydP-Z5wRepq3PS-fOMQvVXZ3XarWxAYBeK_UWZlisRcJksz2GXWtEU8saI-Noa0YaLh_yK0Xx3MoyqRZooxyh_YiG72HoPkP09XLEEt1kHhDm7y8lXKE-OMZ2o7NEVDnOvP4x-P4ZMb_-VbkH3Q")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	// fmt.Printf("Body : %s", body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var response map[string]interface{}
	json.Unmarshal([]byte(body), &response)

	items := response["items"].([]interface{})

	var namespaces []string
	for _, slice := range items {
		metadata := slice.(map[string]interface{})
		metadata1 := metadata["metadata"].(map[string]interface{})
		ns := metadata1["name"].(string)
		if strings.Contains(ns, "dev-") {
			namespaces = append(namespaces, ns)
		} else if strings.Contains(ns, "xblox-dev-") {
			namespaces = append(namespaces, ns)
		}

	}

	// fmt.Println(namespaces)

	for _, namespace := range namespaces {
		go getDeploymentsInANamespace(wg,channel,namespace,url)
	}
	// close(channel)

	wg.Wait()

}


func getDeploymentsInANamespace(wg *sync.WaitGroup,channel chan<- map[string]string,namespace string,url string) {
	wg.Add(1)
	client := http.Client{}
	req, err := http.NewRequest("GET", url+"/apis/apps/v1/namespaces/"+namespace+"/deployments", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6Imd1Q3FoRXZvUEhRcHp6ZTRzU2JlLWF6Y085Mzd3Zl9oRmlPTEVKYU5KQ3MifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJjbGVhbmVyIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6ImNsZWFuZXIiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiY2xlYW5lciIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6ImM0N2I2YjcwLTJjYWYtNGNhMS1iMGY4LTRlNTQ0NzJkN2UxZCIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpjbGVhbmVyOmNsZWFuZXIifQ.ESeIMkUMq_hokC41iNbW0xBw1m_pU-5fAkSAvsNUpk35W_sKxa9oOgjMrLpOGZ9rOWvgxBPy0swyeWwGq7qkhuOkLBfznpsFeAb8as-oCP0q_FxPC_g3N1TIkxX56nObBAJZgh75YTIkhP_uMiTbi69jfOwizA1wnYs8-QWYY07nqD7dxsqkQ-MubdPbypQ-9BCydP-Z5wRepq3PS-fOMQvVXZ3XarWxAYBeK_UWZlisRcJksz2GXWtEU8saI-Noa0YaLh_yK0Xx3MoyqRZooxyh_YiG72HoPkP09XLEEt1kHhDm7y8lXKE-OMZ2o7NEVDnOvP4x-P4ZMb_-VbkH3Q")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	var response map[string]interface{}
	json.Unmarshal([]byte(body), &response)
	deploys := response["items"].([]interface{})
	for _, deploy := range deploys {
		deploy1 := deploy.(map[string]interface{})
		metadata := deploy1["metadata"].(map[string]interface{})
		data := map[string]string{"namespace": metadata["namespace"].(string), "deploymentName": metadata["name"].(string)}
		channel <- data
	}
	wg.Done()
}