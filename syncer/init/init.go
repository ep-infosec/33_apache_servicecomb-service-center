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

package init

import (
	// handlers
	_ "github.com/go-chassis/go-chassis/v2/middleware/monitoring"
	_ "github.com/go-chassis/go-chassis/v2/middleware/ratelimiter"

	//grpc plugin
	_ "github.com/go-chassis/go-chassis-extension/protocol/grpc/server"

	// kie db
	_ "github.com/apache/servicecomb-kie/server/datasource/etcd"

	"github.com/apache/servicecomb-service-center/pkg/log"
	syncv1 "github.com/apache/servicecomb-service-center/syncer/api/v1"
	"github.com/apache/servicecomb-service-center/syncer/config"
	"github.com/apache/servicecomb-service-center/syncer/rpc"
	"github.com/go-chassis/go-chassis/v2"

	//codec
	_ "github.com/apache/servicecomb-service-center/server/plugin"
	_ "github.com/go-chassis/go-chassis-extension/codec/gojson"

	chassisServer "github.com/go-chassis/go-chassis/v2/core/server"
)

func init() {
	chassis.RegisterSchema("grpc", rpc.NewServer(),
		chassisServer.WithRPCServiceDesc(&syncv1.EventService_ServiceDesc))

	if err := chassis.Init(); err != nil {
		log.Warn(err.Error())
	}
	if err := config.Init(); err != nil {
		log.Error("syncer config init failed", err)
	}
}
