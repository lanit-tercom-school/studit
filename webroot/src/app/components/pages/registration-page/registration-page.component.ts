import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../../services/api.service';
import { Router } from '@angular/router';
import { UserInfo } from 'models/user-info';

@Component({
  selector: 'app-registration-page',
  templateUrl: './registration-page.component.html',
  styleUrls: ['./registration-page.component.css']
})
export class RegistrationPageComponent implements OnInit {

  private user: UserInfo = { login: "", nickname: "", password: "" };
  private error: string;

  constructor(private api: ApiService, private router: Router) { }

  ngOnInit() {
  }

  register() {
    this.api.register(this.user).subscribe(
      data => {
        this.router.navigate(['/registration/validate']);
      },
      error => {
        console.log(error);
        this.error = error;
      });
  }

}
