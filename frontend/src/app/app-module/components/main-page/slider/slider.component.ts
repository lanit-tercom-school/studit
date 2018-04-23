import { Component, OnInit } from '@angular/core';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { DataService } from 'services/data.service';
import { ProjectItem } from 'models/project-item';

@Component({
  selector: 'app-slider',
  templateUrl: './slider.component.html',
  styleUrls: ['./slider.component.css']
})
export class SliderComponent implements OnInit {

  config: Object = {
    pagination: '.swiper-pagination',
    paginationClickable: true,
    nextButton: '.swiper-button-next',
    prevButton: '.swiper-button-prev',
    spaceBetween: 30
  };

  public ProjectList: Observable<ProjectItem[]>;

  private loading: boolean;
  constructor(private data: DataService) { }

  ngOnInit() {
    this.ProjectList = this.data.ProjectsForMainPage;
  }

}
