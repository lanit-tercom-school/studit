import { Routes, RouterModule } from '@angular/router';

import { PathManager } from 'services/path.manager';

import { MainComponent } from './components/main-page/main.component';
//import { ProjectListComponent } from './components/project-list/project-list.component';
//import { NewsListComponent } from './components/news-list/news-list.component';
import { ProjectListPageComponent } from './components/project-list-page/project-list-page.component';
import { ProjectPageComponent } from './components/project-page/project-page.component';
import { AuthorizationPageComponent } from './components/authorization-page/authorization-page.component';
import { RegistrationPageComponent } from './components/registration-page/registration-page.component';
import { ValidationPageComponent } from './components/registration-page/validation-page/validation-page.component';
import { UserPublicPageComponent } from './components/user-public-page/user-public-page.component';
import { MainNewsPageComponent } from './components/main-news-page/main-news-page.component';
import { MainFullNewsPageComponent } from './components/main-full-news-page/main-full-news-page.component';
import { ErrorPageComponent } from './components/error-page/error-page.component';

export const routes: Routes = [
  { path: '', redirectTo: 'main', pathMatch: 'full', canActivate: [PathManager] },
  //{ path: 'home',  component: UserPublicPageComponent, canActivate: [PathManager] },
  { path: 'main', component: MainComponent, canActivate: [PathManager] },
  { path: 'projects', component: ProjectListPageComponent },
  { path: 'project/:id', component: ProjectPageComponent, canActivate: [PathManager] },
  { path: 'auth', component: AuthorizationPageComponent, canActivate: [PathManager] },
  { path: 'student', loadChildren: 'student-module/student.module#StudentModule',canActivate: [PathManager] },
  { path: 'teacher', loadChildren: 'teacher-module/teacher.module#TeacherModule',canActivate: [PathManager] },
  // { path: 'admin', loadChildren: 'admin-module/admin.module#AdminModule',canActivate: [PathManager] },
  { path: 'registration', component: RegistrationPageComponent, canActivate: [PathManager] },
  { path: 'registration/validate', component: ValidationPageComponent, canActivate: [PathManager] },
  // { path: 'user/:id', component: UserPublicPageComponent,  canActivate: [PathManager] },
  { path: 'news', component: MainNewsPageComponent },
  { path: 'news/:id', component: MainFullNewsPageComponent },
  { path: '**', component: ErrorPageComponent },
];

export const AppRouterProvider = RouterModule.forRoot(routes);
