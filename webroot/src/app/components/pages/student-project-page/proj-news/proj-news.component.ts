import { Component, OnInit } from '@angular/core';
import {Input} from "@angular/core";
import {ProjectNewsItem} from "models/proj-news-item";

@Component({
  selector: 'app-proj-news',
  templateUrl: './proj-news.component.html',
  styleUrls: ['./proj-news.component.css']
})
export class ProjNewsComponent implements OnInit {
  @Input() public ProjectNewsList:ProjectNewsItem[];
  constructor() { }

  ngOnInit() {
  }

}
