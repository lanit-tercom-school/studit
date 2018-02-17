import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { StudentRoutingModule } from './student-routing.module';

import { TasksItemComponent } from "./components/student-project-tasks/tasks-item/tasks-item.component";
import { StudentProjectTasksComponent } from "./components/student-project-tasks/student-project-task.component";
import { HomePageStudentComponent } from './components/home-page-student/home-page-student.component'
import { HomeStudentProjectsViewComponent } from './components/home-student-projects-view/home-student-projects-view.component'
import { StudentProjectPageComponent } from './components/student-project-page/student-project-page.component';
import { StudentPublicPageComponent } from './components/student-public-page/student-public-page.component';
//import { ProjectTaskItemComponent } from './components/student-project-page/project-task-list/project-task-item/project-task-item.component';
//import { ProjectTaskListComponent } from './components/student-project-page/project-task-list/project-task-list.component';
import { StudentSettingsPageComponent } from './components/student-settings-page/student-settings-page.component';
@NgModule({
  imports: [
    CommonModule,
    StudentRoutingModule,
    FormsModule,
  ],
  declarations: [
    StudentProjectTasksComponent,
    TasksItemComponent,
    HomeStudentProjectsViewComponent,
    HomePageStudentComponent,
    StudentProjectPageComponent,
    StudentPublicPageComponent,
    StudentSettingsPageComponent
  ]
})
export class StudentModule { }
