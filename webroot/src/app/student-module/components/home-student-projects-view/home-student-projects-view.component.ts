import { Component, OnInit, OnDestroy } from '@angular/core';
import { Observable } from "rxjs/Observable";

import { ProjectItem } from 'models/project-item';
import { DataService } from 'services/data.service';

@Component({
  selector: 'app-home-student-projects-view',
  templateUrl: './home-student-projects-view.component.html',
  styleUrls: ['./home-student-projects-view.component.css']
})
export class HomeStudentProjectsViewComponent implements OnInit, OnDestroy {
  private userId: string;
  private ProjectList: Observable<ProjectItem[]>;
  private ProjectEnrollList: Observable<ProjectItem[]>;

  constructor( private data: DataService) { }

  ngOnInit() {
    this.ProjectList = this.data.UserProjects;
    this.ProjectEnrollList = this.data.UserEnrolledProjects;
  }
  ngOnDestroy() {

  }
}
