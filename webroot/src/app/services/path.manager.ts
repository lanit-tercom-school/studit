import { Injectable } from '@angular/core';
import { CanActivate, Router, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';

import { AuthService } from 'services/auth.service'

@Injectable()
export class PathManager implements CanActivate {

    constructor(private router: Router) { }

    canActivate(next: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
        if (next.url[0].path === 'main') {
            if (window.localStorage.getItem('current_user')) {
                console.log('You are already logged in. Main page is forbidden');
                this.router.navigate(['/home']);
                return false;
            }
            else
                return true;
        }
        else if (next.url[0].path === 'auth') {
            if (window.localStorage.getItem('current_user')) {
                console.log('You are already logged in');
                this.router.navigate(['/home']);
                return false;
            }
            else
                return true;
        }
        else if (next.url[0].path === 'registration') {
            if (window.localStorage.getItem('current_user')) {
                console.log('You are already logged in');
                this.router.navigate(['/home']);
                return false;
            }
            else
                return true;
        }
        else if (next.url[0].path === 'registration' && next.url[1].path === 'validate') {
            if (window.localStorage.getItem('validation_code')) {
                return true;
            }
            else {
                console.log('Restricted');
                this.router.navigate(['/registration']);
                return false;
            }
        }
        else if (next.url[0].path === 'home') {
            if (window.localStorage.getItem('current_user')) {
                if (JSON.parse(localStorage.getItem('current_user')).perm_lvl === 0) {
                    this.router.navigate(['student/home']);
                }
                else {
                    this.router.navigate(['teacher/home']);
                }
            }
            else {
                console.log('You must be logged in');
                return false;
            }
        }
        else if (next.url[0].path === 'teacher') {
            if (window.localStorage.getItem('current_user')) {
                if (JSON.parse(localStorage.getItem('current_user')).perm_lvl === 0) {
                    return false;
                }
                else if (JSON.parse(localStorage.getItem('current_user')).perm_lvl === 1) {
                    return true;
                }
                else if (JSON.parse(localStorage.getItem('current_user')).perm_lvl === 2) {
                    return false;
                }
            }
            this.router.navigate(['/error']);
        }
        else if (next.url[0].path === 'student') {
            if (window.localStorage.getItem('current_user')) {
                if (JSON.parse(localStorage.getItem('current_user')).perm_lvl === 0) {
                    return true;
                }
                else if (JSON.parse(localStorage.getItem('current_user')).perm_lvl === 1) {
                    return false;
                }
                else if (JSON.parse(localStorage.getItem('current_user')).perm_lvl === 2) {
                    return false;
                }
            }
            this.router.navigate(['/error']);
        }
        else if (next.url[0].path === 'admin') {
            if (window.localStorage.getItem('current_user')) {
                if (JSON.parse(localStorage.getItem('current_user')).perm_lvl === 0) {
                    return false;
                }
                else if (JSON.parse(localStorage.getItem('current_user')).perm_lvl === 1) {
                    return false;
                }
                else if (JSON.parse(localStorage.getItem('current_user')).perm_lvl === 2) {
                    return true
                }
            }
            this.router.navigate(['/error']);
        }
        return true;
    }
}
