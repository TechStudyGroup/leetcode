package vpc

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

// ModifyVRouterAttribute invokes the vpc.ModifyVRouterAttribute API synchronously
// api document: https://help.aliyun.com/api/vpc/modifyvrouterattribute.html
func (client *Client) ModifyVRouterAttribute(request *ModifyVRouterAttributeRequest) (response *ModifyVRouterAttributeResponse, err error) {
	response = CreateModifyVRouterAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVRouterAttributeWithChan invokes the vpc.ModifyVRouterAttribute API asynchronously
// api document: https://help.aliyun.com/api/vpc/modifyvrouterattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyVRouterAttributeWithChan(request *ModifyVRouterAttributeRequest) (<-chan *ModifyVRouterAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyVRouterAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVRouterAttribute(request)
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

// ModifyVRouterAttributeWithCallback invokes the vpc.ModifyVRouterAttribute API asynchronously
// api document: https://help.aliyun.com/api/vpc/modifyvrouterattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyVRouterAttributeWithCallback(request *ModifyVRouterAttributeRequest, callback func(response *ModifyVRouterAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVRouterAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyVRouterAttribute(request)
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

// ModifyVRouterAttributeRequest is the request struct for api ModifyVRouterAttribute
type ModifyVRouterAttributeRequest struct {
	*requests.RpcRequest
	VRouterName          string           `position:"Query" name:"VRouterName"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	VRouterId            string           `position:"Query" name:"VRouterId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Description          string           `position:"Query" name:"Description"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyVRouterAttributeResponse is the response struct for api ModifyVRouterAttribute
type ModifyVRouterAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyVRouterAttributeRequest creates a request to invoke ModifyVRouterAttribute API
func CreateModifyVRouterAttributeRequest() (request *ModifyVRouterAttributeRequest) {
	request = &ModifyVRouterAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVRouterAttribute", "vpc", "openAPI")
	return
}

// CreateModifyVRouterAttributeResponse creates a response to parse from ModifyVRouterAttribute response
func CreateModifyVRouterAttributeResponse() (response *ModifyVRouterAttributeResponse) {
	response = &ModifyVRouterAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
