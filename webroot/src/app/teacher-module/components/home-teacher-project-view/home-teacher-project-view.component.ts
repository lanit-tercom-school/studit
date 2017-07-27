import { Component, OnInit, OnDestroy } from '@angular/core';
import { Observable } from "rxjs/Observable";

import { TeacherService } from 'services/teacher.service';
import { StudentService } from 'services/student.service';
import { DataService } from 'services/data.service';
import { AlertService } from 'services/alert.service';

import { EnrollItem } from 'models/enroll-item';
import { ProjectItem } from 'models/project-item';

@Component({
  selector: 'app-home-teacher-project-view',
  templateUrl: './home-teacher-project-view.component.html',
  styleUrls: ['./home-teacher-project-view.component.css']
})

export class HomeTeacherProjectViewComponent implements OnInit {
  private ProjectList: Observable<ProjectItem[]>;
  private EnrollList: Observable<EnrollItem[]>;

  constructor(private teacherService: TeacherService,
   private data: DataService,
   private alert: AlertService,
   private studentService: StudentService) { }

  ngOnInit() {
    this.ProjectList = this.data.UserProjects;
    this.EnrollList = this.data.EnrollsForTeacher;
  }
  ngOnDestroy() {
  }
  accept(enroll:EnrollItem) {
    this.studentService.unenrollToProject(enroll.Id, this.data.UserToken).subscribe(r => {
      this.teacherService.postUserToProject(+enroll.User.Id, enroll.Project.Id, this.data.UserToken).subscribe(res => {
        this.data.loadEnrollsForTeacher();
      },
        error => {
          this.alert.alertError(error, 'accept() -> unenrollToProject() -> postUserToProject()');
        });
    },
      error => {
        this.alert.alertError(error, 'accept() -> unenrollToProject()');
      });
  }
}
