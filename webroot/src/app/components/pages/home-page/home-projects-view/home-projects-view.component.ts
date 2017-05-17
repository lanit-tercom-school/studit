import { Component, OnInit, OnDestroy } from '@angular/core';
import { Observable } from "rxjs/Observable";
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
  private ProjectList: Observable<ProjectItem[]>;
  private allProjects;
  private userProjects;
  constructor(private api: ApiService, private data: DataService) { }

  ngOnInit() {
    this.ProjectList = this.data.UserProjects;
  }
  ngOnDestroy() {
    
  }
  choseProjects() {
  }
}
