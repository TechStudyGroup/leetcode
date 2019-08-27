package cloudwf

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

// GetApgroupScanConfigSaveProgress invokes the cloudwf.GetApgroupScanConfigSaveProgress API synchronously
// api document: https://help.aliyun.com/api/cloudwf/getapgroupscanconfigsaveprogress.html
func (client *Client) GetApgroupScanConfigSaveProgress(request *GetApgroupScanConfigSaveProgressRequest) (response *GetApgroupScanConfigSaveProgressResponse, err error) {
	response = CreateGetApgroupScanConfigSaveProgressResponse()
	err = client.DoAction(request, response)
	return
}

// GetApgroupScanConfigSaveProgressWithChan invokes the cloudwf.GetApgroupScanConfigSaveProgress API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getapgroupscanconfigsaveprogress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetApgroupScanConfigSaveProgressWithChan(request *GetApgroupScanConfigSaveProgressRequest) (<-chan *GetApgroupScanConfigSaveProgressResponse, <-chan error) {
	responseChan := make(chan *GetApgroupScanConfigSaveProgressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetApgroupScanConfigSaveProgress(request)
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

// GetApgroupScanConfigSaveProgressWithCallback invokes the cloudwf.GetApgroupScanConfigSaveProgress API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getapgroupscanconfigsaveprogress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetApgroupScanConfigSaveProgressWithCallback(request *GetApgroupScanConfigSaveProgressRequest, callback func(response *GetApgroupScanConfigSaveProgressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetApgroupScanConfigSaveProgressResponse
		var err error
		defer close(result)
		response, err = client.GetApgroupScanConfigSaveProgress(request)
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

// GetApgroupScanConfigSaveProgressRequest is the request struct for api GetApgroupScanConfigSaveProgress
type GetApgroupScanConfigSaveProgressRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Query" name:"Id"`
}

// GetApgroupScanConfigSaveProgressResponse is the response struct for api GetApgroupScanConfigSaveProgress
type GetApgroupScanConfigSaveProgressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateGetApgroupScanConfigSaveProgressRequest creates a request to invoke GetApgroupScanConfigSaveProgress API
func CreateGetApgroupScanConfigSaveProgressRequest() (request *GetApgroupScanConfigSaveProgressRequest) {
	request = &GetApgroupScanConfigSaveProgressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GetApgroupScanConfigSaveProgress", "cloudwf", "openAPI")
	return
}

// CreateGetApgroupScanConfigSaveProgressResponse creates a response to parse from GetApgroupScanConfigSaveProgress response
func CreateGetApgroupScanConfigSaveProgressResponse() (response *GetApgroupScanConfigSaveProgressResponse) {
	response = &GetApgroupScanConfigSaveProgressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
