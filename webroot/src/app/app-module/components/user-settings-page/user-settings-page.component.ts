import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params, Router } from '@angular/router';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { ApiService } from "services/api.service";
import { CurrentUser } from 'models/current-user';

@Component({
  selector: 'app-user-settings-page',
  templateUrl: './user-settings-page.component.html',
  styleUrls: ['./user-settings-page.component.css']
})
export class UserSettingsPageComponent implements OnInit {

  private currentUser: BehaviorSubject<CurrentUser> = new BehaviorSubject(new CurrentUser());
  private clicked = false;
  private isChanged = false;
  private passwords = { old: '', new: '' };
  private NewPasswordAgain = '';
  private error: any;

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
  ShowHide() {
    if (!this.clicked) {
      this.clicked = true;
      this.isChanged = false;
    }
    else
      this.clicked = false;
  }

  ChangePassword() {
    if (this.passwords.new != this.NewPasswordAgain) {
      alert('Пароли не совпадают!');
      this.ClearPasswords();
    }
    else {
      this.apiService.changePasswordForUser(JSON.parse(localStorage.getItem('current_user')).token, this.passwords)
        .subscribe(res => {
          this.isChanged = true;
          this.clicked = false;
          this.ClearPasswords();
        },
        error => {
          //this.error = error;
          console.debug('ERROR: status ' + error.status + ' ' + error.statusText);
          console.debug('apiService: getPublicStudentInfoById()');
          alert('Ошибка! ' + error);
          this.ClearPasswords();
          this.isChanged = false;
        });
    }
  }

  ClearPasswords() {
    this.passwords.old = '';
    this.passwords.new = '';
    this.NewPasswordAgain = '';
  }
}
