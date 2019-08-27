package httpdns

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

// GetResolveStatistics invokes the httpdns.GetResolveStatistics API synchronously
// api document: https://help.aliyun.com/api/httpdns/getresolvestatistics.html
func (client *Client) GetResolveStatistics(request *GetResolveStatisticsRequest) (response *GetResolveStatisticsResponse, err error) {
	response = CreateGetResolveStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// GetResolveStatisticsWithChan invokes the httpdns.GetResolveStatistics API asynchronously
// api document: https://help.aliyun.com/api/httpdns/getresolvestatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetResolveStatisticsWithChan(request *GetResolveStatisticsRequest) (<-chan *GetResolveStatisticsResponse, <-chan error) {
	responseChan := make(chan *GetResolveStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetResolveStatistics(request)
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

// GetResolveStatisticsWithCallback invokes the httpdns.GetResolveStatistics API asynchronously
// api document: https://help.aliyun.com/api/httpdns/getresolvestatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetResolveStatisticsWithCallback(request *GetResolveStatisticsRequest, callback func(response *GetResolveStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetResolveStatisticsResponse
		var err error
		defer close(result)
		response, err = client.GetResolveStatistics(request)
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

// GetResolveStatisticsRequest is the request struct for api GetResolveStatistics
type GetResolveStatisticsRequest struct {
	*requests.RpcRequest
	Granularity  string           `position:"Query" name:"Granularity"`
	ProtocolName string           `position:"Query" name:"ProtocolName"`
	DomainName   string           `position:"Query" name:"DomainName"`
	TimeSpan     requests.Integer `position:"Query" name:"TimeSpan"`
}

// GetResolveStatisticsResponse is the response struct for api GetResolveStatistics
type GetResolveStatisticsResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	DataPoints DataPoints `json:"DataPoints" xml:"DataPoints"`
}

// CreateGetResolveStatisticsRequest creates a request to invoke GetResolveStatistics API
func CreateGetResolveStatisticsRequest() (request *GetResolveStatisticsRequest) {
	request = &GetResolveStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Httpdns", "2016-02-01", "GetResolveStatistics", "httpdns", "openAPI")
	return
}

// CreateGetResolveStatisticsResponse creates a response to parse from GetResolveStatistics response
func CreateGetResolveStatisticsResponse() (response *GetResolveStatisticsResponse) {
	response = &GetResolveStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
