import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { ApiService } from 'services/api.service';
import { CurrentUser } from 'models/current-user';

@Component({
  selector: 'app-admin-public-page',
  templateUrl: './admin-public-page.component.html',
  styleUrls: ['./admin-public-page.component.css']
})
export class AdminPublicPageComponent implements OnInit {

  private currentUser: BehaviorSubject<CurrentUser> = new BehaviorSubject(new CurrentUser());

  constructor(private apiService: ApiService, private route: ActivatedRoute) { }

  ngOnInit() {
    this.route.params
      .subscribe(params => {
        this.apiService.getUserById(JSON.parse(localStorage.getItem('current_user')).user.id)
          .subscribe(res => this.currentUser.next(res));
      });
  }
}
