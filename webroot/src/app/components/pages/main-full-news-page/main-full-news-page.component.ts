import { Component, OnInit } from '@angular/core';
import { NewsItem } from 'models/news-item';
import { ApiService } from './../../../services/api.service';
import { Router, Route, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-main-full-news-page',
  templateUrl: './main-full-news-page.component.html',
  styleUrls: ['./main-full-news-page.component.css']
})
export class MainFullNewsPageComponent implements OnInit {
private readingNews;

  constructor(private apiService: ApiService,private router: Router) { }

  ngOnInit() {
  this.getReadingNews();
  }

  getReadingNews(){
   var id_ = 1;
    console.log("id ", id_);
    this.apiService.getNewsById(id_).subscribe(res => this.readingNews = res);
  }
}
