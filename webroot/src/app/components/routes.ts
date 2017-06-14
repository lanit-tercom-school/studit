import { Routes, RouterModule } from '@angular/router';

import { PathManager } from 'services/path.manager';
import { MainComponent } from './main/main.component';
import { ProjectListComponent } from './shared/project-list/project-list.component';
import { NewsListComponent } from './shared/news-list/news-list.component';
import { HomePageTeacherComponent } from './pages/home-page-teacher/home-page-teacher.component'
import { HomePageStudentComponent } from './pages/home-page-student/home-page-student.component'
import { StudentProjectPageComponent } from './pages/student-project-page/student-project-page.component';
import { ProjectListPageComponent } from './pages/project-list-page/project-list-page.component';
import { AuthorizationPageComponent } from './pages/authorization-page/authorization-page.component';
import { RegistrationPageComponent } from './pages/registration-page/registration-page.component';
import { ValidationPageComponent } from './pages/registration-page/validation-page/validation-page.component';
import { AuthorPublicPageComponent } from './pages/author-public-page/author-public-page.component';
import { StudentPublicPageComponent } from './pages/student-public-page/student-public-page.component';
import { UserSettingsPageComponent } from './pages/user-settings-page/user-settings-page.component';
import { MainNewsPageComponent } from './pages/main-news-page/main-news-page.component';
import { ProjectTasksPageComponent } from './pages/project-tasks-page/project-tasks-page.component';
import { HomeProjectsViewComponent } from './pages/home-projects-view/home-projects-view.component';
import { MainFullNewsPageComponent } from './pages/main-full-news-page/main-full-news-page.component';
import { ErrorPageComponent } from './pages/error-page/error-page.component';

export const routes: Routes = [
  { path: '', redirectTo: 'main', pathMatch: 'full', canActivate: [PathManager] },
  { path: 'home', component: ErrorPageComponent, canActivate: [PathManager] },
  { path: 'main', component: MainComponent, canActivate: [PathManager] },
  { path: 'projects', component: ProjectListPageComponent },
  { path: 'auth', component: AuthorizationPageComponent, canActivate: [PathManager] },
  { path: 'student', loadChildren: 'student/student.module#StudentModule' },
  { path: 'teacher', loadChildren: 'teacher/teacher.module#TeacherModule' },
  { path: 'admin', loadChildren: 'admin/admin.module#AdminModule' },
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
  { path: 'students/:id', component: StudentPublicPageComponent },
  { path: 'students/:id/settings', component: UserSettingsPageComponent },
  { path: 'news', component: MainNewsPageComponent },
  { path: 'news/:id', component: MainFullNewsPageComponent },
  { path: '**', component: ErrorPageComponent },
];

export const AppRouterProvider = RouterModule.forRoot(routes);
