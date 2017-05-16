import { Component, OnInit } from '@angular/core';
import { DataService } from '../../../services/data.service'
import { ProjectItem } from '../../shared/project-list/project-item/project-item';
import { ApiService } from './../../../services/api.service';

@Component({
  selector: 'app-project-list-page',
  templateUrl: './project-list-page.component.html',
  styleUrls: ['./project-list-page.component.css']
})
export class ProjectListPageComponent implements OnInit {

  private ProjectList;

  constructor(private apiService: ApiService, private data: DataService) { }

  ngOnInit() {
    if (this.data.isProjectsUploaded()) {
      this.ProjectList = this.data.Projects;
    } else {
      this.data.ProjectsUploaded.subscribe(res => {
        this.ProjectList = this.data.Projects;
      });
    }
  }
}
