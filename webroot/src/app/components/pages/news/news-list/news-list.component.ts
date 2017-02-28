import { Component, OnInit } from '@angular/core';
import {NewsItem} from "../news-item";
import {NewsService} from "../news.service";

@Component({
  selector: 'app-news-list',
  templateUrl: 'news-list.component.html',
  styleUrls: ['news-list.component.css'],
  providers: [NewsService]
})
export class NewsListComponent implements OnInit {

  constructor(private newsService: NewsService) { }

  newsItems: NewsItem[];

  ngOnInit() {
    this.newsItems = this.newsService.getNews();
  }

}
