import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Observable } from "rxjs/Observable";

import { TeacherService } from "services/teacher.service";

@Component({
  selector: 'app-admin-project-conrol',
  templateUrl: './admin-project-conrol.component.html',
  styleUrls: ['./admin-project-conrol.component.css']
})
export class AdminProjectConrolComponent implements OnInit {

  private createdProject = {
    name: "", description: "", status: 0,
    logo: "https://yegitsin.com/admin/pages/pictures/empty.jpg",
    id: 0
  };
  private projectId: number;
  private isCreated = false;
  constructor(private router: Router, private teacherService: TeacherService) { }

  ngOnInit() {
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
      this.createdProject.logo = promptValue;
  }

}
