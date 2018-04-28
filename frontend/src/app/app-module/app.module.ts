import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { NgxPaginationModule } from 'ngx-pagination'; 
import { SwiperModule } from 'angular2-useful-swiper';

import { AppRouterProvider } from './routes';
import { PathManager } from 'services/path.manager';

import { DataService } from 'services/data.service';
import { TeacherService } from 'services/teacher.service';
import { StudentService } from 'services/student.service';
import { ProjectService } from 'services/project.service';
import { NewsService } from 'services/news.service';
import { UserService } from 'services/user.service';
import { TaskService } from 'services/task.service';
import { AuthService } from 'services/auth.service';
import { FileService } from 'services/file.service';

import { AppComponent } from './components/app.component';
import { MainComponent } from './components/main-page/main.component';
import { AboutComponent } from './components/main-page/about/about.component';
import { EnrolmentComponent } from './components/main-page/enrolment/enrolment.component';
import { FeaturesComponent } from './components/main-page/features/features.component';
import { PartnersComponent } from './components/main-page/partners/partners.component';
import { PartnerItemComponent } from './components/main-page/partners/partner-item/partner-item.component';
import { SliderComponent } from './components/main-page/slider/slider.component';

import { TopPanelComponent } from './components/top-panel/top-panel.component';
import { FooterComponent } from "./components/footer/footer.component";

import { ProjectListPageComponent } from './components/project-list-page/project-list-page.component';
import { ProjectListComponent } from './components/project-list-page/project-list/project-list.component';
import { ProjectItemComponent } from './components/project-list-page/project-list/project-item/project-item.component';

import { MainNewsPageComponent } from './components/main-news-page/main-news-page.component';
import { NewsListComponent } from './components/main-news-page/news-list/news-list.component';
import { NewsItemComponent } from './components/main-news-page/news-list/news-item/news-item.component';
import { MainFullNewsPageComponent } from './components/main-full-news-page/main-full-news-page.component';

import { ProjectPageComponent } from './components/project-page/project-page.component';
import { MaterialsComponent } from '../shared-components/project-items/materials/materials.component';
import { MaterialsItemComponent } from '../shared-components/project-items/materials/materials-item/materials-item.component';
import { ProjNewsComponent } from '../shared-components/project-items/proj-news/proj-news.component';
import { ProjNewsItemComponent } from "../shared-components/project-items/proj-news/proj-news-item/proj-news-item.component";

import { AuthorizationPageComponent } from './components/authorization-page/authorization-page.component';
import { RegistrationPageComponent } from './components/registration-page/registration-page.component';
import { ValidationPageComponent } from './components/registration-page/validation-page/validation-page.component';

import { UserPublicPageComponent } from './components/user-public-page/user-public-page.component';

import { ErrorPageComponent } from './components/error-page/error-page.component';
import { TechnologiesComponent } from './components/technologies/technologies.component';


@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    AppRouterProvider,
    NgxPaginationModule,
    SwiperModule
  ],
  declarations: [
    AppComponent,
    MainComponent,
    AboutComponent,
    EnrolmentComponent,
    FeaturesComponent,
    PartnersComponent,
    PartnerItemComponent,
    SliderComponent,
    TopPanelComponent,
    FooterComponent,
    ProjectListPageComponent,
    ProjectListComponent,
    ProjectItemComponent,
    MainNewsPageComponent,
    NewsListComponent,
    NewsItemComponent,
    MainFullNewsPageComponent,
    ProjectPageComponent,
    MaterialsComponent,
    MaterialsItemComponent,
    ProjNewsComponent,
    ProjNewsItemComponent,
    AuthorizationPageComponent,
    RegistrationPageComponent,
    ValidationPageComponent,
    UserPublicPageComponent,
    ErrorPageComponent,
    TechnologiesComponent,
  ],
  providers: [
    FileService,
    TeacherService,
    StudentService,
    ProjectService,
    NewsService,
    UserService,
    TaskService,
    AuthService,
    DataService,
    PathManager
    ],
  bootstrap: [AppComponent]
})
export class AppModule { }
