import { Component, OnInit, Input } from '@angular/core';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { ProjectShort } from 'models/project-short';
import { TestImageService } from 'services/testImage.service';
import { Observable } from "rxjs/Observable";
import { DataService } from 'services/data.service';
import { ProjectItem } from 'models/project-item';
import { ActivatedRoute } from '@angular/router';
@Component({
  selector: 'app-user-projects-view',
  templateUrl: './user-projects-view.component.html',
  styleUrls: ['./user-projects-view.component.css']
})
export class UserProjectsViewComponent implements OnInit {
  @Input() public UsersProjectList: Observable<ProjectItem[]>;

  constructor(private data: DataService, private testImageService: TestImageService, private route: ActivatedRoute) {

   }

  ngOnInit() {
      this.UsersProjectList.subscribe(data => {
        console.log(data)
        data.forEach(item => {
          this.testImageService.testImage(item.Logo, () => {
            item.Logo = "";
          });
        })
      });
  }
}
