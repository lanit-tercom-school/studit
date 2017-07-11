import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { StudentRoutingModule } from './student-routing.module';

import { HomePageStudentComponent } from './components/home-page-student/home-page-student.component'
import { HomeStudentProjectsViewComponent } from './components/home-student-projects-view/home-student-projects-view.component'


@NgModule({
  imports: [
    CommonModule,
    StudentRoutingModule,
    FormsModule,
  ],
  declarations: [
    HomeStudentProjectsViewComponent,
    HomePageStudentComponent,
  ]
})
export class StudentModule { }
