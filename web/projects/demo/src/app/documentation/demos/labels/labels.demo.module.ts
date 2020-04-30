import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { LabelsDemoComponent } from './labels.demo';
import { ApiLabelsDemoComponent } from './api-labels.demo';
import { AngularLabelsDemoComponent } from './angular-labels.demo';

import { SharedModule } from '../../../../../../../src/app/modules/shared/shared.module';
import { UtilsModule } from '../../../utils/utils.module';

@NgModule({
  imports: [
    UtilsModule,
    SharedModule,
    CommonModule,
    FormsModule,
    FormsModule,
    RouterModule.forChild([{ path: '', component: LabelsDemoComponent }]),
  ],
  declarations: [
    AngularLabelsDemoComponent,
    LabelsDemoComponent,
    ApiLabelsDemoComponent,
  ],
  exports: [LabelsDemoComponent],
})
export class LabelsDemoModule {}
