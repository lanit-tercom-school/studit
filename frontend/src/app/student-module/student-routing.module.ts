import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { StudentSettingsPageComponent } from './components/student-settings-page/student-settings-page.component';

import { HomePageStudentComponent } from './components/home-page-student/home-page-student.component'
import { HomeStudentProjectsViewComponent } from './components/home-student-projects-view/home-student-projects-view.component';
import { StudentProjectPageComponent } from './components/student-project-page/student-project-page.component';
import { StudentPublicPageComponent } from './components/student-public-page/student-public-page.component';
import { HomeStudentEnrollingPageComponent } from './components/home-student-enrolling-page/home-student-enrolling-page.component';

const routes: Routes = [
  { path: '', redirectTo: 'home', pathMatch: 'full' },
  { path: 'project/:id', component: StudentProjectPageComponent },
  { path: 'profile', pathMatch: 'full', component: StudentPublicPageComponent },
  { path: 'profile/settings', component: StudentSettingsPageComponent },
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
        component: HomeStudentProjectsViewComponent,
      },
      {
        path: 'enrollings',
        component: HomeStudentEnrollingPageComponent,
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class StudentRoutingModule { }
