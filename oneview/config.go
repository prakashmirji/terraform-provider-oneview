// (C) Copyright 2016 Hewlett Packard Enterprise Development LP
//
// Licensed under the Apache License, Version 2.0 (the "License");
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed
// under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package oneview 

import (
  "github.com/HewlettPackard/oneview-golang/ov"
  "github.com/HewlettPackard/oneview-golang/icsp"
)

type Config struct {
  OVDomain     string
  OVUsername   string
  OVPassword   string
  OVEndpoint   string
  OVSSLVerify  bool
  OVAPIVersion int

  ICSPDomain     string
  ICSPUsername   string
  ICSPPassword   string
  ICSPEndpoint   string
  ICSPSSLVerify  bool
  ICSPAPIVersion int

  ovClient *ov.OVClient
  icspClient *icsp.ICSPClient
}

func (c *Config) loadAndValidate() error {
  var client2 *ov.OVClient 

  client:= client2.NewOVClient(c.OVUsername, c.OVPassword, c.OVDomain, c.OVEndpoint, c.OVSSLVerify, c.OVAPIVersion)

  c.ovClient = client

  session, error := c.ovClient.SessionLogin()
  session = session
  if error != nil {
    return error
  }

  return nil
}

func (c *Config) loadAndValidateICSP() error{
  
  var client2 *icsp.ICSPClient 
 
  client := client2.NewICSPClient(c.ICSPUsername, c.ICSPPassword, c.ICSPDomain, c.ICSPEndpoint, c.ICSPSSLVerify, c.ICSPAPIVersion)
 
  c.icspClient = client
 
  session, error := c.icspClient.SessionLogin()
  session = session
  if error != nil {
    return error
  }
  
  return nil
}