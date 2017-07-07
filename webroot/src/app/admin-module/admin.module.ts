import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { AdminRoutingModule } from './admin-routing.module';
import { HomePageAdminComponent } from './components/home-page-admin/home-page-admin.component';
import { AdminUserConrolComponent } from './components/admin-user-conrol/admin-user-conrol.component';
import { AdminProjectConrolComponent } from './components/admin-project-conrol/admin-project-conrol.component';
import { AdminNewsConrolComponent } from './components/admin-news-conrol/admin-news-conrol.component';


@NgModule({
  imports: [
    CommonModule,
    AdminRoutingModule,
    FormsModule
  ],
  declarations: [HomePageAdminComponent, AdminUserConrolComponent, AdminProjectConrolComponent, AdminNewsConrolComponent, ]
})
export class AdminModule { }
