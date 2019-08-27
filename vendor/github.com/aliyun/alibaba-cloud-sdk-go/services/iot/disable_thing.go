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

// DisableThing invokes the iot.DisableThing API synchronously
// api document: https://help.aliyun.com/api/iot/disablething.html
func (client *Client) DisableThing(request *DisableThingRequest) (response *DisableThingResponse, err error) {
	response = CreateDisableThingResponse()
	err = client.DoAction(request, response)
	return
}

// DisableThingWithChan invokes the iot.DisableThing API asynchronously
// api document: https://help.aliyun.com/api/iot/disablething.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableThingWithChan(request *DisableThingRequest) (<-chan *DisableThingResponse, <-chan error) {
	responseChan := make(chan *DisableThingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableThing(request)
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

// DisableThingWithCallback invokes the iot.DisableThing API asynchronously
// api document: https://help.aliyun.com/api/iot/disablething.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableThingWithCallback(request *DisableThingRequest, callback func(response *DisableThingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableThingResponse
		var err error
		defer close(result)
		response, err = client.DisableThing(request)
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

// DisableThingRequest is the request struct for api DisableThing
type DisableThingRequest struct {
	*requests.RpcRequest
	AccessKeyId   string `position:"Query" name:"AccessKeyId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	IotId         string `position:"Query" name:"IotId"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	DeviceName    string `position:"Query" name:"DeviceName"`
}

// DisableThingResponse is the response struct for api DisableThing
type DisableThingResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateDisableThingRequest creates a request to invoke DisableThing API
func CreateDisableThingRequest() (request *DisableThingRequest) {
	request = &DisableThingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "DisableThing", "iot", "openAPI")
	return
}

// CreateDisableThingResponse creates a response to parse from DisableThing response
func CreateDisableThingResponse() (response *DisableThingResponse) {
	response = &DisableThingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
