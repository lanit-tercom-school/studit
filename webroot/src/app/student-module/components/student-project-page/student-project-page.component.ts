import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';

import { ProjectService } from 'services/project.service';
import { StudentService } from 'services/student.service';
import { ProjectItem } from 'models/project-item';
import { ProjectTaskItem } from "models/project-task-item";

@Component({
  selector: 'app-student-project-page',
  templateUrl: './student-project-page.component.html',
  styleUrls: ['./student-project-page.component.css']
})
export class StudentProjectPageComponent implements OnInit {

private project;

constructor(
  private route: ActivatedRoute,
  private projectService: ProjectService,
  private studentService: StudentService) { }

ngOnInit() {
  this.route.params
    .subscribe(params => {
    this.project = this.projectService.getProjectById(+params['id'])
      .subscribe(res => this.project = res.json());
    });
}

getProjectAllTaskItem() {
  return this.projectService.getProjectAllTaskItem(1);
}
getProjectStudentTaskItem() {
  return this.studentService.getProjectStudentTaskItem(1);
}
}
