import { Component, OnInit } from '@angular/core';
import {ProjectItem} from '../../../shared/project-list/project-item/project-item'
import {ApiService } from '../../../../services/api.service'
@Component({
  selector: 'app-home-projects-view',
  templateUrl: './home-projects-view.component.html',
  styleUrls: ['./home-projects-view.component.css']
})
export class HomeProjectsViewComponent implements OnInit {
  
  private ProjectList:ProjectItem[];
  constructor(private api:ApiService) { }

  ngOnInit() {
    this.api.getProjectItems().subscribe(res => { this.ProjectList = res });
  }

}
