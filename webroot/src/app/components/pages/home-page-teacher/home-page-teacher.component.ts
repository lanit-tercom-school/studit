import { Component, OnInit } from '@angular/core';
import {ApiService} from '../../../services/api.service'

@Component({
  selector: 'app-home-page-teacher',
  templateUrl: './home-page-teacher.component.html',
  styleUrls: ['./home-page-teacher.component.css']
})
export class HomePageTeacherComponent implements OnInit {
  private description:string;
  private id:number;
  private logo: string;
  private name:string;
  constructor(private api:ApiService) { }

  ngOnInit() {

  }
  makeProject()
  {
    this.name=prompt("Name of project?");
    this.description=prompt("Description of project?");
    this.logo=prompt("Logo of project?");
    this.id=Number(prompt("Id of project?"));
    let project={'description':this.description,'id':this.id,'logo':this.logo,'name':this.name};
    this.api.postProject(project,JSON.parse(localStorage.getItem('current_user')).token).subscribe(res=>{});
  }
}
