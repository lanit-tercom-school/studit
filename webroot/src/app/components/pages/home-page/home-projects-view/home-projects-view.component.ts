import { Component, OnInit, OnDestroy } from '@angular/core';
import { ProjectItem } from '../../../shared/project-list/project-item/project-item'
import { ApiService } from '../../../../services/api.service'
import { DataService } from '../../../../services/data.service'
@Component({
  selector: 'app-home-projects-view',
  templateUrl: './home-projects-view.component.html',
  styleUrls: ['./home-projects-view.component.css']
})
export class HomeProjectsViewComponent implements OnInit, OnDestroy {
  private userId: string;
  private ProjectList: ProjectItem[];
  private allProjects;
  private userProjects;
  constructor(private api: ApiService, private data: DataService) { }

  ngOnInit() {
    this.ProjectList = new Array<any>();
    this.userId = JSON.parse(localStorage.getItem('current_user')).id;
    if (this.data.isProjectsUploaded()) {
      this.allProjects = this.data.Projects;
      if (this.data.isUsersProjectsUploaded()) {
        this.userProjects = this.data.UsersProjects;
        this.choseProjects();
      } else {
        this.data.UsersProjectsUploaded.subscribe(res => {
          this.userProjects = this.data.UsersProjects;
          this.choseProjects();
        });
      }
    } else {
      this.data.ProjectsUploaded.subscribe(res => {
        this.allProjects = this.data.Projects;
        if (this.data.isUsersProjectsUploaded()) {
          this.userProjects = this.data.UsersProjects;
          this.choseProjects();
        } else {
          this.data.UsersProjectsUploaded.subscribe(res => {
            this.userProjects = this.data.UsersProjects;
            this.choseProjects();
          });
        }
      });
    }
  }
  ngOnDestroy() {
    //this.data.UsersProjectsUploaded.unsubscribe();
    //this.data.ProjectsUploaded.unsubscribe();
  }
  choseProjects() {
    for (let i = 0; i < this.userProjects.length; i++) {
      for (let j = 0; j < this.allProjects.length; j++) {
        if (this.allProjects[j].id === this.userProjects[i]) { this.ProjectList.push(this.allProjects[j]); }
      }
    }
  }
}
