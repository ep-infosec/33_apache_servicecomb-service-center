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

package plugin

// Kind is an alias, it represents a plugin interface.
type Kind string

// ImplName is an alias，it represents a plugin interface implementation.
type ImplName string

// Instance is an instance of a plugin interface which is represented by
// Kind.
type Instance interface{}

// String implements fmt.Stringer.
func (pn Kind) String() string {
	return string(pn)
}

// Plugin generates a plugin instance
// Plugin holds the 'Kind' and 'ImplName'
// to manage the plugin instance generation.
type Plugin struct {
	Kind Kind
	Name ImplName
	// New news an instance of 'Kind' represented plugin interface
	New func() Instance
}

// Configurator provides the interfaces for obtaining configs in
// plugin mgr's operations, user can customize the impl
type Configurator interface {
	GetImplName(kind Kind) string
	GetPluginDir() string
}
