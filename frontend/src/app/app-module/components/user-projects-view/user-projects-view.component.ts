import { Component, OnInit, Input } from '@angular/core';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { ProjectShort } from 'models/project-short';

@Component({
  selector: 'app-user-projects-view',
  templateUrl: './user-projects-view.component.html',
  styleUrls: ['./user-projects-view.component.css']
})
export class UserProjectsViewComponent implements OnInit {
 @Input() public UsersProjectList: BehaviorSubject<ProjectShort[]>;

  constructor() { }

  ngOnInit() {
    this.UsersProjectList.subscribe(data=>console.log(data));
  }

}
