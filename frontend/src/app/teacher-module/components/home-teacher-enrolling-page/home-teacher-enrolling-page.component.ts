import { Component, OnInit } from '@angular/core';
import { Observable } from "rxjs/Observable";

import { EnrollItem } from 'models/enroll-item';
import { DataService } from 'services/data.service';
import { TeacherService } from 'services/teacher.service';
import { StudentService } from 'services/student.service';

@Component({
  selector: 'app-home-teacher-enrolling-page',
  templateUrl: './home-teacher-enrolling-page.component.html',
  styleUrls: ['./home-teacher-enrolling-page.component.css']
})
export class HomeTeacherEnrollingPageComponent implements OnInit {

  public EnrollList: Observable<EnrollItem[]>;

  constructor(private teacherService: TeacherService, private data: DataService, private studentService: StudentService) { }

  ngOnInit() {
    this.EnrollList = this.data.EnrollsForTeacher;
  }
  accept(enroll: EnrollItem) {
    this.studentService.unenrollToProject(enroll.Id, this.data.UserToken).subscribe(r => {
      this.teacherService.postUserToProject(+enroll.User.Id, enroll.Project.Id, this.data.UserToken).subscribe(res => {
        this.data.loadEnrollsForTeacher();
      });
    });
  }
}
