import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { NewsItem } from 'models/news-item';
import { DataService } from 'services/data.service';
import { TestImageService } from 'services/testImage.service';

@Component({
  selector: 'app-main-full-news-page',
  templateUrl: './main-full-news-page.component.html',
  styleUrls: ['./main-full-news-page.component.css'],

})
export class MainFullNewsPageComponent implements OnInit {
  public ReadingNews: BehaviorSubject<NewsItem> = new BehaviorSubject(null);
  newsID: number;
  private sub: any;

  constructor(
    private data: DataService,
    private route: ActivatedRoute,
    private router: Router,
    private testImageService: TestImageService) { }

  ngOnInit() {
    this.sub = this.route.params.subscribe(params => {
      this.newsID = +params['id'];
    });
    this.getReadingNews();
  }

  getReadingNews() {
    this.data.loadNewsByID(this.newsID);
    this.data.NewsForViewing.subscribe(res => {
      if (res != null){
        this.ReadingNews.next(res);
        this.testImageService.testImage(res.Image, () => res.Image = "");
      }
    });
  }
}
