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
   private p: number = 1;
  private limit: number = 2;
  private totalObs: Observable<number>;

  private loading: boolean;
  constructor(private data: DataService) { }

  ngOnInit() {
    this.totalObs = this.data.ProjectCountObs;
    this.getPage(1);
  }

  getPage(page: number) {
    this.loading = true;
    let offset = 0;
    if (page > 1)
      offset = (page - 1) * this.limit;
    console.log('offset' + offset)
    this.data.loadProjects(this.limit, offset);
   // console.log('on page' + this.data.NumberOfProjectsOnPage);
    this.ProjectList = this.data.Projects;
   // console.log('total '+ this.totalObs);
    this.p = page;
    this.loading = false;
    window.scrollTo(0, 0);
  }
}
