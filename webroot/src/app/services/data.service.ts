import { Injectable, EventEmitter } from '@angular/core';
import { ApiService } from './api.service';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';
import { NewsItem } from "models/news-item";

@Injectable()
export class DataService {
  // Data
  public Projects;
  public UsersProjects;
  public EnrolledUsersProject;
  public UserId: number;
  // Events and boolean that show what is loaded.
  public ProjectsUploaded: EventEmitter<any> = new EventEmitter();
  public UsersProjectsUploaded: EventEmitter<any> = new EventEmitter();
  public NewsUploaded: EventEmitter<any> = new EventEmitter();
  public EnrolledUsersProjectUploaded: EventEmitter<any> = new EventEmitter();

  private news: BehaviorSubject<NewsItem[]> = <BehaviorSubject<NewsItem[]>>new BehaviorSubject([]);
  private dataStore: {
    news: NewsItem[];
  } =  { news: [] };
  public get News(){
    return this.news.asObservable();
  }

  private projectsUploaded = false;
  private usesProjectsUploaded = false;
  private newsUploaded = false;
  private enrolledUsersProjectUploaded: boolean;
  // Functions for testing is data uploaded?
  isUsersProjectsUploaded() {
    if (this.usesProjectsUploaded) { return true; } else { return false; }
  }
  isProjectsUploaded() {
    if (this.projectsUploaded) { return true; } else { return false; }
  }
  isErolledUsersProjectUploaded() {
    if (this.enrolledUsersProjectUploaded) { return true; } else { return false; }
  }
  isNewsUploaded() {
    if (this.newsUploaded) { return true; } else { return false; }
  }
  // Constructor.
  constructor(private api: ApiService) { }
  // Load data functions.
  loadAll() {
    console.log('Data.service ->loadAll');
    this.loadProjects();
    this.loadNews();
    if (localStorage.getItem('current_user')) {
      this.UserId = JSON.parse(localStorage.getItem('current_user')).id;
      this.loadUsersProjects();
      this.loadEnrolledUsersProject();
    }
  }
  // Load from server functions
  loadProjects() {
    this.api.getProjectItems().subscribe(res => {
      if (res != null) {
        this.Projects = res;
        this.projectsUploaded = true;
        this.ProjectsUploaded.emit();
      }
    });
  }
  
  loadUsersProjects() {
    if (localStorage.getItem('current_user')) {
      this.UsersProjects = new Array<number>();
      this.api.getProjectsOfUser(this.UserId).subscribe(res => {
        for (let i = 0; i < res.json().length; i++) {
          this.UsersProjects.push(res.json()[i].id);
        }
        this.usesProjectsUploaded = true;
        this.UsersProjectsUploaded.emit();
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
      this.dataStore.news = res;
      this.news.next(Object.assign({}, this.dataStore).news);
      this.newsUploaded = true;
      this.NewsUploaded.emit();
    });
  }
  //Get-functions
  

  getProjectById(id: number) {
    if (this.projectsUploaded) {
      for (let i = 0; i < this.Projects.length; i++) {
        if (+this.Projects[i].id === id) {
          return this.Projects[i];
        }
      }
    }
  }
}
