import { Component, OnInit } from '@angular/core';
import {TasksItem} from "./tasks-item";
import {Input} from "@angular/core/src/metadata/directives";

@Component({
  selector: 'app-tasks-item',
  templateUrl: './tasks-item.component.html',
  styleUrls: ['./tasks-item.component.css']
})
export class TasksItemComponent implements OnInit {
  @Input() public TasksItem:TasksItem;
  constructor() { }

  ngOnInit() {
  }

}
