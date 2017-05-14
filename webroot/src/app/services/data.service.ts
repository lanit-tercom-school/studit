import { Injectable, EventEmitter } from '@angular/core';
import { ApiService } from './api.service'

@Injectable()
export class DataService {
  //Events and boolean that show what is loaded.
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
  //Load data functions.
  loadAll() {
    this.loadProjects();
    if (localStorage.getItem('current_user')) {
      this.loadUsersProjects();
    }
  }
  //Load projects =)
  loadProjects() {
    this.api.getProjectItems().subscribe(res => {
      if (res != null) {
        this.projects = res;
        this.projectsUploaded.emit();
        this._projectsUploaded = true;
      }
    })
  }
  //load information about user`s projects
  loadUsersProjects() {
    if (localStorage.getItem('current_user')) {
      this.usersProjects = new Array<number>();
      this.api.getProjectsOfUser(this.userId).subscribe(res => {
        for (let i = 0; i < res.json().length; i++) {
          this.usersProjects.push(res.json()[i].id);
        }
        this._usesprojectUploaded = true;
        this.usersprojectUploaded.emit();
      });
    }
    else {
      console.log('Error in data.service: can not load usersProject without auth');
    }
  }
  loadEnrolledUsersProject() {
    if (localStorage.getItem('current_user')) {

    }
    else {
      console.log('Error in data.service: can not load enrolledUsersProject without auth');
    }
  }
  getProjectsOfUser() {
    return this.projects;
  }
  getEnrollingProjectsOfUser() {
    return this.enrolledUsersProject;
  }
}
