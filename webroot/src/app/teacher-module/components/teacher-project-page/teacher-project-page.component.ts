import { Component, OnInit, OnChanges, ViewContainerRef, OnDestroy } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { TeacherService } from 'services/teacher.service';
import { DataService } from 'services/data.service';
import { ProjectService } from 'services/project.service';

import { MaterialsItem } from 'models/materials-item';
import { ProjectItem } from 'models/project-item';
import { ProjectNewsItem } from 'models/proj-news-item';
import { TasksItem } from 'models/tasks-item';

import 'rxjs/add/operator/filter';


@Component({
  selector: 'app-teacher-project-page',
  templateUrl: './teacher-project-page.component.html',
  styleUrls: ['./teacher-project-page.component.css']
})
export class TeacherProjectPageComponent implements OnInit, OnDestroy {

  private projectObs: BehaviorSubject<ProjectItem> = new BehaviorSubject({
    id: 0, name: "Loading...", description: "Loading...", logo: "dsasda"
  });
  private projectId;
  private authorized = false;
  private isTeacher = false;
  private isSuccess = false;
  private tasks = [];
  private message = 'Please write back soon!';
  private enrollButtonStatus: number = 0;//0 - enrolling,1 - you are in project, 2 - unenrolling
  constructor(private teacherService: TeacherService,
    private route: ActivatedRoute,
    private http: Http,
    private data: DataService,
     private projectService: ProjectService
  )
  { }

  ngOnInit() {
    if (localStorage.getItem('current_user')) { this.authorized = true; }
    this.route.params
      .subscribe(params => {
        this.projectId = params['id'];
        this.getProjectInfo();
        this.choseButtonStatus();
      });
    this.getTaskItems();
    if (this.data.PermLvl === 1)
      this.isTeacher = true;

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
    return this.projectService.getMaterialsItems(1);
  }
  getProjectNewsItem(): ProjectNewsItem[] {
    return this.projectService.getProjectNewsItem(1);
  }
  getTaskItems() {
    this.http.get('https://api.github.com/repos/lanit-tercom-school/studit/issues')
      .map((response: Response) => {
        let res = response.json().slice(0, 4);
        return res;
      }).subscribe(res => this.tasks = res);
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
