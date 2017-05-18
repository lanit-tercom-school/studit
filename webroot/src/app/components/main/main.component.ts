import { Component, OnInit } from '@angular/core';
import { DataService } from "services/data.service";
import { ProjectItem } from 'models/project-item';
import { ApiService } from './../../services/api.service';
import { Observable } from "rxjs/Observable";

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class MainComponent implements OnInit {

  private projects: Observable<ProjectItem[]>;
  private partners;

  constructor(private apiService: ApiService, private data:DataService) {
  }

  ngOnInit() {
    this.getMainProjectList();
    this.partners = [{
      src: "https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Microsoft_logo_(2012).svg/2000px-Microsoft_logo_(2012).svg.png",
      link: "https://www.microsoft.com/ru-ru/"
    },
    {
      src: "https://upload.wikimedia.org/wikipedia/ru/0/00/The_Faculty_of_Mathematics_and_Mechanics_Logo.png",
      link: "http://www.math.spbu.ru/rus/"
    },
    {
      src: "http://www.sporos.narod.ru/images/logo_spbgu.jpg",
      link: "http://spbu.ru/"
    }];
  }

  getMainProjectList(){
    this.projects=this.data.ProjectsForMainPage;
  }


}
