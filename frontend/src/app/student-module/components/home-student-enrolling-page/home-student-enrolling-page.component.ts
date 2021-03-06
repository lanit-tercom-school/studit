import { Component, OnInit } from '@angular/core';
import { Observable } from "rxjs/Observable";

import { EnrollItem } from 'models/enroll-item';
import { DataService } from 'services/data.service';
import { StudentService } from 'services/student.service';
import { TeacherService } from 'services/teacher.service';
import { TestImageService } from 'services/testImage.service';

@Component({
  selector: 'app-home-student-enrolling-page',
  templateUrl: './home-student-enrolling-page.component.html',
  styleUrls: ['./home-student-enrolling-page.component.css']
})
export class HomeStudentEnrollingPageComponent implements OnInit {

  public ProjectEnrollList: Observable<EnrollItem[]>;

  constructor(private data: DataService, private studentService: StudentService, private teacherService: TeacherService, private testImageService: TestImageService) { }

  ngOnInit() {
    this.ProjectEnrollList = this.data.UserEnrolledProjects;
    this.ProjectEnrollList.subscribe(data => {
      data.forEach(item => {
        this.testImageService.testImage(item.Project.Logo, ()=> {
          item.Project.Logo = "";
        });
      })
    });
  }
}
