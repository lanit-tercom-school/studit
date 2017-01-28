import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProgressComponent } from './progress.component';
import { TaskListComponent } from './task-list/task-list.component';
import { ProjectCardListComponent } from './project-card-list/project-card-list.component';
import { NotificationListComponent } from './notification-list/notification-list.component';

@NgModule({
  imports: [
    CommonModule
  ],
  exports: [
    ProgressComponent,
  ],
  declarations: [ProgressComponent, TaskListComponent, ProjectCardListComponent, NotificationListComponent]
})
export class ProgressModule { }
