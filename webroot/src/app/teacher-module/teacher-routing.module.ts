import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { HomePageTeacherComponent } from './components/home-page-teacher/home-page-teacher.component';
import { HomeTeacherProjectViewComponent } from './components/home-teacher-project-view/home-teacher-project-view.component'
import { CreateProjectPageComponent } from 'shared-components/create-project-page/create-project-page.component';
import { TeacherNotePageComponent } from './components/teacher-notification-page/teacher-notification-page.component';

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
        component: HomeTeacherProjectViewComponent,
      },
    ]
  },

  { path: 'notifications', component: TeacherNotePageComponent, },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class TeacherRoutingModule { }
