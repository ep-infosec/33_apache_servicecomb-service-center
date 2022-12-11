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

package tlsconf

import (
	"crypto/tls"

	"github.com/apache/servicecomb-service-center/pkg/plugin"
)

const TLS plugin.Kind = "ssl"

var options Options

type TLSConfig interface {
	ClientConfig() (*tls.Config, error)
	ServerConfig() (*tls.Config, error)
}

func ClientConfig() (*tls.Config, error) {
	return plugin.Plugins().Instance(TLS).(TLSConfig).ClientConfig()
}

func ServerConfig() (*tls.Config, error) {
	return plugin.Plugins().Instance(TLS).(TLSConfig).ServerConfig()
}

func Init(opts Options) error {
	options = opts
	return nil
}

func GetOptions() Options {
	return options
}
