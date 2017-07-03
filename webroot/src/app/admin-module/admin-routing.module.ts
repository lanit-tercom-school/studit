import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { HomePageAdminComponent } from './components/home-page-admin/home-page-admin.component';
import { ProjectsViewAdminComponent } from './components/projects-view-admin/projects-view-admin.component';
import { UsersViewAdminComponent } from './components/users-view-admin/users-view-admin.component';
const routes: Routes = [

  { path: '', redirectTo: '', pathMatch: 'full' },
  {
    path: 'home',
    component: HomePageAdminComponent,
    children: [
      {
        path: '',
        redirectTo: 'projects',
        pathMatch: 'full',
      },
      {
        path: 'projects',
        component: ProjectsViewAdminComponent,
      },
      {
        path: 'users',
        component: UsersViewAdminComponent,
      }
    ]
  }
];


@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AdminRoutingModule { }
