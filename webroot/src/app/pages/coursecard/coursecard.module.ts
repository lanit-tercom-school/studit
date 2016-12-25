import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CoursecardComponent } from './coursecard.component';
import { CourseListComponent } from './course-list/course-list.component';

@NgModule({
  imports: [
    CommonModule
  ],
  exports: [
    CoursecardComponent
  ],
  declarations: [CoursecardComponent, CourseListComponent]
})
export class CoursecardModule { }
