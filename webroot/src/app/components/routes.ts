import { Routes, RouterModule } from '@angular/router';

import { AuthManager } from './../managers/authmanager';

import { MainComponent } from './main/main.component';
import { ProjectListComponent } from './shared/project-list/project-list.component';
import { StudentProjectPageComponent } from './pages/student-project-page/student-project-page.component';
import { ProjectListPageComponent } from './pages/project-list-page/project-list-page.component';
import { AuthorizationPageComponent } from './pages/authorization-page/authorization-page.component';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { RegistrationPageComponent } from './pages/registration-page/registration-page.component';
import { ValidationPageComponent } from './pages/registration-page/validation-page/validation-page.component';
import { AuthorPublicPageComponent } from './pages/author-public-page/author-public-page.component';
import { StudentPublicPageComponent } from './pages/student-public-page/student-public-page.component';

import { MainNewsPageComponent } from './pages/main-news-page/main-news-page.component';

export const routes: Routes = [
  { path: '', redirectTo: 'main', pathMatch: 'full' },
  { path: 'main', component: MainComponent },
  { path: 'projects', component: ProjectListPageComponent },
  { path: 'project/:id', component: StudentProjectPageComponent },
  { path: 'auth', component: AuthorizationPageComponent, canActivate: [AuthManager] },
  { path: 'home', component: HomePageComponent, canActivate: [AuthManager] },
  { path: 'registration', component: RegistrationPageComponent },
  { path: 'registration/validate', component: ValidationPageComponent, canActivate: [AuthManager] },
  { path: 'author/:id', component: AuthorPublicPageComponent },
  { path: 'student/:id', component: StudentPublicPageComponent },
  { path: 'news', component: MainNewsPageComponent },
  //otherwise main
  { path: '**', redirectTo: 'main' }
];

export const AppRouterProvider = RouterModule.forRoot(routes);