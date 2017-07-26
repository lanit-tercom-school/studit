import { Component, OnInit } from '@angular/core';

import { AuthService } from 'services/auth.service';
import { AlertService } from 'services/alert.service';

@Component({
  selector: 'app-validation-page',
  templateUrl: './validation-page.component.html',
  styleUrls: ['./validation-page.component.css']
})
export class ValidationPageComponent implements OnInit {

  private validationCode: string;
  private message: string;
  private isValidated: boolean;


  constructor(private auth: AuthService,
  private alert: AlertService) { }

  validate() {
    this.auth.validate(this.validationCode)
      .subscribe(
      () => {
        this.message = 'Registered!';
        localStorage.removeItem("validation_code");
      },
      error => {
        this.message = error;
        this.alert.alertError(error, 'ERROR: validate() -> auth.validate()');
      });
  }

  ngOnInit() {
    this.validationCode = localStorage.getItem("validation_code")
  }

}
