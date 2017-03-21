import { Component, OnInit } from '@angular/core';

import { ProjectItem } from '../shared/project-list/project-item/project-item';
import { ApiService } from './../../services/api.service';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class MainComponent implements OnInit {

  private projects;

  constructor(private apiService: ApiService) { }

  ngOnInit() {
    this.getMainProjectList
  }

  getMainProjectList(){
    this.apiService.getProjectItems().subscribe(res => this.projects = res.json());
  }
}
