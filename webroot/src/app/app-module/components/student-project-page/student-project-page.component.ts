import { Component, OnInit, OnChanges, DoCheck, OnDestroy } from '@angular/core';
import { ActivatedRoute, Router, Params } from '@angular/router';
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
    id: 0, name: "Loading...", description: "Loading...",
    logo: "https://web2.hirez.com/smite-esports/wp-content/uploads/2017/03/404.png"
  });
  private projectId;
  private authorized = false;
  private tasks = [];
  private enrollMessage = 'Please write back soon!';
  private enrollButtonStatus: number = 0;//0 - enrolling,1 - you are in project, 2 - unenrolling
  
  constructor(private apiService: ApiService,
    private data: DataService,
    private route: ActivatedRoute,
    private router: Router,
    private http: Http) { }

  ngOnInit() {
    if (localStorage.getItem('current_user')) { this.authorized = true; }
    this.route.params
      .subscribe(params => {
        this.projectId = params['id'];
        this.getProjectInfo();
        this.choseButtonStatus();
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
      else
       this.router.navigate(['/error']);
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
    this.apiService.enrollToProject(this.projectId,
      JSON.parse(localStorage.getItem('current_user')).bearer_token, this.enrollMessage).subscribe(res => {
        this.enrollButtonStatus = 2;
        this.data.loadEnrolledUsersProject();
      },
      error => {
        alert('Ошибка! ' + error.status + ' ' + error.statusText);
        console.debug('ERROR: status ' + error.status + ' ' + error.statusText);
        console.debug('ERROR: apiService: enrollToProject()');
      });
  }
  unenroll() {
    this.apiService.unenrollToProject(this.projectId,
      JSON.parse(localStorage.getItem('current_user')).bearer_token).subscribe(res => {
        this.enrollButtonStatus = 0;
        this.data.loadEnrolledUsersProject();
      },
      error => {
          alert('Ошибка! ' + error.status + ' ' + error.statusText);
          console.debug('ERROR: status ' + error.status + ' ' + error.statusText);
          console.debug('ERROR: apiService: unenrollToProject()');
        });
  }

  choseButtonStatus() {
    this.enrollButtonStatus = 0;
    this.data.UserProjects.subscribe(res => {
      if (res != null && res.find(pr => pr.id == this.projectId)) {
        this.enrollButtonStatus = 1;
      }
    })
    this.data.UserEnrolledProjects.subscribe(res => {
      if (res != null && res.find(pr => pr.id == this.projectId)) {
        this.enrollButtonStatus = 2;
      }
    })
  }
}
