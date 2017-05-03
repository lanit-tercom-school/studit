import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ApiService } from "../../../services/api.service";
//import { NgForm} from '@angular/forms';

import { ProjectItem } from '../../../models/project-item';
@Component({
  selector: 'create-project-page',
  templateUrl: './create-project-page.component.html',
  styleUrls: ['./create-project-page.component.css']
})
export class CreateProjectPageComponent implements OnInit {
    private createdProject  =  { name: "", description: "", logo: "", id : 0 };
    private projectId : number;

  constructor(private router : Router, private api: ApiService) { }

  ngOnInit() {
    }

    makeProject(){
      this.api.postProject(this.createdProject, JSON.parse(localStorage.getItem('current_user')).token).subscribe();
      console.log('Project was added');
      //this.router.navigate(['/home']);
    }

}
