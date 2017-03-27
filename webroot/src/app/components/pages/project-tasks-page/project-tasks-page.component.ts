import { Component, OnInit } from '@angular/core';
import { ApiService } from './../../../services/api.service';
import { ActivatedRoute, Params } from '@angular/router';
import { ProjectItem } from './../../shared/project-list/project-item/project-item';
///import { ProjectTaskItem } from './proj-news/proj-news-item/proj-news-item';
import { ProjectTaskItem } from "./project-all-tasks/project-task-item/project-task-item";

@Component({
  selector: 'app-project-tasks-page',
  templateUrl: './project-tasks-page.component.html',
  styleUrls: ['./project-tasks-page.component.css']
})
export class ProjectTasksPageComponent implements OnInit {

private project;

constructor(private apiService: ApiService,
  private route: ActivatedRoute) { }

ngOnInit() {
  this.route.params
    .subscribe(params => {
    this.project = this.apiService.getProjectById(+params['id'])
      .subscribe(res => this.project = res.json());
    });
}

getProjectTaskItem(): ProjectTaskItem[] {
  return this.apiService.getProjectTaskItem(1);
}
}
