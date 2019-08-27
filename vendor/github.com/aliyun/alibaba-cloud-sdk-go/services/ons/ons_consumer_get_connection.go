package ons

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

// OnsConsumerGetConnection invokes the ons.OnsConsumerGetConnection API synchronously
// api document: https://help.aliyun.com/api/ons/onsconsumergetconnection.html
func (client *Client) OnsConsumerGetConnection(request *OnsConsumerGetConnectionRequest) (response *OnsConsumerGetConnectionResponse, err error) {
	response = CreateOnsConsumerGetConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// OnsConsumerGetConnectionWithChan invokes the ons.OnsConsumerGetConnection API asynchronously
// api document: https://help.aliyun.com/api/ons/onsconsumergetconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsConsumerGetConnectionWithChan(request *OnsConsumerGetConnectionRequest) (<-chan *OnsConsumerGetConnectionResponse, <-chan error) {
	responseChan := make(chan *OnsConsumerGetConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsConsumerGetConnection(request)
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

// OnsConsumerGetConnectionWithCallback invokes the ons.OnsConsumerGetConnection API asynchronously
// api document: https://help.aliyun.com/api/ons/onsconsumergetconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsConsumerGetConnectionWithCallback(request *OnsConsumerGetConnectionRequest, callback func(response *OnsConsumerGetConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsConsumerGetConnectionResponse
		var err error
		defer close(result)
		response, err = client.OnsConsumerGetConnection(request)
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

// OnsConsumerGetConnectionRequest is the request struct for api OnsConsumerGetConnection
type OnsConsumerGetConnectionRequest struct {
	*requests.RpcRequest
	PreventCache requests.Integer `position:"Query" name:"PreventCache"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	GroupId      string           `position:"Query" name:"GroupId"`
}

// OnsConsumerGetConnectionResponse is the response struct for api OnsConsumerGetConnection
type OnsConsumerGetConnectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateOnsConsumerGetConnectionRequest creates a request to invoke OnsConsumerGetConnection API
func CreateOnsConsumerGetConnectionRequest() (request *OnsConsumerGetConnectionRequest) {
	request = &OnsConsumerGetConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsConsumerGetConnection", "ons", "openAPI")
	return
}

// CreateOnsConsumerGetConnectionResponse creates a response to parse from OnsConsumerGetConnection response
func CreateOnsConsumerGetConnectionResponse() (response *OnsConsumerGetConnectionResponse) {
	response = &OnsConsumerGetConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
