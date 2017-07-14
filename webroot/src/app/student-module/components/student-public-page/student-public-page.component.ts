import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { UserService } from 'services/user.service';
import { CurrentUser } from 'models/current-user';

@Component({
  selector: 'app-student-public-page',
  templateUrl: './student-public-page.component.html',
  styleUrls: ['./student-public-page.component.css']
})
export class StudentPublicPageComponent implements OnInit {

  private currentUser: BehaviorSubject<CurrentUser> = new BehaviorSubject(new CurrentUser());

  constructor(private userService: UserService, private route: ActivatedRoute) { }

  ngOnInit() {
    this.route.params
      .subscribe(params => {
        this.userService.getUserById(JSON.parse(localStorage.getItem('current_user')).user.id)
          .subscribe(res => this.currentUser.next(res));
      });
  }
}
