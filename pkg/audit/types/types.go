// Copyright © 2018 Heptio
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

// Event is an ModSecurity audit log event in JSON (`SecAuditLogFormat JSON`)
// Note: this is somewhat particular to the set of SecAuditLogParts we have configured.
type Event struct {
	Transaction Transaction `json:"transaction"`
}

// Transaction reprsents a single ModSecurity transaction (request->processing->response)
type Transaction struct {
	ClientIP   string    `json:"client_ip"`
	ClientPort int       `json:"client_port"`
	HostIP     string    `json:"host_ip"`
	HostPort   int       `json:"host_port"`
	ID         string    `json:"id"`
	Messages   []Message `json:"messages"`
	Producer   Producer  `json:"producer"`
	Request    Request   `json:"request"`
	Response   Response  `json:"response"`
	ServerID   string    `json:"server_id"`
	TimeStamp  string    `json:"time_stamp"`
}

// Message is the data generated by a single rule firing on a transaction
type Message struct {
	Details Details `json:"details"`
	Message string  `json:"message"`
}

// Details contains the details related to a single rule firing on a transaction
type Details struct {
	Accuracy   string   `json:"accuracy"`
	Data       string   `json:"data"`
	File       string   `json:"file"`
	LineNumber string   `json:"lineNumber"`
	Match      string   `json:"match"`
	Maturity   string   `json:"maturity"`
	Reference  string   `json:"reference"`
	Rev        string   `json:"rev"`
	RuleID     string   `json:"ruleId"`
	Severity   string   `json:"severity"`
	Tags       []string `json:"tags"`
	Ver        string   `json:"ver"`
}

// Producer identifies the components/versions of the software that generated the alert
type Producer struct {
	Components     []string `json:"components"`
	Connector      string   `json:"connector"`
	Modsecurity    string   `json:"modsecurity"`
	SecrulesEngine string   `json:"secrules_engine"`
}

// Request contains details about the HTTP request in the transaction
type Request struct {
	Headers     map[string]string `json:"headers"`
	HTTPVersion float64           `json:"http_version"`
	Method      string            `json:"method"`
	URI         string            `json:"uri"`
}

// Response contains details about the HTTP response in the transaction
type Response struct {
	HTTPCode int `json:"http_code"`
}
