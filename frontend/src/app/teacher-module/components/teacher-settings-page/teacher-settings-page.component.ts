import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';
import { FileService } from 'services/file.service';
import { UserService } from "services/user.service";
import { CurrentUser } from 'models/current-user';

@Component({
  selector: 'app-teacher-settings-page',
  templateUrl: './teacher-settings-page.component.html',
  styleUrls: ['./teacher-settings-page.component.css']
})
export class TeacherSettingsPageComponent implements OnInit {

  public CurrentUser: CurrentUser = {
    User: {
      Avatar: './assets/no_image.png',
      Description: '',
      Id: -1,
      Login: '',
      Nickname: '',
    },
    DataOfExpiration: undefined,
    PermissionLevel: undefined,
    Token: undefined
  };
  public Clicked = false;
  public IsChanged = false;
  public Passwords = { old: '', new: '' };
  public NewPasswordAgain = '';
  private error: any;
  private email: string;

  constructor(private userService: UserService, private route: ActivatedRoute, private fileService: FileService) { }

  ngOnInit() {
    this.CurrentUser.User.Login = JSON.parse(localStorage.getItem('current_user')).User.Login;
    this.route.params
      .subscribe(params => {
        this.userService.getUserById(JSON.parse(localStorage.getItem('current_user')).User.Id)
          .subscribe(res => {
            this.CurrentUser.User.Avatar = res.Avatar;
            this.CurrentUser.User.Id = +res.Id;
            this.CurrentUser.User.Description = res.Description;
            this.CurrentUser.User.Nickname = res.Nickname;
          });
      });
    this.email = JSON.parse(localStorage.getItem('current_user')).User.Id;
  }
  load(event) {
  this.fileService.uploadFiles(event.target.files).subscribe(res => {
    this.CurrentUser.User.Avatar = res;
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
