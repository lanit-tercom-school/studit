import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { DataService } from 'services/data.service';
import { ProjectService } from 'services/project.service';
import { StudentService } from 'services/student.service';

import { MaterialsItem } from 'models/materials-item';
import { ProjectItem } from 'models/project-item';
import { ProjectNewsItem } from 'models/proj-news-item';
import { TasksItem } from 'models/tasks-item';

type StatusEnroll = "Enrolling" | "InProject" | "Unenrolling";

@Component({
  selector: 'app-student-project-page',
  templateUrl: './student-project-page.component.html',
  styleUrls: ['./student-project-page.component.css']
})
export class StudentProjectPageComponent implements OnInit, OnDestroy {

  private projectObs: BehaviorSubject<ProjectItem> = new BehaviorSubject({
    Description: 'string',
    DateOfCreation: 'string',
    Logo: 'string',
    Tags: {},
    Id: 0,
    Name: 'string'
  });

  private projectId: number;
  private projectEnrollId: number;
  private authorized = false;
  private isTeacher = false;
  private tasks = [];
  private enrollMessage = 'Please write back soon!';
  private enrollButtonStatus = "Enrolling";
  constructor(private route: ActivatedRoute,
    private http: Http,
    private data: DataService,
    private studentService: StudentService,
    private projectService: ProjectService) { }

  ngOnInit() {
    if (localStorage.getItem('current_user')) { this.authorized = true; }
    this.route.params
      .subscribe(params => {
        this.projectId = params['id'];
        this.getProjectInfo();
        this.choseButtonStatus();
        //console.log(this.enrollButtonStatus);
      });
    //this.getTaskItems();

  }

  ngOnDestroy() {
  }

  getProjectInfo() {
    this.data.loadProjectByID(this.projectId);
    console.log('page: getProjectInfo');
    this.data.MissedProject.subscribe(res => {
      console.log(res);
      if (res != null)
        this.projectObs.next(res);
    });

  }

  /* getMaterialsItems(): MaterialsItem[] {
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
   }*/
  enroll() {
    this.studentService.enrollToProject(this.data.UserId, this.projectId,
      JSON.parse(localStorage.getItem('current_user')).Token, this.enrollMessage)
      .subscribe(res => {
        this.enrollButtonStatus = "Unenrolling";
        this.data.loadEnrolledUsersProject();
      });
  }
  unenroll() {
    this.studentService.unenrollToProject(this.projectEnrollId,
      JSON.parse(localStorage.getItem('current_user')).Token).subscribe(res => {
        this.enrollButtonStatus = "Enrolling";
        this.data.loadEnrolledUsersProject();
      });
  }

  choseButtonStatus() {
    this.enrollButtonStatus = "Enrolling";
    this.data.UserProjects.subscribe(res => {
      if (res != null && res.find(pr => pr.Id == this.projectId)) {
        this.enrollButtonStatus = "InProject";
      }
    })
    this.data.UserEnrolledProjects.subscribe(res => {
      if (res != null && res.find(pr => pr.Project.Id == this.projectId)) {
        this.enrollButtonStatus = "Unenrolling";
        res.forEach(p => {
          if (p.Project.Id == this.projectId){
            this.projectEnrollId=p.Id;
          }
        })
      }
    })
  }
}