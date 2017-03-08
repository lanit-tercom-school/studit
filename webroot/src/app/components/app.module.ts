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

import { AboutComponent } from './main/about/about.component';
import { EnrolmentComponent } from './main/enrolment/enrolment.component';
import { FeaturesComponent } from './main/features/features.component';

import { ApiService } from '.././services/api.service';
import { AuthService } from '.././services/auth.service';
import { PartnersComponent } from './main/partners/partners.component';
import { PartnerItemComponent } from './main/partners/partner-item/partner-item.component';
import { SProjectPageComponent } from './pages/s-project-page/s-project-page.component';
import { MaterialsComponent } from './pages/s-project-page/materials/materials.component';
import { TasksComponent } from './pages/s-project-page/tasks/tasks.component';
import { ProjNewsComponent } from './pages/s-project-page/proj-news/proj-news.component';
import { MaterialsItemComponent } from './pages/s-project-page/materials/materials-item/materials-item.component';
import { AuthorizationComponent } from './pages/authorization/authorization.component';
import { HomePageComponent } from './pages/home-page/home-page.component';

import {FooterComponent} from "./footer/footer.component";

import { AppRouterProvider } from './routes';
import {AuthManager} from './authmanager';

@NgModule({
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
    AppRouterProvider,
  ],
  declarations: [
    AppComponent,
    MainComponent,
    TopPanelComponent,
    ProjectListComponent,
    ProjectItemComponent,
    AboutComponent,
    EnrolmentComponent,
    FeaturesComponent,
    PartnersComponent,
    PartnerItemComponent,
    SProjectPageComponent,
    MaterialsComponent,
    TasksComponent,
    ProjNewsComponent,
    MaterialsItemComponent,
    AuthorizationComponent,
    FooterComponent,
    HomePageComponent,
  ],
  providers: [ApiService, AuthService, AuthManager],
  bootstrap: [AppComponent]
})
export class AppModule { }
