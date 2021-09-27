import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { NotebooksComponent } from './notebooks.component';
import { ListComponent } from './list/list.component'
import { CreateInsranceComponent } from './create-insrance/create-insrance.component'
const routes: Routes = [
  {
    path: '',
    component: NotebooksComponent,
    data: {
      keep: true
    },
    children: [
      {
        path: '',
        redirectTo: 'list'
      },
      {
        path: 'list',
        component: ListComponent,
        data: {
          keep: true
        }
      },
      {
        path: 'newInstance',
        component: CreateInsranceComponent,
        data: {
          keep: false
        }
      }
    ],
  },
];

export const declarations = [
];

export const providers = [
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class NotebooksRoutingModule {}