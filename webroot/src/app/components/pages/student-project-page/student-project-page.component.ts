import { Component, OnInit, OnChanges, DoCheck } from '@angular/core';
import { ApiService } from './../../../services/api.service';
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
  private projectId: number;
  private tasks = [];
  private subscribedUsers = [];
  private authorized: boolean;
  private enrollButtonStatus: number;
  constructor(private apiService: ApiService,
    private route: ActivatedRoute, private http: Http) { }

  ngOnInit() {
    this.enrollButtonStatus = 0;
    this.route.params
      .subscribe(params => {
        this.projectId = params['id'];
        this.apiService.getProjectUsers(params['id']).subscribe(res => {
          this.subscribedUsers = res.json();
          if (this.subscribedUsers != null) {
            for (let a of this.subscribedUsers) {
              if (a === JSON.parse(localStorage.getItem('current_user')).id) {
                this.enrollButtonStatus = 1;
                break;
              }
            }
          }
        });
        this.apiService.getEnrolledUsersToProject(params['id']).subscribe(res => {
          this.subscribedUsers = res.json();
          if (this.subscribedUsers != null) {
            for (let a of this.subscribedUsers) {
              if (a === JSON.parse(localStorage.getItem('current_user')).id) {
                this.enrollButtonStatus = 2;
                break;
              }
            }
          }
        });
        this.project = this.apiService.getProjectById(+params['id']).subscribe(res => this.project = res.json());
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
  }
  unenroll() {
    this.apiService.unenrollToProject(this.projectId, JSON.parse(localStorage.getItem('current_user')).token).subscribe(res => { });
    this.enrollButtonStatus = 0;
  }
}
