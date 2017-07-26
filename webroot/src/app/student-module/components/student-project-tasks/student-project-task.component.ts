import { Component, OnInit, Input } from '@angular/core';

import { TasksItem } from "models/tasks-item";

@Component({
  selector: 'app-student-project-tasks',
  templateUrl: './student-project-task.component.html',
  styleUrls: ['./student-project-task.component.css']
})
export class StudentProjectTasksComponent implements OnInit {
  @Input() public TasksItemList;
  constructor() { }

  ngOnInit() {
  }

}
