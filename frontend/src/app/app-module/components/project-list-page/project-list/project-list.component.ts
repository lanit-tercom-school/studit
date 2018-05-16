import { Component, Input, OnInit } from '@angular/core';
import { Observable } from "rxjs/Observable";

import { DataService } from 'services/data.service';
import { ProjectItem } from 'models/project-item';

@Component({
  selector: 'app-project-list',
  templateUrl: './project-list.component.html',
  styleUrls: ['./project-list.component.css']
})
export class ProjectListComponent implements OnInit {

  public ProjectList: Observable<ProjectItem[]>;
  public CurrentPage: number = 1;
  public Limit: number = 2;
  public TotalCount: Observable<number>;

  public Loading: boolean;
  constructor(private data: DataService) { }

  ngOnInit() {
    this.TotalCount = this.data.ProjectCountObs;
    this.getPage(1);
  }

  getPage(page: number) {
    this.Loading = true;
    let offset = 0;
    if (page > 1)
      offset = (page - 1) * this.Limit;
    this.data.loadProjects(this.Limit, offset);
    this.ProjectList = this.data.Projects;
    this.CurrentPage = page;
    this.Loading = false;
    window.scrollTo(0, 0);
  }
}
