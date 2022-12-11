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
import { Component, OnInit, ViewChild } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';
import { cloneDeep, map, uniqBy } from 'lodash';
import { ICategorySearchTagItem } from 'ng-devui';
import { DataTableComponent, TableWidthConfig } from 'ng-devui/data-table';
import { ModalService } from 'ng-devui/modal';
import { envOptions } from 'src/config/global.config';
import { ServiceService } from '../../common/service.service';
import {
  FilterItem,
  filterTabDataByCategory,
  getTabelData,
} from '../shared/toolFunction/tabel.pagination';
import { CreateComponent } from './modal/create/create.component';
import { DeleteComponent } from './modal/delete/delete.component';

@Component({
  selector: 'app-service-list',
  templateUrl: './service-list.component.html',
  styleUrls: ['./service-list.component.less'],
})
export class ServiceListComponent implements OnInit {
  constructor(
    private modalService: ModalService,
    private service: ServiceService,
    private translate: TranslateService
  ) {}
  @ViewChild(DataTableComponent, { static: true })
  datatable!: DataTableComponent;

  private basicDataSource: any;
  dataSource: any; // 展示的数据

  tableWidthConfig: TableWidthConfig[] = [
    {
      field: 'name',
      width: '150px',
    },
    {
      field: 'environment',
      width: '100px',
    },
    {
      field: 'version',
      width: '100px',
    },
    {
      field: 'instances',
      width: '150px',
    },
    {
      field: 'operation',
      width: '150px',
    },
  ];
  private totalDataChecked = false;
  pager = {
    total: 0,
    pageIndex: 1,
    pageSize: 10,
    pageSizeOptions: [5, 10, 20, 50],
  };

  // todo ui框架问题设置为any解决
  category: Array<ICategorySearchTagItem> | any;

  ngOnInit(): void {
    this.initData();
  }

  async initData(): Promise<void> {
    this.dataSource = [];
    this.category = [];
    const common = await this.translate.get('common').toPromise();
    const columns = await this.translate.get('service.columns').toPromise();
    this.category = [
      {
        field: 'serviceName',
        label: columns.serviceName,
        type: 'textInput',
      },
      {
        field: 'environment',
        label: columns.environment,
        type: 'label',
        options: cloneDeep(
          envOptions.map((item) => {
            item.label = item.label || common.empty;
            return item;
          })
        ),
      },
      {
        field: 'version',
        label: columns.version,
        type: 'label',
        options: [],
      },
      {
        field: 'appId',
        label: columns.app,
        type: 'textInput',
      },
    ];
    this.service.getServiceByGovern().subscribe(
      (data) => {
        this.basicDataSource = (data?.allServicesDetail || [])
          .map((item: any) => {
            if (!item.microService?.environment) {
              item.microService.environment = common.empty;
            }
            return item.microService;
          })
          .sort((a: any, b: any) => b.modTimestamp - a.modTimestamp);
        this.pager.total = this.basicDataSource.length;
        this.dataSource = getTabelData(this.basicDataSource, this.pager);
        this.category[2].options = uniqBy(
          map(this.basicDataSource, (item: any) => ({
            label: item.version,
          })),
          'label'
        );
      },
      (err) => {
        // todo 提示
      }
    );
  }

  public deleteItem(rowItems: any): void {
    const result = this.modalService.open({
      id: 'delte-service',
      width: '750px',
      component: DeleteComponent,
      data: {
        services: rowItems,
        onCancel: (data?: any) => {
          if (data) {
            this.initData();
          }
          result.modalInstance.hide();
        },
      },
    });
  }

  public onPaginationChange(pageIndex: number, pageSize: number): void {
    this.dataSource = getTabelData(this.basicDataSource, {
      ...cloneDeep(this.pager),
      pageIndex,
      pageSize,
    });
    setTimeout(() => {
      if (this.totalDataChecked) {
        this.datatable.setTableCheckStatus({
          pageAllChecked: true,
        });
      } else {
        this.datatable.setTableCheckStatus({
          pageAllChecked: false,
        });
      }
    });
  }

  onCreateService(): void {
    const resulse = this.modalService.open({
      id: 'create-service',
      showAnimation: true,
      width: '750px',
      component: CreateComponent,
      data: {
        onClose: (res?: any) => {
          if (res) {
            this.initData();
          }
          resulse.modalInstance.hide();
        },
      },
    });
  }

  onSelectedTagsChange(e: FilterItem[]): void {
    const { data, tableData, pageination } = filterTabDataByCategory(
      this.basicDataSource,
      this.pager,
      e
    );
    this.pager = pageination;
    this.dataSource = tableData;
  }

  onRefresh(): void {
    this.initData();
  }
}
