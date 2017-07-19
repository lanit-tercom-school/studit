import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { TeacherRoutingModule } from './teacher-routing.module';

import { HomePageTeacherComponent } from './components/home-page-teacher/home-page-teacher.component'
import { CreateProjectPageComponent } from './components/create-project-page/create-project-page.component';
import { TeacherNotePageComponent } from './components/teacher-notification-page/teacher-notification-page.component';
import { HomeTeacherProjectViewComponent } from './components/home-teacher-project-view/home-teacher-project-view.component'
import { TeacherProjectPageComponent } from './components/teacher-project-page/teacher-project-page.component';
import { TeacherSettingsPageComponent } from './components/teacher-settings-page/teacher-settings-page.component';
import { TeacherPublicPageComponent } from './components/teacher-public-page/teacher-public-page.component';

@NgModule({
  imports: [
    CommonModule,
    TeacherRoutingModule,
    FormsModule
  ],
  declarations: [
    HomePageTeacherComponent,
    CreateProjectPageComponent,
    TeacherNotePageComponent,
    HomeTeacherProjectViewComponent,
    TeacherProjectPageComponent,
    TeacherSettingsPageComponent,
    TeacherPublicPageComponent
  ]
})
export class TeacherModule { }
