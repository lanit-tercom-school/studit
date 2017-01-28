import { Component, OnInit } from '@angular/core';
import {Coursecard} from "../coursecard";
import {CoursecardService} from "../coursecard.service";

@Component({
  selector: 'app-course-list',
  templateUrl: './course-list.component.html',
  styleUrls: ['./course-list.component.css']
})
export class CourseListComponent implements OnInit {

  courses: Coursecard[];

  constructor(private courseService: CoursecardService) { }

  ngOnInit() {
    this.courses = this.courseService.getCoursecards();
  }

}
