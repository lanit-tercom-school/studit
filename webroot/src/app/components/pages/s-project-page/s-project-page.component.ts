///<reference path="proj-news/proj-news-item/proj-news-item.ts"/>
import { Component, OnInit } from '@angular/core';
import { ApiService } from './../../../services/api.service';
import { MaterialsItem } from './materials/materials-item/materials-item';
import { ActivatedRoute, Params } from '@angular/router';
import { ProjectItem } from './../../shared/project-list/project-item/project-item';
import {ProjectNewsItem} from './proj-news/proj-news-item/proj-news-item';
import {TasksItem} from "./tasks/tasks-item/tasks-item";

@Component({
  selector: 'app-s-project-page',
  templateUrl: './s-project-page.component.html',
  styleUrls: ['./s-project-page.component.css']
})
export class SProjectPageComponent implements OnInit {

  project: ProjectItem;

  constructor(private apiService: ApiService,
    private route: ActivatedRoute) { }

  ngOnInit() {
    this.route.params
      .subscribe(params => { this.project = this.apiService.getProjectById(+params['id']); });
  }

  getMaterialsItems(): MaterialsItem[] {
    return this.apiService.getMaterialsItems(1);
  }
  getProjectNewsItem (): ProjectNewsItem [] {
    return this.apiService.getProjectNewsItem(1);
  }
  getTaskItem(): TasksItem[]{
    return this.apiService.getTaskItem(1);
  }
}
