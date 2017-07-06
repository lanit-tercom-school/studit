import { Component, OnInit, OnDestroy } from '@angular/core';
import { Observable } from "rxjs/Observable";

import { ProjectItem } from 'models/project-item';
import { ApiService } from 'services/api.service';
import { DataService } from 'services/data.service';
import { EnrollItem } from 'models/enroll-item';

@Component({
  selector: 'app-home-teacher-project-view',
  templateUrl: './home-teacher-project-view.component.html',
  styleUrls: ['./home-teacher-project-view.component.css']
})
export class HomeTeacherProjectViewComponent implements OnInit {
  private ProjectList: Observable<ProjectItem[]>;
  private ProjectEnrollList: Observable<ProjectItem[]>;
  private EnrollList: EnrollItem[];

  constructor(private api: ApiService, private data: DataService) { }

  ngOnInit() {
    this.ProjectList = this.data.UserProjects;
    this.ProjectEnrollList = this.data.UserEnrolledProjects;
    this.api.getEnrollsForTeacher(JSON.parse(window.localStorage.getItem("current_user")).bearer_token).subscribe(res => {
      console.log(res);
      this.EnrollList = res;
    });
  }
  ngOnDestroy() {

  }
}
