import { Component, OnInit } from '@angular/core';
import { Observable } from "rxjs/Observable";
import { DataService } from 'services/data.service';
import { ApiService } from 'services/api.service';
import { ProjectItem } from 'models/project-item';


@Component({
  selector: 'app-project-list-page',
  templateUrl: './project-list-page.component.html',
  styleUrls: ['./project-list-page.component.css']
})
export class ProjectListPageComponent implements OnInit {

  private ProjectList: Observable<ProjectItem[]>;

  constructor(private apiService: ApiService, private data: DataService) { }

  ngOnInit() {
    this.ProjectList=this.data.Projects;
  }
}
