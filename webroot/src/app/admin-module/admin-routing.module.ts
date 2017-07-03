import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {HomePageAdminComponent} from './components/home-page-admin/home-page-admin.component'

const routes: Routes = [

  { path: '', redirectTo: 'home', pathMatch: 'full' },
  {
    path: 'home',
    component: HomePageAdminComponent,
    children: [ ]
  }
];


@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AdminRoutingModule { }
