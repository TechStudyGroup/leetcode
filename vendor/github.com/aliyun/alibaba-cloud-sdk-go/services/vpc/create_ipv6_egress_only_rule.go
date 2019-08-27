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

// CreateIpv6EgressOnlyRule invokes the vpc.CreateIpv6EgressOnlyRule API synchronously
// api document: https://help.aliyun.com/api/vpc/createipv6egressonlyrule.html
func (client *Client) CreateIpv6EgressOnlyRule(request *CreateIpv6EgressOnlyRuleRequest) (response *CreateIpv6EgressOnlyRuleResponse, err error) {
	response = CreateCreateIpv6EgressOnlyRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateIpv6EgressOnlyRuleWithChan invokes the vpc.CreateIpv6EgressOnlyRule API asynchronously
// api document: https://help.aliyun.com/api/vpc/createipv6egressonlyrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateIpv6EgressOnlyRuleWithChan(request *CreateIpv6EgressOnlyRuleRequest) (<-chan *CreateIpv6EgressOnlyRuleResponse, <-chan error) {
	responseChan := make(chan *CreateIpv6EgressOnlyRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateIpv6EgressOnlyRule(request)
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

// CreateIpv6EgressOnlyRuleWithCallback invokes the vpc.CreateIpv6EgressOnlyRule API asynchronously
// api document: https://help.aliyun.com/api/vpc/createipv6egressonlyrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateIpv6EgressOnlyRuleWithCallback(request *CreateIpv6EgressOnlyRuleRequest, callback func(response *CreateIpv6EgressOnlyRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateIpv6EgressOnlyRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateIpv6EgressOnlyRule(request)
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

// CreateIpv6EgressOnlyRuleRequest is the request struct for api CreateIpv6EgressOnlyRule
type CreateIpv6EgressOnlyRuleRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Description          string           `position:"Query" name:"Description"`
	InstanceType         string           `position:"Query" name:"InstanceType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	Ipv6GatewayId        string           `position:"Query" name:"Ipv6GatewayId"`
	Name                 string           `position:"Query" name:"Name"`
}

// CreateIpv6EgressOnlyRuleResponse is the response struct for api CreateIpv6EgressOnlyRule
type CreateIpv6EgressOnlyRuleResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	Ipv6EgressRuleId string `json:"Ipv6EgressRuleId" xml:"Ipv6EgressRuleId"`
}

// CreateCreateIpv6EgressOnlyRuleRequest creates a request to invoke CreateIpv6EgressOnlyRule API
func CreateCreateIpv6EgressOnlyRuleRequest() (request *CreateIpv6EgressOnlyRuleRequest) {
	request = &CreateIpv6EgressOnlyRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateIpv6EgressOnlyRule", "vpc", "openAPI")
	return
}

// CreateCreateIpv6EgressOnlyRuleResponse creates a response to parse from CreateIpv6EgressOnlyRule response
func CreateCreateIpv6EgressOnlyRuleResponse() (response *CreateIpv6EgressOnlyRuleResponse) {
	response = &CreateIpv6EgressOnlyRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
