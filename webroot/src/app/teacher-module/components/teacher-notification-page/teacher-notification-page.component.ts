import { Component, OnInit, OnDestroy } from '@angular/core';
import { Observable } from "rxjs/Observable";

import { EnrollItem } from 'models/enroll-item';
import { ApiService } from 'services/api.service';
import { DataService } from 'services/data.service';

@Component({
  selector: 'app-teacher-note',
  templateUrl: './teacher-notification-page.component.html',
  styleUrls: ['./teacher-notification-page.component.css']
})
export class TeacherNotePageComponent implements OnInit, OnDestroy {
  private EnrollList: EnrollItem[];
  constructor(private api: ApiService, private data: DataService) { }

  ngOnInit() {
    this.api.getEnrollsForTeacher(JSON.parse(window.localStorage.getItem("current_user")).bearer_token).subscribe(res => {
      console.log(res);
      this.EnrollList = res;
    });
  }
  ngOnDestroy() {

  }
}
