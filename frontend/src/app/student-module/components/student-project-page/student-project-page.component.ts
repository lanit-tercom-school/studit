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
import { UserInfo } from 'models/user-info';

type StatusEnroll = "Enrolling" | "InProject" | "Unenrolling";

@Component({
  selector: 'app-student-project-page',
  templateUrl: './student-project-page.component.html',
  styleUrls: ['./student-project-page.component.css']
})
export class StudentProjectPageComponent implements OnInit, OnDestroy {

  public ProjectObs: BehaviorSubject<ProjectItem> = new BehaviorSubject(new ProjectItem());
  public TasksObs: Observable<TasksItem[]>;
  public ProjectUsers: BehaviorSubject<UserInfo[]> = new BehaviorSubject<UserInfo[]>([]);
  public EnrollMessage = 'Please write back soon!';
  public EnrollButtonStatus = "Enrolling";


  private projectId: number;
  private projectEnrollId: number;
  private authorized = false;
  private isTeacher = false;
  constructor(private route: ActivatedRoute,
    private http: Http,
    private data: DataService,
    private studentService: StudentService,
    private projectService: ProjectService,
    private taskService: TaskService,
  ) { }

  ngOnInit() {
    if (localStorage.getItem('current_user')) { this.authorized = true; }
    this.route.params.subscribe(p => {
      this.projectId = +p['id'];
      this.getProjectInfo();
      this.choseButtonStatus();
    });
    let a: UserInfo[] = [];
    a.push(
      {
        Avatar: "a",
        Description: "dd",
        Id: "22",
        Nickname: "Happy penguin"
      }
    );
    a.push(
      {
        Avatar: "a",
        Description: "dd",
        Id: "22",
        Nickname: "Sad eagle"
      }
    );
    a.push(
      {
        Avatar: "a",
        Description: "dd",
        Id: "22",
        Nickname: "Indifferent boa"
      }
    );

    this.ProjectUsers.next(a)
  }

  ngOnDestroy() {
  }

  getProjectTasks(gitHubUrl: string) {
    this.data.loadTaskByGitHubUrl(gitHubUrl);
    this.TasksObs = this.data.TasksForViewing;
  }

  getProjectInfo() {
    this.data.loadProjectByID(this.projectId);
    this.data.ProjectForViewing.subscribe(res => {
      if (res != null) {
        this.ProjectObs.next(res);
        this.getProjectTasks(res.GitHubUrl)
      }
    });

  }


  enroll() {
    this.studentService.enrollToProject(this.data.UserId, this.projectId,
      JSON.parse(localStorage.getItem('current_user')).Token, this.EnrollMessage)
      .subscribe(res => {
        this.EnrollButtonStatus = "Unenrolling";
        this.data.loadEnrolledUsersProject();
      });
  }
  unenroll() {
    this.studentService.unenrollToProject(this.projectEnrollId,
      JSON.parse(localStorage.getItem('current_user')).Token).subscribe(res => {
        this.EnrollButtonStatus = "Enrolling";
        this.data.loadEnrolledUsersProject();
      });
  }

  choseButtonStatus() {
    this.EnrollButtonStatus = "Enrolling";
    this.data.UserProjects.subscribe(res => {
      if (res != null) {
        if (res.find(pr => pr.Id == this.projectId)) {
          this.EnrollButtonStatus = "InProject";
        }
      }
    });
    this.data.UserEnrolledProjects.subscribe(res => {
      if (res != null && res.find(pr => pr.Project.Id == this.projectId)) {
        this.EnrollButtonStatus = "Unenrolling";
        res.forEach(p => {
          if (p.Project.Id == this.projectId) {
            this.projectEnrollId = p.Id;
          }
        })
      }
    })
  }
} 