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

package disco

import (
	"context"
	"fmt"

	"github.com/apache/servicecomb-service-center/datasource"
	"github.com/apache/servicecomb-service-center/datasource/schema"
	"github.com/apache/servicecomb-service-center/pkg/log"
)

func RetireService(ctx context.Context, plan *datasource.RetirePlan) error {
	return datasource.GetMetadataManager().RetireService(ctx, plan)
}

func RetireSchema(ctx context.Context) error {
	n, err := schema.Instance().DeleteNoRefContents(ctx)
	if err != nil {
		log.Error("delete no ref contents failed", err)
		return err
	}
	if n > 0 {
		log.Warn(fmt.Sprintf("%d schema-contents retired", n))
	}
	return nil
}
