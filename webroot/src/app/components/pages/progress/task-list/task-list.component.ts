import { Component, OnInit } from '@angular/core';
import {ProgressService} from "../progress.service";
import {TaskItem} from "../task";

@Component({
  selector: 'app-task-list',
  templateUrl: './task-list.component.html',
  styleUrls: ['./task-list.component.css']
})
export class TaskListComponent implements OnInit {

  constructor(private progressService: ProgressService) { }

  taskItems: TaskItem[];

  ngOnInit() {
    this.taskItems = this.progressService.getTasks();
  }

}
