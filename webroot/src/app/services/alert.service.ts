import { Injectable } from '@angular/core';
import { Response } from '@angular/http';

@Injectable()
export class AlertService {

  constructor() { }

  public alertError(error: any, stackFunction: string) {
    var alertMessage = 'ERROR! ';
    if (error.status)
      alertMessage += error.status + ' ' + error.statusText + ' \n\n';
    if (error.message)
      alertMessage += error.message + '\n\n';
    alertMessage += 'STACKTRACE: ' + stackFunction + '\n';
    alert(alertMessage);
    console.debug(alertMessage);
  }

  public checkGraphQLResponse(response: Response) {
    var res = response.json();
    if (res.errors) {
      var err = new Error();
      err.message = res.errors[0].message;
      throw err;
    }
  }
}
