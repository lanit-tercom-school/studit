import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { AppComponent } from './app.component';
import { NewsModule } from "./pages/news/news.module";
import { SettingsModule } from "./settings/settings.module";
import { ProfileModule } from "./profile/profile.module";
import { ProgressModule } from "./pages/progress/progress.module";
import { CoursecardComponent } from "./pages/coursecard/coursecard.component";
import { CoursecardModule } from "./pages/coursecard/coursecard.module";
import { MainComponent } from './main/main.component';
import { TopPanelComponent } from './top-panel/top-panel.component';
import { ProjectListComponent } from './shared/project-list/project-list.component';
import { ProjectItemComponent } from './shared/project-list/project-item/project-item.component';

import { ApiService } from '.././services/api.service';

@NgModule({
  declarations: [
    AppComponent,
    MainComponent,
    TopPanelComponent,
    ProjectListComponent,
    ProjectItemComponent,
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    NewsModule,
    SettingsModule,
    ProfileModule,
    ProgressModule,
    ProfileModule,
    CoursecardModule,
  ],
  providers: [ApiService],
  bootstrap: [AppComponent]
})
export class AppModule { }
