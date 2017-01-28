import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { AppComponent } from './app.component';
import {NewsModule} from "./pages/news/news.module";
<<<<<<< HEAD
import {SettingsModule} from "./settings/settings.module";
import {ProfileModule} from "./profile/profile.module";
=======
import {ProgressModule} from "./pages/progress/progress.module";
>>>>>>> origin/feature/progress

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
    ProfileModule
=======
    ProgressModule
>>>>>>> origin/feature/progress
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
