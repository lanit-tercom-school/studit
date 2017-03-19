import { Component, OnInit } from '@angular/core';
import {Input} from "@angular/core/src/metadata/directives";
import {TasksItem} from "./tasks-item/tasks-item";

@Component({
  selector: 'app-tasks',
  templateUrl: './tasks.component.html',
  styleUrls: ['./tasks.component.css']
})
export class TasksComponent implements OnInit {
  @Input() public TasksItemList:TasksItem[];
  constructor() { }

  ngOnInit() {
  }

}
