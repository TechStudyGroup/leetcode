package mts

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

// UnbindOutputBucket invokes the mts.UnbindOutputBucket API synchronously
// api document: https://help.aliyun.com/api/mts/unbindoutputbucket.html
func (client *Client) UnbindOutputBucket(request *UnbindOutputBucketRequest) (response *UnbindOutputBucketResponse, err error) {
	response = CreateUnbindOutputBucketResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindOutputBucketWithChan invokes the mts.UnbindOutputBucket API asynchronously
// api document: https://help.aliyun.com/api/mts/unbindoutputbucket.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnbindOutputBucketWithChan(request *UnbindOutputBucketRequest) (<-chan *UnbindOutputBucketResponse, <-chan error) {
	responseChan := make(chan *UnbindOutputBucketResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindOutputBucket(request)
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

// UnbindOutputBucketWithCallback invokes the mts.UnbindOutputBucket API asynchronously
// api document: https://help.aliyun.com/api/mts/unbindoutputbucket.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnbindOutputBucketWithCallback(request *UnbindOutputBucketRequest, callback func(response *UnbindOutputBucketResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindOutputBucketResponse
		var err error
		defer close(result)
		response, err = client.UnbindOutputBucket(request)
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

// UnbindOutputBucketRequest is the request struct for api UnbindOutputBucket
type UnbindOutputBucketRequest struct {
	*requests.RpcRequest
	Bucket               string           `position:"Query" name:"Bucket"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// UnbindOutputBucketResponse is the response struct for api UnbindOutputBucket
type UnbindOutputBucketResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnbindOutputBucketRequest creates a request to invoke UnbindOutputBucket API
func CreateUnbindOutputBucketRequest() (request *UnbindOutputBucketRequest) {
	request = &UnbindOutputBucketRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "UnbindOutputBucket", "mts", "openAPI")
	return
}

// CreateUnbindOutputBucketResponse creates a response to parse from UnbindOutputBucket response
func CreateUnbindOutputBucketResponse() (response *UnbindOutputBucketResponse) {
	response = &UnbindOutputBucketResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
