import { Component, OnInit, OnDestroy } from '@angular/core';
import { Observable } from "rxjs/Observable";

import { ProjectItem } from 'models/project-item';
import { ApiService } from 'services/api.service';
import { DataService } from 'services/data.service';

@Component({
  selector: 'app-home-projects-view',
  templateUrl: './home-projects-view.component.html',
  styleUrls: ['./home-projects-view.component.css']
})
export class HomeProjectsViewComponent implements OnInit, OnDestroy {
  private userId: string;
  private ProjectList: Observable<ProjectItem[]>;

  constructor(private api: ApiService, private data: DataService) { }

  ngOnInit() {
    this.ProjectList = this.data.UserProjects;
  }
  ngOnDestroy() {

  }
}
