import { Component, OnInit, OnChanges, DoCheck } from '@angular/core';
import { ApiService } from './../../../services/api.service';
import { DataService } from './../../../services/data.service';
import { MaterialsItem } from './materials/materials-item/materials-item';
import { ActivatedRoute, Params } from '@angular/router';
import { ProjectItem } from './../../shared/project-list/project-item/project-item';
import { ProjectNewsItem } from './proj-news/proj-news-item/proj-news-item';
import { TasksItem } from "./tasks/tasks-item/tasks-item";
import { Http, Headers, RequestOptions, Response } from '@angular/http';


@Component({
  selector: 'app-student-project-page',
  templateUrl: './student-project-page.component.html',
  styleUrls: ['./student-project-page.component.css']
})
export class StudentProjectPageComponent implements OnInit, DoCheck {

  private project;
  private projectId;
  private tasks = [];
  private subscribedUsers = [];
  private authorized: boolean;
  private enrollButtonStatus: number;//0 - enrolling,1 - you are in project, 2 - unenrolling
  constructor(private apiService: ApiService, private route: ActivatedRoute, private http: Http, private data: DataService) { }

  ngOnInit() {
    this.enrollButtonStatus = 3;
    this.route.params
      .subscribe(params => {
        this.projectId = params['id'];
        console.log(this.projectId);
        console.log(this.data.getProjectsOfUser());
        this.project = this.apiService.getProjectById(+params['id']).subscribe(res => this.project = res.json());
        if (this.data.getProjectsOfUser().indexOf(+this.projectId) !== -1) {
          this.enrollButtonStatus = 1;
        }
        else if (this.data.getEnrollingProjectsOfUser().indexOf(+this.projectId) !== -1) {
          this.enrollButtonStatus = 2;
        }
        else {
          this.enrollButtonStatus = 0;
        }
      });

    this.getTaskItems();
    if (localStorage.getItem('current_user')) {
      this.authorized = true;
    }
    else {
      this.authorized = false;
    }
  }
  ngDoCheck() {

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
    this.apiService.enrollToProject(this.projectId, JSON.parse(localStorage.getItem('current_user')).token).subscribe(res => { });
    this.enrollButtonStatus = 2;
    this.data.loadEnrollingAndUserProjects();
  }
  unenroll() {
    this.apiService.unenrollToProject(this.projectId, JSON.parse(localStorage.getItem('current_user')).token).subscribe(res => { });
    this.enrollButtonStatus = 0;
    this.data.loadEnrollingAndUserProjects();
  }
}
