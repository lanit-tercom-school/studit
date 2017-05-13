import { Injectable, EventEmitter } from '@angular/core';
import { ApiService } from './api.service'

@Injectable()
export class DataService {
  //Events and boolean that show what is loaded
  public projectsUploaded: EventEmitter<any>;
  private _projectsUploaded: boolean;
  public usersprojectUploaded: EventEmitter<any>;
  private _usesprojectUploaded: boolean;
  public newsUploaded: EventEmitter<any>;
  private _newsUploaded: boolean;
  //Data
  private projects;
  private news;
  private usersProjects;
  private enrolledUsersProject;
  private userId: number;
  //Constructor.
  constructor(private api: ApiService) { }
  //Load data function.
  loadEnrollingAndUserProjects() {
    this.projects = new Array<number>();
    this.enrolledUsersProject = new Array<number>();
    this.userId = JSON.parse(localStorage.getItem('current_user')).id;
    this.api.getProjectItems().subscribe(res => {
      console.log(res);
      this.projects = res;
      for (let i = 0; i < this.projects.length; i++) {
        this.api.getEnrolledUsersToProject(this.projects[i].id).subscribe(res => {
          if (res.json() != null) {
            for (let a of res.json()) {
              if (a === this.userId) {
                this.enrolledUsersProject.push(this.projects[i].id);
                break;
              }
            }
          }
        });
        this.api.getProjectUsers(this.projects[i].id).subscribe(res => {
          if (res.json() != null) {
            for (let a of res.json()) {
              if (a === this.userId) {
                this.projects.push(this.projects[i].id);
                break;
              }
            }
          }
        });
      }
    });
  }
  loadProjects() {
    this.api.postProject
  }
  getProjectsOfUser() {
    return this.projects;
  }
  getEnrollingProjectsOfUser() {
    return this.enrolledUsersProject;
  }
}
