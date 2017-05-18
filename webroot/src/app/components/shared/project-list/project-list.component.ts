import { Component, Input, OnInit } from '@angular/core';
import { ProjectItem } from '../../../models/project-item';
import { Observable } from "rxjs/Observable";

@Component({
  selector: 'app-project-list',
  templateUrl: './project-list.component.html',
  styleUrls: ['./project-list.component.css']
})
export class ProjectListComponent implements OnInit {

  @Input() public ProjectList: Observable<ProjectItem[]>;

  constructor() { }

  ngOnInit() {
  }

}
