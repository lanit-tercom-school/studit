import { Component, OnInit } from '@angular/core';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { ActivatedRoute, Params } from '@angular/router';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { ProjectService } from 'services/project.service';
import { StudentService } from 'services/student.service';
import { DataService } from 'services/data.service';

import { ProjectItem } from 'models/project-item';
import { ProjectTaskItem } from "models/project-task-item";

type StatusEnroll = "Enrolling" | "InProject" | "Unenrolling";

@Component({
  selector: 'app-student-project-page',
  templateUrl: './student-project-page.component.html',
  styleUrls: ['./student-project-page.component.css']
})
export class StudentProjectPageComponent implements OnInit {

  private projectObs: BehaviorSubject<ProjectItem> = new BehaviorSubject(null);
  private projectId;
  private isTeacher = false;
  private isSuccess = false;
  private tasks = [];
  private message = 'Please write back soon!';
  //0 - enrolling,1 - you are in project, 2 - unenrolling
  private enrollButtonStatus = "Enrolling";
  constructor(private route: ActivatedRoute,
    private http: Http,
    private data: DataService,
    private studentService: StudentService) { }

  ngOnInit() {
    this.route.params
      .subscribe(params => {
        this.projectId = params['id'];
        this.getProjectInfo();
        this.choseButtonStatus();
      });
    //this.getTaskItems();

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

  /*getMaterialsItems(): MaterialsItem[] {
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
  }*/
  enroll() {
    this.isSuccess = false;
    this.studentService.enrollToProject(this.projectId,
      JSON.parse(localStorage.getItem('current_user')).bearer_token, this.message).subscribe(res => {
        this.enrollButtonStatus = "Unenrolling";
        this.data.loadEnrolledUsersProject();
      });
    this.isSuccess = true;
  }
  unenroll() {
    this.isSuccess = false;
    this.studentService.unenrollToProject(this.projectId,
      JSON.parse(localStorage.getItem('current_user')).bearer_token).subscribe(res => {
        this.enrollButtonStatus = "Enrolling";
        this.data.loadEnrolledUsersProject();
      });
  }

  choseButtonStatus() {
    this.enrollButtonStatus = "Enrolling";
    this.data.UserProjects.subscribe(res => {
      if (res != null && res.find(pr => pr.id == this.projectId)) {
        this.enrollButtonStatus = "InProject";
      }
    })
    this.data.UserEnrolledProjects.subscribe(res => {
      if (res != null && res.find(pr => pr.id == this.projectId)) {
        this.enrollButtonStatus = "Unenrolling";
      }
    });
  }
}