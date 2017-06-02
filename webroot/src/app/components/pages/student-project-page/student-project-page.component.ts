import { Component, OnInit, OnChanges, DoCheck, OnDestroy } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { ApiService } from 'services/api.service';
import { DataService } from 'services/data.service';
import { MaterialsItem } from 'models/materials-item';
import { ProjectItem } from 'models/project-item';
import { ProjectNewsItem } from 'models/proj-news-item';
import { TasksItem } from 'models/tasks-item';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import 'rxjs/add/operator/filter';


@Component({
  selector: 'app-student-project-page',
  templateUrl: './student-project-page.component.html',
  styleUrls: ['./student-project-page.component.css']
})
export class StudentProjectPageComponent implements OnInit, OnDestroy {

  private projectObs: BehaviorSubject<ProjectItem> = new BehaviorSubject({
    id: 0, name: "Loading...", description: "Loading...", logo: "dsasda"
  });
  private projectId;
  private authorized = false;
  private tasks = [];
  private enrollButtonStatus: number;//0 - enrolling,1 - you are in project, 2 - unenrolling
  constructor(private apiService: ApiService, private route: ActivatedRoute, private http: Http, private data: DataService) { }

  ngOnInit() {
    if (localStorage.getItem('current_user')) { this.authorized = true; }
    this.route.params
      .subscribe(params => {
        this.projectId = params['id'];
        this.getProjectInfo();
        this.choseButtonStatus();
        console.log(this.enrollButtonStatus);
      });
    this.getTaskItems();
  }

  ngOnDestroy() {
  }

  getProjectInfo() {
    this.data.Projects.subscribe(projects => {
      if (projects.find(res => res.id == this.projectId)) {
        this.projectObs.next(projects.find(res => res.id == this.projectId));
      }
      else {
      }
    });
  }

  getMaterialsItems(): MaterialsItem[] {
    return this.apiService.getMaterialsItems(1);
  }
  getProjectNewsItem(): ProjectNewsItem[] {
    return this.apiService.getProjectNewsItem(1);
  }
  getTaskItems() {
    this.http.get('https://api.github.com/repos/lanit-tercom-school/studit/issues')
      .map((response: Response) => {
        let res = response.json().slice(0, 4);
        return res;
      }).subscribe(res => this.tasks = res);
  }
  enroll() {
    this.apiService.enrollToProject(this.projectId, JSON.parse(localStorage.getItem('current_user')).bearer_token).subscribe(res => { });
    this.enrollButtonStatus = 2;
    this.data.loadEnrolledUsersProject();
  }
  unenroll() {
    this.apiService.unenrollToProject(this.projectId, JSON.parse(localStorage.getItem('current_user')).bearer_token).subscribe(res => { });
    this.enrollButtonStatus = 0;
    this.data.loadEnrolledUsersProject();
  }
  choseButtonStatus() {
    this.data.UserProjects.subscribe(res => {
      for (let i = 0; i < res.length; i++) {
        if (res[i].id === +this.projectId) {
          this.enrollButtonStatus = 1;
          break;
        }
      }
    });

  }
}
