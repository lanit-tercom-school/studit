import { ProjectShort } from './project-short'
import { PermLevel } from './permission-level.enum';

export class CurrentUser {

    constructor() {
        this.User = {
            Id: 0,
            Nickname: "",
            Description: "",
            Avatar: "",
            Login: ""
        };
        this.Token = "";
        this.PermissionLevel = PermLevel.Student;
        this.DataOfExpiration = ""; 
    }

    User: {
        Id: number,
        Nickname: string,
        Description: string,
        Avatar: string,
        Login: string
    };
    Token: String;
    PermissionLevel: PermLevel;
    DataOfExpiration: String;

}