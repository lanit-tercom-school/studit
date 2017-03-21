import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent implements OnInit {

  private currentUser;
  constructor() {
    this.currentUser = JSON.parse(localStorage.getItem('current_user'));
   }

  ngOnInit() {
  }

}
