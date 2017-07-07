import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';

import { ApiService } from "services/api.service";

@Component({
  selector: 'error-page',
  templateUrl: './error-page.component.html',
  styleUrls: ['./error-page.component.css']
})
export class ErrorPageComponent implements OnInit {
  private isAuthorised: boolean = false;

  constructor() { }

  ngOnInit() {
    if (localStorage.getItem('current_user'))
        this.isAuthorised = true;
  }

}
