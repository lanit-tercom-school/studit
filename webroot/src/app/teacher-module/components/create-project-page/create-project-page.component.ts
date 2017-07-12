import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Observable } from "rxjs/Observable";

import { ApiService } from "services/api.service";

@Component({
  selector: 'create-project-page',
  templateUrl: './create-project-page.component.html',
  styleUrls: ['./create-project-page.component.css']
})
export class CreateProjectPageComponent implements OnInit {
    private createdProject  =  { name: "", description: "", status: 0,
     logo: "https://yegitsin.com/admin/pages/pictures/empty.jpg",
     id : 0 };
    private projectId : number;
    private isCreated = false;
  constructor(private router : Router, private api: ApiService) { }

  ngOnInit() {
    }

    makeProject(){
      this.api.postProject(this.createdProject, JSON.parse(localStorage.getItem('current_user')).bearer_token)
      .subscribe(()  => {
      console.log('Project was added');
      this.isCreated = true;
      //this.router.navigate(['/home']);
      });
    }

    addLogo(){
    var promptValue = prompt('Укажите адрес картинки.', '');
    if (promptValue != null && promptValue != '')
      this.createdProject.logo = promptValue;
    }
}
