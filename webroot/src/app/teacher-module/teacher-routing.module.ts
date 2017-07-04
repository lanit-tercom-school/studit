import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { HomePageTeacherComponent } from './components/home-page-teacher/home-page-teacher.component';
import { HomeProjectsViewComponent } from 'shared-components/home-projects-view/home-projects-view.component';
import { CreateProjectPageComponent } from 'shared-components/create-project-page/create-project-page.component';
import { ProjectViewTeacherComponent } from './components/project-view-teacher/project-view-teacher.component';

const routes: Routes = [

  { path: '', redirectTo: 'home', pathMatch: 'full' },
  { path: 'createproject', component: CreateProjectPageComponent, },
  {
    path: 'home',
    component: HomePageTeacherComponent,
    children: [
      {
        path: '',
        redirectTo: 'projects',
        pathMatch: 'full',
      },
      {
        path: 'projects',
        component: ProjectViewTeacherComponent,
      },
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class TeacherRoutingModule { }
