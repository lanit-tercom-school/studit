import { Component, OnInit } from '@angular/core';
import { Observable } from "rxjs/Observable";

import { NewsItem } from 'models/news-item';
import { DataService } from 'services/data.service';

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
