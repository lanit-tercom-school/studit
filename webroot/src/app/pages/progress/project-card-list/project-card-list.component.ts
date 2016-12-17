import { Component, OnInit } from '@angular/core';
import {ProgressService} from "../progress.service";
import {ProjectCardItem} from "../project-card-item";

@Component({
  selector: 'app-project-card-list',
  templateUrl: './project-card-list.component.html',
  styleUrls: ['./project-card-list.component.css']
})
export class ProjectCardListComponent implements OnInit {

  constructor(private progressService: ProgressService) { }

  projectCardItems: ProjectCardItem[];

  ngOnInit() {
    this.projectCardItems = this.progressService.getProjectCards();
  }

}
