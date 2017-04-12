import { Component, Input, OnInit } from '@angular/core';
import { NewsItem } from 'models/news-item';

@Component({
  selector: 'app-news-list',
  templateUrl: './news-list.component.html',
  styleUrls: ['./news-list.component.css']
})
export class NewsListComponent implements OnInit {

  @Input() public NewsList: NewsItem[];

  constructor() { }

  ngOnInit() {
  }

}
