import { Component, OnInit } from '@angular/core';
import { Observable } from "rxjs/Observable";

import { EnrollItem } from 'models/enroll-item';
import { DataService } from 'services/data.service';

@Component({
  selector: 'app-home-student-enrolling-page',
  templateUrl: './home-student-enrolling-page.component.html',
  styleUrls: ['./home-student-enrolling-page.component.css']
})
export class HomeStudentEnrollingPageComponent implements OnInit {

  private ProjectEnrollList: Observable<EnrollItem[]>;

  constructor(private data: DataService) { }

  ngOnInit() {
    this.ProjectEnrollList = this.data.UserEnrolledProjects;
  }

}
