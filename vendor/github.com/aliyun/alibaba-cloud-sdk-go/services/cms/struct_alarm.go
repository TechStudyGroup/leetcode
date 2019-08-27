package cms

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

// Alarm is a nested struct in cms response
type Alarm struct {
	Resources           string      `json:"Resources" xml:"Resources"`
	RuleName            string      `json:"RuleName" xml:"RuleName"`
	ContactGroups       string      `json:"ContactGroups" xml:"ContactGroups"`
	EffectiveInterval   string      `json:"EffectiveInterval" xml:"EffectiveInterval"`
	GroupName           string      `json:"GroupName" xml:"GroupName"`
	Dimensions          string      `json:"Dimensions" xml:"Dimensions"`
	RuleId              string      `json:"RuleId" xml:"RuleId"`
	ComparisonOperator  string      `json:"ComparisonOperator" xml:"ComparisonOperator"`
	MailSubject         string      `json:"MailSubject" xml:"MailSubject"`
	EnableState         bool        `json:"EnableState" xml:"EnableState"`
	GroupId             string      `json:"GroupId" xml:"GroupId"`
	EvaluationCount     string      `json:"EvaluationCount" xml:"EvaluationCount"`
	SilenceTime         string      `json:"SilenceTime" xml:"SilenceTime"`
	Threshold           string      `json:"Threshold" xml:"Threshold"`
	Period              string      `json:"Period" xml:"Period"`
	StartTime           string      `json:"StartTime" xml:"StartTime"`
	State               string      `json:"State" xml:"State"`
	EndTime             string      `json:"EndTime" xml:"EndTime"`
	Namespace           string      `json:"Namespace" xml:"Namespace"`
	Enable              string      `json:"Enable" xml:"Enable"`
	MetricName          string      `json:"MetricName" xml:"MetricName"`
	AlertState          string      `json:"AlertState" xml:"AlertState"`
	Statistics          string      `json:"Statistics" xml:"Statistics"`
	Webhook             string      `json:"Webhook" xml:"Webhook"`
	NoEffectiveInterval string      `json:"NoEffectiveInterval" xml:"NoEffectiveInterval"`
	Escalations         Escalations `json:"Escalations" xml:"Escalations"`
}
