import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-main-news-page',
  templateUrl: './main-news-page.component.html',
  styleUrls: ['./main-news-page.component.css']
})
export class MainNewsPageComponent implements OnInit {

   constructor() { }

  ngOnInit() {
    window.scrollTo(0,0);
  }
 
}
