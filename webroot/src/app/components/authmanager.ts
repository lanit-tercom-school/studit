import { Injectable } from '@angular/core';

import { CanActivate, Router, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';

import { AuthService } from './../services/auth.service'

@Injectable()
export class AuthManager implements CanActivate {

    constructor(private router: Router) { }

    canActivate(next: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
        if (next.url[0].path == 'auth') {
            if (window.localStorage.getItem('auth_key')) {
                console.log('You are already logged in');
                this.router.navigate(['/home']);
                return false;
            }
            else
                return true;
        }
        else if (next.url[0].path == 'home') {
            if (window.localStorage.getItem('auth_key')) {
                return true;
            }
            else {
                console.log('You must be logged in');
                this.router.navigate(['/auth']);
                return false;
            }
        }

        return true;
    }
}