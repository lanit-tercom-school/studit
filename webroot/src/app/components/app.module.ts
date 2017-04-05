import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { AppComponent } from './app.component';

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
import { StudentProjectPageComponent } from './pages/student-project-page/student-project-page.component';
import { MaterialsComponent } from './pages/student-project-page/materials/materials.component';
import { TasksComponent } from './pages/student-project-page/tasks/tasks.component';
import { ProjNewsComponent } from './pages/student-project-page/proj-news/proj-news.component';
import { MaterialsItemComponent } from './pages/student-project-page/materials/materials-item/materials-item.component';
import { AuthorizationPageComponent } from './pages/authorization-page/authorization-page.component';
import { RegistrationPageComponent } from './pages/registration-page/registration-page.component';
import { ValidationPageComponent } from './pages/registration-page/validation-page/validation-page.component';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { ProjectListPageComponent } from './pages/project-list-page/project-list-page.component';

import { FooterComponent } from "./footer/footer.component";
import { TasksItemComponent } from "./pages/student-project-page/tasks/tasks-item/tasks-item.component";
import { ProjNewsItemComponent } from "./pages/student-project-page/proj-news/proj-news-item/proj-news-item.component";


import { AppRouterProvider, AppProjectRouterProvider } from './routes';
import { AuthManager } from './../managers/authmanager';

import { AuthorPublicPageComponent } from './pages/author-public-page/author-public-page.component';
import { StudentPublicPageComponent } from './pages/student-public-page/student-public-page.component';

import { MainNewsPageComponent } from './pages/main-news-page/main-news-page.component';
import { MainNewsItemComponent } from './pages/main-news-page/main-news-item/main-news-item.component';

import { ProjectTasksPageComponent } from './pages/project-tasks-page/project-tasks-page.component';
import { ProjectTaskItemComponent } from './pages/project-tasks-page/project-task-list/project-task-item/project-task-item.component';
import { ProjectTaskListComponent } from './pages/project-tasks-page/project-task-list/project-task-list.component'
@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    AppRouterProvider,
    AppProjectRouterProvider,
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
    HomePageComponent,
    ProjectListPageComponent,
    RegistrationPageComponent,
    ValidationPageComponent,
    AuthorPublicPageComponent,
    StudentPublicPageComponent,
    MainNewsPageComponent,
    MainNewsItemComponent,
    ProjectTasksPageComponent,
    ProjectTaskItemComponent,
    ProjectTaskListComponent,
  ],
  providers: [ApiService, AuthService, AuthManager],
  bootstrap: [AppComponent]
})
export class AppModule { }
