import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params  } from '@angular/router';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { NewsItem } from 'models/news-item';
import { DataService } from 'services/data.service';
import { AlertService } from 'services/alert.service';

@Component({
  selector: 'app-main-full-news-page',
  templateUrl: './main-full-news-page.component.html',
  styleUrls: ['./main-full-news-page.component.css'],

})
export class MainFullNewsPageComponent implements OnInit {
 private readingNews: BehaviorSubject<NewsItem> = new BehaviorSubject(null);
newsID: number;
private sub: any;

constructor(
  private data: DataService,
  private alert: AlertService,
  private route: ActivatedRoute,
  private router: Router) { }

  ngOnInit() {
  this.sub = this.route.params.subscribe(params => {
         this.newsID = +params['id'];
      });
  this.getReadingNews();
  }

  getReadingNews(){
    this.data.loadNewsByID(this.newsID);
    this.data.NewsForViewing.subscribe(res => {
      if (res != null)
        this.readingNews.next(res);
    },
      error => {
        this.alert.alertError(error, 'ERROR: getReadingNews() -> MissedNews');
      });
  }
}
