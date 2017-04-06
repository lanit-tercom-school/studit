import { Component, OnInit } from '@angular/core';
import { ProjectNewsItem} from "./proj-news-item";
import {Input} from "@angular/core";

@Component({
  selector: 'app-proj-news-item',
  templateUrl: './proj-news-item.component.html',
  styleUrls: ['./proj-news-item.component.css']
})
export class ProjNewsItemComponent implements OnInit {
  @Input() public  ProjectNewsItem:  ProjectNewsItem;
  constructor() { }

  ngOnInit() {
  }

}
