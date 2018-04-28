import { Component, Input,OnInit } from '@angular/core';
import {NewsItem} from "models/news-item";

@Component({
  host: {'class': 'newscard'},
  selector: 'app-news-item',
  templateUrl: './news-item.component.html',
  styleUrls: ['./news-item.component.css']
})
export class NewsItemComponent implements OnInit {
  @Input() public NewsItem: NewsItem;
  constructor() { }

  ngOnInit() {
  }

}
