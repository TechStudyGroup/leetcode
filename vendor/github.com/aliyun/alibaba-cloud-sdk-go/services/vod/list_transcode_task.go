package vod

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

// ListTranscodeTask invokes the vod.ListTranscodeTask API synchronously
// api document: https://help.aliyun.com/api/vod/listtranscodetask.html
func (client *Client) ListTranscodeTask(request *ListTranscodeTaskRequest) (response *ListTranscodeTaskResponse, err error) {
	response = CreateListTranscodeTaskResponse()
	err = client.DoAction(request, response)
	return
}

// ListTranscodeTaskWithChan invokes the vod.ListTranscodeTask API asynchronously
// api document: https://help.aliyun.com/api/vod/listtranscodetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTranscodeTaskWithChan(request *ListTranscodeTaskRequest) (<-chan *ListTranscodeTaskResponse, <-chan error) {
	responseChan := make(chan *ListTranscodeTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTranscodeTask(request)
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

// ListTranscodeTaskWithCallback invokes the vod.ListTranscodeTask API asynchronously
// api document: https://help.aliyun.com/api/vod/listtranscodetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTranscodeTaskWithCallback(request *ListTranscodeTaskRequest, callback func(response *ListTranscodeTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTranscodeTaskResponse
		var err error
		defer close(result)
		response, err = client.ListTranscodeTask(request)
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

// ListTranscodeTaskRequest is the request struct for api ListTranscodeTask
type ListTranscodeTaskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PageNo               requests.Integer `position:"Query" name:"PageNo"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	EndTime              string           `position:"Query" name:"EndTime"`
	VideoId              string           `position:"Query" name:"VideoId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ListTranscodeTaskResponse is the response struct for api ListTranscodeTask
type ListTranscodeTaskResponse struct {
	*responses.BaseResponse
	RequestId         string          `json:"RequestId" xml:"RequestId"`
	TranscodeTaskList []TranscodeTask `json:"TranscodeTaskList" xml:"TranscodeTaskList"`
}

// CreateListTranscodeTaskRequest creates a request to invoke ListTranscodeTask API
func CreateListTranscodeTaskRequest() (request *ListTranscodeTaskRequest) {
	request = &ListTranscodeTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "ListTranscodeTask", "vod", "openAPI")
	return
}

// CreateListTranscodeTaskResponse creates a response to parse from ListTranscodeTask response
func CreateListTranscodeTaskResponse() (response *ListTranscodeTaskResponse) {
	response = &ListTranscodeTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
