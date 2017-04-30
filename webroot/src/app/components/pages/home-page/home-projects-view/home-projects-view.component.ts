import { Component, OnInit } from '@angular/core';
import { ProjectItem } from '../../../../models/project-item';
import {ApiService } from '../../../../services/api.service';
@Component({
  selector: 'app-home-projects-view',
  templateUrl: './home-projects-view.component.html',
  styleUrls: ['./home-projects-view.component.css']
})
export class HomeProjectsViewComponent implements OnInit {
  private userId:string;
  private ProjectList:ProjectItem[];
  constructor(private api:ApiService) { }

  ngOnInit() {
    this.userId=JSON.parse(localStorage.getItem("current_user")).id;
    this.api.getProjectItemsByUserId(this.userId).subscribe(res => { this.ProjectList = res });
  }

}
