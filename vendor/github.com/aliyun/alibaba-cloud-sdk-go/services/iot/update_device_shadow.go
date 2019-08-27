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

// UpdateDeviceShadow invokes the iot.UpdateDeviceShadow API synchronously
// api document: https://help.aliyun.com/api/iot/updatedeviceshadow.html
func (client *Client) UpdateDeviceShadow(request *UpdateDeviceShadowRequest) (response *UpdateDeviceShadowResponse, err error) {
	response = CreateUpdateDeviceShadowResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDeviceShadowWithChan invokes the iot.UpdateDeviceShadow API asynchronously
// api document: https://help.aliyun.com/api/iot/updatedeviceshadow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateDeviceShadowWithChan(request *UpdateDeviceShadowRequest) (<-chan *UpdateDeviceShadowResponse, <-chan error) {
	responseChan := make(chan *UpdateDeviceShadowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDeviceShadow(request)
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

// UpdateDeviceShadowWithCallback invokes the iot.UpdateDeviceShadow API asynchronously
// api document: https://help.aliyun.com/api/iot/updatedeviceshadow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateDeviceShadowWithCallback(request *UpdateDeviceShadowRequest, callback func(response *UpdateDeviceShadowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDeviceShadowResponse
		var err error
		defer close(result)
		response, err = client.UpdateDeviceShadow(request)
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

// UpdateDeviceShadowRequest is the request struct for api UpdateDeviceShadow
type UpdateDeviceShadowRequest struct {
	*requests.RpcRequest
	AccessKeyId   string `position:"Query" name:"AccessKeyId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	DeviceName    string `position:"Query" name:"DeviceName"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	ShadowMessage string `position:"Query" name:"ShadowMessage"`
}

// UpdateDeviceShadowResponse is the response struct for api UpdateDeviceShadow
type UpdateDeviceShadowResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateUpdateDeviceShadowRequest creates a request to invoke UpdateDeviceShadow API
func CreateUpdateDeviceShadowRequest() (request *UpdateDeviceShadowRequest) {
	request = &UpdateDeviceShadowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "UpdateDeviceShadow", "iot", "openAPI")
	return
}

// CreateUpdateDeviceShadowResponse creates a response to parse from UpdateDeviceShadow response
func CreateUpdateDeviceShadowResponse() (response *UpdateDeviceShadowResponse) {
	response = &UpdateDeviceShadowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
