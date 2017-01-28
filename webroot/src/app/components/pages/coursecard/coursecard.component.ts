import { Component, OnInit } from '@angular/core';
import {CoursecardService} from "./coursecard.service";

@Component({
  selector: 'app-coursecard',
  templateUrl: './coursecard.component.html',
  styleUrls: ['./coursecard.component.css'],
  providers: [CoursecardService]
})
export class CoursecardComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

}
