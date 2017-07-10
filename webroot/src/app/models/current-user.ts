import { ProjectShort } from './project-short'
import { PermLevel } from '../shared-components/permission-level.enum';

export class CurrentUser {

    constructor() {
        this.user = {
            id: 0,
            nickname: "",
            description: "",
            avatar: "",
            permission_level: PermLevel.Student,
        };
        this.member_of = [];
        this.master_of = [];
    }

    user: {
        id: number,
        nickname: string,
        description: string,
        avatar: string,
        permission_level: PermLevel
    };
    member_of: ProjectShort[];
    master_of: ProjectShort[];
}