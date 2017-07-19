import { Injectable, EventEmitter } from '@angular/core';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { TeacherService } from 'services/teacher.service';
import { StudentService } from 'services/student.service';
import { ProjectService } from 'services/project.service';
import { NewsService } from 'services/news.service';
import { UserService } from 'services/user.service';

import { NewsItem } from "models/news-item";
import { ProjectItem } from 'models/project-item';
import { EnrollItem } from 'models/enroll-item';
import { environment } from '../../environments/environment';
import { PermLevel } from 'models/permission-level.enum';

import 'rxjs/add/operator/filter';

@Injectable()
export class DataService {
  private userId: number;
  private numberOfNewsOnPage: number;
  private userToken: string;
  private userPermLvl: PermLevel;
  private news: BehaviorSubject<NewsItem[]> = <BehaviorSubject<NewsItem[]>>new BehaviorSubject([]);
  private projects: BehaviorSubject<ProjectItem[]> = <BehaviorSubject<ProjectItem[]>>new BehaviorSubject([]);
  private userProjects: BehaviorSubject<ProjectItem[]> = <BehaviorSubject<ProjectItem[]>>new BehaviorSubject([]);
  private userEnrolledProjects: BehaviorSubject<ProjectItem[]> = <BehaviorSubject<ProjectItem[]>>new BehaviorSubject([]);
  private projectsForMainPage: BehaviorSubject<ProjectItem[]> = <BehaviorSubject<ProjectItem[]>>new BehaviorSubject([]);
  private enrollsForTeacher: BehaviorSubject<EnrollItem[]> = <BehaviorSubject<EnrollItem[]>>new BehaviorSubject([]);

  private newsCount: number;
  private newsCountObs: BehaviorSubject<number> = new BehaviorSubject<number>(0);
  private dataStore: {
    news: NewsItem[];
    projects: ProjectItem[];
    userProjects: ProjectItem[];
    userEnrolledProjects: ProjectItem[];
    projectsForMainPage: ProjectItem[];
    enrollsForTeacher: EnrollItem[];
  } = {
    news: [], projects: [], userProjects: [],
    userEnrolledProjects: [], projectsForMainPage: [], enrollsForTeacher: []
  };

constructor(
   private teacherService: TeacherService,
   private studentService: StudentService,
   private newsService: NewsService,
   private projectService: ProjectService,
   private userService: UserService
   ) { }

  public get News() {
    return this.news.asObservable();
  }
  public get NewsCountObs() {
    return this.newsCountObs.asObservable();
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
  public get EnrollsForTeacher() {
    return this.enrollsForTeacher.asObservable();
  }

  public get ProjectsForMainPage() {
    return this.projectsForMainPage.asObservable();
  }
  public set NumberOfNewsOnPage(value: number) {
    this.numberOfNewsOnPage = value;
  }
  public get PermLvl()
  {
    return this.userPermLvl;
  }
  
  loadAll() {
    /*console.log('Data.service ->loadAll');
    this.loadProjects();
    this.loadProjectsForMainPage();
    if (localStorage.getItem('current_user')) {
      this.userToken = JSON.parse(localStorage.getItem('current_user')).bearer_token;
      this.userId = JSON.parse(localStorage.getItem('current_user')).user.id;
      this.userPermLvl = JSON.parse(localStorage.getItem('current_user')).perm_lvl;
      this.loadUsersProjects();
      if (this.userPermLvl === PermLevel.Student) {
        this.loadEnrolledUsersProject();
      }
      if (this.userPermLvl === PermLevel.Teacher) {
        this.loadEnrollsForTeacher();
      }
    }*/
  }

  loadProjects() {
    this.projectService.getProjectItems().subscribe(res => {
      if (res != null) {
        this.dataStore.projects = res;
        this.dataStore.projects.forEach(a => { a.logo = this.addApiUrl(a.logo); })
        this.projects.next(Object.assign({}, this.dataStore).projects);
      }
    });
  }

  loadProjectsForMainPage() {
    this.projectService.getMainPageProjects().subscribe(res => {
      this.dataStore.projectsForMainPage = res;
      this.dataStore.projectsForMainPage.forEach(a => { a.logo = this.addApiUrl(a.logo); })
      this.projectsForMainPage.next(Object.assign({}, this.dataStore).projectsForMainPage);
    })
  }

  loadUsersProjects() {
    if (localStorage.getItem('current_user')) {
      this.userService.getProjectsOfUser(this.userId).subscribe(res => {
        if (res != null) {
          this.dataStore.userProjects = res;
          this.dataStore.userProjects.forEach(a => { a.logo = this.addApiUrl(a.logo); })
          this.userProjects.next(Object.assign({}, this.dataStore).userProjects);
        }

      });
    } else {
      console.log('Error in data.service: can not load usersProject without auth');
    }
  }

  loadEnrolledUsersProject() {
    if (localStorage.getItem('current_user')) {
      this.studentService.getEnrolledUsersProject(this.userId, this.userToken).subscribe(res => {
        this.dataStore.userEnrolledProjects = res;
        this.userEnrolledProjects.next(Object.assign({}, this.dataStore).userEnrolledProjects);
      })
    } else {
      console.log('Error in data.service: can not load enrolledUsersProject without auth');
    }
  }

// значения по умолчанию
  loadNews(offset: number) {
    this.newsService.getNewsPage(this.numberOfNewsOnPage, offset).subscribe(res => {
      this.newsCount = res.total_count;
      this.newsCountObs.next(Object.assign(res.total_count));
      this.dataStore.news = res.news_list;
      this.news.next(Object.assign({}, this.dataStore).news);

    });
  }

  loadEnrollsForTeacher() {
    this.teacherService.getEnrollsForTeacher(this.userToken).subscribe(res => {
      this.dataStore.enrollsForTeacher = res;
      this.enrollsForTeacher.next(Object.assign({}, this.dataStore).enrollsForTeacher);
    });
  }

  // TODO: сделать метод для проверки наличия новости в dataService


  addApiUrl(url: string): string {
    //return environment.apiUrl + url;
    return url;
  }
}
