package smartag

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

// DeleteSAGLinkLevelHa invokes the smartag.DeleteSAGLinkLevelHa API synchronously
// api document: https://help.aliyun.com/api/smartag/deletesaglinklevelha.html
func (client *Client) DeleteSAGLinkLevelHa(request *DeleteSAGLinkLevelHaRequest) (response *DeleteSAGLinkLevelHaResponse, err error) {
	response = CreateDeleteSAGLinkLevelHaResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSAGLinkLevelHaWithChan invokes the smartag.DeleteSAGLinkLevelHa API asynchronously
// api document: https://help.aliyun.com/api/smartag/deletesaglinklevelha.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteSAGLinkLevelHaWithChan(request *DeleteSAGLinkLevelHaRequest) (<-chan *DeleteSAGLinkLevelHaResponse, <-chan error) {
	responseChan := make(chan *DeleteSAGLinkLevelHaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSAGLinkLevelHa(request)
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

// DeleteSAGLinkLevelHaWithCallback invokes the smartag.DeleteSAGLinkLevelHa API asynchronously
// api document: https://help.aliyun.com/api/smartag/deletesaglinklevelha.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteSAGLinkLevelHaWithCallback(request *DeleteSAGLinkLevelHaRequest, callback func(response *DeleteSAGLinkLevelHaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSAGLinkLevelHaResponse
		var err error
		defer close(result)
		response, err = client.DeleteSAGLinkLevelHa(request)
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

// DeleteSAGLinkLevelHaRequest is the request struct for api DeleteSAGLinkLevelHa
type DeleteSAGLinkLevelHaRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteSAGLinkLevelHaResponse is the response struct for api DeleteSAGLinkLevelHa
type DeleteSAGLinkLevelHaResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteSAGLinkLevelHaRequest creates a request to invoke DeleteSAGLinkLevelHa API
func CreateDeleteSAGLinkLevelHaRequest() (request *DeleteSAGLinkLevelHaRequest) {
	request = &DeleteSAGLinkLevelHaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DeleteSAGLinkLevelHa", "smartag", "openAPI")
	return
}

// CreateDeleteSAGLinkLevelHaResponse creates a response to parse from DeleteSAGLinkLevelHa response
func CreateDeleteSAGLinkLevelHaResponse() (response *DeleteSAGLinkLevelHaResponse) {
	response = &DeleteSAGLinkLevelHaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
