import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { AppComponent } from './app.component';
import {NewsModule} from "./pages/news/news.module";
<<<<<<< HEAD
import {SettingsModule} from "./settings/settings.module";
import {ProfileModule} from "./profile/profile.module";
<<<<<<< HEAD
=======
import {ProgressModule} from "./pages/progress/progress.module";
>>>>>>> origin/feature/progress
=======
import {CoursecardComponent} from "./pages/coursecard/coursecard.component";
import {CoursecardModule} from "./pages/coursecard/coursecard.module";
>>>>>>> origin/feature/courselist

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    NewsModule,
<<<<<<< HEAD
    SettingsModule,
<<<<<<< HEAD
    ProfileModule
=======
    ProgressModule
>>>>>>> origin/feature/progress
=======
    ProfileModule,
    CoursecardModule
>>>>>>> origin/feature/courselist
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
