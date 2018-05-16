import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { AuthService } from 'services/auth.service';

@Component({
  selector: 'app-validation-page',
  templateUrl: './validation-page.component.html',
  styleUrls: ['./validation-page.component.css']
})
export class ValidationPageComponent implements OnInit {

  public ValidationCode: string ='';
  private isValidated: boolean;


  constructor(private auth: AuthService, private router: Router) { }

  validate() {
    this.auth.validate(this.ValidationCode)
      .subscribe(
      data => {
        this.router.navigateByUrl("/auth")
      },
      error => {
        console.log(error);
        //TODO: Make notification for user here.
      });
  }

  ngOnInit() {
  }

}
