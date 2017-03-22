///<reference path="proj-news/proj-news-item/proj-news-item.ts"/>
import { Component, OnInit } from '@angular/core';
import { ApiService } from './../../../services/api.service';
import { MaterialsItem } from './materials/materials-item/materials-item';
import { ActivatedRoute, Params } from '@angular/router';
import { ProjectItem } from './../../shared/project-list/project-item/project-item';
import { ProjectNewsItem } from './proj-news/proj-news-item/proj-news-item';
import { TasksItem } from "./tasks/tasks-item/tasks-item";

@Component({
  selector: 'app-student-project-page',
  templateUrl: './student-project-page.component.html',
  styleUrls: ['./student-project-page.component.css']
})
export class SProjectPageComponent implements OnInit {

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

  getMaterialsItems(): MaterialsItem[] {
    return this.apiService.getMaterialsItems(1);
  }
  getProjectNewsItem(): ProjectNewsItem[] {
    return this.apiService.getProjectNewsItem(1);
  }
  getTaskItem(): TasksItem[] {
    return this.apiService.getTaskItem(1);
  }
}
