import { Component, OnInit, OnDestroy } from '@angular/core';
import { Observable } from "rxjs/Observable";

import { EnrollItem } from 'models/enroll-item';
import { TeacherService } from 'services/teacher.service';
//import { DataService } from 'services/data.service';

@Component({
  selector: 'app-teacher-note',
  templateUrl: './teacher-notification-page.component.html',
  styleUrls: ['./teacher-notification-page.component.css']
})
export class TeacherNotePageComponent implements OnInit, OnDestroy {
  private EnrollList: EnrollItem[];
  constructor(private teacherService: TeacherService) { }

  ngOnInit() {
    this.teacherService.getEnrollsForTeacher(JSON.parse(window.localStorage.getItem("current_user")).Token).subscribe(res => {
      console.log(res);
      this.EnrollList = res;
    });
  }
  ngOnDestroy() {

  }
}
