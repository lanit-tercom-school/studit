import { Component, OnInit, Input } from '@angular/core';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { TasksItem } from "models/tasks-item";

@Component({
  selector: 'app-teacher-project-tasks',
  templateUrl: './teacher-project-task.component.html',
  styleUrls: ['./teacher-project-task.component.css']
})
export class TeacherProjectTasksComponent  {
  @Input() public TasksItemList: BehaviorSubject<TasksItem[]>;
}
