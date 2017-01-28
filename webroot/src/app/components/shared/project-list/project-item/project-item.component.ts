import { Component, Input, OnInit } from '@angular/core';
import { ProjectItem } from './project-item'

@Component({
  selector: 'app-project-item',
  templateUrl: './project-item.component.html',
  styleUrls: ['./project-item.component.css']
})
export class ProjectItemComponent implements OnInit {

  @Input() ProjectItem: ProjectItem;

  constructor() { }

  ngOnInit() {
  }

}
