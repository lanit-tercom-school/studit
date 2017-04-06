import { Component, OnInit } from '@angular/core';

import { ProjectItem } from '../shared/project-list/project-item/project-item';
import { ApiService } from './../../services/api.service';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class MainComponent implements OnInit {

  private projects;
  private partners;

  constructor(private apiService: ApiService) { }

  ngOnInit() {
    this.getMainProjectList();
    this.partners = [{
      src: "https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Microsoft_logo_(2012).svg/2000px-Microsoft_logo_(2012).svg.png",
      link: "#"
    },
    {
      src: "https://upload.wikimedia.org/wikipedia/ru/0/00/The_Faculty_of_Mathematics_and_Mechanics_Logo.png",
      link: "#"
    },
    {
      src: "http://www.sporos.narod.ru/images/logo_spbgu.jpg",
      link: "#"
    }];
  }

  getMainProjectList(){
    this.apiService.getMainPageProjects().subscribe(res => this.projects = res);
  }
}
