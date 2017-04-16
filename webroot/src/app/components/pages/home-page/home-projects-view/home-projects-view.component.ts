import { Component, OnInit } from '@angular/core';
import {ProjectItem} from '../../../shared/project-list/project-item/project-item'
@Component({
  selector: 'app-home-projects-view',
  templateUrl: './home-projects-view.component.html',
  styleUrls: ['./home-projects-view.component.css']
})
export class HomeProjectsViewComponent implements OnInit {
  private id:string;
  private ProjectList:ProjectItem[];
  constructor() { }

  ngOnInit() {
    this.getId();
    this.ProjectList = this.getProjectsByUserId(this.id);
  }
  getId()
  {
    this.id=JSON.parse(localStorage.getItem("current_user")).id;
  }
  getProjectsByUserId(id:string)
  {
    return [
    {
    "Id": 3,
    "Name": "Оригинальное название",
    "Description": "Click-bait описание Bla bla bla bla bla bla bla Bla bla bla bla bla bla bla Bla bla bla bla bla bla bla Bla bla bla bla bla bla bla",
    "Logo": "/assets/project.jpg"
    },
    {
    "Id": 2,
    "Name": "Модный фрилансер",
    "Description": "Какие же стрелочки вокруг ноубука! Bla bla bla bla bla bla bla Bla bla bla bla bla bla bla Bla bla bla bla bla bla bla Bla bla bla bla bla bla bla",
    "Logo": "/assets/project.jpg"
    },
    {
      "Id": 2,
      "Name": "Модный фрилансер",
      "Description": "Какие же стрелочки вокруг ноубука! Bla bla bla bla bla bla bla Bla bla bla bla bla bla bla Bla bla bla bla bla bla bla Bla bla bla bla bla bla bla",
      "Logo": "/assets/project.jpg"
    },
    {
      "Id": 2,
      "Name": "Модный фрилансер",
      "Description": "Какие же стрелочки вокруг ноубука! Bla bla bla bla bla bla bla Bla bla bla bla bla bla bla Bla bla bla bla bla bla bla Bla bla bla bla bla bla bla",
      "Logo": "/assets/project.jpg"
    },
    {
    "Id": 3,
    "Name": "Оригинальное название",
    "Description": "Click-bait описание Bla bla bla bla bla bla bla Bla bla bla bla bla bla bla Bla bla bla bla bla bla bla Bla bla bla bla bla bla bla",
    "Logo": "/assets/project.jpg"
    } 
    ]
  }
}
