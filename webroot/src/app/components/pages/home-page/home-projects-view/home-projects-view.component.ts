import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-home-projects-view',
  templateUrl: './home-projects-view.component.html',
  styleUrls: ['./home-projects-view.component.css']
})
export class HomeProjectsViewComponent implements OnInit {
  private id:string;
  constructor() { }

  ngOnInit() {
  }
  getId()
  {
    this.id=JSON.parse(localStorage.getItem("current_user")).token;
    console.log(this.id);
  }
  getProjectsById(id:string)
  {

  }
}
