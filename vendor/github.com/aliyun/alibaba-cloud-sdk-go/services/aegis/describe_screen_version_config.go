package aegis

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeScreenVersionConfig invokes the aegis.DescribeScreenVersionConfig API synchronously
// api document: https://help.aliyun.com/api/aegis/describescreenversionconfig.html
func (client *Client) DescribeScreenVersionConfig(request *DescribeScreenVersionConfigRequest) (response *DescribeScreenVersionConfigResponse, err error) {
	response = CreateDescribeScreenVersionConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScreenVersionConfigWithChan invokes the aegis.DescribeScreenVersionConfig API asynchronously
// api document: https://help.aliyun.com/api/aegis/describescreenversionconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScreenVersionConfigWithChan(request *DescribeScreenVersionConfigRequest) (<-chan *DescribeScreenVersionConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeScreenVersionConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScreenVersionConfig(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeScreenVersionConfigWithCallback invokes the aegis.DescribeScreenVersionConfig API asynchronously
// api document: https://help.aliyun.com/api/aegis/describescreenversionconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScreenVersionConfigWithCallback(request *DescribeScreenVersionConfigRequest, callback func(response *DescribeScreenVersionConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScreenVersionConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeScreenVersionConfig(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeScreenVersionConfigRequest is the request struct for api DescribeScreenVersionConfig
type DescribeScreenVersionConfigRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
}

// DescribeScreenVersionConfigResponse is the response struct for api DescribeScreenVersionConfig
type DescribeScreenVersionConfigResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	AssetLevel     int    `json:"AssetLevel" xml:"AssetLevel"`
	AvdsFlag       int    `json:"AvdsFlag" xml:"AvdsFlag"`
	CreateTime     int    `json:"CreateTime" xml:"CreateTime"`
	Flag           int    `json:"Flag" xml:"Flag"`
	InstanceId     string `json:"InstanceId" xml:"InstanceId"`
	IsSasOpening   bool   `json:"IsSasOpening" xml:"IsSasOpening"`
	IsTrialVersion int    `json:"IsTrialVersion" xml:"IsTrialVersion"`
	LogCapacity    int    `json:"LogCapacity" xml:"LogCapacity"`
	LogTime        int    `json:"LogTime" xml:"LogTime"`
	ReleaseTime    int    `json:"ReleaseTime" xml:"ReleaseTime"`
	SasLog         int    `json:"SasLog" xml:"SasLog"`
	SasScreen      int    `json:"SasScreen" xml:"SasScreen"`
	Version        int    `json:"Version" xml:"Version"`
}

// CreateDescribeScreenVersionConfigRequest creates a request to invoke DescribeScreenVersionConfig API
func CreateDescribeScreenVersionConfigRequest() (request *DescribeScreenVersionConfigRequest) {
	request = &DescribeScreenVersionConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeScreenVersionConfig", "vipaegis", "openAPI")
	return
}

// CreateDescribeScreenVersionConfigResponse creates a response to parse from DescribeScreenVersionConfig response
func CreateDescribeScreenVersionConfigResponse() (response *DescribeScreenVersionConfigResponse) {
	response = &DescribeScreenVersionConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
