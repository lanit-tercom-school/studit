import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

import { UserService } from "services/user.service";
import { CurrentUser } from 'models/current-user';

@Component({
  selector: 'app-teacher-settings-page',
  templateUrl: './teacher-settings-page.component.html',
  styleUrls: ['./teacher-settings-page.component.css']
})
export class TeacherSettingsPageComponent implements OnInit {

  public CurrentUser: BehaviorSubject<CurrentUser> = new BehaviorSubject(new CurrentUser());
  public Clicked = false;
  public IsChanged = false;
  public Passwords = { old: '', new: '' };
  public NewPasswordAgain = '';
  private error: any;

  constructor(private userService: UserService, private route: ActivatedRoute) { }

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
            this.CurrentUser.next(c);
          });
      });
  }

  ShowHide() {
    if (!this.Clicked) {
      this.Clicked = true;
      this.IsChanged = false;
    }
    else
      this.Clicked = false;
  }

  ChangePassword() {
    if (this.Passwords.new != this.NewPasswordAgain) {
      alert('Пароли не совпадают!');
      this.ClearPasswords();
    }
    else {
      this.userService.changePasswordForUser(JSON.parse(localStorage.getItem('current_user')).token, this.Passwords)
        .subscribe(res => {
          this.IsChanged = true;
          this.Clicked = false;
          this.ClearPasswords();
        },
          error => {
            this.error = error;
            alert('Ошибка! ' + this.error);
            this.ClearPasswords();
            this.IsChanged = false;
          });
    }
  }

  ClearPasswords() {
    this.Passwords.old = '';
    this.Passwords.new = '';
    this.NewPasswordAgain = '';
  }
}
