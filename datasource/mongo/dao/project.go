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

package dao

import (
	"context"
	"fmt"

	"github.com/go-chassis/cari/db/mongo"

	"github.com/apache/servicecomb-service-center/datasource/mongo/model"
	"github.com/apache/servicecomb-service-center/pkg/log"
)

func AddProject(ctx context.Context, project model.Project) error {
	result, err := mongo.GetClient().GetDB().Collection(model.CollectionProject).InsertOne(ctx, project)
	if err == nil {
		log.Info(fmt.Sprintf("insert project to mongodb success %s", result.InsertedID))
	}
	return err
}

func ExistProject(ctx context.Context, filter interface{}) (bool, error) {
	result := mongo.GetClient().GetDB().Collection(model.CollectionProject).FindOne(ctx, filter)
	if result.Err() != nil {
		return false, nil
	}
	return true, nil
}
