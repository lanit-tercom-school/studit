import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params  } from '@angular/router';

import { NewsItem } from 'models/news-item';
import { ApiService } from 'services/api.service';

@Component({
  selector: 'app-main-full-news-page',
  templateUrl: './main-full-news-page.component.html',
  styleUrls: ['./main-full-news-page.component.css'],

})
export class MainFullNewsPageComponent implements OnInit {
private readingNews;
id_: number;
private sub: any;

  constructor(private apiService: ApiService,
    private route: ActivatedRoute,
    private router: Router) { }

  ngOnInit() {
  this.sub = this.route.params.subscribe(params => {
         this.id_ = +params['id'];
      });
  this.getReadingNews();
  }

  getReadingNews(){
   this.apiService.getNewsById(this.id_).subscribe(res => this.readingNews = res);
  }
}
