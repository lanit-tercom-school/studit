import { Component, OnInit, OnDestroy } from '@angular/core';
import { Observable } from "rxjs/Observable";
import { Router } from '@angular/router';
import { Http } from '@angular/http';

import { ProjectItem } from 'models/project-item';
import { TeacherService } from 'services/teacher.service';
import { StudentService } from 'services/student.service';
import { DataService } from 'services/data.service';
import { FileService } from 'services/file.service';
import { TestImageService } from 'services/testImage.service';

@Component({
  selector: 'app-home-teacher-project-view',
  templateUrl: './home-teacher-project-view.component.html',
  styleUrls: ['./home-teacher-project-view.component.css']
})
export class HomeTeacherProjectViewComponent implements OnInit {
  public ProjectList: Observable<ProjectItem[]>;
  public CreatedProject = new ProjectItem();
  public IsCreated = false;
  constructor(
    private router: Router,
    private teacherService: TeacherService,
    private data: DataService,
    private http: Http,
    private fileService: FileService,
    private studentService: StudentService,
    private testImageService: TestImageService
   ) {};

  ngOnInit() {
    this.ProjectList = this.data.UserProjects;
    this.ProjectList.subscribe(data => {
      data.forEach(item => {
        this.testImageService.testImage(item.Logo, ()=> {
          item.Logo = "";
        });
      })
    });
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
        // this.router.navigate(['/home']);
      });
    }
}
