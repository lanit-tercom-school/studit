import { Injectable } from '@angular/core';

@Injectable()
export class AlertService {

  constructor() { }

  public alertError(error: any, stackFunction: string) {
    if (error.status) {
      alert('Ошибка! ' + error.status + ' ' + error.statusText);
      console.debug('ERROR: status ' + error.status + ' ' + error.statusText);
    }
    if (error.message) {
      alert('Ошибка! ' + error.message);
      console.debug('ERROR: status ' + error.message);
    }
    console.debug(stackFunction);
  }

}
