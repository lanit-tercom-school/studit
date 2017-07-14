import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-enrolment',
  templateUrl: './enrolment.component.html',
  styleUrls: ['./enrolment.component.css']
})
export class EnrolmentComponent implements OnInit {
  private currentUser;
  constructor() {
  this.currentUser = JSON.parse(localStorage.getItem('current_user'));
   }

  ngOnInit() {
  }

}
