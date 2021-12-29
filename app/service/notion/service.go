package notion

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/cangyan/notion-youtube-feed/app/service/notion/types/filter_object"
)

type Service interface {
	FindNotionPageExistedById(ids []string) ([]string, error)
	CreatePage(body string) error
}
type service struct {
	NotionToken      string
	NotionDatabaseId string
}

func NewService(token, databaseId string) Service {
	return &service{
		NotionToken:      token,
		NotionDatabaseId: databaseId,
	}
}

func (s *service) FindNotionPageExistedById(ids []string) ([]string, error) {
	ret := make([]string, 0)
	url := fmt.Sprintf("https://api.notion.com/v1/databases/%v/query", s.NotionDatabaseId)
	method := "POST"

	payload := strings.NewReader(`{"filter":` + filter_object.GenerateTextOrFilterObject("ID", ids) + `}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ret, err
	}
	req.Header.Add("Notion-Version", "2021-08-16")
	req.Header.Add("Authorization", "Bearer "+s.NotionToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ret, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ret, err
	}

	// fmt.Println(string(body))
	resp := NotionDatabaseQueryResp{}

	_ = json.Unmarshal(body, &resp)

	return resp.GetArticleIds(), nil
}

func (s *service) CreatePage(body string) error {
	url := "https://api.notion.com/v1/pages"
	method := "POST"

	payload := strings.NewReader(body)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Notion-Version", "2021-08-16")
	req.Header.Add("Authorization", "Bearer "+s.NotionToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// fmt.Println(string(b))

	return nil
}
