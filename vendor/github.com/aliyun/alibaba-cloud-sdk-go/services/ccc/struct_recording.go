package ccc

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

// Recording is a nested struct in ccc response
type Recording struct {
	AgentId         string `json:"AgentId" xml:"AgentId"`
	Duration        int    `json:"Duration" xml:"Duration"`
	CallingNumber   string `json:"CallingNumber" xml:"CallingNumber"`
	ContactType     string `json:"ContactType" xml:"ContactType"`
	FilePath        string `json:"FilePath" xml:"FilePath"`
	FileDescription string `json:"FileDescription" xml:"FileDescription"`
	Channel         string `json:"Channel" xml:"Channel"`
	StartTime       int    `json:"StartTime" xml:"StartTime"`
	CalledNumber    string `json:"CalledNumber" xml:"CalledNumber"`
	ContactId       string `json:"ContactId" xml:"ContactId"`
	FileName        string `json:"FileName" xml:"FileName"`
	InstanceId      string `json:"InstanceId" xml:"InstanceId"`
	AgentName       string `json:"AgentName" xml:"AgentName"`
}
