import { Injectable, EventEmitter } from '@angular/core';
import { ApiService } from './api.service';

@Injectable()
export class DataService {
  // Data
  private projects;
  private news;
  private usersProjects;
  private enrolledUsersProject;
  private userId: number;
  // Events and boolean that show what is loaded.
  public projectsUploaded: EventEmitter<any> = new EventEmitter();
  private _projectsUploaded = false;
  public usersProjectsUploaded: EventEmitter<any> = new EventEmitter();
  private _usesProjectsUploaded = false;
  public newsUploaded: EventEmitter<any> = new EventEmitter();
  private _newsUploaded = false;
  public enrolledUsersProjectUploaded: EventEmitter<any> = new EventEmitter();
  private _enrolledUsersProjectUploaded: boolean;
  // Functions for testing is data uploaded?
  isUsersProjectsUploaded() {
    if (this._usesProjectsUploaded) { return true; } else { return false; }
  }
  isProjectsUploaded() {
    if (this._projectsUploaded) { return true; } else { return false; }
  }
  isErolledUsersProjectUploaded() {
    if (this._enrolledUsersProjectUploaded) { return true; } else { return false; }
  }
  isNewsUploaded() {
    if (this._newsUploaded) { return true; } else { return false; }
  }
  // Constructor.
  constructor(private api: ApiService) { }
  // Load data functions.
  loadAll() {
    console.log('Data.service ->loadAll');
    this.loadProjects();
    this.loadNews();
    if (localStorage.getItem('current_user')) {
      this.userId = JSON.parse(localStorage.getItem('current_user')).id;
      this.loadUsersProjects();
      this.loadEnrolledUsersProject();
    }
  }
  // Load from server functions
  loadProjects() {
    this.api.getProjectItems().subscribe(res => {
      if (res != null) {
        this.projects = res;
        this._projectsUploaded = true;
        this.projectsUploaded.emit();
      }
    });
  }
  
  loadUsersProjects() {
    if (localStorage.getItem('current_user')) {
      this.usersProjects = new Array<number>();
      this.api.getProjectsOfUser(this.userId).subscribe(res => {
        for (let i = 0; i < res.json().length; i++) {
          this.usersProjects.push(res.json()[i].id);
        }
        this._usesProjectsUploaded = true;
        this.usersProjectsUploaded.emit();
      });
    } else {
      console.log('Error in data.service: can not load usersProject without auth');
    }
  }
  loadEnrolledUsersProject() {
    if (localStorage.getItem('current_user')) {

    } else {
      console.log('Error in data.service: can not load enrolledUsersProject without auth');
    }
  }
  loadNews() {
    this.api.getNewsPage().subscribe(res => {
      this.news = res;
      this._newsUploaded = true;
      this.newsUploaded.emit();
    });
  }
  //Get-functions
  getUsersProjects() {
    return this.usersProjects;
  }
  getEnrolledUsersProject() {
    return this.enrolledUsersProject;
  }
  getProjects() {
    return this.projects;
  }
  getNews(){
    return this.news;
  }
  getProjectById(id: number) {
    if (this._projectsUploaded) {
      for (let i = 0; i < this.projects.length; i++) {
        if (+this.projects[i].id === id) {
          return this.projects[i];
        }
      }
    }
  }
}
