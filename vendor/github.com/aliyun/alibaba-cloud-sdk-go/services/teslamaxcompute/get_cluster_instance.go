package teslamaxcompute

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

// GetClusterInstance invokes the teslamaxcompute.GetClusterInstance API synchronously
// api document: https://help.aliyun.com/api/teslamaxcompute/getclusterinstance.html
func (client *Client) GetClusterInstance(request *GetClusterInstanceRequest) (response *GetClusterInstanceResponse, err error) {
	response = CreateGetClusterInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// GetClusterInstanceWithChan invokes the teslamaxcompute.GetClusterInstance API asynchronously
// api document: https://help.aliyun.com/api/teslamaxcompute/getclusterinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetClusterInstanceWithChan(request *GetClusterInstanceRequest) (<-chan *GetClusterInstanceResponse, <-chan error) {
	responseChan := make(chan *GetClusterInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetClusterInstance(request)
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

// GetClusterInstanceWithCallback invokes the teslamaxcompute.GetClusterInstance API asynchronously
// api document: https://help.aliyun.com/api/teslamaxcompute/getclusterinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetClusterInstanceWithCallback(request *GetClusterInstanceRequest, callback func(response *GetClusterInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetClusterInstanceResponse
		var err error
		defer close(result)
		response, err = client.GetClusterInstance(request)
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

// GetClusterInstanceRequest is the request struct for api GetClusterInstance
type GetClusterInstanceRequest struct {
	*requests.RpcRequest
	Cluster  string           `position:"Query" name:"Cluster"`
	PageSize requests.Integer `position:"Query" name:"PageSize"`
	PageNum  requests.Integer `position:"Query" name:"PageNum"`
	Region   string           `position:"Query" name:"Region"`
	Status   string           `position:"Query" name:"Status"`
}

// GetClusterInstanceResponse is the response struct for api GetClusterInstance
type GetClusterInstanceResponse struct {
	*responses.BaseResponse
	Code      int                      `json:"Code" xml:"Code"`
	Message   string                   `json:"Message" xml:"Message"`
	RequestId string                   `json:"RequestId" xml:"RequestId"`
	Data      DataInGetClusterInstance `json:"Data" xml:"Data"`
}

// CreateGetClusterInstanceRequest creates a request to invoke GetClusterInstance API
func CreateGetClusterInstanceRequest() (request *GetClusterInstanceRequest) {
	request = &GetClusterInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("TeslaMaxCompute", "2018-01-04", "GetClusterInstance", "teslamaxcompute", "openAPI")
	return
}

// CreateGetClusterInstanceResponse creates a response to parse from GetClusterInstance response
func CreateGetClusterInstanceResponse() (response *GetClusterInstanceResponse) {
	response = &GetClusterInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
