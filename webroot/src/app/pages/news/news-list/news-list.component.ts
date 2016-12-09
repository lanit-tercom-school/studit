import { Component, OnInit } from '@angular/core';
import {NewsItem} from "../news-item";

@Component({
  selector: 'app-news-item-list',
  templateUrl: 'news-list.component.html',
  styleUrls: ['news-list.component.css']
})
export class NewsItemListComponent implements OnInit {

  constructor() { }

  newsItems: NewsItem

  ngOnInit() {
  }

}
