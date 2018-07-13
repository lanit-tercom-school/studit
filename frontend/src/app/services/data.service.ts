import { Injectable, EventEmitter } from '@angular/core';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { TeacherService } from 'services/teacher.service';
import { StudentService } from 'services/student.service';
import { ProjectService } from 'services/project.service';
import { NewsService } from 'services/news.service';
import { UserService } from 'services/user.service';
import { TaskService } from 'services/task.service';

import { NewsItem } from "models/news-item";
import { TasksItem } from "models/tasks-item";
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
  private tasksForViewing: BehaviorSubject<TasksItem[]> = <BehaviorSubject<TasksItem[]>>new BehaviorSubject([]);
  private newsCount: number;
  private projectCount: number;
  private newsCountObs: BehaviorSubject<number> = new BehaviorSubject<number>(0);

  private numberOfProjectsOnPage: number=2;
  private projectsCountObs: BehaviorSubject<number> = new BehaviorSubject<number>(7);
  private projectForViewing: BehaviorSubject<ProjectItem> = new BehaviorSubject<ProjectItem>(null);
  private newsForViewing: BehaviorSubject<NewsItem> = new BehaviorSubject<NewsItem>(null);

  private dataStore: {
    tasks: Map<string, TasksItem[]>;
    news: NewsItem[];
    projects: ProjectItem[];
    userProjects: ProjectItem[];
    userEnrolledProjects: EnrollItem[];
    projectsForMainPage: ProjectItem[];
    enrollsForTeacher: EnrollItem[];
  } = {
    news: [], projects: [], userProjects: [],
    userEnrolledProjects: [], projectsForMainPage: [], enrollsForTeacher: [],
    tasks: new Map<string, TasksItem[]>([]),
  };

  constructor(
    private teacherService: TeacherService,
    private studentService: StudentService,
    private newsService: NewsService,
    private projectService: ProjectService,
    private userService: UserService,
    private taskService: TaskService,
  ) { }

  public get News() {
    return this.news.asObservable();
  }
  public get TasksForViewing() {
    return this.tasksForViewing.asObservable();
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
  public get ProjectForViewing() {
    return this.projectForViewing.asObservable();
  }

  public get NewsForViewing() {
    return this.newsForViewing.asObservable();
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

  public set NumberOfProjectOnPage(value: number) {
    this.numberOfProjectsOnPage = value;
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


  loadTaskByGitHubUrl(gitHubUrl: string) {
    let url: string = gitHubUrl;
    let tasks: TasksItem[] = this.dataStore.tasks[gitHubUrl];
    if (tasks == null) {
      this.taskService.getTaskItemsFromGitHub(gitHubUrl)
        .subscribe(res => {
          if (res != null) {
            this.dataStore.tasks[url] = res;
            this.tasksForViewing.next(Object.assign({}, this.dataStore).tasks[url]);
          }
        });
    } else {
      this.tasksForViewing.next(tasks);
      this.taskService.getTaskItemsFromGitHub(gitHubUrl)
        .subscribe(res => {
          if (res != null && tasks !== res) {
            this.dataStore.tasks[url] = res;
            this.tasksForViewing.next(Object.assign({}, this.dataStore).tasks[url]);
          }
        });
    }
  }

  // для подгрузки проекта в ProjectForViewing
  loadProjectByID(id: number) {
    console.debug('data: load Project by ID');
    let foundproject = this.dataStore.projects.find(item => item.Id == id);
    if (foundproject) {
      this.projectForViewing.next(foundproject);
      console.debug('load from data');
    }
    else {
      this.projectService.getProjectById(id).subscribe(res => {
        if (res != null) {
          // дописываем в конец массива            
          this.dataStore.projects.push(res);
          this.projectForViewing.next(res);
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

  loadProjects(offset: number) {
    this.projectService.getProjectItems(this.numberOfProjectsOnPage, offset).subscribe(res => {
      this.projectCount = res.TotalCount;
      this.projectsCountObs.next(this.projectCount);
      this.dataStore.projects = res.ProjectList;
      this.projects.next(Object.assign({}, this.dataStore).projects);

    });
  }

  loadEnrollsForTeacher() {
    this.teacherService.getEnrollsForTeacher(this.userToken, this.userId).subscribe(res => {
      this.dataStore.enrollsForTeacher = res;
      this.enrollsForTeacher.next(Object.assign({}, this.dataStore).enrollsForTeacher);
    });
  }

  // TODO: сделать метод для проверки наличия новости в dataService



  // для подгрузки новости в NewsForViewing
  loadNewsByID(id: number) {
    console.debug('data: load News by ID');
    let foundnews = this.dataStore.news.find(item => item.Id == id);
    if (foundnews) {
      this.newsForViewing.next(foundnews);
      console.debug('load from data');
    }
    else {
      console.debug('can not find');
      this.newsService.getNewsById(id).subscribe(res => {
        if (res != null) {
          console.debug('NEW NEWS');
          // дописываем в конец массива            
          this.dataStore.news.push(res);
          this.newsForViewing.next(Object.assign({}, res));
        }
      });
    }

  }

  addApiUrl(url: string): string {
    //return environment.apiUrl + url;
    return url;
  }
}
