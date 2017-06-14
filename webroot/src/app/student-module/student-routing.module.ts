import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { HomePageStudentComponent } from './components/home-page-student/home-page-student.component'
import { HomeProjectsViewComponent } from 'shared-components/home-projects-view/home-projects-view.component';

const routes: Routes = [
  { path: '', redirectTo: 'home', pathMatch: 'full' },
  {
    path: 'home',
    component: HomePageStudentComponent,
    children: [
      {
        path: '',
        redirectTo: 'projects',
        pathMatch: 'full',
      },
      {
        path: 'projects',
        component: HomeProjectsViewComponent,
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class StudentRoutingModule { }
