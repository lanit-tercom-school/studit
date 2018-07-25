import { Component, OnInit, OnDestroy } from '@angular/core';
import { Observable } from "rxjs/Observable";

import { ProjectItem } from 'models/project-item';
import { EnrollItem } from 'models/enroll-item';
import { DataService } from 'services/data.service';
import { TestImageService } from 'services/testImage.service';

@Component({
  selector: 'app-home-student-projects-view',
  templateUrl: './home-student-projects-view.component.html',
  styleUrls: ['./home-student-projects-view.component.css']
})
export class HomeStudentProjectsViewComponent implements OnInit, OnDestroy {
  private userId: string;
  public ProjectList: Observable<ProjectItem[]>;

  constructor(private data: DataService, private testImageService: TestImageService) { }

  ngOnInit() {
    this.ProjectList = this.data.UserProjects;
    this.ProjectList.subscribe(data => {
      data.forEach(item => {
        this.testImageService.testImage(item.Logo, ()=> {
          item.Logo = "";
        });
      })
    });
  }
  ngOnDestroy() {

  }
}
