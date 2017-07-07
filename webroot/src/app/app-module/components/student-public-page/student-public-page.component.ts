import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router, Params } from '@angular/router';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { ApiService } from 'services/api.service';
import { CurrentUser } from 'models/current-user';

@Component({
  selector: 'app-student-public-page',
  templateUrl: './student-public-page.component.html',
  styleUrls: ['./student-public-page.component.css']
})
export class StudentPublicPageComponent implements OnInit {

  private currentUser: BehaviorSubject<CurrentUser> = new BehaviorSubject(new CurrentUser());

  constructor(private apiService: ApiService, private router: Router, private route: ActivatedRoute) { }

  ngOnInit() {
    this.route.params
      .subscribe(params => {
        this.apiService.getPublicStudentInfoById(+params['id'])
          .subscribe(res => this.currentUser.next(res.json()),
          error => {
            // сейчас на бэке по отстутствию пользователя ошибка 400
            if ((error.status === 404) || (error.status === 400))
              this.router.navigate(['/error']);
            else {
              alert('Ошибка! ' + error.status + ' ' + error.statusText);
              console.debug('ERROR: status ' + error.status + ' ' + error.statusText);
              console.debug('apiService: getPublicStudentInfoById()');
            }
          });
      });
  }
}
