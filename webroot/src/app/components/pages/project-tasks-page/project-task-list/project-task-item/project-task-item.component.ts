import { Component, OnInit } from '@angular/core';
import {ProjectTaskItem} from "models/project-task-item";
import {Input} from "@angular/core";

@Component({
  selector: 'app-project-task-item',
  templateUrl: './project-task-item.component.html',
  styleUrls: ['./project-task-item.component.css']
})
export class ProjectTaskItemComponent implements OnInit {
  @Input() public ProjectTaskItem;
  constructor() { }

  ngOnInit() {
  }

}
