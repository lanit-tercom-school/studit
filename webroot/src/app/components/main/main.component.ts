import { Component, OnInit } from '@angular/core';

import { ProjectItem } from '../shared/project-list/project-item/project-item';
import { ApiService } from './../../services/api.service';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class MainComponent implements OnInit {

  constructor(private apiService: ApiService) { }

  ngOnInit() {
  }

  getMainProjectList() : ProjectItem [] {
    return this.apiService.getProjectItems();
  }
}
