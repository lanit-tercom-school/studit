import { Component, OnInit } from '@angular/core';
import { AuthService } from './../../../services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-registration-page',
  templateUrl: './registration-page.component.html',
  styleUrls: ['./registration-page.component.css']
})
export class RegistrationPageComponent implements OnInit {

  private fullName: string;
  private password: string;
  private email: string;

  constructor(private auth: AuthService, private router: Router) { }

  ngOnInit() {
  }

}
