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
  public Partners = [
    {
      image: "./assets/partners/microsoft.png",
      url: "https://www.microsoft.com/ru-ru"
    },
    {
      image: "./assets/partners/mm.jpg",
      url: "http://www.math.spbu.ru/rus/"
    },
    {
      image: "./assets/partners/SPBSU3.jpg",
      url: "https://spbu.ru/"
    }
  ];;

  constructor(private data: DataService) {
  }

  ngOnInit() {
    this.getMainProjectList();
  }

  getMainProjectList() {
    this.projects = this.data.ProjectsForMainPage;
  }


}
