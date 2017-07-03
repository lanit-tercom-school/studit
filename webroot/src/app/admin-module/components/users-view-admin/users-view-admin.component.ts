import { Component, OnInit, OnDestroy } from '@angular/core';
import { Observable } from "rxjs/Observable";

//import { UserItem } from 'models/project-item';
import { ApiService } from 'services/api.service';
import { DataService } from 'services/data.service';

@Component({
  selector: 'app-users-view-admin',
  templateUrl: './users-view-admin.component.html',
  styleUrls: ['./users-view-admin.component.css']
})
export class UsersViewAdminComponent implements OnInit, OnDestroy {
  //private userId: string;
  //private UserList: Observable<ProjectItem[]>;

  constructor(private api: ApiService, private data: DataService) { }

  ngOnInit() {
    //this.UserList = this.data.UserProjects;
  }
  ngOnDestroy() {

  }
}
