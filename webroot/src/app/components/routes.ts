import { Routes, RouterModule } from '@angular/router';

import { AuthManager } from './authmanager';

import { MainComponent } from './main/main.component';
import { ProjectListComponent } from './shared/project-list/project-list.component';
import { SProjectPageComponent } from './pages/s-project-page/s-project-page.component';
import { AuthorizationComponent } from './pages/authorization/authorization.component';
import { HomePageComponent } from './pages/home-page/home-page.component';

export const routes: Routes = [
  { path: '', redirectTo: 'main', pathMatch: 'full' },
  { path: 'main', component: MainComponent },
  { path: 'projects', component: ProjectListComponent },
  { path: 'project', component: SProjectPageComponent},
  { path: 'auth', component: AuthorizationComponent, canActivate: [AuthManager]},
  { path: 'home', component: HomePageComponent, canActivate: [AuthManager]}
];

export const AppRouterProvider = RouterModule.forRoot(routes);