/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package buildin

import (
	"crypto/tls"

	"github.com/apache/servicecomb-service-center/pkg/plugin"
	"github.com/apache/servicecomb-service-center/server/plugin/security/tlsconf"
)

func init() {
	plugin.RegisterPlugin(plugin.Plugin{Kind: tlsconf.TLS, Name: "buildin", New: New})
}

func New() plugin.Instance {
	return &DefaultTLS{}
}

// DefaultTLS support new the *tls.Config object from certs and private key with password
type DefaultTLS struct {
}

func (c *DefaultTLS) ClientConfig() (*tls.Config, error) {
	df, ok := plugin.DynamicPluginFunc(tlsconf.TLS, "ClientConfig").(func() (*tls.Config, error))
	if ok {
		return df()
	}
	return GetClientTLSConfig()
}

func (c *DefaultTLS) ServerConfig() (*tls.Config, error) {
	df, ok := plugin.DynamicPluginFunc(tlsconf.TLS, "ServerConfig").(func() (*tls.Config, error))
	if ok {
		return df()
	}
	return GetServerTLSConfig()
}
