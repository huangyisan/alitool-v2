package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
	"github.com/spf13/cobra"
)

func NewDomainClient(regionId, accessKeyId, accessKeySecret string) *domain.Client {
	// 1. 初始化配置
	client, err := domain.NewClientWithAccessKey("cn-shanghai", accessKeyId, accessKeySecret)
	if err != nil {
		return nil
	}
	return client
}

// func ListDomainInfoByAccountName(client *domain.Client, accountName string) ()
func getDomainInfo(client *domain.Client, domainName string) (*domain.QueryDomainByDomainNameResponse, error) {
	// 2. 初始化请求参数
	request := domain.CreateQueryDomainByDomainNameRequest()
	request.Scheme = "https"
	request.DomainName = domainName
	// 3. 发起请求
	response, err := client.QueryDomainByDomainName(request)
	if err != nil {
		cobra.CheckErr(err)
	}
	return response, nil
}

func getDomainList(client *domain.Client) ([]domain.Domain, error) {
	response := make([]domain.Domain, 0)
	var pageStartNumber = 1
	var nextFlag = true
	var pageSize = 20
	request := domain.CreateQueryDomainListRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(pageSize)
	// 3. 发起请求
	for nextFlag {
		request.PageNum = requests.NewInteger(pageStartNumber)
		res, err := client.QueryDomainList(request)
		if err != nil {
			return nil, err
		}
		response = append(response, res.Data.Domain...)
		if res.NextPage == false {
			nextFlag = false
		}
		pageStartNumber += 1
	}
	return response, nil
}
