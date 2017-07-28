import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Observable } from "rxjs/Observable";
import { Http, Headers, RequestOptions, Response } from '@angular/http';

import { FileService } from "services/file.service";
import { TeacherService } from "services/teacher.service";
import { DataService } from "services/data.service";
import { ProjectItem } from "models/project-item";

@Component({
  selector: 'create-project-page',
  templateUrl: './create-project-page.component.html',
  styleUrls: ['./create-project-page.component.css']
})


export class CreateProjectPageComponent implements OnInit {

  private createdProject = new ProjectItem();
  private isCreated = false;
  constructor(
    private router: Router,
    private teacherService: TeacherService,
    private data: DataService,
    private http: Http,
    private fileService: FileService
  ) { }

  ngOnInit() {
  }

  load(event) {
    this.fileService.uploadFiles(event.target.files).subscribe(res => {
      this.createdProject.Logo = res;
    });
  }

  makeProject() {
    this.teacherService.postProject(this.createdProject, JSON.parse(localStorage.getItem('current_user')).Token)
      .subscribe(() => {
        console.log('Project was added');
        this.isCreated = true;
        //this.router.navigate(['/home']);
      });
  }

  addLogo() {
    var promptValue = prompt('Укажите адрес картинки.', '');
    if (promptValue != null && promptValue != '')
      this.createdProject.Logo = promptValue;
  }
}
