import { ProjectShort } from './project-short'

export class CurrentUser {

    constructor() {
        this.user = {
            id: 0,
            nickname: "",
            description: "",
            avatar: "",
            permission_level: 0,
        };
        this.member_of = [];
        this.master_of = [];
    }

    user: {
        id: number,
        nickname: string,
        description: string,
        avatar: string,
        permission_level: number
    };
    member_of: ProjectShort[];
    master_of: ProjectShort[];
}