package emr

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

// SaveParagraph invokes the emr.SaveParagraph API synchronously
// api document: https://help.aliyun.com/api/emr/saveparagraph.html
func (client *Client) SaveParagraph(request *SaveParagraphRequest) (response *SaveParagraphResponse, err error) {
	response = CreateSaveParagraphResponse()
	err = client.DoAction(request, response)
	return
}

// SaveParagraphWithChan invokes the emr.SaveParagraph API asynchronously
// api document: https://help.aliyun.com/api/emr/saveparagraph.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveParagraphWithChan(request *SaveParagraphRequest) (<-chan *SaveParagraphResponse, <-chan error) {
	responseChan := make(chan *SaveParagraphResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveParagraph(request)
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

// SaveParagraphWithCallback invokes the emr.SaveParagraph API asynchronously
// api document: https://help.aliyun.com/api/emr/saveparagraph.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveParagraphWithCallback(request *SaveParagraphRequest, callback func(response *SaveParagraphResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveParagraphResponse
		var err error
		defer close(result)
		response, err = client.SaveParagraph(request)
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

// SaveParagraphRequest is the request struct for api SaveParagraph
type SaveParagraphRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	NoteId          string           `position:"Query" name:"NoteId"`
	Id              string           `position:"Query" name:"Id"`
	Text            string           `position:"Query" name:"Text"`
}

// SaveParagraphResponse is the response struct for api SaveParagraph
type SaveParagraphResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSaveParagraphRequest creates a request to invoke SaveParagraph API
func CreateSaveParagraphRequest() (request *SaveParagraphRequest) {
	request = &SaveParagraphRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "SaveParagraph", "emr", "openAPI")
	return
}

// CreateSaveParagraphResponse creates a response to parse from SaveParagraph response
func CreateSaveParagraphResponse() (response *SaveParagraphResponse) {
	response = &SaveParagraphResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
