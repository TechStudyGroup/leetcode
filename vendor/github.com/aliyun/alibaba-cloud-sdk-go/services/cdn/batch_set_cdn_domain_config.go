package cdn

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

// BatchSetCdnDomainConfig invokes the cdn.BatchSetCdnDomainConfig API synchronously
// api document: https://help.aliyun.com/api/cdn/batchsetcdndomainconfig.html
func (client *Client) BatchSetCdnDomainConfig(request *BatchSetCdnDomainConfigRequest) (response *BatchSetCdnDomainConfigResponse, err error) {
	response = CreateBatchSetCdnDomainConfigResponse()
	err = client.DoAction(request, response)
	return
}

// BatchSetCdnDomainConfigWithChan invokes the cdn.BatchSetCdnDomainConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/batchsetcdndomainconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchSetCdnDomainConfigWithChan(request *BatchSetCdnDomainConfigRequest) (<-chan *BatchSetCdnDomainConfigResponse, <-chan error) {
	responseChan := make(chan *BatchSetCdnDomainConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchSetCdnDomainConfig(request)
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

// BatchSetCdnDomainConfigWithCallback invokes the cdn.BatchSetCdnDomainConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/batchsetcdndomainconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchSetCdnDomainConfigWithCallback(request *BatchSetCdnDomainConfigRequest, callback func(response *BatchSetCdnDomainConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchSetCdnDomainConfigResponse
		var err error
		defer close(result)
		response, err = client.BatchSetCdnDomainConfig(request)
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

// BatchSetCdnDomainConfigRequest is the request struct for api BatchSetCdnDomainConfig
type BatchSetCdnDomainConfigRequest struct {
	*requests.RpcRequest
	Functions     string           `position:"Query" name:"Functions"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainNames   string           `position:"Query" name:"DomainNames"`
	OwnerAccount  string           `position:"Query" name:"OwnerAccount"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// BatchSetCdnDomainConfigResponse is the response struct for api BatchSetCdnDomainConfig
type BatchSetCdnDomainConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBatchSetCdnDomainConfigRequest creates a request to invoke BatchSetCdnDomainConfig API
func CreateBatchSetCdnDomainConfigRequest() (request *BatchSetCdnDomainConfigRequest) {
	request = &BatchSetCdnDomainConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "BatchSetCdnDomainConfig", "", "")
	return
}

// CreateBatchSetCdnDomainConfigResponse creates a response to parse from BatchSetCdnDomainConfig response
func CreateBatchSetCdnDomainConfigResponse() (response *BatchSetCdnDomainConfigResponse) {
	response = &BatchSetCdnDomainConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
