import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home-page-teacher',
  templateUrl: './home-page-teacher.component.html',
  styleUrls: ['./home-page-teacher.component.css']
})
export class HomePageTeacherComponent implements OnInit {

  PROJECT_SECTION_URL: string = '/teacher/home/projects';
  ENROLLINGS_SECTION_URL: string = '/teacher/home/enrollings';

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
