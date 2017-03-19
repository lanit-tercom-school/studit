import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../../../services/api.service';

@Component({
  selector: 'app-validation-page',
  templateUrl: './validation-page.component.html',
  styleUrls: ['./validation-page.component.css']
})
export class ValidationPageComponent implements OnInit {

  private validationCode: string;
  private message: string;
  private isValidated: boolean;


  constructor(private api: ApiService) { }

  validate() {
    this.api.validate(this.validationCode)
      .subscribe(
      data => {
        this.message = data;
      },
      error => {
        console.log(error);
        this.message = error;
      });
  }

  ngOnInit() {
    this.validationCode = localStorage.getItem("validation_key")
  }

}
