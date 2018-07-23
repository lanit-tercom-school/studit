import { Component, OnInit, Input } from '@angular/core';

import { ProjectNewsItem } from "models/proj-news-item";


@Component({
  selector: 'app-proj-news-item',
  templateUrl: './proj-news-item.component.html',
  styleUrls: ['./proj-news-item.component.css']
})
export class ProjNewsItemComponent implements OnInit {
  @Input() public ProjectNewsItem: ProjectNewsItem;
  constructor() { }

  ngOnInit() {
  }

}
