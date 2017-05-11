import { Injectable } from '@angular/core';
import { ApiService } from './api.service'

@Injectable()
export class DataService {
  private projectItems;
  private projects;
  private enrollingProjects;
  private userId: number;
  constructor(private api: ApiService) { }
  loadEnrollingAndUserProjects() {
    this.projects = new Array<number>();
    this.enrollingProjects = new Array<number>();
    this.userId = JSON.parse(localStorage.getItem('current_user')).id;
    this.api.getProjectItems().subscribe(res => {
      console.log(res);
      this.projectItems = res;
      for (let i = 0; i < this.projectItems.length; i++) {
        this.api.getEnrolledUsersToProject(this.projectItems[i].id).subscribe(res => {
          if (res.json() != null) {
            for (let a of res.json()) {
              if (a === this.userId) {
                this.enrollingProjects.push(this.projectItems[i].id);
                break;
              }
            }
          }
        });
        this.api.getProjectUsers(this.projectItems[i].id).subscribe(res => {
          if (res.json() != null) {
            for (let a of res.json()) {
              if (a === this.userId) {
                this.projects.push(this.projectItems[i].id);
                break;
              }
            }
          }
        });
      }
    });
  }
  getProjectsOfUser() {
    return this.projects;
  }
  getEnrollingProjectsOfUser() {
    return this.enrollingProjects;
  }
}
