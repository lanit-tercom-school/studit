import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProfileComponent } from './profile.component';
import { ActivityComponent } from './activity/activity.component';
import { UserComponent } from './user/user.component';

@NgModule({
  imports: [
    CommonModule
  ],
  exports: [  ProfileComponent ],
  declarations: [ProfileComponent, ActivityComponent, UserComponent]
})
export class ProfileModule { }
