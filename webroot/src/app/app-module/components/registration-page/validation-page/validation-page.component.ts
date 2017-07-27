import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { AuthService } from 'services/auth.service';
import { AlertService } from 'services/alert.service';

@Component({
  selector: 'app-validation-page',
  templateUrl: './validation-page.component.html',
  styleUrls: ['./validation-page.component.css']
})
export class ValidationPageComponent implements OnInit {

  private validationCode: string;
  private isValidated: boolean;


  constructor(private auth: AuthService,
  private alert: AlertService,
  private router: Router) { }

  validate() {
    this.auth.validate(this.validationCode)
      .subscribe(
      () => {
        localStorage.removeItem("validation_code");
        alert('Вы успешно зарегистрировались. Добро пожаловать!');
        //TODO   Добавить сюда автоматическую авторизацию и вход пользователя
        
        /*localStorage.setItem('current_user', user);
        this.router.navigate(['/home']);*/
      },
      error => {
        this.alert.alertError(error, 'validate() -> auth.validate()');
      });
  }

  ngOnInit() {
    this.validationCode = localStorage.getItem("validation_code")
  }

}
