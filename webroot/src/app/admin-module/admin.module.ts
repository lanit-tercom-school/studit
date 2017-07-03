import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { AdminRoutingModule } from './admin-routing.module';
import { HomePageAdminComponent } from './components/home-page-admin/home-page-admin.component';
import { ProjectsViewAdminComponent } from './components/projects-view-admin/projects-view-admin.component';
import { UsersViewAdminComponent } from './components/users-view-admin/users-view-admin.component';

@NgModule({
  imports: [
    CommonModule,
    AdminRoutingModule,
    FormsModule
  ],
  declarations: [HomePageAdminComponent, ProjectsViewAdminComponent, UsersViewAdminComponent]
})
export class AdminModule { }
