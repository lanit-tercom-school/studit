import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent implements OnInit {

  private token: string = "";
  constructor() { }

  ngOnInit() {
      this.token = window.localStorage.getItem('auth_key')
  }

}
