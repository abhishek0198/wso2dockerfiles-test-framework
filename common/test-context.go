/*
 ------------------------------------------------------------------------

 Copyright 2016 WSO2, Inc. (http://wso2.com)

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License

 ------------------------------------------------------------------------
*/

package common

import "log"

const (
	BuildLogs = "buildlogs.log"
	RunLogs = "runlogs.log"
	TestConfigPath = "test-config.json"
)

var (
	Testconfig TestConfig
	Logger     *log.Logger
	DockerfilesHome string
)

type Product struct {
	Enabled             string
	Name                string
	Version             string
	Provisioning_method []string
	Organization        string
	Platform            string
	Profile             []string
	Smoke_Test_Path     string
}

type TestConfig struct {
	Products               []Product
	Output_file            string
	Docker_Container_Ip    string
	Carbon_Server_Port     string
	Carbon_Server_Username string
	Carbon_Server_Password string
	Silent_Build_Output    string
}
