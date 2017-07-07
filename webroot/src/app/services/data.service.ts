import { Injectable, EventEmitter } from '@angular/core';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';
import { Router } from '@angular/router';

import { ApiService } from 'services/api.service';
import { NewsItem } from "models/news-item";
import { ProjectItem } from 'models/project-item';
import { environment } from '../../environments/environment'

import 'rxjs/add/operator/filter';

@Injectable()
export class DataService {
  private userId: number;
  private userToken: string;
  private news: BehaviorSubject<NewsItem[]> = <BehaviorSubject<NewsItem[]>>new BehaviorSubject([]);
  private projects: BehaviorSubject<ProjectItem[]> = <BehaviorSubject<ProjectItem[]>>new BehaviorSubject([]);
  private userProjects: BehaviorSubject<ProjectItem[]> = <BehaviorSubject<ProjectItem[]>>new BehaviorSubject([]);
  private userEnrolledProjects: BehaviorSubject<ProjectItem[]> = <BehaviorSubject<ProjectItem[]>>new BehaviorSubject([]);
  private projectsForMainPage: BehaviorSubject<ProjectItem[]> = <BehaviorSubject<ProjectItem[]>>new BehaviorSubject([]);

  private dataStore: {
    news: NewsItem[];
    projects: ProjectItem[];
    userProjects: ProjectItem[];
    userEnrolledProjects: ProjectItem[];
    projectsForMainPage: ProjectItem[];
  } = { news: [], projects: [], userProjects: [], userEnrolledProjects: [], projectsForMainPage: [], };

  public get News() {
    return this.news.asObservable();
  }
  public get Projects() {
    return this.projects.asObservable();
  }
  public get UserProjects() {
    return this.userProjects.asObservable();
  }
  public get UserEnrolledProjects() {
    return this.userEnrolledProjects.asObservable();
  }
  public get ProjectsForMainPage() {
    return this.projectsForMainPage.asObservable();
  }
  constructor(private api: ApiService, private router: Router) { }

  loadAll() {
    this.loadProjects();
    this.loadNews();
    this.loadProjectsForMainPage();
    if (localStorage.getItem('current_user')) {
      this.userId = JSON.parse(localStorage.getItem('current_user')).user.id;
      this.loadUsersProjects();
      this.loadEnrolledUsersProject();
    }
    console.debug('Data.service ->loadAll');
  }

  loadProjects() {
    this.api.getProjectItems().subscribe(res => {
      if (res != null) {
        this.dataStore.projects = res;
        this.dataStore.projects.forEach(a => { a.logo = this.addApiUrl(a.logo); })
        this.projects.next(Object.assign({}, this.dataStore).projects);
      }
    },
    error => {
      alert('Ошибка! ' + error.status + ' ' + error.statusText);
      console.debug('ERROR: status ' + error.status + ' ' + error.statusText);
      console.debug('ERROR: loadProjects() -> getProjectItems()');
    });
  }

  loadProjectsForMainPage() {
    this.api.getMainPageProjects().subscribe(res => {
      this.dataStore.projectsForMainPage = res;
      this.dataStore.projectsForMainPage.forEach(a => { a.logo = this.addApiUrl(a.logo); })
      this.projectsForMainPage.next(Object.assign({}, this.dataStore).projectsForMainPage);
    },
      error => {
      alert('Ошибка! ' + error.status + ' ' + error.statusText);
      console.debug('ERROR: status ' + error.status + ' ' + error.statusText);
      console.debug('ERROR: loadProjectsForMainPage() -> getMainPageProjects()');
    });
  }

  loadUsersProjects() {
    if (localStorage.getItem('current_user')) {
      this.api.getProjectsOfUser(this.userId).subscribe(res => {
        if (res != null) {
          this.dataStore.userProjects = res;
          this.dataStore.userProjects.forEach(a => { a.logo = this.addApiUrl(a.logo); })
          this.userProjects.next(Object.assign({}, this.dataStore).userProjects);
        }

      },
      error => {
      alert('Ошибка! ' + error.status + ' ' + error.statusText);
      console.debug('ERROR: status ' + error.status + ' ' + error.statusText);
      console.debug('ERROR: loadUsersProjects() -> getProjectsOfUser()');
      });
    } else {
      console.debug('ERROR: data.service: can not load usersProject without auth');
    }
  }

  loadEnrolledUsersProject() {
    if (localStorage.getItem('current_user')) {
      this.api.getEnrolledUsersProject(this.userId, this.userToken).subscribe(res => {
        this.dataStore.userEnrolledProjects = res;
        this.userEnrolledProjects.next(Object.assign({}, this.dataStore).userEnrolledProjects);
      },
       error => {
      alert('Ошибка! ' + error.status + ' ' + error.statusText);
      console.debug('ERROR: status ' + error.status + ' ' + error.statusText);
      console.debug('ERROR: loadEnrolledUsersProject() -> getEnrolledUsersProject()');
      });
    } else {
      console.debug('ERROR: data.service: can not load enrolledUsersProject without auth');
    }
  }

  loadNews() {
    this.api.getNewsPage().subscribe(res => {
      this.dataStore.news = res;
      this.news.next(Object.assign({}, this.dataStore).news);
    },
      error => {
      alert('Ошибка! ' + error.status + ' ' + error.statusText);
      console.debug('ERROR: status ' + error.status + ' ' + error.statusText);
      console.debug('ERROR: loadNews() -> getNewsPage()');
      });
  }

  addApiUrl(url: string): string {
  //return environment.apiUrl + url;
   return url;
  }
}
