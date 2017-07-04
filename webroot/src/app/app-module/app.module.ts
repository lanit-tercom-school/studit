import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import {NgxPaginationModule} from 'ngx-pagination'; 

import { AppRouterProvider } from './routes';
import { PathManager } from 'services/path.manager';

import { DataService } from 'services/data.service';
import { ApiService } from 'services/api.service';
import { AuthService } from 'services/auth.service';
import { AppComponent } from './components/app.component';
import { MainComponent } from './components/main-page/main.component';
import { TopPanelComponent } from './components/top-panel/top-panel.component';
import { ProjectListComponent } from './components/project-list/project-list.component';
import { ProjectItemComponent } from './components/project-list/project-item/project-item.component';
import { NewsListComponent } from './components/news-list/news-list.component';
import { NewsItemComponent } from './components/news-list/news-item/news-item.component';
import { AboutComponent } from './components/main-page/about/about.component';
import { EnrolmentComponent } from './components/main-page/enrolment/enrolment.component';
import { FeaturesComponent } from './components/main-page/features/features.component';
import { PartnersComponent } from './components/main-page/partners/partners.component';
import { PartnerItemComponent } from './components/main-page/partners/partner-item/partner-item.component';
import { StudentProjectPageComponent } from './components/student-project-page/student-project-page.component';
import { MaterialsComponent } from './components/student-project-page/materials/materials.component';
import { TasksComponent } from './components/student-project-page/tasks/tasks.component';
import { ProjNewsComponent } from './components/student-project-page/proj-news/proj-news.component';
import { MaterialsItemComponent } from './components/student-project-page/materials/materials-item/materials-item.component';
import { AuthorizationPageComponent } from './components/authorization-page/authorization-page.component';
import { RegistrationPageComponent } from './components/registration-page/registration-page.component';
import { ValidationPageComponent } from './components/registration-page/validation-page/validation-page.component';
import { UserSettingsPageComponent } from './components/user-settings-page/user-settings-page.component';
import { ProjectListPageComponent } from './components/project-list-page/project-list-page.component';
import { FooterComponent } from "./components/footer/footer.component";
import { TasksItemComponent } from "./components/student-project-page/tasks/tasks-item/tasks-item.component";
import { ProjNewsItemComponent } from "./components/student-project-page/proj-news/proj-news-item/proj-news-item.component";
import { AuthorPublicPageComponent } from './components/author-public-page/author-public-page.component';
import { StudentPublicPageComponent } from './components/student-public-page/student-public-page.component';
import { MainNewsPageComponent } from './components/main-news-page/main-news-page.component';
import { MainFullNewsPageComponent } from './components/main-full-news-page/main-full-news-page.component';
import { ProjectTasksPageComponent } from './components/project-tasks-page/project-tasks-page.component';
import { ProjectTaskItemComponent } from './components/project-tasks-page/project-task-list/project-task-item/project-task-item.component';
import { ProjectTaskListComponent } from './components/project-tasks-page/project-task-list/project-task-list.component';
import { ErrorPageComponent } from './components/error-page/error-page.component';


@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    AppRouterProvider,
    NgxPaginationModule
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
    StudentProjectPageComponent,
    MaterialsComponent,
    TasksComponent,
    TasksItemComponent,
    ProjNewsItemComponent,
    ProjNewsComponent,
    MaterialsItemComponent,
    AuthorizationPageComponent,
    FooterComponent,
    ProjectListPageComponent,
    RegistrationPageComponent,
    ValidationPageComponent,
    AuthorPublicPageComponent,
    StudentPublicPageComponent,
    UserSettingsPageComponent,
    MainNewsPageComponent,
    NewsListComponent,
    NewsItemComponent,
    MainFullNewsPageComponent,
    ProjectTasksPageComponent,
    ProjectTaskItemComponent,
    ProjectTaskListComponent,
    ErrorPageComponent,
  ],
  providers: [ApiService, AuthService, PathManager, DataService],
  bootstrap: [AppComponent]
})
export class AppModule { }
