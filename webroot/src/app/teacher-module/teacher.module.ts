import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { TeacherRoutingModule } from './teacher-routing.module';
import { HomePageTeacherComponent } from './components/home-page-teacher/home-page-teacher.component'
import { HomeProjectsViewComponent } from 'shared-components/home-projects-view/home-projects-view.component';
import { CreateProjectPageComponent } from 'shared-components/create-project-page/create-project-page.component';
import { TeacherNotePageComponent  } from './components/teacher-notification-page/teacher-notification-page.component'

@NgModule({
  imports: [
    CommonModule,
    TeacherRoutingModule,
    FormsModule
  ],
  declarations: [
    HomePageTeacherComponent,
    HomeProjectsViewComponent,
    CreateProjectPageComponent,
    TeacherNotePageComponent,
  ]
})
export class TeacherModule { }
