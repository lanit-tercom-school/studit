import { Component, OnInit } from '@angular/core';
import {Input} from "@angular/core/src/metadata/directives";
import {ProjectTaskItem} from "./project-task-item/project-task-item";

@Component({
  selector: 'app-project-all-tasks',
  templateUrl: './project-all-tasks.component.html',
  styleUrls: ['./project-all-tasks.component.css']
})
export class ProjectAllTasksComponent implements OnInit {
  @Input() public ProjectTaskItemList:ProjectTaskItem[];
  constructor() { }

  ngOnInit() {
  
  }

}
