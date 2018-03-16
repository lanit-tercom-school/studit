import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { TeacherSettingsPageComponent } from './components/teacher-settings-page/teacher-settings-page.component';

import { HomePageTeacherComponent } from './components/home-page-teacher/home-page-teacher.component';
import { HomeTeacherProjectViewComponent } from './components/home-teacher-project-view/home-teacher-project-view.component'
import { CreateProjectPageComponent } from './components/create-project-page/create-project-page.component';
import { TeacherNotePageComponent } from './components/teacher-notification-page/teacher-notification-page.component';
import { TeacherProjectPageComponent } from './components/teacher-project-page/teacher-project-page.component';
import { TeacherPublicPageComponent } from './components/teacher-public-page/teacher-public-page.component';
import { HomeTeacherEnrollingPageComponent } from './components/home-teacher-enrolling-page/home-teacher-enrolling-page.component';


const routes: Routes = [

  { path: '', redirectTo: 'home', pathMatch: 'full' },
  { path: 'createproject', component: CreateProjectPageComponent, },
  { path: 'project/:id', component: TeacherProjectPageComponent, },
  { path: 'profile', component: TeacherPublicPageComponent, },
  { path: 'profile/settings', component: TeacherSettingsPageComponent },
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
      {
        path: 'enrollings',
        component: HomeTeacherEnrollingPageComponent,
      }
    ]
  },

  { path: 'notifications', component: TeacherNotePageComponent, },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class TeacherRoutingModule { }
