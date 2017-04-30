import { Injectable } from '@angular/core';
import { User } from './../components/pages/authorization-page/user';
import { Http, Headers, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import { environment } from '../../environments/environment';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

@Injectable()
export class AuthService {

    constructor(private http: Http) { }

    authenticatenow(user: User) {

        //var headers = new Headers();

        //headers.append('Content-Type', 'application/json');

        return this.http.post(environment.apiUrl + '/v1/auth/signin/', JSON.stringify(user)/*, { headers: headers }*/)
            .map((response: Response) => {
                // successful login => getting jwt
                let res = response.json();
                if (res && res.token) {
                    // save data for keeping user logged in
                    res.login = user.Login;
                    localStorage.setItem('current_user', JSON.stringify(res));
                }
            });
    }

    unauthentificatenow() {
        localStorage.removeItem("current_user");
    }


}
