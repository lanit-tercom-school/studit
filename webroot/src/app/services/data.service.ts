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
  private userEnrolledProjects: BehaviorSubject<EnrollItem[]> = <BehaviorSubject<EnrollItem[]>>new BehaviorSubject([]);
  private projectsForMainPage: BehaviorSubject<ProjectItem[]> = <BehaviorSubject<ProjectItem[]>>new BehaviorSubject([]);
  private enrollsForTeacher: BehaviorSubject<EnrollItem[]> = <BehaviorSubject<EnrollItem[]>>new BehaviorSubject([]);

  private newsCount: number;
  private newsCountObs: BehaviorSubject<number> = new BehaviorSubject<number>(0);

  private projectsCount: number = 7; //заглушка
  private projectsCountObs: BehaviorSubject<number> = new BehaviorSubject<number>(7);
  private missedProject: BehaviorSubject<ProjectItem> = new BehaviorSubject<ProjectItem>(null);
  private missedNews: BehaviorSubject<NewsItem> = new BehaviorSubject<NewsItem>(null);

  private dataStore: {
    news: NewsItem[];
    projects: ProjectItem[];
    userProjects: ProjectItem[];
    userEnrolledProjects: EnrollItem[];
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
  public get ProjectCountObs() {
    return this.projectsCountObs.asObservable();
  }

  //TODO: Change Missed to Viewed
  public get MissedProject() {
    return this.missedProject.asObservable();
  }

  public get MissedNews() {
    return this.missedNews.asObservable();
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

  public get PermLvl() {
    return this.userPermLvl;
  }
  public get UserId() {
    return this.userId;
  }
  public get UserToken() {
    return this.userToken;
  }

  loadAll() {
    console.log('Data.service ->loadAll');
    this.loadProjects(2, 0);
    this.loadProjectsForMainPage();
    if (localStorage.getItem('current_user')) {
      this.userToken = JSON.parse(localStorage.getItem('current_user')).Token;
      this.userId = JSON.parse(localStorage.getItem('current_user')).User.Id;
      this.userPermLvl = JSON.parse(localStorage.getItem('current_user')).PermissionLevel;
      this.loadUsersProjects();
      if (this.userPermLvl === PermLevel.Student) {
        this.loadEnrolledUsersProject();
      }
      if (this.userPermLvl === PermLevel.Teacher) {
        this.loadEnrollsForTeacher();
      }
    }
  }

  loadProjects(limit: number, offset: number) {
    this.projectService.getProjectItems(limit, offset)
      .subscribe(res => {
        if (res != null) {
          this.dataStore.projects = res;
          //this.projectsCountObs.next(Object.assign({},this.projectsCount));
          this.dataStore.projects.forEach(a => { a.Logo = this.addApiUrl(a.Logo); })
          this.projects.next(Object.assign({}, this.dataStore).projects);
        }
      });
  }

  // для подгрузки проекта
  loadProjectByID(id: number) {
    console.debug('data: load Project by ID');
    let foundproject = this.dataStore.projects.find(item => item.Id == id);
    if (foundproject) {
      this.missedProject.next(foundproject);
      console.debug('load from data');
    }
    else {
      this.projectService.getProjectById(id).subscribe(res => {
        if (res != null) {
          // дописываем в конец массива            
          this.dataStore.projects.push(res);
          this.missedProject.next(res);
        }
      });
    }

  }

  loadProjectsForMainPage() {
    this.projectService.getMainPageProjects().subscribe(res => {
      this.dataStore.projectsForMainPage = res;
      this.dataStore.projectsForMainPage.forEach(a => { a.Logo = this.addApiUrl(a.Logo); })
      this.projectsForMainPage.next(Object.assign({}, this.dataStore).projectsForMainPage);
    })
  }

  loadUsersProjects() {
    if (localStorage.getItem('current_user')) {
      this.userService.getProjectsOfUser(this.userToken, this.userId).subscribe(res => {
        if (res != null) {
          this.dataStore.userProjects = res;
          //this.dataStore.userProjects.forEach(a => { a.Logo = this.addApiUrl(a.Logo); })
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
        //console.log(this.dataStore.userEnrolledProjects);
        this.userEnrolledProjects.next(Object.assign({}, this.dataStore).userEnrolledProjects);
      })
    } else {
      console.log('Error in data.service: can not load enrolledUsersProject without auth');
    }
  }

  // значения по умолчанию
  loadNews(offset: number) {
    this.newsService.getNewsPage(this.numberOfNewsOnPage, offset).subscribe(res => {
      this.newsCount = res.TotalCount;
      this.newsCountObs.next(this.newsCount);
      this.dataStore.news = res.NewsList;
      this.news.next(Object.assign({}, this.dataStore).news);

    });
  }

  loadEnrollsForTeacher() {
    this.teacherService.getEnrollsForTeacher(this.userToken, this.userId).subscribe(res => {
      this.dataStore.enrollsForTeacher = res;
      this.enrollsForTeacher.next(Object.assign({}, this.dataStore).enrollsForTeacher);
    });
  }

  // TODO: сделать метод для проверки наличия новости в dataService



  // для подгрузки новости
  loadNewsByID(id: number) {
    console.debug('data: load News by ID');
    let foundnews = this.dataStore.news.find(item => item.Id == id);
    if (foundnews) {
      this.missedNews.next(foundnews);
      console.debug('load from data');
    }
    else {
      console.debug('can not find');
      this.newsService.getNewsById(id).subscribe(res => {
        if (res != null) {
          console.debug('NEW NEWS');
          // дописываем в конец массива            
          this.dataStore.news.push(res);  
          this.missedNews.next(Object.assign({}, res));
        }
      });
    }

  }

  addApiUrl(url: string): string {
    //return environment.apiUrl + url;
    return url;
  }
}
