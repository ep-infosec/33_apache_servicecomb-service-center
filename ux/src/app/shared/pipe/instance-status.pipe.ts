/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
import { Pipe, PipeTransform } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';

@Pipe({ name: 'InstanceStatus' })
export class InstanceStatusPipe implements PipeTransform {
  statusMap!: StatusMap;
  constructor(private translate: TranslateService) {
    // UP在线,OUTOFSERVICE摘机,STARTING正在启动,DOWN下线,TESTING拨测状态。
    this.translate
      .get('instanceStatus')
      .subscribe(({ UP, DOWN, STARTING, TESTING, OUTOFSERVICE }) => {
        this.statusMap = {
          UP,
          DOWN,
          STARTING,
          TESTING,
          OUTOFSERVICE,
        };
      });
  }
  transform(value: State): string {
    return this.statusMap[value] || value;
  }
}
export interface StatusMap {
  UP: string;
  DOWN: string;
  STARTING: string;
  OUTOFSERVICE: string;
  TESTING: string;
}

type State = keyof StatusMap;
