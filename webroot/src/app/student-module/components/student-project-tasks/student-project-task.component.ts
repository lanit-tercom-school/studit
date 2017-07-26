import { Component, OnInit, Input } from '@angular/core';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { TasksItem } from "models/tasks-item";

@Component({
  selector: 'app-student-project-tasks',
  templateUrl: './student-project-task.component.html',
  styleUrls: ['./student-project-task.component.css']
})
export class StudentProjectTasksComponent  {
  @Input() public TasksItemList: BehaviorSubject<TasksItem[]>;
}
