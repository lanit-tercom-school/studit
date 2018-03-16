import { Component, OnInit, OnDestroy } from '@angular/core';
import { Observable } from "rxjs/Observable";

import { ProjectItem } from 'models/project-item';
import { TeacherService } from 'services/teacher.service';
import { StudentService } from 'services/student.service';
import { DataService } from 'services/data.service';
import { EnrollItem } from 'models/enroll-item';

@Component({
  selector: 'app-home-teacher-project-view',
  templateUrl: './home-teacher-project-view.component.html',
  styleUrls: ['./home-teacher-project-view.component.css']
})
export class HomeTeacherProjectViewComponent implements OnInit {
  private ProjectList: Observable<ProjectItem[]>;
  private EnrollList: Observable<EnrollItem[]>;

  constructor(private teacherService: TeacherService, private data: DataService, private studentService: StudentService) { }

  ngOnInit() {
    this.ProjectList = this.data.UserProjects;
    this.EnrollList = this.data.EnrollsForTeacher;
  }
  ngOnDestroy() {
  }
}
