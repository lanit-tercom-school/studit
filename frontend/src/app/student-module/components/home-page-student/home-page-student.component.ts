import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home-page-student',
  templateUrl: './home-page-student.component.html',
  styleUrls: ['./home-page-student.component.css']
})
export class HomePageStudentComponent implements OnInit {

  PROJECT_SECTION_URL: string = '/student/home/projects';
  ENROLLINGS_SECTION_URL: string = '/student/home/enrollings';

  CurrentUrl: string = ''

  constructor(private router: Router) { }

  ngOnInit() {
    this.CurrentUrl = this.router.url.valueOf();
  }

  public chooseSection(section: string) {
    this.CurrentUrl = section;
    this.router.navigateByUrl(section);
  }


}
