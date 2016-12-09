import { Component, OnInit } from '@angular/core';
import {NewsItem} from "../news-item";

@Component({
  selector: 'app-news-item-list',
  templateUrl: './news-item-list.component.html',
  styleUrls: ['./news-item-list.component.css']
})
export class NewsItemListComponent implements OnInit {

  constructor() { }

  newsItems: NewsItem

  ngOnInit() {
  }

}
