package scdn

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

// SetScdnDomainCertificate invokes the scdn.SetScdnDomainCertificate API synchronously
// api document: https://help.aliyun.com/api/scdn/setscdndomaincertificate.html
func (client *Client) SetScdnDomainCertificate(request *SetScdnDomainCertificateRequest) (response *SetScdnDomainCertificateResponse, err error) {
	response = CreateSetScdnDomainCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// SetScdnDomainCertificateWithChan invokes the scdn.SetScdnDomainCertificate API asynchronously
// api document: https://help.aliyun.com/api/scdn/setscdndomaincertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetScdnDomainCertificateWithChan(request *SetScdnDomainCertificateRequest) (<-chan *SetScdnDomainCertificateResponse, <-chan error) {
	responseChan := make(chan *SetScdnDomainCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetScdnDomainCertificate(request)
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

// SetScdnDomainCertificateWithCallback invokes the scdn.SetScdnDomainCertificate API asynchronously
// api document: https://help.aliyun.com/api/scdn/setscdndomaincertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetScdnDomainCertificateWithCallback(request *SetScdnDomainCertificateRequest, callback func(response *SetScdnDomainCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetScdnDomainCertificateResponse
		var err error
		defer close(result)
		response, err = client.SetScdnDomainCertificate(request)
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

// SetScdnDomainCertificateRequest is the request struct for api SetScdnDomainCertificate
type SetScdnDomainCertificateRequest struct {
	*requests.RpcRequest
	ForceSet      string           `position:"Query" name:"ForceSet"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	CertType      string           `position:"Query" name:"CertType"`
	SSLPub        string           `position:"Query" name:"SSLPub"`
	CertName      string           `position:"Query" name:"CertName"`
	SSLProtocol   string           `position:"Query" name:"SSLProtocol"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	Region        string           `position:"Query" name:"Region"`
	SSLPri        string           `position:"Query" name:"SSLPri"`
}

// SetScdnDomainCertificateResponse is the response struct for api SetScdnDomainCertificate
type SetScdnDomainCertificateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetScdnDomainCertificateRequest creates a request to invoke SetScdnDomainCertificate API
func CreateSetScdnDomainCertificateRequest() (request *SetScdnDomainCertificateRequest) {
	request = &SetScdnDomainCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "SetScdnDomainCertificate", "scdn", "openAPI")
	return
}

// CreateSetScdnDomainCertificateResponse creates a response to parse from SetScdnDomainCertificate response
func CreateSetScdnDomainCertificateResponse() (response *SetScdnDomainCertificateResponse) {
	response = &SetScdnDomainCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
