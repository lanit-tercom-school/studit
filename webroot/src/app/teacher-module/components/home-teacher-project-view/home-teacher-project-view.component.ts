import { Component, OnInit, OnDestroy } from '@angular/core';
import { Observable } from "rxjs/Observable";

import { ProjectItem } from 'models/project-item';
import { TeacherService } from 'services/teacher.service';
import { DataService } from 'services/data.service';
import { EnrollItem } from 'models/enroll-item';

@Component({
  selector: 'app-home-teacher-project-view',
  templateUrl: './home-teacher-project-view.component.html',
  styleUrls: ['./home-teacher-project-view.component.css']
})
export class HomeTeacherProjectViewComponent implements OnInit {
  private ProjectList: Observable<ProjectItem[]>;
  private ProjectEnrollList: Observable<ProjectItem[]>;
  private EnrollList: Observable<EnrollItem[]>;

  constructor(private teacherService: TeacherService, private data: DataService) { }

  ngOnInit() {
    this.ProjectList = this.data.UserProjects;
    this.ProjectEnrollList = this.data.UserEnrolledProjects;
    this.EnrollList = this.data.EnrollsForTeacher;
  }
  ngOnDestroy() {
  }
  accept(user_id: number, project_id: number) {
    this.teacherService.postUserToProject(user_id, project_id, JSON.parse(localStorage.getItem('current_user')).bearer_token).subscribe(res => {
       this.data.loadEnrollsForTeacher();
      });
  }
}
