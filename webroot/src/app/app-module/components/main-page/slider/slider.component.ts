import { Component, OnInit } from '@angular/core';
import { Observable } from "rxjs/Observable";

import { DataService } from 'services/data.service';
import { ProjectItem } from 'models/project-item';

@Component({
  selector: 'app-slider',
  templateUrl: './slider.component.html',
  styleUrls: ['./slider.component.css']
})
export class SliderComponent implements OnInit {

    public ProjectList: Observable<ProjectItem[]>;
   private p: number = 1;
  private limit: number = 1;
  private total: number = 3;

  private loading: boolean;
  constructor(private data: DataService) { }

  ngOnInit() {
    this.ProjectList = this.data.ProjectsForMainPage;
    this.getPage(1);
  }

  getPage(page: number) {
    this.loading = true;
    let offset = 0;
    if (page > 1)
      offset = (page - 1) * this.limit;
    this.p = page;
    this.loading = false;
  }

}
