import { Component, Input, OnInit } from '@angular/core';

import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { NewsItem } from 'models/news-item';
//import { DataService } from 'services/data.service';
import { ApiService } from 'services/api.service';

@Component({
  selector: 'app-news-list',
  templateUrl: './news-list.component.html',
  styleUrls: ['./news-list.component.css']
})
export class NewsListComponent implements OnInit {

 private p: number = 1;
  total: number = 10;  //сюда вставить число всех новостей
  loading: boolean;
 private NewsList: BehaviorSubject<NewsItem[]> = <BehaviorSubject<NewsItem[]>>new BehaviorSubject([]);

  constructor(private api: ApiService) { }

  ngOnInit() {
    this.getPage(1);
  }

  getPage(page: number) {
    this.loading = true;
    const limit = 3;
    const offset = (page - 1) * limit;
    this.api.getNewsPage(limit, offset).subscribe(res => {
      this.NewsList.next(res);
      this.p = page;
      this.loading = false;
      //this.total = res.total;
    });
  }
 
}
