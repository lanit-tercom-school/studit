import { Component, OnInit } from '@angular/core';
import {Input} from "@angular/core";
import {TasksItem} from "./tasks-item/tasks-item";

@Component({
  selector: 'app-tasks',
  templateUrl: './tasks.component.html',
  styleUrls: ['./tasks.component.css']
})
export class TasksComponent implements OnInit {
  @Input() public TasksItemList;
  constructor() { }

  ngOnInit() {
  }

}
