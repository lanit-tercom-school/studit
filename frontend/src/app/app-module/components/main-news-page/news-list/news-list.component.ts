import { Component, Input, OnInit } from '@angular/core';

import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { NewsItem } from 'models/news-item';
import { DataService } from 'services/data.service';
import { TestImageService } from 'services/testImage.service';
//import { ApiService } from 'services/api.service';

import { Router, ActivatedRoute, Params } from '@angular/router';

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
  constructor(private data: DataService, private testImageService: TestImageService,private activatedRoute: ActivatedRoute, private router: Router) { }

  ngOnInit() {
    this.data.NumberOfNewsOnPage = 3;
    //this.getPage(1);
   this.activatedRoute.queryParams.subscribe((params: Params) => {
      let a = params['page'];
      console.log(a);
      //this.getPage(a); 
      if (a) {this.getPage(a,false);
      }else {
        this.getPage(1,true);
      }
    }); 
  }

  getPage(page: number, shouldChangeQuery: boolean) {
    if (shouldChangeQuery){
      this.router.navigate(['news'], {queryParams:{'page':page}});
    }
    this.TotalNumberOfNews = this.data.NewsCountObs;
    this.Loading = true;
    let offset = 0;
    if (page > 1)
      offset = (page - 1) * this.Limit;
    this.data.loadNews(offset);
    this.NewsList = this.data.News;
     this.NewsList.subscribe(data => {
      data.forEach(item => {
        this.testImageService.testImage(item.Image, ()=> {
          item.Image = "";
        });
      })
    });   //зачем это?
    this.PageNumber = page;
    this.Loading = false;
    window.scrollTo(0, 0);
  }

}
