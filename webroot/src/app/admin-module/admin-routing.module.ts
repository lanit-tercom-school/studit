import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { HomePageAdminComponent } from './components/home-page-admin/home-page-admin.component';
import { AdminUserConrolComponent } from './components/admin-user-conrol/admin-user-conrol.component';
import { AdminProjectConrolComponent } from './components/admin-project-conrol/admin-project-conrol.component';
import { AdminNewsConrolComponent } from './components/admin-news-conrol/admin-news-conrol.component';

const routes: Routes = [

  { path: '', redirectTo: '', pathMatch: 'full' },
  {
    path: 'home',
    component: HomePageAdminComponent,
    children: [
      {
        path: '',
        redirectTo: 'user',
        pathMatch: 'full',
      },
      {
        path: 'users',
        component: AdminUserConrolComponent,
      },
      {
        path: 'projects',
        component: AdminProjectConrolComponent,
      },
      {
        path: 'news',
        component: AdminNewsConrolComponent,
      },
    ]
  }
];


@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AdminRoutingModule { }
