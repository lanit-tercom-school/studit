import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { AppRouterProvider } from './routes';
import { PathManager } from 'services/path.manager';

import { DataService } from 'services/data.service';
import { ApiService } from 'services/api.service';
import { AuthService } from 'services/auth.service';
import { AppComponent } from './app.component';
import { MainComponent } from './main/main.component';
import { TopPanelComponent } from './top-panel/top-panel.component';
import { ProjectListComponent } from './shared/project-list/project-list.component';
import { ProjectItemComponent } from './shared/project-list/project-item/project-item.component';
import { NewsListComponent } from './shared/news-list/news-list.component';
import { NewsItemComponent } from './shared/news-list/news-item/news-item.component';
import { AboutComponent } from './main/about/about.component';
import { EnrolmentComponent } from './main/enrolment/enrolment.component';
import { FeaturesComponent } from './main/features/features.component';
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
import { UserSettingsPageComponent } from './pages/user-settings-page/user-settings-page.component';
import { ProjectListPageComponent } from './pages/project-list-page/project-list-page.component';
import { FooterComponent } from "./footer/footer.component";
import { TasksItemComponent } from "./pages/student-project-page/tasks/tasks-item/tasks-item.component";
import { ProjNewsItemComponent } from "./pages/student-project-page/proj-news/proj-news-item/proj-news-item.component";
import { AuthorPublicPageComponent } from './pages/author-public-page/author-public-page.component';
import { StudentPublicPageComponent } from './pages/student-public-page/student-public-page.component';
import { MainNewsPageComponent } from './pages/main-news-page/main-news-page.component';
import { MainFullNewsPageComponent } from './pages/main-full-news-page/main-full-news-page.component';
import { ProjectTasksPageComponent } from './pages/project-tasks-page/project-tasks-page.component';
import { ProjectTaskItemComponent } from './pages/project-tasks-page/project-task-list/project-task-item/project-task-item.component';
import { ProjectTaskListComponent } from './pages/project-tasks-page/project-task-list/project-task-list.component';
import { ErrorPageComponent } from './pages/error-page/error-page.component';


@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
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
