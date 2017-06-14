import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { StudentRoutingModule } from './student-routing.module';

import { HomePageStudentComponent } from './components/home-page-student/home-page-student.component'
import { HomeProjectsViewComponent } from 'shared-components/home-projects-view/home-projects-view.component';


@NgModule({
  imports: [
    CommonModule,
    StudentRoutingModule,
    FormsModule,
  ],
  declarations: [
    HomeProjectsViewComponent,
    HomePageStudentComponent,
  ]
})
export class StudentModule { }
