import { Component, OnInit } from '@angular/core';

import { ProjectItem } from '../../shared/project-list/project-item/project-item';
import { ApiService } from './../../../services/api.service';

@Component({
  selector: 'app-project-list-page',
  templateUrl: './project-list-page.component.html',
  styleUrls: ['./project-list-page.component.css']
})
export class ProjectListPageComponent implements OnInit {

  private ProjectList;

  constructor(private apiService: ApiService) { }

  ngOnInit() {
    this.apiService.getProjectItems().subscribe(res => { this.ProjectList = res.json() });
  }
}
