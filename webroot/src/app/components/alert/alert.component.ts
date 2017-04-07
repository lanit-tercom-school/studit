import { Component, OnInit } from '@angular/core';


import { AlertService } from 'services/alert.service';

@Component({
  moduleId: module.id,
  selector: 'app-alert',
  templateUrl: './alert.component.html',
  styleUrls: ['./alert.component.css']
})
export class AlertComponent implements OnInit {

  message: any;
 
  constructor(private alertService: AlertService) { }

//  constructor() { }

  ngOnInit() {
  	this.alertService.getMessage().subscribe(message => { this.message = message; });
  }

}
