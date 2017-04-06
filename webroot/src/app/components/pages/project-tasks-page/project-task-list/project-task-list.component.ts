import { Component, OnInit } from '@angular/core';
import {Input} from "@angular/core";
import {ProjectTaskItem} from "models/project-task-item";

@Component({
  selector: 'app-project-task-list',
  templateUrl: './project-task-list.component.html',
  styleUrls: ['./project-task-list.component.css']
})
export class ProjectTaskListComponent implements OnInit {
  @Input() public ProjectTaskItemList:ProjectTaskItem[];
  constructor() { }

  ngOnInit() {

  }

}
