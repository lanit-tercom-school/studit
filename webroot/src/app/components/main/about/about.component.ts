import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserInfo } from '../../../models/user-info';
import { ApiService } from '../../../services/api.service';
@Component({
  selector: 'app-about',
  templateUrl: './about.component.html',
  styleUrls: ['./about.component.css']
})
export class AboutComponent implements OnInit {
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
