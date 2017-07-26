import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { UserService } from "services/user.service";
import { AlertService } from "services/alert.service";

import { CurrentUser } from 'models/current-user';

@Component({
  selector: 'app-student-settings-page',
  templateUrl: './student-settings-page.component.html',
  styleUrls: ['./student-settings-page.component.css']
})
export class StudentSettingsPageComponent implements OnInit {

  private currentUser: BehaviorSubject<CurrentUser> = new BehaviorSubject(new CurrentUser());
  private clicked = false;
  private isChanged = false;
  private passwords = { old: '', new: '' };
  private NewPasswordAgain = '';
  private error: any;

  constructor(private userService: UserService,
  private alert: AlertService,
    private route: ActivatedRoute) { }

  ngOnInit() {
    this.route.params
      .subscribe(params => {
        this.userService.getUserById(JSON.parse(localStorage.getItem('current_user')).User.Id)
          .subscribe(res => {
            let c: CurrentUser = new (CurrentUser);
            c.User.Avatar = res.Avatar;
            c.User.Id = +res.Id;
            c.User.Description = res.Description;
            c.User.Nickname = res.Nickname;
            this.currentUser.next(c);
          },
           error => {
            this.alert.alertError(error, 'ngOnInit() -> getUserById');
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
      this.userService.changePasswordForUser(JSON.parse(localStorage.getItem('current_user')).Token, this.passwords)
        .subscribe(res => {
          this.isChanged = true;
          this.clicked = false;
          this.ClearPasswords();
        },
        error => {
          this.error = error;
          alert('Ошибка! ' + this.error);
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
