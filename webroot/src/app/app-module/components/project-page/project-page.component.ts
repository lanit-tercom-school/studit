import { Component, OnInit, OnChanges, ViewContainerRef, OnDestroy } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';


import { DataService } from 'services/data.service';
import { ProjectService } from 'services/project.service';
import { MaterialsItem } from 'models/materials-item';
import { ProjectItem } from 'models/project-item';
import { ProjectNewsItem } from 'models/proj-news-item';
import { TasksItem } from 'models/tasks-item';

@Component({
  selector: 'app-project-page',
  templateUrl: './project-page.component.html',
  styleUrls: ['./project-page.component.css']
})
export class ProjectPageComponent implements OnInit, OnDestroy {

  private projectObs: BehaviorSubject<ProjectItem> = new BehaviorSubject(new ProjectItem());
  private projectId;
  private tasks = [];
  private message = 'Please write back soon!';
  constructor(
    private route: ActivatedRoute,
    private http: Http,
    private data: DataService,
    private projectService: ProjectService
  )
  { }

  ngOnInit() {
    this.route.params
      .subscribe(params => {
        this.projectId = params['id'];
        this.getProjectInfo();
      });
    this.getTaskItems();

  }

  ngOnDestroy() {
  }

  getProjectInfo() {
    this.data.loadProjectByID(this.projectId);
    this.data.ProjectForViewing.subscribe(res => {
      if (res != null)
        this.projectObs.next(res);
    },
      error => {
        this.data.alertError(error, 'ERROR: getProjectInfo() -> MissedProject');
      });

  }

  getMaterialsItems(): MaterialsItem[] {
    return this.projectService.getMaterialsItems(1);
  }

  getProjectNewsItem(): ProjectNewsItem[] {
    return this.projectService.getProjectNewsItem(1);
  }

  getTaskItems() {
    this.http.get('https://api.github.com/repos/lanit-tercom-school/studit/issues')
      .map((response: Response) => {
        let res = response.json().slice(0, 4);
        return res;
      }).subscribe(res => {
        this.tasks = res
      },
      error => {
        this.data.alertError(error, 'ERROR: getTaskItems()');
      });
  }
}
