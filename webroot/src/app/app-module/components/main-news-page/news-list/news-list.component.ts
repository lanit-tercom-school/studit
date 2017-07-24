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

  private pageNumber: number = 1;
  private limit: number = 3;
  private totalNumberOfNews: Observable<number>;

  private loading: boolean;
  private NewsList: Observable<NewsItem[]>;
  constructor(private data: DataService) { }

  ngOnInit() {
    this.data.NumberOfNewsOnPage = 3;
    this.getPage(1);
  }

  getPage(page: number) {
    this.totalNumberOfNews = this.data.NewsCountObs;
    this.loading = true;
    let offset = 0;
    if (page > 1)
      offset = (page - 1) * this.limit;
    this.data.loadNews(offset);
    this.NewsList = this.data.News;
    this.pageNumber = page;
    this.loading = false;
    window.scrollTo(0, 0);
  }

}
