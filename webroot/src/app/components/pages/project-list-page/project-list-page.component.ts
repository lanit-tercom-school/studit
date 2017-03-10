import { Component, OnInit } from '@angular/core';

import { ProjectItem } from '../../shared/project-list/project-item/project-item';
import { ApiService } from './../../../services/api.service';

@Component({
  selector: 'app-project-list-page',
  templateUrl: './project-list-page.component.html',
  styleUrls: ['./project-list-page.component.css']
})
export class ProjectListPageComponent implements OnInit {

  constructor(private apiService: ApiService) { }

  ngOnInit() {
  }

  getProjectList() : ProjectItem [] {
    return this.apiService.getProjectItems();
  }

}
