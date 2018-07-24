import { Component, OnInit, Input } from '@angular/core';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { UserInfo } from 'models/user-info';
import { TestImageService } from 'services/testImage.service';

@Component({
  selector: 'app-project-users-view',
  templateUrl: './project-users-view.component.html',
  styleUrls: ['./project-users-view.component.css']
})
export class ProjectUsersViewComponent implements OnInit {
  @Input() public ProjectUsersList: BehaviorSubject<UserInfo[]>;

  constructor(private testImageService: TestImageService) { }

  ngOnInit() {
    this.ProjectUsersList.subscribe(data => {
      data.forEach(item => {
        this.testImageService.testImage(item.Avatar, ()=> {
          item.Avatar = "";
        });
      })
    });
  }

}
