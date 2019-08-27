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

package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// QueryTopicRouteTable invokes the iot.QueryTopicRouteTable API synchronously
// api document: https://help.aliyun.com/api/iot/querytopicroutetable.html
func (client *Client) QueryTopicRouteTable(request *QueryTopicRouteTableRequest) (response *QueryTopicRouteTableResponse, err error) {
	response = CreateQueryTopicRouteTableResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTopicRouteTableWithChan invokes the iot.QueryTopicRouteTable API asynchronously
// api document: https://help.aliyun.com/api/iot/querytopicroutetable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTopicRouteTableWithChan(request *QueryTopicRouteTableRequest) (<-chan *QueryTopicRouteTableResponse, <-chan error) {
	responseChan := make(chan *QueryTopicRouteTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTopicRouteTable(request)
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

// QueryTopicRouteTableWithCallback invokes the iot.QueryTopicRouteTable API asynchronously
// api document: https://help.aliyun.com/api/iot/querytopicroutetable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTopicRouteTableWithCallback(request *QueryTopicRouteTableRequest, callback func(response *QueryTopicRouteTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTopicRouteTableResponse
		var err error
		defer close(result)
		response, err = client.QueryTopicRouteTable(request)
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

// QueryTopicRouteTableRequest is the request struct for api QueryTopicRouteTable
type QueryTopicRouteTableRequest struct {
	*requests.RpcRequest
	AccessKeyId   string `position:"Query" name:"AccessKeyId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	RegionId      string `position:"Query" name:"RegionId"`
	Topic         string `position:"Query" name:"Topic"`
}

// QueryTopicRouteTableResponse is the response struct for api QueryTopicRouteTable
type QueryTopicRouteTableResponse struct {
	*responses.BaseResponse
	RequestId    string   `json:"RequestId" xml:"RequestId"`
	Success      bool     `json:"Success" xml:"Success"`
	Code         string   `json:"Code" xml:"Code"`
	ErrorMessage string   `json:"ErrorMessage" xml:"ErrorMessage"`
	DstTopics    []string `json:"DstTopics" xml:"DstTopics"`
}

// CreateQueryTopicRouteTableRequest creates a request to invoke QueryTopicRouteTable API
func CreateQueryTopicRouteTableRequest() (request *QueryTopicRouteTableRequest) {
	request = &QueryTopicRouteTableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryTopicRouteTable", "iot", "openAPI")
	return
}

// CreateQueryTopicRouteTableResponse creates a response to parse from QueryTopicRouteTable response
func CreateQueryTopicRouteTableResponse() (response *QueryTopicRouteTableResponse) {
	response = &QueryTopicRouteTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
