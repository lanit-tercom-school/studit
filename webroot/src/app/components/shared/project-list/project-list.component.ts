import { Component, Input, OnInit } from '@angular/core';
import { ProjectItem } from './project-item/project-item';

@Component({
  host: {'class': 'container'},
  selector: 'app-project-list',
  templateUrl: './project-list.component.html',
  styleUrls: ['./project-list.component.css']
})
export class ProjectListComponent implements OnInit {

  @Input() public ProjectList: ProjectItem[];

  constructor() { }

  ngOnInit() {
  }

}
