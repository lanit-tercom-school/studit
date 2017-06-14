import { Routes, RouterModule } from '@angular/router';

import { PathManager } from 'services/path.manager';
import { MainComponent } from 'components/main/main.component';
import { ProjectListComponent } from 'components/shared/project-list/project-list.component';
import { NewsListComponent } from 'components/shared/news-list/news-list.component';
import { HomePageTeacherComponent } from 'components/pages/home-page-teacher/home-page-teacher.component'
import { HomePageStudentComponent } from 'components/pages/home-page-student/home-page-student.component'
import { StudentProjectPageComponent } from 'components/pages/student-project-page/student-project-page.component';
import { ProjectListPageComponent } from 'components/pages/project-list-page/project-list-page.component';
import { AuthorizationPageComponent } from 'components/pages/authorization-page/authorization-page.component';
import { RegistrationPageComponent } from 'components/pages/registration-page/registration-page.component';
import { ValidationPageComponent } from 'components/pages/registration-page/validation-page/validation-page.component';
import { AuthorPublicPageComponent } from 'components/pages/author-public-page/author-public-page.component';
import { StudentPublicPageComponent } from 'components/pages/student-public-page/student-public-page.component';
import { UserSettingsPageComponent } from 'components/pages/user-settings-page/user-settings-page.component';
import { MainNewsPageComponent } from 'components/pages/main-news-page/main-news-page.component';
import { ProjectTasksPageComponent } from 'components/pages/project-tasks-page/project-tasks-page.component';
import { HomeProjectsViewComponent } from 'components/pages/home-projects-view/home-projects-view.component';
import { MainFullNewsPageComponent } from 'components/pages/main-full-news-page/main-full-news-page.component';
import { ErrorPageComponent } from 'components/pages/error-page/error-page.component';

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
