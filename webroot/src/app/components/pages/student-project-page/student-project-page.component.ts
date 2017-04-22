import { Component, OnInit, OnChanges } from '@angular/core';
import { ApiService } from './../../../services/api.service';
import { MaterialsItem } from './materials/materials-item/materials-item';
import { ActivatedRoute, Params } from '@angular/router';
import { ProjectItem } from './../../shared/project-list/project-item/project-item';
import { ProjectNewsItem } from './proj-news/proj-news-item/proj-news-item';
import { TasksItem } from "./tasks/tasks-item/tasks-item";
import { Http, Headers, RequestOptions, Response } from '@angular/http';


@Component({
  selector: 'app-student-project-page',
  templateUrl: './student-project-page.component.html',
  styleUrls: ['./student-project-page.component.css']
})
export class StudentProjectPageComponent implements OnInit {

  private project;
  private tasks = [];
  private Authorized: boolean;
  constructor(private apiService: ApiService,
    private route: ActivatedRoute, private http: Http) { }

  ngOnInit() {
    this.route.params
      .subscribe(params => {
        this.project = this.apiService.getProjectById(+params['id'])
          .subscribe(res => this.project = res.json());
      });
    this.getTaskItems();
    if (localStorage.getItem('current_user')) {
      this.Authorized = true;
    }
    else {
      this.Authorized = false;
    }
  }

  getMaterialsItems(): MaterialsItem[] {
    return this.apiService.getMaterialsItems(1);
  }
  getProjectNewsItem(): ProjectNewsItem[] {
    return this.apiService.getProjectNewsItem(1);
  }
  getTaskItems() {
    this.http.get('https://api.github.com/repos/lanit-tercom-school/studit/issues')
      .map((response: Response) => {
        var res = response.json().slice(0, 4);
        return res;
      }).subscribe(res => this.tasks = res);
  }
  enroll() {
    
  }
}
