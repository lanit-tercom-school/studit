import { Component, OnInit } from '@angular/core';
import { Observable } from "rxjs/Observable";

import { DataService } from "services/data.service";
import { ProjectItem } from 'models/project-item';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class MainComponent implements OnInit {

  private projects: Observable<ProjectItem[]>;
  public Partners;

  constructor(private data: DataService) {
  }

  ngOnInit() {
    this.getMainProjectList();
    this.Partners = [{
      src: "https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Microsoft_logo_(2012).svg/2000px-Microsoft_logo_(2012).svg.png",
      link: "https://www.microsoft.com/ru-ru/"
    },
    {
      src: "https://sun1-4.userapi.com/c834103/v834103120/94e29/EWgKHEbny-Q.jpg",
      link: "https://www.math.spbu.ru/rus/"
    },
    {
      src: "https://spbu.ru/sites/all/themes/spbgu/markup/dist/img/logo-big-color.svg",
      link: "https://spbu.ru/"
    }];
  }

  getMainProjectList(){
    this.projects = this.data.ProjectsForMainPage;
  }


}
