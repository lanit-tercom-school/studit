import { Component, OnInit, Input } from '@angular/core';

import {TasksItem} from "models/tasks-item";

@Component({
  selector: 'app-tasks-item',
  templateUrl: './tasks-item.component.html',
  styleUrls: ['./tasks-item.component.css']
})
export class TasksItemComponent implements OnInit {
  @Input() public TasksItem;
  constructor() { }

  ngOnInit() {
  	this.TasksItem.body = this.TasksItem.body.slice(0, 100);
  }

}
