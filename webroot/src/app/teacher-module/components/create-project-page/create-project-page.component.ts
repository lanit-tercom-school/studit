import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Observable } from "rxjs/Observable";
import { Http, Headers, RequestOptions, Response } from '@angular/http';

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
  constructor(private router: Router, private teacherService: TeacherService, private data: DataService, private http: Http) { }

  ngOnInit() {
  }

  fileChange(event) {
    let fileList: FileList = event.target.files;
    if(fileList.length > 0) {
        let file: File = fileList[0];
        let formData:FormData = new FormData();
        formData.append('uploadFile', file, file.name);
        let headers = new Headers();
        headers.append('Accept', 'application/json');
        headers.append('Authorization', 'Bearer '+this.data.UserToken);
        this.http.post(`http://localhost:8080/graphql?query=mutation{PostFile{Id Path}}`, formData, { headers: headers })
            .map(res => res.json())
            .catch(error => Observable.throw(error))
            .subscribe(
                data => console.log('success'),
                error => console.log(error)
            )
    }
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
