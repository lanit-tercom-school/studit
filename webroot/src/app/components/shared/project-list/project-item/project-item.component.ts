import { Component, Input, OnInit } from '@angular/core';
import { ProjectItem } from '../../../../models/project-item';

@Component({
  host: {'class': 'card'},
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
