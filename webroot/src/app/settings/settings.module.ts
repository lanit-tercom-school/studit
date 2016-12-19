import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SettingsComponent } from './settings.component';
import { DataComponent } from './data/data.component';
import { SocialProfileComponent } from './social-profile/social-profile.component';

@NgModule({
  imports: [
    CommonModule
  ],
  exports: [  SettingsComponent ],
  declarations: [SettingsComponent, DataComponent, SocialProfileComponent]
})
export class SettingsModule { }
