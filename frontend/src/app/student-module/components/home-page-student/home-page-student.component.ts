import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home-page-student',
  templateUrl: './home-page-student.component.html',
  styleUrls: ['./home-page-student.component.css']
})
export class HomePageStudentComponent implements OnInit {

  CurrentSectionOpened: string = 'projects'

  constructor(private router: Router) { }

  ngOnInit() {
  }

  public chooseSection(section: string) {
    this.CurrentSectionOpened = section;
    switch (section) {
      case 'projects':
        this.router.navigateByUrl('student/home/projects');
        break;
      case 'enrollings':
        this.router.navigateByUrl('student/home/enrollings');
      default:
        break;
    }
  }


}
