import { Routes, RouterModule } from '@angular/router';

import { PathManager } from 'services/path.manager';
import { MainComponent } from './components/main-page/main.component';
import { ProjectListComponent } from './components/project-list/project-list.component';
import { NewsListComponent } from './components/news-list/news-list.component';
import { StudentProjectPageComponent } from './components/student-project-page/student-project-page.component';
import { ProjectListPageComponent } from './components/project-list-page/project-list-page.component';
import { AuthorizationPageComponent } from './components/authorization-page/authorization-page.component';
import { RegistrationPageComponent } from './components/registration-page/registration-page.component';
import { ValidationPageComponent } from './components/registration-page/validation-page/validation-page.component';
import { AuthorPublicPageComponent } from './components/author-public-page/author-public-page.component';
import { StudentPublicPageComponent } from './components/student-public-page/student-public-page.component';
import { UserSettingsPageComponent } from './components/user-settings-page/user-settings-page.component';
import { MainNewsPageComponent } from './components/main-news-page/main-news-page.component';
import { ProjectTasksPageComponent } from './components/project-tasks-page/project-tasks-page.component'; 
import { MainFullNewsPageComponent } from './components/main-full-news-page/main-full-news-page.component';
import { ErrorPageComponent } from './components/error-page/error-page.component';

export const routes: Routes = [
  { path: '', redirectTo: 'main', pathMatch: 'full', canActivate: [PathManager] },
  { path: 'home', component: ErrorPageComponent, canActivate: [PathManager] },
  { path: 'main', component: MainComponent, canActivate: [PathManager] },
  { path: 'projects', component: ProjectListPageComponent },
  { path: 'auth', component: AuthorizationPageComponent, canActivate: [PathManager] },
  { path: 'student', loadChildren: 'student-module/student.module#StudentModule' },
  { path: 'teacher', loadChildren: 'teacher-module/teacher.module#TeacherModule' },
  { path: 'admin', loadChildren: 'admin-module/admin.module#AdminModule' },
  {
    path: 'project/:id',
    children: [{
      path: '',
      pathMatch: 'full',
      component: StudentProjectPageComponent,
    }, {
      path: 'tasks',
      component: ProjectTasksPageComponent,
    }]
  },
  { path: 'registration', component: RegistrationPageComponent, canActivate: [PathManager] },
  { path: 'registration/validate', component: ValidationPageComponent, canActivate: [PathManager] },
  { path: 'author/:id', component: AuthorPublicPageComponent },
  { path: 'author/:id/settings', component: UserSettingsPageComponent },
  { path: 'user/:id', component: StudentPublicPageComponent },
  { path: 'user/:id/settings', component: UserSettingsPageComponent },
  { path: 'news', component: MainNewsPageComponent },
  { path: 'news/:id', component: MainFullNewsPageComponent },
  { path: '**', component: ErrorPageComponent },
];

export const AppRouterProvider = RouterModule.forRoot(routes);
