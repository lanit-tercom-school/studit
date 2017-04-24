import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { ApiService } from "../../../services/api.service";

@Component({
  selector: 'app-user-settings-page',
  templateUrl: './user-settings-page.component.html',
  styleUrls: ['./user-settings-page.component.css']
})
export class UserSettingsPageComponent implements OnInit {

  private currentStudent;
  private clicked = false;
  constructor(private apiService: ApiService,
    private route: ActivatedRoute) { }

  ngOnInit() {
    this.route.params
      .subscribe(params => {
      this.currentStudent = this.apiService.getPublicStudentInfoById(+params['id'])
        .subscribe(res => this.currentStudent = res.json());
      });

  }

ShowHide()
  {
  if (!this.clicked)
     this.clicked = true;
 else
   this.clicked = false;
  }
}
