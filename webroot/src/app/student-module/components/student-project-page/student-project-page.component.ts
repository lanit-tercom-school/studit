import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { DataService } from 'services/data.service';
import { TaskService } from 'services/task.service';
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

  private projectObs: BehaviorSubject<ProjectItem> = new BehaviorSubject(new ProjectItem());
  private tasksObs: BehaviorSubject<TasksItem[]> = new BehaviorSubject(new Array<TasksItem>());


  private projectId: number;
  private projectEnrollId: number;
  private authorized = false;
  private isTeacher = false;
  private enrollMessage = 'Please write back soon!';
  private enrollButtonStatus = "Enrolling";
  constructor(private route: ActivatedRoute,
    private http: Http,
    private data: DataService,
    private studentService: StudentService,
    private projectService: ProjectService,
    private taskService : TaskService,
  ) { }

  ngOnInit() {
    if (localStorage.getItem('current_user')) { this.authorized = true; }
    this.route.params.subscribe(p => {
      this.projectId = +p['id'];
      this.getProjectInfo();
      this.choseButtonStatus();
    });
  }

  ngOnDestroy() {
  }

  getProjectTasks(gitHubUrl:string) {
    this.taskService.getTaskItemsFromGitHub(gitHubUrl).subscribe(res=>{
      this.tasksObs.next(res);
      });;
  }

  getProjectInfo() {
    this.data.loadProjectByID(this.projectId);
    this.data.ProjectForViewing.subscribe(res => {
      if (res != null){
        this.projectObs.next(res);
        this.getProjectTasks(res.GitHubUrl)
      }
    });

  }


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
      if (res != null) {
        if (res.find(pr => pr.Id == this.projectId)) {
          this.enrollButtonStatus = "InProject";
        }
      }
    });
    this.data.UserEnrolledProjects.subscribe(res => {
      if (res != null && res.find(pr => pr.Project.Id == this.projectId)) {
        this.enrollButtonStatus = "Unenrolling";
        res.forEach(p => {
          if (p.Project.Id == this.projectId) {
            this.projectEnrollId = p.Id;
          }
        })
      }
    })
  }
} 