package dcdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dcdn"
)

func NewDCDNClient(regionId, accessKeyId, accessKeySecret string) *dcdn.Client {
	// 1. 初始化配置
	client, err := dcdn.NewClientWithAccessKey("cn-shanghai", accessKeyId, accessKeySecret)
	if err != nil {
		return nil
	}
	return client
}

func getDCDNDomainInfo(client *dcdn.Client, domainName string) (dcdn.DescribeDcdnDomainDetailResponse, error) {
	// 2. 初始化请求参数
	request := dcdn.CreateDescribeDcdnDomainDetailRequest()
	request.DomainName = domainName

	// 3. 发起请求并处理异常
	response, err := client.DescribeDcdnDomainDetail(request)
	if err != nil {
		return *response, err
	}
	return *response, nil
}

func listDCDNDomains(client *dcdn.Client) ([]*dcdn.DescribeDcdnUserDomainsResponse, error) {
	var pageStartNumber = 1
	var totalCount int64
	var pageSize = 20
	nextFlag := true

	response := make([]*dcdn.DescribeDcdnUserDomainsResponse, 0)

	request := dcdn.CreateDescribeDcdnUserDomainsRequest()
	request.Scheme = "https"
	request.PageSize = "100"
	request.DomainStatus = "online"

	for nextFlag {
		request.PageNumber = requests.NewInteger(pageStartNumber)
		res, err := client.DescribeDcdnUserDomains(request)
		if err != nil {
			return response, err
		}
		totalCount = res.TotalCount
		response = append(response, res)
		if pageStartNumber*pageSize >= int(totalCount) {
			nextFlag = false
		}
		pageStartNumber += 1
	}
	return response, nil
}

func getDCDNSSLCertificateList(client *dcdn.Client) (*dcdn.DescribeDcdnHttpsDomainListResponse, error) {
	//var pageStartNumber = 1
	//var totalCount int64
	//var pageSize = 20
	//nextFlag := true
	//response := make([]*dcdn.DescribeDcdnCertificateListResponse, 0)
	//dcdn.DescribeDcdnUsercer
	request := dcdn.CreateDescribeDcdnHttpsDomainListRequest()
	request.Scheme = "https"

	res, err := client.DescribeDcdnHttpsDomainList(request)
	if err != nil {
		return nil, err
	}
	return res, nil
}
