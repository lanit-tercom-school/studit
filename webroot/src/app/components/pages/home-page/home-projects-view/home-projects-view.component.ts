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
  
  }
  ngOnDestroy() {
    
  }
  choseProjects() {
  }
}
