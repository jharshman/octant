// Copyright (c) 2019 the Octant contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
//

import { Component, Input, OnChanges, SimpleChanges } from '@angular/core';
import {
  Confirmation,
  GridAction,
  GridActionsView,
  TableFilters,
  TableRow,
  TableView,
  View,
} from 'src/app/modules/shared/models/content';
import trackByIndex from 'src/app/util/trackBy/trackByIndex';
import trackByIdentity from 'src/app/util/trackBy/trackByIdentity';
import { ViewService } from '../../../services/view/view.service';
import { ActionService } from '../../../services/action/action.service';

@Component({
  selector: 'app-view-datagrid',
  templateUrl: './datagrid.component.html',
  styleUrls: ['./datagrid.component.scss'],
})
export class DatagridComponent implements OnChanges {
  private v: TableView;

  @Input() set view(v: View) {
    this.v = v as TableView;
  }
  get view() {
    return this.v;
  }

  columns: string[];
  rowsWithMetadata: TableRowWithMetadata[];
  title: string;
  placeholder: string;
  lastUpdated: Date;
  filters: TableFilters;
  isModalOpen = false;
  confirmation?: Confirmation;

  private modalAction: GridAction;
  private previousView: SimpleChanges;

  identifyRow = trackByIndex;
  identifyColumn = trackByIdentity;
  loading: boolean;

  constructor(
    private viewService: ViewService,
    private actionService: ActionService
  ) {}

  ngOnChanges(changes: SimpleChanges): void {
    if (changes.view) {
      if (
        JSON.stringify(this.previousView) !==
        JSON.stringify(changes.view.currentValue)
      ) {
        this.title = this.viewService.viewTitleAsText(this.view);

        const current = changes.view.currentValue as TableView;
        this.columns = current.config.columns.map(column => column.name);

        if (current.config.rows) {
          this.rowsWithMetadata = this.getRowsWithMetadata(current.config.rows);
        }

        this.placeholder = current.config.emptyContent;
        this.lastUpdated = new Date();
        this.loading = current.config.loading;
        this.filters = current.config.filters;

        this.previousView = changes.view.currentValue;
      }
    }
  }

  private getRowsWithMetadata(rows: TableRow[]) {
    return rows.map(row => {
      let actions: GridAction[] = [];

      if (row.hasOwnProperty('_action')) {
        actions = (row._action as GridActionsView).config.actions;
      }
      return {
        data: row,
        actions,
      };
    });
  }

  runAction(action: GridAction) {
    if (!action.confirmation) {
      const update = { ...action.payload, action: action.actionPath };
      this.actionService.perform(update);
      return;
    }

    this.modalAction = action;
    this.isModalOpen = true;
    this.confirmation = action.confirmation;
  }

  showTitle() {
    if (this.view) {
      return this.view.totalItems === undefined || this.view.totalItems > 1;
    }
    return true;
  }

  cancelModal() {
    this.resetModal();
  }

  acceptModal() {
    this.resetModal();
    const update = {
      ...this.modalAction.payload,
      action: this.modalAction.actionPath,
    };
    this.actionService.perform(update);
  }

  private resetModal() {
    this.isModalOpen = false;
    this.confirmation = undefined;
  }
}

interface TableRowWithMetadata {
  data: TableRow;
  actions?: GridAction[];
}
