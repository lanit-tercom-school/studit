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

  public CreatedProject = new ProjectItem();
  public IsCreated = false;
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
      this.CreatedProject.Logo = res;
    });
  }

  makeProject() {
    this.teacherService.postProject(this.CreatedProject, JSON.parse(localStorage.getItem('current_user')).Token)
      .subscribe(() => {
        console.log('Project was added');
        this.IsCreated = true;
        //this.router.navigate(['/home']);
      });
  }

  addLogo() {
    var promptValue = prompt('Укажите адрес картинки.', '');
    if (promptValue != null && promptValue != '')
      this.CreatedProject.Logo = promptValue;
  }
}
