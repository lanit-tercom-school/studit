import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { ApiService } from "../../../services/api.service";
//import { NgForm} from '@angular/forms';

import { ProjectItem } from '../../../models/project-item';
@Component({
  selector: 'create-project-page',
  templateUrl: './create-project-page.component.html',
  styleUrls: ['./create-project-page.component.css']
})
export class CreateProjectPageComponent implements OnInit {
    private createdProject:ProjectItem =  { name: "", description: "", logo: "", id: 0 };

  constructor(private api: ApiService) { }

  ngOnInit() {
    }

    makeProject(){
      this.api.postProject(this.createdProject, JSON.parse(localStorage.getItem('current_user')).token).subscribe(res => { });
      console.log('Hurray!');
    }
    deleteProject() {
    //  this.id = Number(prompt("Id of project?"));
    //  this.api.deleteProject("" + this.id, JSON.parse(localStorage.getItem('current_user')).token).subscribe(res=>{});
    }
}
