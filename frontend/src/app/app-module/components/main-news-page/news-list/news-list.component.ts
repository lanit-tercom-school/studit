import { Component, Input, OnInit } from '@angular/core';

import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { NewsItem } from 'models/news-item';
import { DataService } from 'services/data.service';
//import { ApiService } from 'services/api.service';

@Component({
  selector: 'app-news-list',
  templateUrl: './news-list.component.html',
  styleUrls: ['./news-list.component.css']
})
export class NewsListComponent implements OnInit {

  public PageNumber: number = 1;
  public Limit: number = 3;
  public TotalNumberOfNews: Observable<number>;

  public Loading: boolean;
  public NewsList: Observable<NewsItem[]>;
  constructor(private data: DataService) { }

  ngOnInit() {
    this.data.NumberOfNewsOnPage = 3;
    this.getPage(1);
  }

  getPage(page: number) {
    this.TotalNumberOfNews = this.data.NewsCountObs;
    this.Loading = true;
    let offset = 0;
    if (page > 1)
      offset = (page - 1) * this.Limit;
    this.data.loadNews(offset);
    this.NewsList = this.data.News;
    this.PageNumber = page;
    this.Loading = false;
    window.scrollTo(0, 0);
  }

}
